package repositories

import (
	"context"

	"common/ent"
	"common/ent/subscription"
)

// SubscriptionRepositoryInterface defines the contract for the Subscription repository.
type SubscriptionRepositoryInterface interface {
	GetAll(ctx context.Context) ([]*ent.Subscription, error)
	GetById(ctx context.Context, subscriptionId int) (*ent.Subscription, error)
}

// SubscriptionRepository is a concrete implementation of SubscriptionRepositoryInterface.
type SubscriptionRepository struct {
	dbClient *ent.Client
}

// NewSubscriptionRepository creates a new instance of SubscriptionRepository.
func NewSubscriptionRepository(dbClient *ent.Client) *SubscriptionRepository {
	return &SubscriptionRepository{
		dbClient: dbClient,
	}
}

// GetAll retrieves all subscriptions.
func (s *SubscriptionRepository) GetAll(ctx context.Context) ([]*ent.Subscription, error) {
	return s.dbClient.Subscription.Query().All(ctx)
}

// GetById retrieves a subscription by its ID, including related exams.
func (s *SubscriptionRepository) GetById(ctx context.Context, subscriptionId int) (*ent.Subscription, error) {
	return s.dbClient.Subscription.Query().
		Where(subscription.IDEQ(subscriptionId)).
		WithExams(func(seq *ent.SubscriptionExamQuery) {
			seq.WithExam()
		}).
		Only(ctx)
}
