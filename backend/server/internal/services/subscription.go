package services

import (
	"common/ent"
	commonRepositories "common/repositories"
	"context"
	"fmt"

	"github.com/razorpay/razorpay-go"
)

type SubscriptionService struct {
	paymentService             *PaymentService
	subscriptionRepository     *commonRepositories.SubscriptionRepository
	userSubscriptionRepository *commonRepositories.UserSubscriptioRepository
}

func NewSubscriptionService(dbClient *ent.Client, paymentClient *razorpay.Client) *SubscriptionService {
	paymentService := NewPaymentService(paymentClient)
	subscriptionRepository := commonRepositories.NewSubscriptionRepository(dbClient)
	userSubscriptionRepository := commonRepositories.NewUserSubscriptioRepository(dbClient)

	return &SubscriptionService{
		paymentService:             paymentService,
		subscriptionRepository:     subscriptionRepository,
		userSubscriptionRepository: userSubscriptionRepository,
	}
}

func (s *SubscriptionService) GetAll(ctx context.Context) ([]*ent.Subscription, error) {
	return s.subscriptionRepository.GetAll(ctx)
}

func (s *SubscriptionService) StartUserSubscription(ctx context.Context, subscriptionId int, userId string) (*ent.UserSubscription, error) {
	subscription, err := s.subscriptionRepository.GetById(ctx, subscriptionId)
	if err != nil {
		return nil, err
	}

	model := CreateSubscriptionModel{
		ProviderPlanId:     subscription.ProviderPlanID,
		TotalBillingCycles: 3,
	}

	createdSubscription, err := s.paymentService.CreateSubscription(model)
	if err != nil {
		return nil, err
	}

	createdSubscriptionId, ok := createdSubscription["id"].(string)
	if !ok {
		return nil, fmt.Errorf("cound not get id from object %v", createdSubscription)
	}

	userSubscriptionModel := commonRepositories.UserSubscriptionModel{
		SubscriptionId:         subscriptionId,
		UserId:                 userId,
		ProviderSubscriptionId: createdSubscriptionId,
	}

	userSubscription, err := s.userSubscriptionRepository.Create(ctx, userSubscriptionModel)
	if err != nil {
		return nil, err
	}

	return userSubscription, nil
}
