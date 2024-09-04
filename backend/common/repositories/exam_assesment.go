package repositories

import (
	"common/ent"
	"common/ent/examassesment"
	"common/ent/examattempt"
	"common/ent/user"
	"context"

	"github.com/google/uuid"
)

type ExamAssesmentRepository struct {
	dbClient *ent.Client
}

type AssesmentModel struct {
	CompletedSeconds int
}

func NewExamAssesmentRepository(dbClient *ent.Client) *ExamAssesmentRepository {
	return &ExamAssesmentRepository{
		dbClient: dbClient,
	}
}

func (e *ExamAssesmentRepository) Create(ctx context.Context, attemptId int, model AssesmentModel) (*ent.ExamAssesment, error) {
	return e.dbClient.ExamAssesment.Create().
		SetAttemptID(attemptId).
		SetCompletedSeconds(model.CompletedSeconds).
		Save(ctx)
}

func (e *ExamAssesmentRepository) GetById(ctx context.Context, assesmentId int, userId string) (*ent.ExamAssesment, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return e.dbClient.ExamAssesment.Query().
		Where(examassesment.HasAttemptWith(examattempt.HasUserWith(user.ID(userUid))), examassesment.ID(assesmentId)).
		Only(ctx)
}
