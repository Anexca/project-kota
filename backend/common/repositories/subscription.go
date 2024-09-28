package repositories

import (
	"context"

	"common/ent"
	"common/ent/subscription"
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

func (s *SubscriptionRepository) GetById(ctx context.Context, subscriptionId int) (*ent.Subscription, error) {
	return s.dbClient.Subscription.Query().
		Where(subscription.IDEQ(subscriptionId)).
		WithExams(func(seq *ent.SubscriptionExamQuery) {
			seq.WithExam()
		}).
		Only(ctx)
}
