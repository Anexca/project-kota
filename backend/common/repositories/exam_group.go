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

type ExamGroupRepository struct {
	dbClient *ent.Client
}

func NewExamGroupRepository(dbClient *ent.Client) *ExamGroupRepository {
	return &ExamGroupRepository{
		dbClient: dbClient,
	}
}

func (e *ExamGroupRepository) GetById(ctx context.Context, examGroupId int) (*ent.ExamGroup, error) {
	return e.dbClient.ExamGroup.Query().
		Where(examgroup.IDEQ(examGroupId)).
		Only(ctx)
}

func (e *ExamGroupRepository) GetActiveByIdWithExams(ctx context.Context, examGroupId int, isActive bool, examCategory constants.ExamCategoryName, examType constants.ExamType) (*ent.ExamGroup, error) {
	return e.dbClient.ExamGroup.Query().
		Where(examgroup.IDEQ(examGroupId), examgroup.IsActiveEQ(isActive)).
		WithExams(func(eq *ent.ExamQuery) {
			eq.Where(exam.IsActiveEQ(isActive), exam.TypeEQ(exam.Type(examType)), exam.HasCategoryWith(examcategory.NameEQ(examcategory.Name(examCategory))))
			eq.WithSetting()
			eq.WithCategory()
			eq.WithGeneratedexams(func(geq *ent.GeneratedExamQuery) {
				geq.Where(generatedexam.IsActiveEQ(isActive))
			})
		}).
		Only(ctx)
}
