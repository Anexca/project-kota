package repositories

import (
	"common/ent"
	"common/ent/examattempt"
	"common/ent/generatedexam"
	"context"

	"github.com/google/uuid"
)

type ExamAttemptRepository struct {
	dbClient *ent.Client
}

func NewExamAttemptRepository(dbClient *ent.Client) *ExamSettingRepository {
	return &ExamSettingRepository{
		dbClient: dbClient,
	}
}

func (e *ExamAttemptRepository) GetByExam(ctx context.Context, generatedExamId int) ([]*ent.ExamAttempt, error) {
	return e.dbClient.ExamAttempt.Query().
		Where(examattempt.HasGeneratedexamWith(generatedexam.ID(generatedExamId))).
		All(ctx)
}

func (e *ExamAttemptRepository) Create(ctx context.Context, currentAttempt int, generatedExamId int, userId uuid.UUID) (*ent.ExamAttempt, error) {
	return e.dbClient.ExamAttempt.Create().
		SetGeneratedexamID(generatedExamId).
		SetUserID(userId).
		Save(ctx)
}
