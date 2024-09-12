package services

import (
	"common/ent"
	"common/ent/usersubscription"
	commonRepositories "common/repositories"
	"context"
	"errors"
	"fmt"

	"github.com/razorpay/razorpay-go"
)

type SubscriptionService struct {
	paymentService             *PaymentService
	userService                *UserService
	subscriptionRepository     *commonRepositories.SubscriptionRepository
	userSubscriptionRepository *commonRepositories.UserSubscriptioRepository
}

func NewSubscriptionService(dbClient *ent.Client, paymentClient *razorpay.Client) *SubscriptionService {
	paymentService := NewPaymentService(paymentClient)
	userService := NewUserService(dbClient, paymentClient)
	subscriptionRepository := commonRepositories.NewSubscriptionRepository(dbClient)
	userSubscriptionRepository := commonRepositories.NewUserSubscriptioRepository(dbClient)

	return &SubscriptionService{
		paymentService:             paymentService,
		userService:                userService,
		subscriptionRepository:     subscriptionRepository,
		userSubscriptionRepository: userSubscriptionRepository,
	}
}

func (s *SubscriptionService) GetAll(ctx context.Context) ([]*ent.Subscription, error) {
	return s.subscriptionRepository.GetAll(ctx)
}

func (s *SubscriptionService) StartUserSubscription(ctx context.Context, subscriptionId int, userId string) (*ent.UserSubscription, error) {
	user, err := s.userService.GetUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	if user.PaymentProviderCustomerID == "" {
		return nil, fmt.Errorf("user %s profile does not have enough data to start subscription", user.Email)
	}

	subscription, err := s.subscriptionRepository.GetById(ctx, subscriptionId)
	if err != nil {
		return nil, err
	}

	model := CreateSubscriptionModel{
		ProviderPlanId:     subscription.ProviderPlanID,
		TotalBillingCycles: 3,
		CustomerId:         user.PaymentProviderCustomerID,
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

func (s *SubscriptionService) CancelUserSubscription(ctx context.Context, userSubscriptionId int, userId string) (*ent.UserSubscription, error) {
	userSubscriptionToCancel, err := s.userSubscriptionRepository.GetById(ctx, userSubscriptionId, userId)
	if err != nil {
		return nil, err
	}

	if userSubscriptionToCancel.Status == usersubscription.StatusCANCELED {
		return nil, errors.New("user subscrption is already cancelled")
	}

	_, err = s.paymentService.CancelUserSubscription(userSubscriptionToCancel.ProviderSubscriptionID)
	if err != nil {
		return nil, fmt.Errorf("error canceling subscription with payment provider: %v", err)
	}

	userSubscriptionToCancel.Status = usersubscription.StatusCANCELED

	err = s.userSubscriptionRepository.Update(ctx, userSubscriptionToCancel)
	if err != nil {
		return nil, err
	}

	return userSubscriptionToCancel, nil
}
