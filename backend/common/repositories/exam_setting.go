package repositories

import (
	"context"

	"common/ent"
	"common/ent/exam"
	"common/ent/examsetting"
)

type ExamSettingRepository struct {
	dbClient *ent.Client
}

func NewExamSettingRepository(dbClient *ent.Client) *ExamSettingRepository {
	return &ExamSettingRepository{
		dbClient: dbClient,
	}
}

func (e *ExamSettingRepository) GetByExam(ctx context.Context, examId int) (*ent.ExamSetting, error) {
	return e.dbClient.ExamSetting.
		Query().
		Where(examsetting.HasExamWith(exam.ID(examId))).
		WithExam().
		Only(ctx)
}
