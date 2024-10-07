package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"strings"

	"common/constants"
	"common/ent"
	"common/ent/exam"

	commonInterfaces "common/interfaces"
	commonRepositories "common/repositories"
	commonServices "common/services"

	goaway "github.com/TwiN/go-away"
	"github.com/redis/go-redis/v9"

	"server/internal/interfaces"
	"server/pkg/models"
)

type ExamAssesmentService struct {
	accessService           interfaces.AccessServiceInterface
	promptService           interfaces.PromptServiceInterface
	examGenerationService   interfaces.ExamGenerationServiceInterface
	profanityService        commonInterfaces.ProfanityServiceInterface
	generatedExamRepository commonInterfaces.GeneratedExamRepositoryInterface
	examAttemptRepository   commonInterfaces.ExamAttemptRepositoryInterface
	examAssesmentRepository commonInterfaces.ExamAssessmentRepositoryInterface
}

func NewExamAssesmentService(
	accessService interfaces.AccessServiceInterface,
	promptService interfaces.PromptServiceInterface,
	profanityService commonInterfaces.ProfanityServiceInterface,
	generatedExamRepository commonInterfaces.GeneratedExamRepositoryInterface,
	examGenerationService interfaces.ExamGenerationServiceInterface,
	examAttemptRepository commonInterfaces.ExamAttemptRepositoryInterface,
	examAssesmentRepository commonInterfaces.ExamAssessmentRepositoryInterface,
) *ExamAssesmentService {
	return &ExamAssesmentService{
		accessService:           accessService,
		promptService:           promptService,
		profanityService:        profanityService,
		generatedExamRepository: generatedExamRepository,
		examGenerationService:   examGenerationService,
		examAttemptRepository:   examAttemptRepository,
		examAssesmentRepository: examAssesmentRepository,
	}
}

func InitExamAssesmentService(redisClient *redis.Client, dbClient *ent.Client) *ExamAssesmentService {
	accessService := InitAccessService(dbClient)
	promptService := NewPromptService()
	profanityService := commonServices.NewProfanityService()
	generatedExamRepository := commonRepositories.NewGeneratedExamRepository(dbClient)
	examGenerationService := InitExamGenerationService(redisClient, dbClient)
	examAttemptRepository := commonRepositories.NewExamAttemptRepository(dbClient)
	examAssesmentRepository := commonRepositories.NewExamAssessmentRepository(dbClient)

	return NewExamAssesmentService(
		accessService,
		promptService,
		profanityService,
		generatedExamRepository,
		examGenerationService,
		examAttemptRepository,
		examAssesmentRepository,
	)
}

func (e *ExamAssesmentService) StartNewDescriptiveAssesment(ctx context.Context, generatedExamId int, attempt *ent.ExamAttempt, request *models.DescriptiveExamAssesmentRequest, userId string, isOpen bool) (*models.AssessmentDetails, error) {
	generatedExam, err := e.generatedExamRepository.GetOpenById(ctx, generatedExamId, isOpen)
	if err != nil {
		log.Printf("Error getting generated exam: %v", err)
		return nil, err
	}

	if generatedExam == nil {
		return nil, errors.New("generated exam not found")
	}

	if !isOpen {
		hasAccess, err := e.accessService.UserHasAccessToExam(ctx, generatedExam.Edges.Exam.ID, userId)
		if err != nil {
			return nil, fmt.Errorf("failed to check access: %w", err)
		}

		if !hasAccess {
			return nil, errors.New("forbidden")
		}
	}

	userSubmission := map[string]interface{}{
		"content": request.Content,
	}

	assesmentModel := commonRepositories.AssessmentModel{
		CompletedSeconds:  request.CompletedSeconds,
		Status:            constants.ASSESSMENT_PENDING,
		RawUserSubmission: userSubmission,
	}

	assessment, err := e.examAssesmentRepository.Create(ctx, attempt.ID, assesmentModel)
	if err != nil {
		return nil, err
	}

	go func() {
		bgCtx := context.Background()
		e.AssessDescriptiveExam(bgCtx, generatedExamId, assessment.ID, request.Content, userId, isOpen)
	}()

	assessmentModel := &models.AssessmentDetails{
		Id:               assessment.ID,
		CompletedSeconds: assessment.CompletedSeconds,
		Status:           assessment.Status.String(),
		CreatedAt:        assessment.CreatedAt,
		UpdatedAt:        assessment.UpdatedAt,
	}

	return assessmentModel, nil
}

