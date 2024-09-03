package services

import (
	"common/ent"
	commonRepositories "common/repositories"
	commonService "common/services"
	commonUtil "common/util"
	"context"
	"log"
	"sync"
	"time"

	"cloud.google.com/go/vertexai/genai"
	"github.com/redis/go-redis/v9"
)

type ExamService struct {
	genAIService                     *GenAIService
	redisService                     *commonService.RedisService
	examRepository                   *commonRepositories.ExamRepository
	examCategoryRepository           *commonRepositories.ExamCategoryRepository
	examSettingRepository            *commonRepositories.ExamSettingRepository
	cachedQuestionMetaDataRepository *commonRepositories.CachedExamRepository
}

func NewExamService(genAIClient *genai.Client, redisClient *redis.Client, dbClient *ent.Client) *ExamService {
	genAIService := NewGenAIService(genAIClient)
	redisService := commonService.NewRedisService(redisClient)
	examRepository := commonRepositories.NewExamRespository(dbClient)
	examCategoryRepository := commonRepositories.NewExamCategoryRepository(dbClient)
	examSettingRepository := commonRepositories.NewExamSettingRepository(dbClient)
	cachedQuestionMetaDataRepository := commonRepositories.NewCachedExamRepository(dbClient)

	return &ExamService{
		genAIService:                     genAIService,
		redisService:                     redisService,
		examRepository:                   examRepository,
		examCategoryRepository:           examCategoryRepository,
		examSettingRepository:            examSettingRepository,
		cachedQuestionMetaDataRepository: cachedQuestionMetaDataRepository,
	}
}

const GEN_AI_MODEL = "gemini-1.5-pro"
const DEFAULT_CACHE_EXPIRY = 24 * time.Hour

func (q *ExamService) PopulateExamQuestionCache(ctx context.Context) error {
	wg := &sync.WaitGroup{}

	examCategories, err := q.examCategoryRepository.Get(ctx)
	if err != nil {
		return err
	}

	for _, examCategory := range examCategories {
		wg.Add(1)
		go func(cat *ent.ExamCategory) {
			defer wg.Done()

			exams, err := q.examRepository.GetByExamCategory(ctx, cat)
			if err != nil {
				log.Printf("Error getting exams for category %s: %v", cat.Name, err)
				return
			}

			for _, exam := range exams {
				wg.Add(1)
				go func(exam *ent.Exam) {
					defer wg.Done()

					examSetting, err := q.examSettingRepository.GetByExam(ctx, exam.ID)
					if err != nil {
						log.Printf("Error getting exam setting for exam %s: %v", exam.Name, err)
						return
					}

					response, err := q.genAIService.GetContentStream(ctx, examSetting.AiPrompt, GEN_AI_MODEL)
					if err != nil {
						log.Printf("Error generating content for exam %s: %v", exam.Name, err)
						return
					}

					uid := commonUtil.GenerateUUID()
					q.redisService.Store(ctx, uid, response, DEFAULT_CACHE_EXPIRY)
					cacheMetaData, err := q.cachedQuestionMetaDataRepository.Create(ctx, uid, DEFAULT_CACHE_EXPIRY, exam)
					if err != nil {
						log.Printf("Error saving cached meta data for exam %s: %v", exam.Name, err)
						return
					}
					log.Printf("Cached response for exam %s with uid %s, saved its meta data in db with id %d", exam.Name, uid, cacheMetaData.ID)
				}(exam)
			}
		}(examCategory)
	}

	wg.Wait()

	return nil
}
