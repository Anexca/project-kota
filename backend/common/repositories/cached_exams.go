package repositories

import (
	"context"
	"time"

	"common/ent"
	"common/ent/cachedexam"
	"common/ent/exam"
)

// CachedExamRepositoryInterface defines the contract for the cached exam repository.
type CachedExamRepositoryInterface interface {
	Create(ctx context.Context, cacheUID string, expiry time.Duration, exam *ent.Exam) (*ent.CachedExam, error)
	GetByExam(ctx context.Context, ex *ent.Exam) ([]*ent.CachedExam, error)
	MarkAsUsed(ctx context.Context, id int) error
}

// CachedExamRepository is the concrete implementation of the CachedExamRepositoryInterface.
type CachedExamRepository struct {
	dbClient *ent.Client
}

// NewCachedExamRepository creates a new CachedExamRepository.
func NewCachedExamRepository(dbClient *ent.Client) *CachedExamRepository {
	return &CachedExamRepository{
		dbClient: dbClient,
	}
}

// Create stores a new cached exam record with a given expiry duration.
func (c *CachedExamRepository) Create(ctx context.Context, cacheUID string, expiry time.Duration, exam *ent.Exam) (*ent.CachedExam, error) {
	expiresAt := time.Now().Add(expiry)

	return c.dbClient.CachedExam.Create().
		SetCacheUID(cacheUID).
		SetExpiresAt(expiresAt).
		SetExam(exam).
		Save(ctx)
}

// GetByExam retrieves all cached exam entries for a specific exam that are not marked as used.
func (c *CachedExamRepository) GetByExam(ctx context.Context, ex *ent.Exam) ([]*ent.CachedExam, error) {
	return c.dbClient.CachedExam.Query().
		Where(cachedexam.HasExamWith(exam.ID(ex.ID)), cachedexam.IsUsed(false)).
		All(ctx)
}

// MarkAsUsed marks a cached exam entry as used by its ID.
func (c *CachedExamRepository) MarkAsUsed(ctx context.Context, id int) error {
	_, err := c.dbClient.CachedExam.UpdateOneID(id).SetIsUsed(true).Save(ctx)
	return err
}
