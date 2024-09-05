package services

import (
	"common/constants"
	"common/ent"
	commonRepositories "common/repositories"
	"context"
	"encoding/json"
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

func (e *ExamAssesmentService) StartNewDescriptiveAssesment(ctx context.Context, generatedExamId int, attempt *ent.ExamAttempt, request *DescriptiveExamAssesmentRequest) (*ent.ExamAssesment, error) {
	assesmentModel := commonRepositories.AssesmentModel{
		CompletedSeconds: request.CompletedSeconds,
		Status:           constants.ASSESSMENT_PENDING,
	}

	generatedExam, err := e.examGenerationService.GetGeneratedExamById(ctx, generatedExamId)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(generatedExam.RawExamData)
	if err != nil {
		return nil, err
	}

	var descriptiveExam models.DescriptiveExam
	err = json.Unmarshal(jsonData, &descriptiveExam)
	if err != nil {
		return nil, err
	}

	// prompt := fmt.Sprintf(`Can you evaluate %s on topic: %s
	// 				Points to check for:
	// 				- Grammar
	// 				- Punctuation
	// 				- Relevance to the topic
	// 				- Maximum word count: 250. Make sure you count the number of WORDS and not punctuation marks or numbers or anyting
	// 				Provide a rating out of 25 marks based on these criteria.

	// 				Output Requirements:
	// 				- A JSON with keys "rating" with string value, "strengths" with array of string, "weakness"  with array of string, "corrected_version"  with string value
	// 				- Make sure the output is one line string with no formatting such that it shiould be JSON parsed

	// 				Content to evaluate:
	// 				%s`, descriptiveExam.Type, descriptiveExam.Topic, request.Content)

	// response, err := e.promptService.GetPromptResult(ctx, prompt, constants.FLASH_15)

	return e.examAssesmentRepository.Create(ctx, attempt.ID, assesmentModel)
}

func (e *ExamAssesmentService) GetAssesmentById(ctx context.Context, assesmentId int, userId string) (*ent.ExamAssesment, error) {
	return e.examAssesmentRepository.GetById(ctx, assesmentId, userId)
}
