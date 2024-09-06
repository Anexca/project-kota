package services

import (
	"common/constants"
	"common/ent"
	commonRepositories "common/repositories"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"server/pkg/models"

	"github.com/redis/go-redis/v9"
)

type ExamAssesmentService struct {
	promptService           *PromptService
	examGenerationService   *ExamGenerationService
	examAttemptRepository   *commonRepositories.ExamAttemptRepository
	examAssesmentRepository *commonRepositories.ExamAssesmentRepository
}

type DescriptiveExamAssesmentRequest struct {
	CompletedSeconds int    `json:"completed_seconds" validate:"required"`
	Content          string `json:"content" validate:"required"`
}

func NewExamAssesmentService(redisClient *redis.Client, dbClient *ent.Client) *ExamAssesmentService {
	promptService := NewPromptService()
	examGenerationService := NewExamGenerationService(redisClient, dbClient)
	examAttemptRepository := commonRepositories.NewExamAttemptRepository(dbClient)
	examAssesmentRepository := commonRepositories.NewExamAssesmentRepository(dbClient)

	return &ExamAssesmentService{
		promptService:           promptService,
		examGenerationService:   examGenerationService,
		examAttemptRepository:   examAttemptRepository,
		examAssesmentRepository: examAssesmentRepository,
	}
}

func (e *ExamAssesmentService) StartNewDescriptiveAssesment(ctx context.Context, generatedExamId int, attempt *ent.ExamAttempt, request *DescriptiveExamAssesmentRequest) (*models.AssessmentDetails, error) {
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
		e.AssessDescriptiveExam(bgCtx, generatedExamId, assessment.ID, request.Content)
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

func (e *ExamAssesmentService) GetAssesmentById(ctx context.Context, assesmentId int, userId string) (models.AssessmentDetails, error) {
	assessment, err := e.examAssesmentRepository.GetById(ctx, assesmentId, userId)
	if err != nil {
		return models.AssessmentDetails{}, err
	}

	assessmentModel := models.AssessmentDetails{
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

func (e *ExamAssesmentService) GetExamAssessments(ctx context.Context, generatedExamId int, userId string) ([]models.AssessmentDetails, error) {
	assessments, err := e.examAssesmentRepository.GetByExam(ctx, generatedExamId, userId)

	if err != nil {
		return nil, err
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

func (e *ExamAssesmentService) AssessDescriptiveExam(ctx context.Context, generatedExamId, assessmentId int, content string) {
	assesmentModel := &commonRepositories.AssesmentModel{}

	generatedExam, err := e.examGenerationService.GetGeneratedExamById(ctx, generatedExamId)
	if err != nil {
		log.Println("error getting exam", err)
		assesmentModel.Status = constants.ASSESSMENT_REJECTED
		e.examAssesmentRepository.Update(ctx, assessmentId, *assesmentModel)
		return
	}

	jsonData, err := json.Marshal(generatedExam.RawExamData)
	if err != nil {
		log.Println("error processing exam data", err)
		assesmentModel.Status = constants.ASSESSMENT_REJECTED
		e.examAssesmentRepository.Update(ctx, assessmentId, *assesmentModel)
		return
	}

	var descriptiveExam models.DescriptiveExam
	err = json.Unmarshal(jsonData, &descriptiveExam)
	if err != nil {
		log.Println("error processing exam data", err)
		assesmentModel.Status = constants.ASSESSMENT_REJECTED
		e.examAssesmentRepository.Update(ctx, assessmentId, *assesmentModel)
		return
	}

	prompt := fmt.Sprintf(`Can you evaluate %s on topic: %s
					Points to check for:
					- Grammar
					- Punctuation
					- Relevance to the topic
					- Maximum word count: 250. Make sure you count the number of WORDS and not punctuation marks or numbers or anything
					Provide a rating out of 25 marks based on these criteria.
	
					Output Requirements:
					- A JSON with keys "rating" with string value, "strengths" with array of string, "weakness"  with array of string, "corrected_version"  with string value
					- Make sure the output is one line string with no formatting such that it should be JSON parsed
	
					Content to evaluate:
					%s`, descriptiveExam.Type, descriptiveExam.Topic, content)

	response, err := e.promptService.GetPromptResult(ctx, prompt, constants.FLASH_15)
	if err != nil {
		assesmentModel.Status = constants.ASSESSMENT_REJECTED
		e.examAssesmentRepository.Update(ctx, assessmentId, *assesmentModel)
		log.Printf("Error getting prompt result: %v", err)
		return
	}

	var rawJsonData map[string]interface{}
	err = json.Unmarshal([]byte(response), &rawJsonData)
	if err != nil {
		log.Println("error response from AI service", err)
		assesmentModel.Status = constants.ASSESSMENT_REJECTED
		e.examAssesmentRepository.Update(ctx, assessmentId, *assesmentModel)
		return
	}

	var assessmentResult models.DescriptiveExamAssessmentResult
	err = json.Unmarshal([]byte(response), &assessmentResult)
	if err != nil {
		log.Println("response from AI service does not match criteria", err)
		assesmentModel.Status = constants.ASSESSMENT_REJECTED
		e.examAssesmentRepository.Update(ctx, assessmentId, *assesmentModel)
		return
	}

	assesmentModel.Status = constants.ASSESSMENT_COMPLETED
	assesmentModel.RawAssessmentData = rawJsonData
	e.examAssesmentRepository.Update(ctx, assessmentId, *assesmentModel)
	log.Println("Processed Assessment", assessmentId)
}
