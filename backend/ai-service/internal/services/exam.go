package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"common/ent"

	"cloud.google.com/go/vertexai/genai"
	"github.com/redis/go-redis/v9"

	commonConstants "common/constants"
	commonRepositories "common/repositories"
	commonService "common/services"
	commonUtil "common/util"
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
	log.Println("Starting to populate exam question cache...")

	examCategories, err := q.examCategoryRepository.Get(ctx)
	if err != nil {
		log.Printf("Error fetching exam categories: %v", err)
		return err
	}

	for _, examCategory := range examCategories {
		log.Printf("Processing category: %s", examCategory.Name)

		exams, err := q.examRepository.GetByExamCategory(ctx, examCategory)
		if err != nil {
			log.Printf("Error fetching exams for category %s: %v", examCategory.Name, err)
			return err
		}

		for _, exam := range exams {
			log.Printf("Processing exam: %s (ID: %d)", exam.Name, exam.ID)

			if !exam.IsActive {
				log.Printf("Skipping inactive exam: %s (ID: %d)", exam.Name, exam.ID)
				continue
			}

			examSetting, err := q.examSettingRepository.GetByExam(ctx, exam.ID)
			if err != nil {
				log.Printf("Error fetching settings for exam %s: %v", exam.Name, err)
				return err
			}

			if examSetting.AiPrompt == "" {
				log.Printf("Skipping exam %s (ID: %d) due to missing AI prompt", exam.Name, exam.ID)
				continue
			}

			log.Printf("Fetching AI content stream for exam: %s (ID: %d)", exam.Name, exam.ID)
			response, err := q.genAIService.GetContentStream(ctx, examSetting.AiPrompt, commonConstants.PRO_15)
			if err != nil {
				log.Printf("Error generating AI content for exam %s (ID: %d): %v", exam.Name, exam.ID, err)
				return err
			}
			log.Printf("AI content generated for exam %s (ID: %d)", exam.Name, exam.ID)

			log.Printf("Validating AI response for exam %s (ID: %d)", exam.Name, exam.ID)
			validationPrompt := fmt.Sprintf(`Ensure that the following string is a valid JSON string.
											Requirements:
											•	The string must be a valid JSON format.
											•	When parsed as JSON, no errors should occur.
											•	The output should be a single-line string without extra spaces, newlines, or formatting.
											•	Ensure the string is JSON parsable.
											"%s"`, response)

			validationResponse, err := q.genAIService.GetContentStream(ctx, validationPrompt, commonConstants.PRO_15)
			if err != nil {
				log.Printf("Error validating AI content for exam %s (ID: %d): %v", exam.Name, exam.ID, err)
				return err
			}
			log.Printf("Validation successful for exam %s (ID: %d)", exam.Name, exam.ID)

			uid := commonUtil.GenerateUUID()
			log.Printf("Storing AI-generated content for exam %s (ID: %d) with UID %s", exam.Name, exam.ID, uid)
			if err = q.redisService.Store(ctx, uid, validationResponse, DEFAULT_CACHE_EXPIRY); err != nil {
				log.Printf("Error storing AI content for exam %s (ID: %d) with UID %s: %v", exam.Name, exam.ID, uid, err)
				return err
			}

			log.Printf("Saving cached metadata for exam %s (ID: %d)", exam.Name, exam.ID)
			cacheMetaData, err := q.cachedQuestionMetaDataRepository.Create(ctx, uid, DEFAULT_CACHE_EXPIRY, exam)
			if err != nil {
				log.Printf("Error saving cached metadata for exam %s (ID: %d) with UID %s: %v", exam.Name, exam.ID, uid, err)
				return err
			}
			log.Printf("Cached metadata saved for exam %s (ID: %d), metadata ID: %d", exam.Name, exam.ID, cacheMetaData.ID)

			log.Printf("Waiting for 1 minute before processing the next exam...")
			time.Sleep(1 * time.Minute)
		}
	}

	log.Println("Completed populating exam question cache.")
	return nil
}
