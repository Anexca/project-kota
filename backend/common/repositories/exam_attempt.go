package repositories

import (
	"common/ent"
	"common/ent/examattempt"
	"common/ent/generatedexam"
	"common/ent/user"
	"context"

	"github.com/google/uuid"
)

type ExamAttemptRepository struct {
	dbClient *ent.Client
}

func NewExamAttemptRepository(dbClient *ent.Client) *ExamAttemptRepository {
	return &ExamAttemptRepository{
		dbClient: dbClient,
	}
}

func (e *ExamAttemptRepository) GetById(ctx context.Context, attemptId int, userId string) (*ent.ExamAttempt, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return e.dbClient.ExamAttempt.Query().
		Where(examattempt.IDEQ(attemptId), examattempt.HasUserWith(user.IDEQ(userUid))).
		WithGeneratedexam().
		Only(ctx)
}

func (e *ExamAttemptRepository) GetByExam(ctx context.Context, generatedExamId int, userId string) ([]*ent.ExamAttempt, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return e.dbClient.ExamAttempt.Query().
		Where(examattempt.HasGeneratedexamWith(generatedexam.ID(generatedExamId)), examattempt.HasUserWith(user.ID(userUid))).
		WithGeneratedexam().
		All(ctx)
}

func (e *ExamAttemptRepository) Create(ctx context.Context, currentAttempt int, generatedExamId int, userId string) (*ent.ExamAttempt, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return e.dbClient.ExamAttempt.Create().
		SetAttemptNumber(currentAttempt + 1).
		SetGeneratedexamID(generatedExamId).
		SetUserID(userUid).
		Save(ctx)
}
