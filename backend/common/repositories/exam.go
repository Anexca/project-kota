package repositories

import (
	"context"

	"common/constants"
	"common/ent"
	"common/ent/exam"
	"common/ent/examcategory"
	"common/ent/examgroup"
	"common/ent/generatedexam"
)

// ExamRepository is a concrete implementation of ExamRepositoryInterface.
type ExamRepository struct {
	dbClient *ent.Client
}

// NewExamRepository creates a new instance of ExamRepository.
func NewExamRepository(dbClient *ent.Client) *ExamRepository {
	return &ExamRepository{
		dbClient: dbClient,
	}
}

// GetById retrieves an exam by its ID.
func (e *ExamRepository) GetById(ctx context.Context, examId int) (*ent.Exam, error) {
	return e.dbClient.Exam.Query().
		Where(exam.IDEQ(examId)).
		WithSetting().
		WithGeneratedexams().
		Only(ctx)
}

// GetActiveByExamsGroupId retrieves active exams for a specific exam group ID.
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

// GetActiveById retrieves an active exam by its ID.
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

// GetByExamCategory retrieves all exams for a given exam category.
func (e *ExamRepository) GetByExamCategory(ctx context.Context, examCategory *ent.ExamCategory) ([]*ent.Exam, error) {
	return e.dbClient.Exam.
		Query().
		Where(exam.HasCategoryWith(examcategory.ID(examCategory.ID)), exam.IsActiveEQ(true)).
		All(ctx)
}

// GetActiveByType retrieves active exams by type.
func (e *ExamRepository) GetActiveByType(ctx context.Context, examType constants.ExamType) ([]*ent.Exam, error) {
	return e.dbClient.Exam.Query().Where(
		exam.TypeEQ(exam.Type(examType)),
		exam.IsActiveEQ(true),
	).All(ctx)
}

// GetByName retrieves an exam by its name.
func (e *ExamRepository) GetByName(ctx context.Context, examName string) (*ent.Exam, error) {
	return e.dbClient.Exam.Query().
		Where(exam.NameEQ(examName)).
		Only(ctx)
}
