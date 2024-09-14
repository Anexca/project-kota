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
	"sort"

	"github.com/redis/go-redis/v9"
)

type ExamGenerationService struct {
	accessService           *AccessService
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
	accessService := NewAccessService(dbClient)
	examRepository := commonRepositories.NewExamRespository(dbClient)
	examCategoryRepository := commonRepositories.NewExamCategoryRepository(dbClient)
	cachedExamRepository := commonRepositories.NewCachedExamRepository(dbClient)
	generatedExamRepository := commonRepositories.NewGeneratedExamRepository(dbClient)
	examSettingRepository := commonRepositories.NewExamSettingRepository(dbClient)
	examAttemptRepository := commonRepositories.NewExamAttemptRepository(dbClient)

	return &ExamGenerationService{
		accessService:           accessService,
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
		return "", fmt.Errorf("failed to retrieve cached metadata for exam '%s': %w", exam.Name, err)
	}

	if len(cachedMetaData) == 0 {
		return "", fmt.Errorf("no cached metadata found for exam: %s", exam.Name)
	}

	sort.Slice(cachedMetaData, func(i, j int) bool {
		return cachedMetaData[i].UpdatedAt.After(cachedMetaData[j].UpdatedAt)
	})

	latestCachedMeta := cachedMetaData[0]

	cachedData, err := e.redisService.Get(ctx, latestCachedMeta.CacheUID)
	if err != nil {
		return "", fmt.Errorf("failed to fetch cached data from Redis: %w", err)
	}

	if err := e.cachedExamRepository.MarkAsUsed(ctx, latestCachedMeta.ID); err != nil {
		return "", fmt.Errorf("failed to mark cached metadata as used: %w", err)
	}

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

func (e *ExamGenerationService) GetGeneratedExams(ctx context.Context, examType commonConstants.ExamType, userId string) ([]*models.GeneratedExamOverview, error) {
	examName := commonConstants.EXAMS[examType]

	exam, err := e.examRepository.GetByName(ctx, examName)
	if err != nil {
		return nil, fmt.Errorf("failed to get exam by name: %w", err)
	}

	hasAccess, err := e.accessService.UserHasAccessToExam(ctx, exam.ID, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to check access: %w", err)
	}

	if !hasAccess {
		return nil, nil
	}

	sortedExams := e.sortExamsByUpdatedAt(exam.Edges.Generatedexams)

	limit := min(26, len(sortedExams))
	latestExams := sortedExams[:limit]

	return e.buildGeneratedExamOverviewList(ctx, latestExams, exam, userId)
}

func (e *ExamGenerationService) GetGeneratedExamById(ctx context.Context, generatedExamId int, userId string) (*models.GeneratedExamOverview, error) {
	generatedExam, err := e.generatedExamRepository.GetById(ctx, generatedExamId)
	if err != nil {
		return nil, fmt.Errorf("failed to get generated exam: %w", err)
	}

	hasAccess, err := e.accessService.UserHasAccessToExam(ctx, generatedExam.Edges.Exam.ID, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to check access: %w", err)
	}

	if !hasAccess {
		return nil, nil
	}

	userAttempts, err := e.examAttemptRepository.GetByExam(ctx, generatedExam.ID, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user attempts: %w", err)
	}

	examSettings, err := e.examSettingRepository.GetByExam(ctx, generatedExam.Edges.Exam.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get exam settings: %w", err)
	}

	return e.buildGeneratedExamOverview(generatedExam, examSettings, userAttempts), nil
}

func (e *ExamGenerationService) sortExamsByUpdatedAt(exams []*ent.GeneratedExam) []*ent.GeneratedExam {
	sort.Slice(exams, func(i, j int) bool {
		return exams[i].UpdatedAt.After(exams[j].UpdatedAt)
	})
	return exams
}

func (e *ExamGenerationService) buildGeneratedExamOverviewList(ctx context.Context, latestExams []*ent.GeneratedExam, exam *ent.Exam, userId string) ([]*models.GeneratedExamOverview, error) {
	generatedExamOverviewList := make([]*models.GeneratedExamOverview, 0, len(latestExams))

	for _, generatedExam := range latestExams {
		userAttempts, err := e.examAttemptRepository.GetByExam(ctx, generatedExam.ID, userId)
		if err != nil {
			return nil, fmt.Errorf("failed to get user attempts: %w", err)
		}

		overview := e.buildGeneratedExamOverview(generatedExam, exam.Edges.Setting, userAttempts)
		overview.UserAttempts = len(userAttempts)

		generatedExamOverviewList = append(generatedExamOverviewList, overview)
	}

	return generatedExamOverviewList, nil
}

func (e *ExamGenerationService) buildGeneratedExamOverview(generatedExam *ent.GeneratedExam, examSettings *ent.ExamSetting, examAttempts []*ent.ExamAttempt) *models.GeneratedExamOverview {
	return &models.GeneratedExamOverview{
		Id:                generatedExam.ID,
		RawExamData:       generatedExam.RawExamData,
		CreatedAt:         generatedExam.CreatedAt,
		UpdatedAt:         generatedExam.UpdatedAt,
		UserAttempts:      len(examAttempts),
		MaxAttempts:       examSettings.MaxAttempts,
		DurationSeconds:   examSettings.DurationSeconds,
		NumberOfQuestions: examSettings.NumberOfQuestions,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