func (e *ExamAssesmentService) AssessMCQExam(ctx context.Context, generatedExamId int, attempt *ent.ExamAttempt, request *models.MCQExamAssessmentRequest, userId string, isOpen bool) (*models.AssessmentDetails, error) {
	generatedExam, err := e.fetchGeneratedExam(ctx, generatedExamId, isOpen)
	if err != nil {
		return nil, err
	}

	if err := e.checkAccessForExam(ctx, generatedExam, userId, isOpen); err != nil {
		return nil, err
	}

	mcqExam, err := e.parseMCQExamData(generatedExam.RawExamData)
	if err != nil {
		return nil, err
	}

	resultSummary := e.evaluateMCQResponses(mcqExam, request)
	return e.createAndSaveAssessment(ctx, attempt, resultSummary, generatedExam, *request)
}

func (e *ExamAssesmentService) GetAssesmentById(ctx context.Context, assesmentId int, userId string) (*models.AssessmentDetails, error) {
	assessment, err := e.examAssesmentRepository.GetById(ctx, assesmentId, userId)
	if err != nil {
		return nil, err
	}

	assessmentModel := e.mapAssessmentModel(assessment)
	return assessmentModel, nil
}

func (e *ExamAssesmentService) GetExamAssessments(ctx context.Context, generatedExamId int, userId string) ([]models.AssessmentDetails, error) {
	assessments, err := e.examAssesmentRepository.GetByExam(ctx, generatedExamId, userId)

	if err != nil {
		return nil, err
	}

	assessmentsList := make([]models.AssessmentDetails, 0, len(assessments))

	for _, assessment := range assessments {
		assessmentModel := *e.mapAssessmentModel(assessment)
		assessmentModel.RawAssesmentData = nil
		assessmentModel.RawUserSubmission = nil
		assessmentsList = append(assessmentsList, assessmentModel)
	}
	return assessmentsList, nil
}

func (e *ExamAssesmentService) GetUserMCQExamQuestionQueryResponse(ctx context.Context, assessmentId int, request models.MCQExamQuestionQueryRequest, userId string) (map[string]interface{}, error) {
	assessment, err := e.examAssesmentRepository.GetById(ctx, assessmentId, userId)
	if err != nil {
		return nil, err
	}

	if assessment.Edges.Attempt.Edges.Generatedexam.Edges.Exam.Type == exam.TypeDESCRIPTIVE {
		return nil, errors.New("invalid exam type")
	}

	var examData models.GeneratedMCQExam
	jsonString, err := json.Marshal(assessment.Edges.Attempt.Edges.Generatedexam.RawExamData)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonString, &examData)
	if err != nil {
		return nil, err
	}

	var userSubmissions models.MCQExamAssessmentRequest
	jsonString, err = json.Marshal(assessment.RawUserSubmission)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonString, &userSubmissions)
	if err != nil {
		return nil, err
	}

	var question models.MCQExamQuestion
	var userSubmissionForQuestion models.MCQExamAssessmentRequestModel

	allQuestions := getAllQuestions(examData)

	for _, q := range allQuestions {
		if q.QuestionNumber == request.QuestionNumber {
			question = q
		}
	}

	for _, a := range userSubmissions.AttemptedQuestions {
		if a.QuestionNumber == request.QuestionNumber {
			userSubmissionForQuestion = a
		}
	}

	examContent := examData.ContentGroups

	p := fmt.Sprintf(`Respond to the user's query regarding the question '%v'. 
						- User query: '%s'
						- User selected option index: %v from the option array: %v
						- Content ID: %s (to be checked against the content array: %v)

					Considerations:
					1.
						- If the answer is correct, return a positive response.
						- If the answer is incorrect, return a response indicating the correct answer.
					2. If the query is related to the content of the question, check the provided content ID and respond accordingly.
					3. If the query is unrelated to the current question or the data provided, return 'invalid query.'
					4. Ignore the content if it is empty.
					5. If user query is out of context of the exam question and content provided, return "invalid query" as response

					Output should follow this JSON schema:
					Response = {'response': string}

					Return the Response.
				`, question.Question, request.Query, userSubmissionForQuestion.UserSelectedOptionIndex, question.Options, question.ContentReferenceId, examContent)

	response, err := e.promptService.GetStructuredPromptResult(ctx, p, constants.FLASH_15)
	if err != nil {
		return nil, err
	}

	var jsonResponse map[string]interface{}

	err = json.Unmarshal([]byte(response), &jsonResponse)
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}

