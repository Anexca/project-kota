package repositories

import (
	"common/ent"
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

func (c *CachedQuestionMetaDataRepository) Create(ctx context.Context, cacheUID string, expiry time.Duration) (*ent.CachedQuestionMetaData, error) {
	expiresAt := time.Now().Add(expiry)

	return c.dbClient.CachedQuestionMetaData.Create().
		SetCacheUID(cacheUID).
		SetExpiresAt(expiresAt).
		Save(ctx)
}
