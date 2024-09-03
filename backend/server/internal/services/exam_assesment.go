package services

import (
	"common/ent"
	commonRepositories "common/repositories"
	"context"
)

type ExamAssesmentService struct {
	examAttemptRepository   *commonRepositories.ExamAttemptRepository
	examAssesmentRepository *commonRepositories.ExamAssesmentRepository
}

type AssesmentRequest struct {
	CompletedMinutes int `json:"completed_minutes" validate:"required"`
}

func NewExamAssesmentService(dbClient *ent.Client) *ExamAssesmentService {
	examAttemptRepository := commonRepositories.NewExamAttemptRepository(dbClient)
	examAssesmentRepository := commonRepositories.NewExamAssesmentRepository(dbClient)

	return &ExamAssesmentService{
		examAttemptRepository:   examAttemptRepository,
		examAssesmentRepository: examAssesmentRepository,
	}
}

func (e *ExamAssesmentService) StartNewAssesment(ctx context.Context, attempt *ent.ExamAttempt, request *AssesmentRequest) (*ent.ExamAssesment, error) {
	assesmentModel := commonRepositories.AssesmentModel{
		CompletedMinutes: request.CompletedMinutes,
	}

	return e.examAssesmentRepository.Create(ctx, attempt.ID, assesmentModel)

}
