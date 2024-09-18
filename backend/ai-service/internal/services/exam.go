package services

import (
	commonConstants "common/constants"
	"common/ent"
	commonRepositories "common/repositories"
	commonService "common/services"
	commonUtil "common/util"
	"context"
	"fmt"
	"log"
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

const DEFAULT_CACHE_EXPIRY = 24 * time.Hour

func (q *ExamService) PopulateExamQuestionCache(ctx context.Context) error {
	examCategories, err := q.examCategoryRepository.Get(ctx)
	if err != nil {
		return err
	}

	for _, examCategory := range examCategories {
		exams, err := q.examRepository.GetByExamCategory(ctx, examCategory)
		if err != nil {
			log.Printf("Error getting exams for category %s: %v", examCategory.Name, err)
			return err
		}

		for _, exam := range exams {
			if !exam.IsActive {
				continue
			}

			examSetting, err := q.examSettingRepository.GetByExam(ctx, exam.ID)
			if err != nil {
				log.Printf("Error getting exam setting for exam %s: %v", exam.Name, err)
				return err
			}

			response, err := q.genAIService.GetContentStream(ctx, examSetting.AiPrompt, commonConstants.PRO_15)
			if err != nil {
				log.Printf("Error generating content for exam %s: %v", exam.Name, err)
				return err
			}

			validationPrompt := fmt.Sprintf(`Ensure that the following string is a valid JSON string.
											Requirements:

											•	The string must be a valid JSON format.
											•	When parsed as JSON, no errors should occur.
											•	The output should be a single-line string without extra spaces, newlines, or formatting.
											"%s"`, response)

			validationResponse, err := q.genAIService.GetContentStream(ctx, validationPrompt, commonConstants.PRO_15)
			if err != nil {
				log.Printf("Error generating validation content for exam %s: %v", exam.Name, err)
				return err
			}

			uid := commonUtil.GenerateUUID()
			err = q.redisService.Store(ctx, uid, validationResponse, DEFAULT_CACHE_EXPIRY)
			if err != nil {
				log.Printf("Error storing cached response for exam %s with uid %s: %v", exam.Name, uid, err)
				return err
			}

			cacheMetaData, err := q.cachedQuestionMetaDataRepository.Create(ctx, uid, DEFAULT_CACHE_EXPIRY, exam)
			if err != nil {
				log.Printf("Error saving cached meta data for exam %s with uid %s: %v", exam.Name, uid, err)
				return err
			}

			log.Printf("Cached response for exam %s with uid %s, saved its meta data in db with id %d", exam.Name, uid, cacheMetaData.ID)

			log.Printf("Waiting for 1 minute before processing the next exam...")
			time.Sleep(1 * time.Minute)
		}
	}

	return nil
}
