package repositories

import (
	"context"
	"time"

	"common/ent"
	"common/ent/cachedexam"
	"common/ent/exam"
)

type CachedExamRepository struct {
	dbClient *ent.Client
}

func NewCachedExamRepository(dbClient *ent.Client) *CachedExamRepository {
	return &CachedExamRepository{
		dbClient: dbClient,
	}
}

func (c *CachedExamRepository) Create(ctx context.Context, cacheUID string, expiry time.Duration, exam *ent.Exam) (*ent.CachedExam, error) {
	expiresAt := time.Now().Add(expiry)

	return c.dbClient.CachedExam.Create().
		SetCacheUID(cacheUID).
		SetExpiresAt(expiresAt).
		SetExam(exam).
		Save(ctx)
}

func (c *CachedExamRepository) GetByExam(ctx context.Context, ex *ent.Exam) ([]*ent.CachedExam, error) {
	return c.dbClient.CachedExam.Query().
		Where(cachedexam.HasExamWith(exam.ID(ex.ID)), cachedexam.IsUsed(false)).
		All(ctx)
}

func (c *CachedExamRepository) MarkAsUsed(ctx context.Context, id int) error {
	_, err := c.dbClient.CachedExam.UpdateOneID(id).SetIsUsed(true).Save(ctx)
	return err
}
