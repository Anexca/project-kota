package services

import (
	"context"
	"fmt"
	"log"
	"time"

	commonConstants "common/constants"
	"common/ent"
	commonUtil "common/util"
)

// Define interfaces for dependencies to make them mockable

type RedisServiceInterface interface {
	Store(ctx context.Context, key string, value string, expiration time.Duration) error
}

type ExamRepositoryInterface interface {
	GetByExamCategory(ctx context.Context, category *ent.ExamCategory) ([]*ent.Exam, error)
}

type ExamCategoryRepositoryInterface interface {
	Get(ctx context.Context) ([]*ent.ExamCategory, error)
}

type ExamSettingRepositoryInterface interface {
	GetByExam(ctx context.Context, examID int) (*ent.ExamSetting, error)
}

type CachedExamRepositoryInterface interface {
	Create(ctx context.Context, uid string, expiration time.Duration, exam *ent.Exam) (*ent.CachedExam, error)
}

// ExamService uses interfaces to allow testability
type ExamService struct {
	genAIService                     GenAIServiceInterface
	redisService                     RedisServiceInterface
	examRepository                   ExamRepositoryInterface
	examCategoryRepository           ExamCategoryRepositoryInterface
	examSettingRepository            ExamSettingRepositoryInterface
	cachedQuestionMetaDataRepository CachedExamRepositoryInterface
}

// NewExamService constructor with dependencies for both production and test usage
func NewExamService(
	genAIService GenAIServiceInterface,
	redisService RedisServiceInterface,
	examRepository ExamRepositoryInterface,
	examCategoryRepository ExamCategoryRepositoryInterface,
	examSettingRepository ExamSettingRepositoryInterface,
	cachedQuestionMetaDataRepository CachedExamRepositoryInterface,
) *ExamService {
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

// PopulateExamQuestionCache processes exams, fetches AI-generated content, and caches the results
func (q *ExamService) PopulateExamQuestionCache(ctx context.Context) error {
	log.Println("Starting to populate exam question cache...")

	// Fetch exam categories
	examCategories, err := q.examCategoryRepository.Get(ctx)
	if err != nil {
		log.Printf("Error fetching exam categories: %v", err)
		return err
	}

	// Check if examCategories is nil or empty
	if examCategories == nil || len(examCategories) == 0 {
		log.Println("No exam categories found.")
		return nil // or return an error, depending on your requirements
	}

	for _, examCategory := range examCategories {
		log.Printf("Processing category: %s", examCategory.Name)

		// Fetch exams for the current category
		exams, err := q.examRepository.GetByExamCategory(ctx, examCategory)
		if err != nil {
			log.Printf("Error fetching exams for category %s: %v", examCategory.Name, err)
			return err
		}

		for _, exam := range exams {
			log.Printf("Processing exam: %s (ID: %d)", exam.Name, exam.ID)

			// Skip inactive exams
			if !exam.IsActive {
				log.Printf("Skipping inactive exam: %s (ID: %d)", exam.Name, exam.ID)
				continue
			}

			// Fetch exam settings
			examSetting, err := q.examSettingRepository.GetByExam(ctx, exam.ID)
			if err != nil {
				log.Printf("Error fetching settings for exam %s: %v", exam.Name, err)
				return err
			}

			// Skip exams without AI prompt
			if examSetting.AiPrompt == "" {
				log.Printf("Skipping exam %s (ID: %d) due to missing AI prompt", exam.Name, exam.ID)
				continue
			}

			// Fetch AI-generated content
			log.Printf("Fetching AI content stream for exam: %s (ID: %d)", exam.Name, exam.ID)
			response, err := q.genAIService.GetContentStream(ctx, examSetting.AiPrompt, commonConstants.PRO_15)
			if err != nil {
				log.Printf("Error generating AI content for exam %s (ID: %d): %v", exam.Name, exam.ID, err)
				return err
			}

			// Validate AI content
			log.Printf("Validating AI response for exam %s (ID: %d)", exam.Name, exam.ID)
			validationPrompt := fmt.Sprintf(`You are given a JSON string that may contain minor issues... "%s"`, response)
			validationResponse, err := q.genAIService.GetContentStream(ctx, validationPrompt, commonConstants.PRO_15)
			if err != nil {
				log.Printf("Error validating AI content for exam %s (ID: %d): %v", exam.Name, exam.ID, err)
				return err
			}

			// Generate UUID and store in Redis
			uid := commonUtil.GenerateUUID()
			log.Printf("Storing AI-generated content for exam %s (ID: %d) with UID %s", exam.Name, exam.ID, uid)
			if err = q.redisService.Store(ctx, uid, validationResponse, DEFAULT_CACHE_EXPIRY); err != nil {
				log.Printf("Error storing AI content for exam %s (ID: %d) with UID %s: %v", exam.Name, exam.ID, uid, err)
				return err
			}

			// Save metadata
			log.Printf("Saving cached metadata for exam %s (ID: %d)", exam.Name, exam.ID)
			cacheMetaData, err := q.cachedQuestionMetaDataRepository.Create(ctx, uid, DEFAULT_CACHE_EXPIRY, exam)
			if err != nil {
				log.Printf("Error saving cached metadata for exam %s (ID: %d): %v", exam.Name, exam.ID, err)
				return err
			}

			// Log success
			log.Printf("Cached metadata saved for exam %s (ID: %d), metadata ID: %d", exam.Name, exam.ID, cacheMetaData.ID)

			// Sleep before processing the next exam
			time.Sleep(5 * time.Second)
		}
	}

	log.Println("Completed populating exam question cache.")
	return nil
}
