package repositories

import (
	"common/ent"
	"common/ent/exam"
	"common/ent/examsetting"
	"context"
)

type ExamSettingRepository struct {
	dbClient *ent.Client
}

func NewExamSettingRepository(dbClient *ent.Client) *ExamSettingRepository {
	return &ExamSettingRepository{
		dbClient: dbClient,
	}
}

func (e *ExamSettingRepository) GetByExam(ctx context.Context, ex *ent.Exam) (*ent.ExamSetting, error) {
	return e.dbClient.ExamSetting.
		Query().
		Where(examsetting.HasExamWith(exam.ID(ex.ID))).
		Only(ctx)
}
