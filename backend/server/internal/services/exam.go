package services

import (
	commonConstants "common/constants"
	"common/ent"
	commonRepositories "common/repositories"
	commonServices "common/services"
	"context"
	"encoding/json"
	"fmt"
	"server/pkg/models"

	"github.com/redis/go-redis/v9"
)

type ExamService struct {
	redisService                     *commonServices.RedisService
	examRepository                   *commonRepositories.ExamRepository
	generatedExamRepository          *commonRepositories.GeneratedExamRepository
	examCategoryRepository           *commonRepositories.ExamCategoryRepository
	examSettingRepository            *commonRepositories.ExamSettingRepository
	examAttemptRepository            *commonRepositories.ExamAttemptRepository
	cachedQuestionMetaDataRepository *commonRepositories.CachedQuestionMetaDataRepository
}

func NewExamService(redisClient *redis.Client, dbClient *ent.Client) *ExamService {
	redisService := commonServices.NewRedisService(redisClient)
	examRepository := commonRepositories.NewExamRespository(dbClient)
	examCategoryRepository := commonRepositories.NewExamCategoryRepository(dbClient)
	cachedQuestionMetaDataRepository := commonRepositories.NewCachedQuestionMetaDataRepository(dbClient)
	generatedExamRepository := commonRepositories.NewGeneratedExamRepository(dbClient)
	examSettingRepository := commonRepositories.NewExamSettingRepository(dbClient)
	examAttemptRepository := commonRepositories.NewExamAttemptRepository(dbClient)

	return &ExamService{
		redisService:                     redisService,
		examRepository:                   examRepository,
		examCategoryRepository:           examCategoryRepository,
		cachedQuestionMetaDataRepository: cachedQuestionMetaDataRepository,
		generatedExamRepository:          generatedExamRepository,
		examSettingRepository:            examSettingRepository,
		examAttemptRepository:            examAttemptRepository,
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

	e.cachedQuestionMetaDataRepository.MarkAsUsed(ctx, cachedMetaData[0].ID)
	var questions []any

	err = json.Unmarshal([]byte(cachedData), &questions)
	if err != nil {
		return fmt.Errorf("failed to unmarshal cached data: %w", err)
	}

	e.generatedExamRepository.AddMany(ctx, questions, exam)

	return nil
}

func (e *ExamService) GetGeneratedExams(ctx context.Context, examType commonConstants.ExamType) ([]models.GeneratedExamOverview, error) {
	examName := commonConstants.EXAMS[examType]

	exam, err := e.examRepository.GetByName(ctx, examName)
	if err != nil {
		return nil, fmt.Errorf("failed to get exam by name: %w", err)
	}

	examSettings, err := e.examSettingRepository.GetByExam(ctx, exam)
	if err != nil {
		return nil, fmt.Errorf("failed to get exam settings: %w", err)
	}

	generatedExams, err := e.generatedExamRepository.GetByExam(ctx, exam)
	if err != nil {
		return nil, fmt.Errorf("failed to get generated exams: %w", err)
	}

	generatedExamOverviewList := make([]models.GeneratedExamOverview, 0, len(generatedExams))

	for _, generatedExam := range generatedExams {
		userExamAttempts, err := e.examAttemptRepository.GetByExam(ctx, generatedExam.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get exam attempts for exam ID %d: %w", generatedExam.ID, err)
		}

		userAttempts := len(userExamAttempts)

		generatedExamOverview := models.GeneratedExamOverview{
			Id:                generatedExam.ID,
			RawExamData:       generatedExam.RawExamData,
			CreatedAt:         generatedExam.CreatedAt,
			UpdatedAt:         generatedExam.UpdatedAt,
			UserAttempts:      userAttempts,
			MaxAttempts:       examSettings.MaxAttempts,
			DurationMinutes:   examSettings.DurationMinutes,
			NumberOfQuestions: examSettings.NumberOfQuestions,
		}

		generatedExamOverviewList = append(generatedExamOverviewList, generatedExamOverview)
	}

	return generatedExamOverviewList, nil
}
