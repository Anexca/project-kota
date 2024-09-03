package services

import (
	"common/ent"
	commonRepositories "common/repositories"
	"context"
	"errors"
)

type ExamAttemptService struct {
	examRepository          *commonRepositories.ExamRepository
	examAtemptRepository    *commonRepositories.ExamAttemptRepository
	examSettingRepository   *commonRepositories.ExamSettingRepository
	generatedExamRepository *commonRepositories.GeneratedExamRepository
}

func NewExamAttemptService(dbClient *ent.Client) *ExamAttemptService {
	examAtemptRepository := commonRepositories.NewExamAttemptRepository(dbClient)
	examSettingRepository := commonRepositories.NewExamSettingRepository(dbClient)
	examRepository := commonRepositories.NewExamRespository(dbClient)
	generatedExamRepository := commonRepositories.NewGeneratedExamRepository(dbClient)

	return &ExamAttemptService{
		examAtemptRepository:    examAtemptRepository,
		examSettingRepository:   examSettingRepository,
		examRepository:          examRepository,
		generatedExamRepository: generatedExamRepository,
	}
}

func (e *ExamAttemptService) CheckAndAddAttempt(ctx context.Context, generatedExamId int, userId string) (*ent.ExamAttempt, error) {
	generatedExam, err := e.generatedExamRepository.GetById(ctx, generatedExamId)
	if err != nil {
		return nil, err
	}

	currAttempts := len(generatedExam.Edges.Attempts)

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
