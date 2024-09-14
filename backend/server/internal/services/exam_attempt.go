package services

import (
	"common/ent"
	commonRepositories "common/repositories"
	"context"
	"errors"
	"fmt"
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

func (e *ExamAttemptService) GetAttempts(ctx context.Context, userId string) ([]*ent.GeneratedExam, error) {
	return e.generatedExamRepository.GetByUserId(ctx, userId)
}
