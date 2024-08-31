package repositories

import (
	"common/ent"
	"common/ent/exam"
	"common/ent/examcategory"
	"context"
)

type ExamRepository struct {
	dbClient *ent.Client
}

func NewExamRespository(dbClient *ent.Client) *ExamRepository {
	return &ExamRepository{
		dbClient: dbClient,
	}
}

func (e *ExamRepository) GetByExamCategory(ctx context.Context, examCategor *ent.ExamCategory) ([]*ent.Exam, error) {
	return e.dbClient.Exam.
		Query().
		Where(exam.HasCategoryWith(
			examcategory.ID(examCategor.ID),
		)).
		All(ctx)
}
