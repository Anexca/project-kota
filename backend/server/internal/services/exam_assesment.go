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

type DescriptiveExamAssessmentResult struct {
	Rating           string   `json:"string"`
	Strengths        []string `json:"strengths"`
	Weakness         []string `json:"weakness"`
	CorrectedVersion string   `json:"corrected_version"`
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

func (e *ExamAssesmentService) StartNewDescriptiveAssesment(ctx context.Context, generatedExamId int, attempt *ent.ExamAttempt, request *DescriptiveExamAssesmentRequest) (*ent.ExamAssesment, error) {
	assesmentModel := commonRepositories.AssesmentModel{
		CompletedSeconds: request.CompletedSeconds,
		Status:           constants.ASSESSMENT_PENDING,
	}

	assessment, err := e.examAssesmentRepository.Create(ctx, attempt.ID, assesmentModel)
	if err != nil {
		return nil, err
	}

	go func() {
		bgCtx := context.Background()
		e.AssessDescriptiveExam(bgCtx, generatedExamId, assessment.ID, request.Content)
	}()

	return assessment, nil
}

func (e *ExamAssesmentService) GetAssesmentById(ctx context.Context, assesmentId int, userId string) (*ent.ExamAssesment, error) {
	return e.examAssesmentRepository.GetById(ctx, assesmentId, userId)
}

func (e *ExamAssesmentService) AssessDescriptiveExam(ctx context.Context, generatedExamId, assessmentId int, content string) {
	generatedExam, err := e.examGenerationService.GetGeneratedExamById(ctx, generatedExamId)
	if err != nil {
		log.Println("error getting exam")
		return
	}

	jsonData, err := json.Marshal(generatedExam.RawExamData)
	if err != nil {
		log.Println("error processing exam data")
		return
	}

	var descriptiveExam models.DescriptiveExam
	err = json.Unmarshal(jsonData, &descriptiveExam)
	if err != nil {
		log.Println("error processing exam data")
		return
	}

	assesmentModel := commonRepositories.AssesmentModel{
		Status: constants.ASSESSMENT_PENDING,
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
		log.Printf("Error getting prompt result: %v", err)
		return
	}

	var rawJsonData map[string]interface{}
	err = json.Unmarshal([]byte(response), &rawJsonData)
	if err != nil {
		log.Println("error response from AI service")
		assesmentModel.Status = constants.ASSESSMENT_REJECTED
		e.examAssesmentRepository.Update(ctx, assessmentId, assesmentModel)
		return
	}

	var assessmentResult DescriptiveExamAssessmentResult
	err = json.Unmarshal([]byte(response), &assessmentResult)
	if err != nil {
		log.Println("response from AI service does not match criteria")
		assesmentModel.Status = constants.ASSESSMENT_REJECTED
		e.examAssesmentRepository.Update(ctx, assessmentId, assesmentModel)
	}

	assesmentModel.Status = constants.ASSESSMENT_COMPLETED
	assesmentModel.RawAssessmentData = rawJsonData
	e.examAssesmentRepository.Update(ctx, assessmentId, assesmentModel)
}
