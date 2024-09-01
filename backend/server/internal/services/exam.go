package services

import (
	commonConstants "common/constants"
	"common/ent"
	commonRepositories "common/repositories"
	commonServices "common/services"
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type ExamService struct {
	redisService                     *commonServices.RedisService
	examRepository                   *commonRepositories.ExamRepository
	questionRepository               *commonRepositories.QuestionRepository
	examCategoryRepository           *commonRepositories.ExamCategoryRepository
	cachedQuestionMetaDataRepository *commonRepositories.CachedQuestionMetaDataRepository
}

func NewExamService(redisClient *redis.Client, dbClient *ent.Client) *ExamService {
	redisService := commonServices.NewRedisService(redisClient)
	examRepository := commonRepositories.NewExamRespository(dbClient)
	examCategoryRepository := commonRepositories.NewExamCategoryRepository(dbClient)
	cachedQuestionMetaDataRepository := commonRepositories.NewCachedQuestionMetaDataRepository(dbClient)
	questionRepository := commonRepositories.NewQuestionRepository(dbClient)

	return &ExamService{
		redisService:                     redisService,
		examRepository:                   examRepository,
		examCategoryRepository:           examCategoryRepository,
		cachedQuestionMetaDataRepository: cachedQuestionMetaDataRepository,
		questionRepository:               questionRepository,
	}
}

func (e *ExamService) AddCachedQuestionInDatabase(ctx context.Context, examType commonConstants.ExamType) error {
	examName := commonConstants.EXAMS[examType]

	exam, err := e.examRepository.GetByName(ctx, examName)
	if err != nil {
		return err
	}

	cachedMetaData, err := e.cachedQuestionMetaDataRepository.GetByExam(ctx, exam)
	if err != nil {
		return err
	}

	if len(cachedMetaData) == 0 {
		return fmt.Errorf("no cached metadata found for exam: %s", examName)
	}

	cachedData, err := e.redisService.Get(ctx, cachedMetaData[0].CacheUID)
	if err != nil {
		return err
	}

	var questions []any

	err = json.Unmarshal([]byte(cachedData), &questions)
	if err != nil {
		return fmt.Errorf("failed to unmarshal cached data: %w", err)
	}

	e.questionRepository.AddMany(ctx, questions, exam)

	return nil
}

func (e *ExamService) GetQuestionsForExam(ctx context.Context, examType commonConstants.ExamType) ([]*ent.Question, error) {
	examName := commonConstants.EXAMS[examType]
	exam, err := e.examRepository.GetByName(ctx, examName)
	if err != nil {
		return nil, err
	}

	return e.questionRepository.GetByExam(ctx, exam)
}
