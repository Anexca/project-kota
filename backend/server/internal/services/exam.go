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
	redisService            *commonServices.RedisService
	examRepository          *commonRepositories.ExamRepository
	generatedExamRepository *commonRepositories.GeneratedExamRepository
	examCategoryRepository  *commonRepositories.ExamCategoryRepository
	examSettingRepository   *commonRepositories.ExamSettingRepository
	examAttemptRepository   *commonRepositories.ExamAttemptRepository
	cachedExamRepository    *commonRepositories.CachedExamRepository
}

func NewExamService(redisClient *redis.Client, dbClient *ent.Client) *ExamService {
	redisService := commonServices.NewRedisService(redisClient)
	examRepository := commonRepositories.NewExamRespository(dbClient)
	examCategoryRepository := commonRepositories.NewExamCategoryRepository(dbClient)
	cachedExamRepository := commonRepositories.NewCachedExamRepository(dbClient)
	generatedExamRepository := commonRepositories.NewGeneratedExamRepository(dbClient)
	examSettingRepository := commonRepositories.NewExamSettingRepository(dbClient)
	examAttemptRepository := commonRepositories.NewExamAttemptRepository(dbClient)

	return &ExamService{
		redisService:            redisService,
		examRepository:          examRepository,
		examCategoryRepository:  examCategoryRepository,
		cachedExamRepository:    cachedExamRepository,
		generatedExamRepository: generatedExamRepository,
		examSettingRepository:   examSettingRepository,
		examAttemptRepository:   examAttemptRepository,
	}
}

func (e *ExamService) GenerateExams(ctx context.Context, examType commonConstants.ExamType, modelType models.ExamModelType) error {
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

func (e *ExamService) FetchCachedExamData(ctx context.Context, exam *ent.Exam) ([]byte, error) {

	cachedMetaData, err := e.cachedExamRepository.GetByExam(ctx, exam)
	if err != nil {
		return nil, err
	}

	if len(cachedMetaData) == 0 {
		return nil, fmt.Errorf("no cached metadata found for exam: %s", exam.Name)
	}

	cachedData, err := e.redisService.Get(ctx, cachedMetaData[0].CacheUID)
	if err != nil {
		return nil, err
	}

	e.cachedExamRepository.MarkAsUsed(ctx, cachedMetaData[0].ID)

	return []byte(cachedData), nil
}

func (e *ExamService) ProcessExamData(ctx context.Context, exam *ent.Exam, modelType models.ExamModelType, cachedData []byte) error {
	switch modelType {
	case models.DescriptiveExamType:
		var exams []models.DescriptiveExam
		err := json.Unmarshal(cachedData, &exams)
		if err != nil {
			return fmt.Errorf("failed to unmarshal cached data for DescriptiveExam: %w", err)
		}
		anyExams := make([]any, len(exams))
		for i, exam := range exams {
			anyExams[i] = exam
		}
		_, err = e.generatedExamRepository.AddMany(ctx, anyExams, exam)
		if err != nil {
			return fmt.Errorf("failed to generate DescriptiveExams : %w", err)
		}

	default:
		return fmt.Errorf("unsupported exam model type")
	}

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
			RawExamData:       generatedExam.RawData,
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
