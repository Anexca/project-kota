package repositories

import (
	"common/constants"
	"common/ent"
	"common/ent/exam"
	"common/ent/examcategory"
	"common/ent/examgroup"
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

func (e *ExamRepository) GetActiveByExamsGroupId(ctx context.Context, examGroupId int, isActive bool) ([]*ent.Exam, error) {
	return e.dbClient.Exam.Query().
		Where(exam.IsActiveEQ(isActive), exam.HasGeneratedexamsWith(generatedexam.IsActiveEQ(isActive))).
		WithSetting().
		WithCategory().
		WithGeneratedexams(func(geq *ent.GeneratedExamQuery) {
			geq.Where(generatedexam.IsActiveEQ(isActive))
		}).
		WithGroup(func(egq *ent.ExamGroupQuery) {
			egq.Where(examgroup.IDEQ(examGroupId), examgroup.IsActiveEQ(isActive))
		}).
		Order(ent.Desc(exam.FieldUpdatedAt)).
		All(ctx)
}

func (e *ExamRepository) GetActiveById(ctx context.Context, examId int, isActive bool) (*ent.Exam, error) {
	return e.dbClient.Exam.Query().
		Where(exam.IDEQ(examId), exam.IsActiveEQ(isActive), exam.HasGeneratedexamsWith(generatedexam.IsActiveEQ(isActive))).
		WithSetting().
		WithCategory().
		WithGeneratedexams(func(geq *ent.GeneratedExamQuery) {
			geq.Where(generatedexam.IsActiveEQ(isActive))
		}).
		Order(ent.Desc(exam.FieldUpdatedAt)).
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

func (e *ExamRepository) GetActiveByType(ctx context.Context, examType constants.ExamType) ([]*ent.Exam, error) {
	return e.dbClient.Exam.Query().Where(
		exam.TypeEQ(exam.Type(examType)),
		exam.IsActiveEQ(true),
	).All(ctx)
}

func (e *ExamRepository) GetByName(ctx context.Context, examName string) (*ent.Exam, error) {
	return e.dbClient.Exam.Query().
		Where(exam.NameEQ(examName)).
		Only(ctx)
}
