package services

import (
	commonConstants "common/constants"
	"common/ent"
	"common/ent/exam"
	commonRepositories "common/repositories"
	commonServices "common/services"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"server/pkg/models"
	"sort"

	"github.com/redis/go-redis/v9"
)

type ExamGenerationService struct {
	accessService           *AccessService
	redisService            *commonServices.RedisService
	examRepository          *commonRepositories.ExamRepository
	examGroupRepository     *commonRepositories.ExamGroupRepository
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
	examGroupRepository := commonRepositories.NewExamGroupRepository(dbClient)
	examCategoryRepository := commonRepositories.NewExamCategoryRepository(dbClient)
	cachedExamRepository := commonRepositories.NewCachedExamRepository(dbClient)
	generatedExamRepository := commonRepositories.NewGeneratedExamRepository(dbClient)
	examSettingRepository := commonRepositories.NewExamSettingRepository(dbClient)
	examAttemptRepository := commonRepositories.NewExamAttemptRepository(dbClient)

	return &ExamGenerationService{
		accessService:           accessService,
		redisService:            redisService,
		examRepository:          examRepository,
		examGroupRepository:     examGroupRepository,
		examCategoryRepository:  examCategoryRepository,
		cachedExamRepository:    cachedExamRepository,
		generatedExamRepository: generatedExamRepository,
		examSettingRepository:   examSettingRepository,
		examAttemptRepository:   examAttemptRepository,
	}
}

