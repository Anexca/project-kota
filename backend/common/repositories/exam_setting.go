package repositories

import (
	"context"

	"common/ent"
	"common/ent/exam"
	"common/ent/examsetting"
)

// ExamSettingRepositoryInterface defines the contract for the exam setting repository.
type ExamSettingRepositoryInterface interface {
	GetByExam(ctx context.Context, examId int) (*ent.ExamSetting, error)
}

// ExamSettingRepository is a concrete implementation of ExamSettingRepositoryInterface.
type ExamSettingRepository struct {
	dbClient *ent.Client
}

// NewExamSettingRepository creates a new instance of ExamSettingRepository.
func NewExamSettingRepository(dbClient *ent.Client) *ExamSettingRepository {
	return &ExamSettingRepository{
		dbClient: dbClient,
	}
}

// GetByExam retrieves the exam setting for a specific exam by its ID.
func (e *ExamSettingRepository) GetByExam(ctx context.Context, examId int) (*ent.ExamSetting, error) {
	return e.dbClient.ExamSetting.
		Query().
		Where(examsetting.HasExamWith(exam.ID(examId))).
		WithExam().
		Only(ctx)
}