func (e *ExamAssesmentService) AssessDescriptiveExam(ctx context.Context, generatedExamId, assessmentId int, content string, userId string, isOpen bool) {
	assesmentModel := &commonRepositories.AssessmentModel{}
	generatedExamData, err := e.generatedExamRepository.GetOpenById(ctx, generatedExamId, isOpen)
	if err != nil {
		log.Println("error getting generated exam", err)
		err = e.updateAssessment(ctx, assessmentId, commonRepositories.AssessmentModel{Status: constants.ASSESSMENT_REJECTED, Remarks: fmt.Sprintf("error getting generated exam, %v", err)})
		if err != nil {
			log.Printf("Error updating status %v", err)
		}

		return
	}

	if !isOpen {
		hasAccess, err := e.accessService.UserHasAccessToExam(ctx, generatedExamData.Edges.Exam.ID, userId)
		if err != nil {
			log.Println("error getting exam", err)
			err = e.updateAssessment(ctx, assessmentId, commonRepositories.AssessmentModel{Status: constants.ASSESSMENT_REJECTED, Remarks: fmt.Sprintf("error checking exam, %v", err)})
			if err != nil {
				log.Printf("Error updating status %v", err)
			}

			return
		}

		if !hasAccess {
			log.Println("user does not have access to assess exam", err)
			return
		}
	}

	generatedExam, err := e.examGenerationService.GetGeneratedExamById(ctx, generatedExamId, userId, isOpen)
	if err != nil {
		log.Println("error getting exam", err)
		err = e.updateAssessment(ctx, assessmentId, commonRepositories.AssessmentModel{Status: constants.ASSESSMENT_REJECTED, Remarks: fmt.Sprintf("error getting exam, %v", err)})
		if err != nil {
			log.Printf("Error updating status %v", err)
		}

		return
	}

	jsonData, err := json.Marshal(generatedExam.RawExamData)
	if err != nil {
		log.Println("error processing exam data", err)
		err = e.updateAssessment(ctx, assessmentId, commonRepositories.AssessmentModel{Status: constants.ASSESSMENT_REJECTED, Remarks: fmt.Sprintf("error processing exam data, %v", err)})
		if err != nil {
			log.Printf("Error updating status %v", err)
		}

		return
	}

	var descriptiveExam models.DescriptiveExam
	err = json.Unmarshal(jsonData, &descriptiveExam)
	if err != nil {
		log.Println("error processing exam data", err)
		err = e.updateAssessment(ctx, assessmentId, commonRepositories.AssessmentModel{Status: constants.ASSESSMENT_REJECTED, Remarks: fmt.Sprintf("error processing exam data, %v", err)})
		if err != nil {
			log.Printf("Error updating status %v", err)
		}

		return
	}

	if e.profanityService.IsProfane(content) {
		err = e.updateAssessment(ctx, assessmentId, commonRepositories.AssessmentModel{Status: constants.ASSESSMENT_REJECTED, RawAssessmentData: map[string]interface{}{
			"profanity_check": "detected",
			"profane_content": goaway.ExtractProfanity(content),
		}, Remarks: fmt.Sprintf("profanities detected in content, %v", goaway.ExtractProfanity(content))})
		if err != nil {
			log.Printf("Error updating status %v", err)
		}

		return
	}

	prompt := fmt.Sprintf(`
Evaluate the following "%s" based on the topic: “%s”.
Criteria to consider for evaluation:
	•	Grammar accuracy.
	•	Proper use of punctuation.
	•	Relevance to the given topic.
	•	Check for structure if evaluating "essay", formatting if evaluating "error"
	•	Word count in content provided should not exceed "%s" words (only count words, exclude special characters, spaces and formatting characters like "\n, \t, \r" etc).
	•	Do Not visit any URLs provided in Content.
	•	Make sure rating is based only on content provided, and use the provided criteria to calculate it

Scoring: 
	•	Provide a rating out of "%s" marks based on the above criteria. 
	•	The rating must always be between 0 and the maximum marks, with full marks awarded if the content is relevant to the topic and there are no or minimal errors.

Output Requirements:
	•	Return a valid JSON object with the following keys:
	•	"rating": A string representing the rating. 
	•	"strengths": An array of strings highlighting the content’s strengths.
	•	"weaknesses": An array of strings pointing out the content’s weaknesses.
	•	"corrected_version": Generate a single-line string with the corrected version of the content. There should be no extra quotes inside the string, and the output should match the formatting of the provided content.
	•	The entire output should be a single-line string with no extra spaces, newlines, or formatting, ensuring it can be parsed as valid JSON.

Content to evaluate:

	“%s”
`, descriptiveExam.Type, descriptiveExam.Topic, descriptiveExam.MaxNumberOfWordsAllowed, descriptiveExam.TotalMarks, content)

	response, err := e.promptService.GetPromptResult(ctx, prompt, constants.PRO_15)
	if err != nil {
		err = e.updateAssessment(ctx, assessmentId, commonRepositories.AssessmentModel{Status: constants.ASSESSMENT_REJECTED, Remarks: fmt.Sprintf("error getting prompt results, %v", err)})
		if err != nil {
			log.Printf("Error updating status %v", err)
		}

		log.Printf("Error getting prompt result: %v", err)
		return
	}

	if strings.Contains(response, "FinishReasonSafety") {
		err = e.updateAssessment(ctx, assessmentId, commonRepositories.AssessmentModel{Status: constants.ASSESSMENT_REJECTED, RawAssessmentData: map[string]interface{}{
			"profanity_check": "detected",
		}, Remarks: "profanity detected by AI"})
		if err != nil {
			log.Printf("Error updating status %v", err)
		}

		return
	}

	var rawJsonData map[string]interface{}
	err = json.Unmarshal([]byte(response), &rawJsonData)
	if err != nil {
		log.Println("error response from AI service", err)
		err = e.updateAssessment(ctx, assessmentId, commonRepositories.AssessmentModel{Status: constants.ASSESSMENT_REJECTED, Remarks: fmt.Sprintf("error response from AI Service, %v", err)})
		if err != nil {
			log.Printf("Error updating status %v", err)
		}

		return
	}

	var assessmentResult models.DescriptiveExamAssessmentResult
	err = json.Unmarshal([]byte(response), &assessmentResult)
	if err != nil {
		log.Println("response from AI service does not match criteria", err)
		err = e.updateAssessment(ctx, assessmentId, commonRepositories.AssessmentModel{Status: constants.ASSESSMENT_REJECTED, Remarks: fmt.Sprintf("response from AI service does not match criteria, %v", err)})
		if err != nil {
			log.Printf("Error updating status %v", err)
		}

		return
	}

	assesmentModel.Status = constants.ASSESSMENT_COMPLETED
	assesmentModel.RawAssessmentData = rawJsonData
	err = e.examAssesmentRepository.Update(ctx, assessmentId, *assesmentModel)
	if err != nil {
		log.Printf("Error updating status %v", err)
	}

	log.Println("Processed Assessment", assessmentId)
}

