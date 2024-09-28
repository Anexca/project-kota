package repositories

import (
	"common/constants"
	"common/ent"
	"common/ent/examcategory"
	"common/ent/examgroup"
	"context"
)

type ExamCategoryRepository struct {
	dbClient *ent.Client
}

func NewExamCategoryRepository(dbClient *ent.Client) *ExamCategoryRepository {
	return &ExamCategoryRepository{
		dbClient: dbClient,
	}
}

func (e *ExamCategoryRepository) Get(ctx context.Context) ([]*ent.ExamCategory, error) {
	return e.dbClient.ExamCategory.Query().All(ctx)
}

func (e *ExamCategoryRepository) GetByName(ctx context.Context, categoryName constants.ExamCategoryName) (*ent.ExamCategory, error) {
	return e.dbClient.ExamCategory.Query().
		Where(examcategory.NameEQ(examcategory.Name(categoryName))).
		WithGroups(func(egq *ent.ExamGroupQuery) {
			egq.Order(ent.Desc(examgroup.FieldIsActive), ent.Asc(examgroup.FieldID))
		}).
		WithExams().
		Only(ctx)
}
