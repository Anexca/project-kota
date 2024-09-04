package services

import (
	"common/ent"
	commonRepositories "common/repositories"
	"context"
)

type ExamAssesmentService struct {
	promptService           *PromptService
	examAttemptRepository   *commonRepositories.ExamAttemptRepository
	examAssesmentRepository *commonRepositories.ExamAssesmentRepository
}

type AssesmentRequest struct {
	CompletedMinutes int `json:"completed_minutes" validate:"required"`
}

func NewExamAssesmentService(dbClient *ent.Client) *ExamAssesmentService {
	promptService := NewPromptService()
	examAttemptRepository := commonRepositories.NewExamAttemptRepository(dbClient)
	examAssesmentRepository := commonRepositories.NewExamAssesmentRepository(dbClient)

	return &ExamAssesmentService{
		promptService:           promptService,
		examAttemptRepository:   examAttemptRepository,
		examAssesmentRepository: examAssesmentRepository,
	}
}

func (e *ExamAssesmentService) StartNewAssesment(ctx context.Context, attempt *ent.ExamAttempt, request *AssesmentRequest) (*ent.ExamAssesment, error) {
	assesmentModel := commonRepositories.AssesmentModel{
		CompletedSeconds: request.CompletedMinutes,
	}

	return e.examAssesmentRepository.Create(ctx, attempt.ID, assesmentModel)

}

func (e *ExamAssesmentService) GetAssesmentById(ctx context.Context, assesmentId int, userId string) (*ent.ExamAssesment, error) {
	return e.examAssesmentRepository.GetById(ctx, assesmentId, userId)
}
