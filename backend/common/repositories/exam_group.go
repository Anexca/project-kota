package repositories

import (
	"context"

	"common/ent"
	"common/ent/exam"
	"common/ent/examgroup"
	"common/ent/generatedexam"
)

// ExamGroupRepository is a concrete implementation of ExamGroupRepositoryInterface.
type ExamGroupRepository struct {
	dbClient *ent.Client
}

// NewExamGroupRepository creates a new instance of ExamGroupRepository.
func NewExamGroupRepository(dbClient *ent.Client) *ExamGroupRepository {
	return &ExamGroupRepository{
		dbClient: dbClient,
	}
}

// GetById retrieves an exam group by its ID.
func (e *ExamGroupRepository) GetById(ctx context.Context, examGroupId int) (*ent.ExamGroup, error) {
	return e.dbClient.ExamGroup.Query().
		Where(examgroup.IDEQ(examGroupId)).
		Only(ctx)
}

// GetActiveByIdWithExams retrieves an active exam group by its ID, including active exams and their settings, categories, and generated exams.
func (e *ExamGroupRepository) GetActiveByIdWithExams(ctx context.Context, examGroupId int, isActive bool) (*ent.ExamGroup, error) {
	return e.dbClient.ExamGroup.Query().
		Where(examgroup.IDEQ(examGroupId), examgroup.IsActiveEQ(isActive)).
		WithExams(func(eq *ent.ExamQuery) {
			eq.Where(exam.IsActiveEQ(isActive)).
				WithSetting().
				WithCategory().
				WithGroup().
				WithGeneratedexams(func(geq *ent.GeneratedExamQuery) {
					geq.Where(generatedexam.IsActiveEQ(isActive))
				})
		}).
		Only(ctx)
}
