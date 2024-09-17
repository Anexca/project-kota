package repositories

import (
	"common/constants"
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

func (e *ExamRepository) GetById(ctx context.Context, examId int) (*ent.Exam, error) {
	return e.dbClient.Exam.Query().
		Where(exam.IDEQ(examId)).
		WithSetting().
		WithGeneratedexams().
		Only(ctx)
}

func (e *ExamRepository) GetByExamCategory(ctx context.Context, examCategory *ent.ExamCategory) ([]*ent.Exam, error) {
	return e.dbClient.Exam.
		Query().
		Where(exam.HasCategoryWith(
			examcategory.ID(examCategory.ID),
		), exam.IsActiveEQ(true)).
		All(ctx)
}

func (e *ExamRepository) GetByName(ctx context.Context, name string) (*ent.Exam, error) {
	return e.dbClient.Exam.Query().
		Where(exam.NameEQ(name)).
		WithSetting().
		WithGeneratedexams(func(geq *ent.GeneratedExamQuery) {
			geq.Where(generatedexam.IsActiveEQ(true))
		}).
		First(ctx)
}

func (e *ExamRepository) GetByType(ctx context.Context, name string) ([]*ent.Exam, error) {
	return e.dbClient.Exam.Query().
		Where(exam.NameContains(name)).
		All(ctx)

}

func (e *ExamRepository) GetActiveByType(ctx context.Context, examType constants.ExamType) ([]*ent.Exam, error) {
	return e.dbClient.Exam.Query().Where(
		exam.TypeEQ(exam.Type(examType)),
	).All(ctx)
}
