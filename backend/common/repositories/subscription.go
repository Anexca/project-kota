package repositories

import "common/ent"

type SubscriptionRepository struct {
	dbClient *ent.Client
}

func NewSubscriptionRepository(dbClient *ent.Client) *SubscriptionRepository {
	return &SubscriptionRepository{
		dbClient: dbClient,
	}
}
