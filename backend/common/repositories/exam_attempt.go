package repositories

import (
	"context"

	"github.com/google/uuid"

	"common/ent"
	"common/ent/examattempt"
	"common/ent/generatedexam"
	"common/ent/user"
)

// ExamAttemptRepository is a concrete implementation of ExamAttemptRepositoryInterface.
type ExamAttemptRepository struct {
	dbClient *ent.Client
}

// NewExamAttemptRepository creates a new instance of ExamAttemptRepository.
func NewExamAttemptRepository(dbClient *ent.Client) *ExamAttemptRepository {
	return &ExamAttemptRepository{
		dbClient: dbClient,
	}
}

// GetById retrieves an exam attempt by its ID and user ID.
func (e *ExamAttemptRepository) GetById(ctx context.Context, attemptId int, userId string) (*ent.ExamAttempt, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return e.dbClient.ExamAttempt.Query().
		Where(examattempt.IDEQ(attemptId), examattempt.HasUserWith(user.IDEQ(userUid))).
		WithGeneratedexam(
			func(query *ent.GeneratedExamQuery) {
				query.WithExam()
			},
		).
		Only(ctx)
}

// GetByUserId retrieves all exam attempts for a given user.
func (e *ExamAttemptRepository) GetByUserId(ctx context.Context, userId string) ([]*ent.ExamAttempt, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return e.dbClient.ExamAttempt.Query().
		Where(examattempt.HasUserWith(user.IDEQ(userUid))).
		WithGeneratedexam().
		All(ctx)
}

// GetByExam retrieves all exam attempts for a specific generated exam and user.
func (e *ExamAttemptRepository) GetByExam(ctx context.Context, generatedExamId int, userId string) ([]*ent.ExamAttempt, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return e.dbClient.ExamAttempt.Query().
		Where(examattempt.HasGeneratedexamWith(generatedexam.ID(generatedExamId)), examattempt.HasUserWith(user.ID(userUid))).
		WithGeneratedexam().
		All(ctx)
}

// Create creates a new exam attempt for a given user and generated exam.
func (e *ExamAttemptRepository) Create(ctx context.Context, currentAttempt int, generatedExamId int, userId string) (*ent.ExamAttempt, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return e.dbClient.ExamAttempt.Create().
		SetAttemptNumber(currentAttempt + 1).
		SetGeneratedexamID(generatedExamId).
		SetUserID(userUid).
		Save(ctx)
}
