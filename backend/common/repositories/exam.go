package repositories

import (
	"common/ent"
	"common/ent/exam"
	"common/ent/examcategory"
	"common/ent/generatedexam"
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

func (e *ExamRepository) GetByExamCategory(ctx context.Context, examCategory *ent.ExamCategory) ([]*ent.Exam, error) {
	return e.dbClient.Exam.
		Query().
		Where(exam.HasCategoryWith(
			examcategory.ID(examCategory.ID),
		)).
		All(ctx)
}

func (e *ExamRepository) GetByName(ctx context.Context, name string) (*ent.Exam, error) {
	return e.dbClient.Exam.Query().
		Where(exam.Name(name)).
		WithSetting().
		WithGeneratedexams(func(geq *ent.GeneratedExamQuery) {
			geq.Where(generatedexam.IsActiveEQ(true))
		}).
		First(ctx)
}
