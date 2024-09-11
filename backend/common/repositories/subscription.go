package repositories

import (
	"common/ent"
	"context"
)

type SubscriptionRepository struct {
	dbClient *ent.Client
}

func NewSubscriptionRepository(dbClient *ent.Client) *SubscriptionRepository {
	return &SubscriptionRepository{
		dbClient: dbClient,
	}
}

func (s *SubscriptionRepository) GetAll(ctx context.Context) ([]*ent.Subscription, error) {
	return s.dbClient.Subscription.Query().All(ctx)
}
