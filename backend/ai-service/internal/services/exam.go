package services

import (
	"ai-service/pkg/models"
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	commonConstants "common/constants"
	"common/ent"
	commonInterfaces "common/interfaces"
	"common/repositories"
	commonService "common/services"
	commonUtil "common/util"

	"cloud.google.com/go/vertexai/genai"
	"github.com/redis/go-redis/v9"
)

type ExamService struct {
	genAIService                     GenAIServiceInterface
	redisService                     commonInterfaces.RedisServiceInterface
	examRepository                   commonInterfaces.ExamRepositoryInterface
	examCategoryRepository           commonInterfaces.ExamCategoryRepositoryInterface
	examSettingRepository            commonInterfaces.ExamSettingRepositoryInterface
	cachedQuestionMetaDataRepository commonInterfaces.CachedExamRepositoryInterface
}

// NewExamService constructor with dependencies for both production and test usage
func NewExamService(
	genAIService GenAIServiceInterface,
	redisService commonInterfaces.RedisServiceInterface,
	examRepository commonInterfaces.ExamRepositoryInterface,
	examCategoryRepository commonInterfaces.ExamCategoryRepositoryInterface,
	examSettingRepository commonInterfaces.ExamSettingRepositoryInterface,
	cachedQuestionMetaDataRepository commonInterfaces.CachedExamRepositoryInterface,
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

func InitExamService(genAiClient *genai.Client, redisClient *redis.Client, dbClient *ent.Client) *ExamService {
	genAIService := NewGenAIService(genAiClient)
	redisService := commonService.NewRedisService(redisClient)
	examRepository := repositories.NewExamRepository(dbClient)
	examCategoryRepository := repositories.NewExamCategoryRepository(dbClient)
	examSettingRepository := repositories.NewExamSettingRepository(dbClient)
	cachedQuestionMetaDataRepository := repositories.NewCachedExamRepository(dbClient)

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
	if len(examCategories) == 0 {
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
			response, err := q.genAIService.GetStructuredContentStream(ctx, examSetting.AiPrompt, commonConstants.PRO_15)
			if err != nil {
				log.Printf("Error generating AI content for exam %s (ID: %d): %v", exam.Name, exam.ID, err)
				return err
			}

			// Generate UUID and store in Redis
			uid := commonUtil.GenerateUUID()
			log.Printf("Storing AI-generated content for exam %s (ID: %d) with UID %s", exam.Name, exam.ID, uid)
			if err = q.redisService.Store(ctx, uid, response, DEFAULT_CACHE_EXPIRY); err != nil {
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

var rwMutex = &sync.RWMutex{}

func (q *ExamService) GenerateAllDescriptiveQuestions(ctx context.Context) ([]*models.GenerateQuestionResponse, error) {
	rwMutex.Lock()
	defer rwMutex.Unlock()

	descriptiveExams, err := q.examRepository.GetActiveByType(ctx, commonConstants.ExamTypeDescriptive)
	if err != nil {
		return nil, err
	}

	var descriptiveExamsIds []*models.GenerateQuestionResponse

	for _, descriptiveExam := range descriptiveExams {
		generatedExam, err := q.GenerateExamQuestionAndPopulateCache(ctx, descriptiveExam.ID)
		if err != nil {
			return nil, err
		}

		descriptiveExamsIds = append(descriptiveExamsIds, generatedExam)

		// Sleep before processing the next exam
		time.Sleep(30 * time.Second)
	}

	return descriptiveExamsIds, nil
}

func (q *ExamService) GenerateExamQuestionAndPopulateCache(ctx context.Context, examId int) (*models.GenerateQuestionResponse, error) {
	exam, err := q.examRepository.GetById(ctx, examId)
	if err != nil {
		return nil, err
	}
	log.Printf("Processing exam: %s (ID: %d)", exam.Name, exam.ID)

	if !exam.IsActive {
		return nil, fmt.Errorf("Skipping inactive exam: %s (ID: %d)", exam.Name, exam.ID)
	}

	examSetting, err := q.examSettingRepository.GetByExam(ctx, exam.ID)
	if err != nil {
		log.Printf("Error fetching settings for exam %s: %v", exam.Name, err)
		return nil, err
	}

	if examSetting.AiPrompt == "" {
		return nil, fmt.Errorf("Skipping exam %s (ID: %d) due to missing AI prompt", exam.Name, exam.ID)
	}

	log.Printf("Fetching AI content stream for exam: %s (ID: %d)", exam.Name, exam.ID)
	response, err := q.genAIService.GetStructuredContentStream(ctx, examSetting.AiPrompt, commonConstants.PRO_15)
	if err != nil {
		log.Printf("Error generating AI content for exam %s (ID: %d): %v", exam.Name, exam.ID, err)
		return nil, err
	}

	uid := commonUtil.GenerateUUID()
	log.Printf("Storing AI-generated content for exam %s (ID: %d) with UID %s", exam.Name, exam.ID, uid)
	if err = q.redisService.Store(ctx, uid, response, DEFAULT_CACHE_EXPIRY); err != nil {
		log.Printf("Error storing AI content for exam %s (ID: %d) with UID %s: %v", exam.Name, exam.ID, uid, err)
		return nil, err
	}

	log.Printf("Saving cached metadata for exam %s (ID: %d)", exam.Name, exam.ID)
	cacheMetaData, err := q.cachedQuestionMetaDataRepository.Create(ctx, uid, DEFAULT_CACHE_EXPIRY, exam)
	if err != nil {
		log.Printf("Error saving cached metadata for exam %s (ID: %d): %v", exam.Name, exam.ID, err)
		return nil, err
	}

	log.Printf("Cached metadata saved for exam %s (ID: %d), metadata ID: %d", exam.Name, exam.ID, cacheMetaData.ID)
	return &models.GenerateQuestionResponse{
		ExamName:         exam.Name,
		CachedMetaDataId: cacheMetaData.ID,
		RedisCacheUid:    uid,
	}, nil
}
