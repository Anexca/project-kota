package services

import (
	"common/ent"
	commonRepositories "common/repositories"
	commonServices "common/services"

	"github.com/redis/go-redis/v9"
)

type ExamService struct {
	redisService                     *commonServices.RedisService
	examRepository                   *commonRepositories.ExamRepository
	cachedQuestionMetaDataRepository *commonRepositories.CachedQuestionMetaDataRepository
}

func NewExamService(redisClient *redis.Client, dbClient *ent.Client) *ExamService {
	redisService := commonServices.NewRedisService(redisClient)
	examRepository := commonRepositories.NewExamRespository(dbClient)
	cachedQuestionMetaDataRepository := commonRepositories.NewCachedQuestionMetaDataRepository(dbClient)

	return &ExamService{
		redisService:                     redisService,
		examRepository:                   examRepository,
		cachedQuestionMetaDataRepository: cachedQuestionMetaDataRepository,
	}
}
