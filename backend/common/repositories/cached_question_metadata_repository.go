package repositories

import (
	"common/ent"
	"common/ent/cachedquestionmetadata"
	"common/ent/exam"
	"context"
	"time"
)

type CachedQuestionMetaDataRepository struct {
	dbClient *ent.Client
}

func NewCachedQuestionMetaDataRepository(dbClient *ent.Client) *CachedQuestionMetaDataRepository {
	return &CachedQuestionMetaDataRepository{
		dbClient: dbClient,
	}
}

func (c *CachedQuestionMetaDataRepository) Create(ctx context.Context, cacheUID string, expiry time.Duration, exam *ent.Exam) (*ent.CachedQuestionMetaData, error) {
	expiresAt := time.Now().Add(expiry)

	return c.dbClient.CachedQuestionMetaData.Create().
		SetCacheUID(cacheUID).
		SetExpiresAt(expiresAt).
		SetExam(exam).
		Save(ctx)
}

func (c *CachedQuestionMetaDataRepository) GetByExam(ctx context.Context, ex *ent.Exam) ([]*ent.CachedQuestionMetaData, error) {
	return c.dbClient.CachedQuestionMetaData.Query().
		Where(cachedquestionmetadata.HasExamWith(exam.ID(ex.ID)), cachedquestionmetadata.IsUsed(false)).
		All(ctx)
}

func (c *CachedQuestionMetaDataRepository) MarkAsUsed(ctx context.Context, id int) error {
	_, err := c.dbClient.CachedQuestionMetaData.UpdateOneID(id).SetIsUsed(true).Save(ctx)
	return err
}
