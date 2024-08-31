package repositories

import (
	"common/ent"
	"common/ent/examcategory"
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

func (e *ExamCategoryRepository) GetByName(ctx context.Context, name string) (*ent.ExamCategory, error) {
	return e.dbClient.ExamCategory.Query().Where(examcategory.Name(name)).First(ctx)
}