func (e *ExamGenerationService) GenerateExams(ctx context.Context, examId int, modelType models.ExamModelType) error {

	exam, err := e.examRepository.GetById(ctx, examId)
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

func (e *ExamGenerationService) MarkQuestionsAsOpen(ctx context.Context, examName string) error {
	exam, err := e.examRepository.GetByName(ctx, examName)
	if err != nil {
		return fmt.Errorf("failed to get exam by name: %w", err)
	}

	currentOpenQuestions, err := e.generatedExamRepository.GetByOpenFlag(ctx, exam.ID)
	if err != nil {
		return fmt.Errorf("failed to get currentOpenQuestions: %w", err)
	}

	for _, coe := range currentOpenQuestions {
		coe.IsOpen = false
	}

	err = e.generatedExamRepository.UpdateMany(ctx, currentOpenQuestions)
	if err != nil {
		return fmt.Errorf("failed to mark current open questions closed: %w", err)
	}

	generatedOldExams, err := e.generatedExamRepository.GetByWeekOffset(ctx, exam, 1, 2) // Get last weeks 2 questions
	if err != nil {
		return fmt.Errorf("failed to get exam by name: %w", err)
	}

	for _, goe := range generatedOldExams {
		goe.IsOpen = true
	}

	err = e.generatedExamRepository.UpdateMany(ctx, generatedOldExams)
	if err != nil {
		return fmt.Errorf("failed to create new open exams: %w", err)
	}

	log.Printf("Marked %d open questions for %s exam", len(generatedOldExams), exam.Name)

	return nil
}

func (e *ExamGenerationService) MarkExpiredExamsInactive(ctx context.Context, examId int) error {
	exam, err := e.examRepository.GetById(ctx, examId)
	if err != nil {
		return err
	}

	generatedExams, err := e.generatedExamRepository.GetByExam(ctx, exam)
	if err != nil {
		return err
	}

	sort.SliceStable(generatedExams, func(i, j int) bool {
		return generatedExams[i].UpdatedAt.After(generatedExams[j].UpdatedAt)
	})

	if len(generatedExams) > 30 {
		for _, generatedExam := range generatedExams[30:] { // Skip the first 30 exams
			generatedExam.IsActive = false
		}

		if err := e.generatedExamRepository.UpdateMany(ctx, generatedExams[30:]); err != nil {
			return err
		}
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

	case models.MCQExamType:
		var mcqExams []models.MCQExam

		err := json.Unmarshal([]byte(cachedData), &mcqExams)
		if err != nil {
			return fmt.Errorf("failed to validate cached data for MCQ Exam: %w", err)
		}

		generatedMCQExam := models.GeneratedMCQExam{
			ExamContent: mcqExams,
		}

		jsonData, err := json.Marshal(generatedMCQExam)
		if err != nil {
			log.Fatalf("Failed to marshal struct to JSON: %v", err)
		}

		var result map[string]interface{}
		err = json.Unmarshal(jsonData, &result)
		if err != nil {
			log.Fatalf("Failed to unmarshal JSON into map: %v", err)
		}

		_, err = e.generatedExamRepository.Add(ctx, result, exam.ID)
		if err != nil {
			return fmt.Errorf("failed to generate MCQ Exam : %w", err)
		}

	default:
		return fmt.Errorf("unsupported exam model type")
	}

	return nil
}

func (e *ExamGenerationService) GetExamsByExamGroupIdAndExamType(ctx context.Context, examGroupId int, userId string) ([]*models.GeneratedExamOverview, error) {
	examGroup, err := e.examGroupRepository.GetActiveByIdWithExams(ctx, examGroupId, true)
	if err != nil {
		return nil, err
	}

	accessibleExams, err := e.accessService.GetAccessibleExamsForUser(ctx, examGroup.Edges.Exams, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to check access: %w", err)
	}

	accessibleExamMap := make(map[int]struct{})
	for _, exam := range accessibleExams {
		accessibleExamMap[exam.ID] = struct{}{}
	}

	var generatedExamsOverview []*models.GeneratedExamOverview

	for _, exam := range examGroup.Edges.Exams {
		sortedExams := e.sortExamsByUpdatedAt(exam.Edges.Generatedexams)

		limit := min(26, len(sortedExams))
		latestExams := sortedExams[:limit]

		list, err := e.buildGeneratedExamOverviewList(ctx, latestExams, exam, userId)
		if err != nil {
			return nil, err
		}

		for _, overview := range list {
			if _, found := accessibleExamMap[exam.ID]; found {
				overview.UserHasAccessToExam = true
			} else {
				overview.UserHasAccessToExam = false
			}
		}

		generatedExamsOverview = append(generatedExamsOverview, list...)
	}

	return generatedExamsOverview, nil
}

func (e *ExamGenerationService) GetOpenGeneratedExams(ctx context.Context, examType commonConstants.ExamType, userId string) ([]*models.GeneratedExamOverview, error) {
	exam, err := e.examRepository.GetActiveById(ctx, 1, true)

	if err != nil {
		return nil, fmt.Errorf("failed to get exam by name: %w", err)
	}

	generatedExams, err := e.generatedExamRepository.GetByOpenFlag(ctx, exam.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get exam by name: %w", err)
	}

	return e.buildGeneratedExamOverviewList(ctx, generatedExams, exam, userId)
}

func (e *ExamGenerationService) GetGeneratedExamById(ctx context.Context, generatedExamId int, userId string, isOpen bool) (*models.GeneratedExamOverview, error) {
	generatedExam, err := e.generatedExamRepository.GetOpenById(ctx, generatedExamId, isOpen)
	if err != nil {
		return nil, fmt.Errorf("failed to get generated exam: %w", err)
	}

	if !isOpen {
		hasAccess, err := e.accessService.UserHasAccessToExam(ctx, generatedExam.Edges.Exam.ID, userId)
		if err != nil {
			return nil, fmt.Errorf("failed to check access: %w", err)
		}

		if !hasAccess {
			return nil, errors.New("forbidden")
		}
	}

	userAttempts, err := e.examAttemptRepository.GetByExam(ctx, generatedExam.ID, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user attempts: %w", err)
	}

	examSettings, err := e.examSettingRepository.GetByExam(ctx, generatedExam.Edges.Exam.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get exam settings: %w", err)
	}

	examOverview := e.buildGeneratedExamOverview(generatedExam, examSettings, userAttempts)
	examOverview.ExamName = generatedExam.Edges.Exam.Name
	examOverview.ExamType = generatedExam.Edges.Exam.Type.String()

	return examOverview, nil
}

func (e *ExamGenerationService) GetActiveExams(ctx context.Context, examType commonConstants.ExamType) ([]*ent.Exam, error) {
	return e.examRepository.GetActiveByType(ctx, examType)
}

func (e *ExamGenerationService) sortExamsByUpdatedAt(exams []*ent.GeneratedExam) []*ent.GeneratedExam {
	sort.Slice(exams, func(i, j int) bool {
		return exams[i].UpdatedAt.After(exams[j].UpdatedAt)
	})
	return exams
}

func (e *ExamGenerationService) buildGeneratedExamOverviewList(ctx context.Context, latestExams []*ent.GeneratedExam, ex *ent.Exam, userId string) ([]*models.GeneratedExamOverview, error) {
	generatedExamOverviewList := make([]*models.GeneratedExamOverview, 0, len(latestExams))

	for _, generatedExam := range latestExams {
		userAttempts, err := e.examAttemptRepository.GetByExam(ctx, generatedExam.ID, userId)
		if err != nil {
			return nil, fmt.Errorf("failed to get user attempts: %w", err)
		}

		overview := e.buildGeneratedExamOverview(generatedExam, ex.Edges.Setting, userAttempts)
		overview.ExamName = ex.Name
		overview.ExamType = string(ex.Type)
		overview.UserAttempts = len(userAttempts)

		if ex.Type == exam.TypeMCQ {
			overview.RawExamData = nil
		}

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
		NegativeMarking:   examSettings.NegativeMarking,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
