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
	examCategoryRepository           *commonRepositories.ExamCategoryRepository
	cachedQuestionMetaDataRepository *commonRepositories.CachedQuestionMetaDataRepository
}

func NewExamService(redisClient *redis.Client, dbClient *ent.Client) *ExamService {
	redisService := commonServices.NewRedisService(redisClient)
	examRepository := commonRepositories.NewExamRespository(dbClient)
	examCategoryRepository := commonRepositories.NewExamCategoryRepository(dbClient)
	cachedQuestionMetaDataRepository := commonRepositories.NewCachedQuestionMetaDataRepository(dbClient)

	return &ExamService{
		redisService:                     redisService,
		examRepository:                   examRepository,
		examCategoryRepository:           examCategoryRepository,
		cachedQuestionMetaDataRepository: cachedQuestionMetaDataRepository,
	}
}

func (e *ExamService) GetCachedQuestions(ctx context.Context, examType commonConstants.ExamType, returnType interface{}) (interface{}, error) {
	examName := commonConstants.EXAMS[examType]

	exam, err := e.examRepository.GetByName(ctx, examName)
	if err != nil {
		return nil, err
	}

	cachedMetaData, err := e.cachedQuestionMetaDataRepository.GetByExam(ctx, exam)
	if err != nil {
		return nil, err
	}

	if len(cachedMetaData) == 0 {
		return nil, fmt.Errorf("no cached metadata found for exam: %s", examName)
	}

	cachedData, err := e.redisService.Get(ctx, cachedMetaData[0].CacheUID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(cachedData), returnType)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cached data: %w", err)
	}

	return returnType, nil
}
