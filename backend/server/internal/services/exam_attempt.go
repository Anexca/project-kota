package services

import (
	"common/ent"
	commonRepositories "common/repositories"
	"context"
	"errors"
	"fmt"
	"server/pkg/models"
)

type ExamAttemptService struct {
	accessService           *AccessService
	examRepository          *commonRepositories.ExamRepository
	examAtemptRepository    *commonRepositories.ExamAttemptRepository
	examSettingRepository   *commonRepositories.ExamSettingRepository
	generatedExamRepository *commonRepositories.GeneratedExamRepository
}

func NewExamAttemptService(dbClient *ent.Client) *ExamAttemptService {
	accessService := NewAccessService(dbClient)
	examAtemptRepository := commonRepositories.NewExamAttemptRepository(dbClient)
	examSettingRepository := commonRepositories.NewExamSettingRepository(dbClient)
	examRepository := commonRepositories.NewExamRespository(dbClient)
	generatedExamRepository := commonRepositories.NewGeneratedExamRepository(dbClient)

	return &ExamAttemptService{
		accessService:           accessService,
		examAtemptRepository:    examAtemptRepository,
		examSettingRepository:   examSettingRepository,
		examRepository:          examRepository,
		generatedExamRepository: generatedExamRepository,
	}
}

func (e *ExamAttemptService) CheckAndAddAttempt(ctx context.Context, generatedExamId int, userId string) (*ent.ExamAttempt, error) {
	userExamAttempts, err := e.examAtemptRepository.GetByExam(ctx, generatedExamId, userId)
	if err != nil {
		return nil, err
	}

	generatedExam, err := e.generatedExamRepository.GetById(ctx, generatedExamId)
	if err != nil {
		return nil, err
	}

	hasAccess, err := e.accessService.UserHasAccessToExam(ctx, generatedExam.Edges.Exam.ID, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to check access: %w", err)
	}

	if !hasAccess {
		return nil, errors.New("forbidden")
	}

	currAttempts := len(userExamAttempts)

	examSettings, err := e.examSettingRepository.GetByExam(ctx, generatedExam.Edges.Exam.ID)
	if err != nil {
		return nil, err
	}

	if examSettings.MaxAttempts <= currAttempts {
		return nil, errors.New("max attempts for exam exceeded")
	}

	currentAttempt, err := e.examAtemptRepository.Create(ctx, currAttempts, generatedExamId, userId)
	if err != nil {
		return nil, err
	}

	return currentAttempt, nil
}

func (e *ExamAttemptService) GetAttempts(ctx context.Context, userId string) ([]*models.UserExamAttempt, error) {
	examWithAttempts, err := e.generatedExamRepository.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	var userExamAttempts []*models.UserExamAttempt

	for _, generatedExam := range examWithAttempts {
		userExamAttempt := &models.UserExamAttempt{
			Id:           generatedExam.ID,
			IsActive:     generatedExam.IsActive,
			ExamName:     generatedExam.Edges.Exam.Name,
			ExamCategory: generatedExam.Edges.Exam.Edges.Category.Name,
			Topic:        generatedExam.RawExamData["topic"].(string),
			Type:         generatedExam.RawExamData["type"].(string),
			Attempts:     []models.Attempt{},
		}

		for i, attempt := range generatedExam.Edges.Attempts {
			attemptModel := models.Attempt{
				AttemptId:     attempt.ID,
				AttemptNumber: i + 1,
				AttemptDate:   attempt.UpdatedAt,
			}

			if attempt.Edges.Assesment != nil {
				attemptModel.AssessmentId = attempt.Edges.Assesment.ID
			}

			userExamAttempt.Attempts = append(userExamAttempt.Attempts, attemptModel)
		}

		userExamAttempts = append(userExamAttempts, userExamAttempt)
	}

	return userExamAttempts, nil
}
