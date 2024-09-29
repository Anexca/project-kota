package repositories

import (
	"context"

	"common/constants"
	"common/ent"
	"common/ent/examcategory"
	"common/ent/examgroup"
)

// ExamCategoryRepositoryInterface defines the contract for the exam category repository.
type ExamCategoryRepositoryInterface interface {
	Get(ctx context.Context) ([]*ent.ExamCategory, error)
	GetByName(ctx context.Context, categoryName constants.ExamCategoryName) (*ent.ExamCategory, error)
}

// ExamCategoryRepository is a concrete implementation of ExamCategoryRepositoryInterface.
type ExamCategoryRepository struct {
	dbClient *ent.Client
}

// NewExamCategoryRepository creates a new instance of ExamCategoryRepository.
func NewExamCategoryRepository(dbClient *ent.Client) *ExamCategoryRepository {
	return &ExamCategoryRepository{
		dbClient: dbClient,
	}
}

// Get retrieves all exam categories from the database.
func (e *ExamCategoryRepository) Get(ctx context.Context) ([]*ent.ExamCategory, error) {
	return e.dbClient.ExamCategory.Query().All(ctx)
}

// GetByName retrieves an exam category by name, along with its associated exam groups and exams.
func (e *ExamCategoryRepository) GetByName(ctx context.Context, categoryName constants.ExamCategoryName) (*ent.ExamCategory, error) {
	return e.dbClient.ExamCategory.Query().
		Where(examcategory.NameEQ(examcategory.Name(categoryName))).
		WithGroups(func(egq *ent.ExamGroupQuery) {
			egq.Order(ent.Desc(examgroup.FieldIsActive), ent.Asc(examgroup.FieldID))
		}).
		WithExams().
		Only(ctx)
}
