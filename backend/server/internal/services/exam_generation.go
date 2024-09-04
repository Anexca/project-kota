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

type ExamGenerationService struct {
	redisService            *commonServices.RedisService
	examRepository          *commonRepositories.ExamRepository
	generatedExamRepository *commonRepositories.GeneratedExamRepository
	examCategoryRepository  *commonRepositories.ExamCategoryRepository
	examSettingRepository   *commonRepositories.ExamSettingRepository
	examAttemptRepository   *commonRepositories.ExamAttemptRepository
	cachedExamRepository    *commonRepositories.CachedExamRepository
}

func NewExamGenerationService(redisClient *redis.Client, dbClient *ent.Client) *ExamGenerationService {
	redisService := commonServices.NewRedisService(redisClient)
	examRepository := commonRepositories.NewExamRespository(dbClient)
	examCategoryRepository := commonRepositories.NewExamCategoryRepository(dbClient)
	cachedExamRepository := commonRepositories.NewCachedExamRepository(dbClient)
	generatedExamRepository := commonRepositories.NewGeneratedExamRepository(dbClient)
	examSettingRepository := commonRepositories.NewExamSettingRepository(dbClient)
	examAttemptRepository := commonRepositories.NewExamAttemptRepository(dbClient)

	return &ExamGenerationService{
		redisService:            redisService,
		examRepository:          examRepository,
		examCategoryRepository:  examCategoryRepository,
		cachedExamRepository:    cachedExamRepository,
		generatedExamRepository: generatedExamRepository,
		examSettingRepository:   examSettingRepository,
		examAttemptRepository:   examAttemptRepository,
	}
}

func (e *ExamGenerationService) GenerateExams(ctx context.Context, examType commonConstants.ExamType, modelType models.ExamModelType) error {
	examName := commonConstants.EXAMS[examType]

	exam, err := e.examRepository.GetByName(ctx, examName)
	if err != nil {
		return err
	}

	cachedData, err := e.FetchCachedExamData(ctx, exam)
	if err != nil {
		return err
	}

	err = e.ProcessExamData(ctx, exam, modelType, cachedData)
	if err != nil {
		return err
	}

	return nil
}

func (e *ExamGenerationService) FetchCachedExamData(ctx context.Context, exam *ent.Exam) (string, error) {

	cachedMetaData, err := e.cachedExamRepository.GetByExam(ctx, exam)
	if err != nil {
		return "", err
	}

	if len(cachedMetaData) == 0 {
		return "", fmt.Errorf("no cached metadata found for exam: %s", exam.Name)
	}

	cachedData, err := e.redisService.Get(ctx, cachedMetaData[0].CacheUID)
	if err != nil {
		return "", err
	}

	e.cachedExamRepository.MarkAsUsed(ctx, cachedMetaData[0].ID)

	return cachedData, nil
}

func (e *ExamGenerationService) ProcessExamData(ctx context.Context, exam *ent.Exam, modelType models.ExamModelType, cachedData string) error {
	switch modelType {
	case models.DescriptiveExamType:
		var descriptiveExams []models.DescriptiveExam

		err := json.Unmarshal([]byte(cachedData), &descriptiveExams)
		if err != nil {
			return fmt.Errorf("failed to validate cached data for DescriptiveExam: %w", err)
		}

		var exams []any

		err = json.Unmarshal([]byte(cachedData), &exams)
		if err != nil {
			return fmt.Errorf("failed to generate DescriptiveExam: %w", err)
		}

		_, err = e.generatedExamRepository.AddMany(ctx, exams, exam)
		if err != nil {
			return fmt.Errorf("failed to generate DescriptiveExams : %w", err)
		}

	default:
		return fmt.Errorf("unsupported exam model type")
	}

	return nil
}

func (e *ExamGenerationService) GetGeneratedExams(ctx context.Context, examType commonConstants.ExamType) ([]models.GeneratedExamOverview, error) {
	examName := commonConstants.EXAMS[examType]

	exam, err := e.examRepository.GetByName(ctx, examName)
	if err != nil {
		return nil, fmt.Errorf("failed to get exam by name: %w", err)
	}

	generatedExamOverviewList := make([]models.GeneratedExamOverview, 0, len(exam.Edges.Generatedexams))

	for _, generatedExam := range exam.Edges.Generatedexams {

		generatedExamOverview := models.GeneratedExamOverview{
			Id:                generatedExam.ID,
			RawExamData:       generatedExam.RawExamData,
			CreatedAt:         generatedExam.CreatedAt,
			UpdatedAt:         generatedExam.UpdatedAt,
			UserAttempts:      len(generatedExam.Edges.Attempts),
			MaxAttempts:       exam.Edges.Setting.MaxAttempts,
			DurationMinutes:   exam.Edges.Setting.DurationSeconds,
			NumberOfQuestions: exam.Edges.Setting.NumberOfQuestions,
		}

		generatedExamOverviewList = append(generatedExamOverviewList, generatedExamOverview)
	}

	return generatedExamOverviewList, nil
}

func (e *ExamGenerationService) GetGeneratedExamById(ctx context.Context, generatedExamId int) (models.GeneratedExamOverview, error) {
	generatedExam, err := e.generatedExamRepository.GetById(ctx, generatedExamId)
	if err != nil {
		return models.GeneratedExamOverview{}, fmt.Errorf("failed to get generated exam: %w", err)
	}

	examSettings, err := e.examSettingRepository.GetByExam(ctx, generatedExam.Edges.Exam.ID)
	if err != nil {
		return models.GeneratedExamOverview{}, fmt.Errorf("failed to get exam settings: %w", err)
	}

	generatedExamOverview := models.GeneratedExamOverview{
		Id:                generatedExam.ID,
		RawExamData:       generatedExam.RawExamData,
		CreatedAt:         generatedExam.CreatedAt,
		UpdatedAt:         generatedExam.UpdatedAt,
		UserAttempts:      len(generatedExam.Edges.Attempts),
		MaxAttempts:       examSettings.MaxAttempts,
		DurationMinutes:   examSettings.DurationSeconds,
		NumberOfQuestions: examSettings.NumberOfQuestions,
	}

	return generatedExamOverview, nil
}
