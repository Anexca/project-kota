package repositories

import (
	"common/ent"
	"context"
)

type ExamAssesmentRepository struct {
	dbClient *ent.Client
}

type AssesmentModel struct {
	CompletedMinutes int
}

func NewExamAssesmentRepository(dbClient *ent.Client) *ExamAssesmentRepository {
	return &ExamAssesmentRepository{
		dbClient: dbClient,
	}
}

func (e *ExamAssesmentRepository) Create(ctx context.Context, attemptId int, model AssesmentModel) (*ent.ExamAssesment, error) {
	return e.dbClient.ExamAssesment.Create().
		SetAttemptID(attemptId).
		SetCompletedMinutes(model.CompletedMinutes).
		Save(ctx)
}
