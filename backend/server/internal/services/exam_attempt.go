package services

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"common/ent"
	commonInterfaces "common/interfaces"
	commonRepositories "common/repositories"

	"server/internal/interfaces"
	"server/pkg/models"
)

// AccessServiceInterface defines the contract for AccessService
type AccessServiceInterface interface {
	UserHasAccessToExam(ctx context.Context, examId int, userId string) (bool, error)
}

// ExamRepositoryInterface defines the contract for ExamRepository
type ExamRepositoryInterface interface {
	// Define methods needed from the ExamRepository
}

// ExamAttemptRepositoryInterface defines the contract for ExamAttemptRepository
type ExamAttemptRepositoryInterface interface {
	GetByExam(ctx context.Context, generatedExamId int, userId string) ([]*ent.ExamAttempt, error)
	Create(ctx context.Context, currAttempts int, generatedExamId int, userId string) (*ent.ExamAttempt, error)
}

// ExamSettingRepositoryInterface defines the contract for ExamSettingRepository
type ExamSettingRepositoryInterface interface {
	GetByExam(ctx context.Context, examId int) (*ent.ExamSetting, error)
}

// ExamAttemptService is the service for handling exam attempts
type ExamAttemptService struct {
	accessService           interfaces.AccessServiceInterface
	examRepository          commonInterfaces.ExamRepositoryInterface
	examAttemptRepository   commonInterfaces.ExamAttemptRepositoryInterface
	examSettingRepository   commonInterfaces.ExamSettingRepositoryInterface
	generatedExamRepository commonInterfaces.GeneratedExamRepositoryInterface
}

// NewExamAttemptService initializes a new ExamAttemptService
func NewExamAttemptService(
	accessService interfaces.AccessServiceInterface,
	examRepository commonInterfaces.ExamRepositoryInterface,
	examAttemptRepository commonInterfaces.ExamAttemptRepositoryInterface,
	examSettingRepository commonInterfaces.ExamSettingRepositoryInterface,
	generatedExamRepository commonInterfaces.GeneratedExamRepositoryInterface,
) *ExamAttemptService {
	return &ExamAttemptService{
		accessService:           accessService,
		examRepository:          examRepository,
		examAttemptRepository:   examAttemptRepository,
		examSettingRepository:   examSettingRepository,
		generatedExamRepository: generatedExamRepository,
	}
}

func InitExamAttemptService(dbClient *ent.Client) *ExamAttemptService {
	accessService := InitAccessService(dbClient)
	examAttemptRepository := commonRepositories.NewExamAttemptRepository(dbClient)
	examSettingRepository := commonRepositories.NewExamSettingRepository(dbClient)
	examRepository := commonRepositories.NewExamRepository(dbClient)
	generatedExamRepository := commonRepositories.NewGeneratedExamRepository(dbClient)

	return NewExamAttemptService(
		accessService,
		examRepository,
		examAttemptRepository,
		examSettingRepository,
		generatedExamRepository,
	)
}

func (e *ExamAttemptService) CheckAndAddAttempt(ctx context.Context, generatedExamId int, userId string, isOpen bool) (*ent.ExamAttempt, error) {
	userExamAttempts, err := e.examAttemptRepository.GetByExam(ctx, generatedExamId, userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if !errors.As(err, &notFoundError) {
			return nil, err
		}
	}

	generatedExam, err := e.generatedExamRepository.GetOpenById(ctx, generatedExamId, isOpen)
	if err != nil {
		return nil, err
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

	currAttempts := len(userExamAttempts)

	examSettings, err := e.examSettingRepository.GetByExam(ctx, generatedExam.Edges.Exam.ID)
	if err != nil {
		return nil, err
	}

	if examSettings.MaxAttempts <= currAttempts {
		return nil, errors.New("max attempts for exam exceeded")
	}

	currentAttempt, err := e.examAttemptRepository.Create(ctx, currAttempts, generatedExamId, userId)
	if err != nil {
		return nil, err
	}

	return currentAttempt, nil
}

// GetAttempts retrieves exam attempts for the user with pagination
func (e *ExamAttemptService) GetAttempts(ctx context.Context, userId string, page, limit int, from, to *time.Time, examTypeId, categoryId *int) (*models.PaginatedData, error) {
	examWithAttempts, err := e.generatedExamRepository.GetPaginatedExamsByUserAndDate(ctx, userId, page, limit, from, to, examTypeId, categoryId)
	if err != nil {
		return nil, err
	}

	var userExamAttempts []*models.UserExamAttempt

	for _, generatedExam := range examWithAttempts {
		userExamAttempt := &models.UserExamAttempt{
			AttemptedExamId: generatedExam.ID,
			IsActive:        generatedExam.IsActive,
			ExamName:        generatedExam.Edges.Exam.Name,
			ExamType:        generatedExam.Edges.Exam.Type.String(),
			ExamTypeId:      generatedExam.Edges.Exam.ID,
			ExamCategory:    string(generatedExam.Edges.Exam.Edges.Category.Name),
			ExamCategoryId:  generatedExam.Edges.Exam.Edges.Category.ID,
			Topic:           generatedExam.RawExamData["topic"].(string),
			Type:            generatedExam.RawExamData["type"].(string),
			Attempts:        []models.Attempt{},
		}

		for i, attempt := range generatedExam.Edges.Attempts {
			attemptModel := models.Attempt{
				AttemptId:     attempt.ID,
				AttemptNumber: i + 1,
				AttemptDate:   attempt.UpdatedAt,
			}

			if attempt.Edges.Assesment != nil {
				attemptModel.AssessmentId = attempt.Edges.Assesment.ID
				attemptModel.AssessmentStatus = string(attempt.Edges.Assesment.Status)
			}

			userExamAttempt.Attempts = append(userExamAttempt.Attempts, attemptModel)
		}

		userExamAttempts = append(userExamAttempts, userExamAttempt)
	}

	totalCount, err := e.generatedExamRepository.GetCountOfFilteredExamsDataByUserAndDate(ctx, userId, from, to, examTypeId, categoryId)
	if err != nil {
		return nil, err
	}

	paginatedData := &models.PaginatedData{
		CurrentPage: page,
		TotalItems:  totalCount,
		TotalPages:  int(math.Ceil(float64(totalCount) / float64(limit))),
		Data:        userExamAttempts,
		PerPage:     limit,
	}

	return paginatedData, nil
}
