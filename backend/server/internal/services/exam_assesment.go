package services

import (
	"common/constants"
	"common/ent"
	commonRepositories "common/repositories"
	commonServices "common/services"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"server/pkg/models"
	"strings"

	goaway "github.com/TwiN/go-away"
	"github.com/redis/go-redis/v9"
)

type ExamAssesmentService struct {
	accessService           *AccessService
	promptService           *PromptService
	examGenerationService   *ExamGenerationService
	profanityService        *commonServices.ProfanityService
	generatedExamRepository *commonRepositories.GeneratedExamRepository
	examAttemptRepository   *commonRepositories.ExamAttemptRepository
	examAssesmentRepository *commonRepositories.ExamAssesmentRepository
}

type DescriptiveExamAssesmentRequest struct {
	CompletedSeconds int    `json:"completed_seconds" validate:"required"`
	Content          string `json:"content" validate:"required"`
}

func NewExamAssesmentService(redisClient *redis.Client, dbClient *ent.Client) *ExamAssesmentService {
	accessService := NewAccessService(dbClient)
	promptService := NewPromptService()
	profanityService := commonServices.NewProfanityService()
	generatedExamRepository := commonRepositories.NewGeneratedExamRepository(dbClient)
	examGenerationService := NewExamGenerationService(redisClient, dbClient)
	examAttemptRepository := commonRepositories.NewExamAttemptRepository(dbClient)
	examAssesmentRepository := commonRepositories.NewExamAssesmentRepository(dbClient)

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

func (e *ExamAssesmentService) StartNewDescriptiveAssesment(ctx context.Context, generatedExamId int, attempt *ent.ExamAttempt, request *DescriptiveExamAssesmentRequest, userId string, isOpen bool) (*models.AssessmentDetails, error) {
	generatedExam, err := e.generatedExamRepository.GetById(ctx, generatedExamId, isOpen)
	if err != nil {
		return nil, err
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

	assesmentModel := commonRepositories.AssesmentModel{
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

func (e *ExamAssesmentService) GetAssesmentById(ctx context.Context, assesmentId int, userId string, isOpen bool) (*models.AssessmentDetails, error) {
	assessment, err := e.examAssesmentRepository.GetById(ctx, assesmentId, userId)
	if err != nil {
		return nil, err
	}

	attempt, err := e.examAttemptRepository.GetById(ctx, assessment.Edges.Attempt.ID, userId, isOpen)
	if err != nil {
		return nil, err
	}

	if !isOpen {
		hasAccess, err := e.accessService.UserHasAccessToExam(ctx, attempt.Edges.Generatedexam.Edges.Exam.ID, userId)
		if err != nil {
			return nil, fmt.Errorf("failed to check access: %w", err)
		}

		if !hasAccess {
			return nil, errors.New("forbidden")
		}
	}

	assessmentModel := &models.AssessmentDetails{
		Id:                assessment.ID,
		CompletedSeconds:  assessment.CompletedSeconds,
		Status:            assessment.Status.String(),
		RawUserSubmission: assessment.RawUserSubmission,
		CreatedAt:         assessment.CreatedAt,
		UpdatedAt:         assessment.UpdatedAt,
	}

	if assessment.RawAssesmentData == nil {
		return assessmentModel, nil
	}

	assessmentModel.RawAssesmentData = assessment.RawAssesmentData

	return assessmentModel, nil
}

func (e *ExamAssesmentService) GetExamAssessments(ctx context.Context, generatedExamId int, userId string, isOpen bool) ([]models.AssessmentDetails, error) {
	assessments, err := e.examAssesmentRepository.GetByExam(ctx, generatedExamId, userId)

	if err != nil {
		return nil, err
	}

	generatedExam, err := e.generatedExamRepository.GetById(ctx, generatedExamId, isOpen)
	if err != nil {
		return nil, err
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

	assessmentsList := make([]models.AssessmentDetails, 0, len(assessments))

	for _, assessment := range assessments {
		assessmentModel := models.AssessmentDetails{
			Id:               assessment.ID,
			CompletedSeconds: assessment.CompletedSeconds,
			Status:           assessment.Status.String(),
			CreatedAt:        assessment.CreatedAt,
			UpdatedAt:        assessment.UpdatedAt,
		}

		assessmentsList = append(assessmentsList, assessmentModel)
	}
	return assessmentsList, nil
}

func (e *ExamAssesmentService) AssessDescriptiveExam(ctx context.Context, generatedExamId, assessmentId int, content string, userId string, isOpen bool) {
	assesmentModel := &commonRepositories.AssesmentModel{}
	generatedExamData, err := e.generatedExamRepository.GetById(ctx, generatedExamId, isOpen)
	if err != nil {
		log.Println("error getting generated exam", err)
		e.updateAssessment(ctx, assessmentId, commonRepositories.AssesmentModel{Status: constants.ASSESSMENT_REJECTED})
		return
	}

	if !isOpen {
		hasAccess, err := e.accessService.UserHasAccessToExam(ctx, generatedExamData.Edges.Exam.ID, userId)
		if err != nil {
			log.Println("error getting exam", err)
			e.updateAssessment(ctx, assessmentId, commonRepositories.AssesmentModel{Status: constants.ASSESSMENT_REJECTED})
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
		e.updateAssessment(ctx, assessmentId, commonRepositories.AssesmentModel{Status: constants.ASSESSMENT_REJECTED})
		return
	}

	jsonData, err := json.Marshal(generatedExam.RawExamData)
	if err != nil {
		log.Println("error processing exam data", err)
		e.updateAssessment(ctx, assessmentId, commonRepositories.AssesmentModel{Status: constants.ASSESSMENT_REJECTED})
		return
	}

	var descriptiveExam models.DescriptiveExam
	err = json.Unmarshal(jsonData, &descriptiveExam)
	if err != nil {
		log.Println("error processing exam data", err)
		e.updateAssessment(ctx, assessmentId, commonRepositories.AssesmentModel{Status: constants.ASSESSMENT_REJECTED})
		return
	}

	if e.profanityService.IsProfane(content) {
		e.updateAssessment(ctx, assessmentId, commonRepositories.AssesmentModel{Status: constants.ASSESSMENT_REJECTED, RawAssessmentData: map[string]interface{}{
			"profanity_check": "detected",
			"profane_content": goaway.ExtractProfanity(content),
		}})
		return
	}

	prompt := fmt.Sprintf(`
Evaluate the following "%s" based on the topic: “%s”.
Criteria to consider:

	•	Grammar accuracy.
	•	Proper use of punctuation.
	•	Relevance to the given topic.
	•	Word count should not exceed "%s" words (only count words, exclude special characters).
	•	Do Not visit any URLs provided in Content.
	•	Make sure rating is based only on content provided, and use the provided criteria to calculate it

Scoring: Provide a rating out of "%s" marks based on the above criteria. Should always be between 0 and maximum marks

Output Requirements:

	•	Return a valid JSON object with the following keys:
	•	"rating": A string representing the rating. 
	•	"strengths": An array of strings highlighting the content’s strengths.
	•	"weaknesses": An array of strings pointing out the content’s weaknesses.
	•	"corrected_version": A string with the corrected version of the content.
	•	The entire output should be a single-line string with no extra spaces, newlines, or formatting, ensuring it can be parsed as valid JSON.

Content to evaluate:

	“%s”
`, descriptiveExam.Type, descriptiveExam.Topic, descriptiveExam.MaxNumberOfWordsAllowed, descriptiveExam.TotalMarks, content)

	response, err := e.promptService.GetPromptResult(ctx, prompt, constants.PRO_15)
	if err != nil {
		e.updateAssessment(ctx, assessmentId, commonRepositories.AssesmentModel{Status: constants.ASSESSMENT_REJECTED})
		log.Printf("Error getting prompt result: %v", err)
		return
	}

	if strings.Contains(response, "FinishReasonSafety") {
		e.updateAssessment(ctx, assessmentId, commonRepositories.AssesmentModel{Status: constants.ASSESSMENT_REJECTED, RawAssessmentData: map[string]interface{}{
			"profanity_check": "detected",
		}})

		return
	}

	var rawJsonData map[string]interface{}
	err = json.Unmarshal([]byte(response), &rawJsonData)
	if err != nil {
		log.Println("error response from AI service", err)
		e.updateAssessment(ctx, assessmentId, commonRepositories.AssesmentModel{Status: constants.ASSESSMENT_REJECTED})
		return
	}

	var assessmentResult models.DescriptiveExamAssessmentResult
	err = json.Unmarshal([]byte(response), &assessmentResult)
	if err != nil {
		log.Println("response from AI service does not match criteria", err)
		e.updateAssessment(ctx, assessmentId, commonRepositories.AssesmentModel{Status: constants.ASSESSMENT_REJECTED})
		return
	}

	assesmentModel.Status = constants.ASSESSMENT_COMPLETED
	assesmentModel.RawAssessmentData = rawJsonData
	e.examAssesmentRepository.Update(ctx, assessmentId, *assesmentModel)
	log.Println("Processed Assessment", assessmentId)
}

func (e *ExamAssesmentService) updateAssessment(ctx context.Context, assessmentId int, assesmentModel commonRepositories.AssesmentModel) {

	update_err := e.examAssesmentRepository.Update(ctx, assessmentId, assesmentModel)
	if update_err != nil {
		log.Printf("Error updating status %v", update_err)
	}
}
