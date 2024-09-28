package repositories

import (
	"common/ent"
	"common/ent/exam"
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

func (e *ExamGroupRepository) GetActiveByIdWithExams(ctx context.Context, examGroupId int, isActive bool) (*ent.ExamGroup, error) {
	return e.dbClient.ExamGroup.Query().
		Where(examgroup.IDEQ(examGroupId), examgroup.IsActiveEQ(isActive)).
		WithExams(func(eq *ent.ExamQuery) {
			eq.Where(exam.IsActiveEQ(isActive))
			eq.WithSetting()
			eq.WithCategory()
			eq.WithGeneratedexams(func(geq *ent.GeneratedExamQuery) {
				geq.Where(generatedexam.IsActiveEQ(isActive))
			})
		}).
		Only(ctx)
}