func (e *ExamAssesmentService) mapAssessmentModel(assessment *ent.ExamAssesment) *models.AssessmentDetails {
	if assessment == nil {
		return nil
	}

	totalMarks := 0
	if assessment.Edges.Attempt != nil &&
		assessment.Edges.Attempt.Edges.Generatedexam != nil &&
		assessment.Edges.Attempt.Edges.Generatedexam.Edges.Exam != nil &&
		assessment.Edges.Attempt.Edges.Generatedexam.Edges.Exam.Edges.Setting != nil {

		totalMarks = assessment.Edges.Attempt.Edges.Generatedexam.Edges.Exam.Edges.Setting.TotalMarks
	}

	status := assessment.Status.String()

	return &models.AssessmentDetails{
		Id:                assessment.ID,
		CompletedSeconds:  assessment.CompletedSeconds,
		ObtainedMarks:     assessment.ObtainedMarks,
		TotalMarks:        totalMarks,
		Status:            status,
		RawAssesmentData:  assessment.RawAssesmentData,
		RawUserSubmission: assessment.RawUserSubmission,
		CreatedAt:         assessment.CreatedAt,
		UpdatedAt:         assessment.UpdatedAt,
	}
}

func (e *ExamAssesmentService) updateAssessment(ctx context.Context, assessmentId int, assesmentModel commonRepositories.AssessmentModel) error {
	return e.examAssesmentRepository.Update(ctx, assessmentId, assesmentModel)
}

func (e *ExamAssesmentService) fetchGeneratedExam(ctx context.Context, examId int, isOpen bool) (*ent.GeneratedExam, error) {
	generatedExam, err := e.generatedExamRepository.GetOpenById(ctx, examId, isOpen)
	if err != nil {
		log.Printf("Error getting generated exam: %v", err)
		return nil, err
	}
	if generatedExam == nil {
		return nil, errors.New("generated exam not found")
	}
	return generatedExam, nil
}

