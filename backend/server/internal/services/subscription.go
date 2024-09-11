package services

import (
	"common/ent"
	commonRepositories "common/repositories"
	"context"
)

type SubscriptionService struct {
	subscriptionRepository *commonRepositories.SubscriptionRepository
}

func NewSubscriptionService(dbClient *ent.Client) *SubscriptionService {
	subscriptionRepository := commonRepositories.NewSubscriptionRepository(dbClient)

	return &SubscriptionService{
		subscriptionRepository: subscriptionRepository,
	}
}

func (s *SubscriptionService) GetAll(ctx context.Context) ([]*ent.Subscription, error) {
	return s.subscriptionRepository.GetAll(ctx)
}