func (e *ExamAssesmentService) checkAccessForExam(ctx context.Context, exam *ent.GeneratedExam, userId string, isOpen bool) error {
	if !isOpen {
		hasAccess, err := e.accessService.UserHasAccessToExam(ctx, exam.Edges.Exam.ID, userId)
		if err != nil {
			return fmt.Errorf("failed to check access: %w", err)
		}
		if !hasAccess {
			return errors.New("forbidden")
		}
	}
	return nil
}

func (e *ExamAssesmentService) parseMCQExamData(rawData map[string]interface{}) (*models.GeneratedMCQExam, error) {
	jsonData, err := json.Marshal(rawData)
	if err != nil {
		return nil, err
	}

	var mcqExam models.GeneratedMCQExam
	err = json.Unmarshal(jsonData, &mcqExam)
	if err != nil {
		return nil, err
	}
	return &mcqExam, nil
}

func (e *ExamAssesmentService) evaluateMCQResponses(mcqExam *models.GeneratedMCQExam, request *models.MCQExamAssessmentRequest) models.MCQExamAssessmentResultSummary {
	resultSummary := models.MCQExamAssessmentResultSummary{
		Attempted: 0,
		Correct:   0,
		Incorrect: 0,
		Accuracy:  0,
	}

	allQuestions := getAllQuestions(*mcqExam)

	for _, aq := range request.AttemptedQuestions {
		for _, eq := range allQuestions {
			if aq.QuestionNumber == eq.QuestionNumber {
				resultSummary.Attempted++
				if isCorrect(aq.UserSelectedOptionIndex, eq.Answer) {
					resultSummary.Correct++
				} else {
					resultSummary.Incorrect++
				}
				break
			}
		}
	}

	if resultSummary.Attempted > 0 {
		accuracy := float64(resultSummary.Attempted-resultSummary.Incorrect) / float64(resultSummary.Attempted) * 100
		resultSummary.Accuracy = math.Round(accuracy*100) / 100 // Round to two decimal places
	}

	return resultSummary
}

func (e *ExamAssesmentService) createAndSaveAssessment(ctx context.Context, attempt *ent.ExamAttempt, summary models.MCQExamAssessmentResultSummary, exam *ent.GeneratedExam, request models.MCQExamAssessmentRequest) (*models.AssessmentDetails, error) {
	obtainedMarks := calculateObtainedMarks(exam.Edges.Exam.Edges.Setting, summary, exam.Edges.Exam.Edges.Setting.NegativeMarking)
	assessmentModel := &commonRepositories.AssessmentModel{
		RawAssessmentData: map[string]interface{}{"summary": summary},
		RawUserSubmission: map[string]interface{}{"attempted_questions": request.AttemptedQuestions},
		ObtainedMarks:     obtainedMarks,
		Status:            constants.ASSESSMENT_COMPLETED,
		CompletedSeconds:  request.CompletedSeconds,
	}

	assessment, err := e.examAssesmentRepository.Create(ctx, attempt.ID, *assessmentModel)
	if err != nil {
		return nil, err
	}

	assessmentMappedModel := e.mapAssessmentModel(assessment)

	if assessmentMappedModel.TotalMarks == 0 {
		assessmentMappedModel.TotalMarks = exam.Edges.Exam.Edges.Setting.TotalMarks
	}

	return assessmentMappedModel, nil
}

func calculateObtainedMarks(examSettings *ent.ExamSetting, summary models.MCQExamAssessmentResultSummary, negativeMarking float64) float64 {
	markPerQuestion := float64(examSettings.TotalMarks) / float64(examSettings.NumberOfQuestions)
	return (float64(summary.Correct) * markPerQuestion) - (negativeMarking * float64(summary.Incorrect))
}

func isCorrect(selected, correct []int) bool {
	correctMap := make(map[int]bool)
	for _, v := range correct {
		correctMap[v] = true
	}
	for _, v := range selected {
		if !correctMap[v] {
			return false
		}
	}
	return true
}

func getAllQuestions(exam models.GeneratedMCQExam) []models.MCQExamQuestion {
	var allQuestions []models.MCQExamQuestion

	for _, questions := range exam.Sections {
		allQuestions = append(allQuestions, questions...)
	}

	return allQuestions
}
