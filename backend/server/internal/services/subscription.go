package services

import (
	"common/ent"
	"common/ent/usersubscription"
	commonRepositories "common/repositories"
	"context"
	"errors"
	"fmt"
	"log"
	"server/pkg/config"
	"server/pkg/models"
	"time"

	cashfree_pg "github.com/cashfree/cashfree-pg/v4"
)

type SubscriptionService struct {
	environment                *config.Environment
	paymentService             *PaymentService
	userRepository             *commonRepositories.UserRepository
	subscriptionRepository     *commonRepositories.SubscriptionRepository
	userSubscriptionRepository *commonRepositories.UserSubscriptioRepository
	paymentrepository          *commonRepositories.PaymentRepository
}

func NewSubscriptionService(dbClient *ent.Client) *SubscriptionService {
	environment, _ := config.LoadEnvironment()
	paymentService := NewPaymentService()
	userRepository := commonRepositories.NewUserRepository(dbClient)
	subscriptionRepository := commonRepositories.NewSubscriptionRepository(dbClient)
	userSubscriptionRepository := commonRepositories.NewUserSubscriptioRepository(dbClient)
	paymentrepository := commonRepositories.NewPaymentRepository(dbClient)

	return &SubscriptionService{
		environment:                environment,
		paymentService:             paymentService,
		userRepository:             userRepository,
		subscriptionRepository:     subscriptionRepository,
		userSubscriptionRepository: userSubscriptionRepository,
		paymentrepository:          paymentrepository,
	}
}

func (s *SubscriptionService) GetAll(ctx context.Context) ([]models.SubscriptionOverview, error) {
	subscriptions, err := s.subscriptionRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var subscriptionOverviews []models.SubscriptionOverview

	for _, subscription := range subscriptions {
		subscriptionOverview := models.SubscriptionOverview{
			Id:                  subscription.ID,
			ProviderPlanID:      subscription.ProviderPlanID,
			Name:                subscription.Name,
			IsActive:            subscription.IsActive,
			RawSubscriptionData: subscription.RawSubscriptionData,
			DurationInMonths:    subscription.DurationInMonths,
			Price:               subscription.Price,
		}

		subscriptionOverviews = append(subscriptionOverviews, subscriptionOverview)
	}

	return subscriptionOverviews, nil
}

func (s *SubscriptionService) StartUserSubscription(ctx context.Context, subscriptionId int, userId string) (*models.SubscriptionToActivate, error) {
	user, err := s.userRepository.Get(ctx, userId)
	if err != nil {
		return nil, err
	}

	if user.PaymentProviderCustomerID == "" {
		return nil, fmt.Errorf("user %s's profile does not have enough data to start subscription", user.Email)
	}

	subscription, err := s.subscriptionRepository.GetById(ctx, subscriptionId)
	if err != nil {
		return nil, err
	}

	// userHasSubscription, err := s.UserHasActiveSubscription(ctx, subscription, user)
	// if err != nil {
	// 	var notFoundError *ent.NotFoundError
	// 	if !errors.As(err, &notFoundError) {
	// 		return nil, err
	// 	}
	// }

	// if userHasSubscription {
	// 	return nil, fmt.Errorf("user already has active subscription")
	// }

	model := CreateOrderModel{
		Amount:              subscription.Price,
		CustomerId:          user.PaymentProviderCustomerID,
		CustomerPhoneNumber: user.PhoneNumber,
		CustomerEmail:       user.Email,
		UserId:              userId,
		CustomerName:        fmt.Sprintf("%s %s", user.FirstName, user.LastName),
	}

	createdSubscription, err := s.paymentService.CreateOrder(model)
	if err != nil {
		return nil, err
	}

	userSubscriptionModel := commonRepositories.UserSubscriptionModel{
		SubscriptionId:         subscriptionId,
		UserId:                 userId,
		ProviderSubscriptionId: *createdSubscription.OrderId,
	}

	userSubscription, err := s.userSubscriptionRepository.Create(ctx, userSubscriptionModel)
	if err != nil {
		return nil, err
	}

	subscriptionToActivate := &models.SubscriptionToActivate{
		Id:               userSubscription.ID,
		Status:           string(userSubscription.Status),
		SubscriptionId:   userSubscription.ProviderSubscriptionID,
		PaymentSessionId: *createdSubscription.PaymentSessionId,
	}

	return subscriptionToActivate, nil
}

func (s *SubscriptionService) ActivateUserSubscription(ctx context.Context, userSubscriptionId int, userId string) (*ent.UserSubscription, error) {
	userSubscriptionToUpdate, err := s.userSubscriptionRepository.GetById(ctx, userSubscriptionId, userId)
	if err != nil {
		return nil, err
	}

	isSuccessful, transaction, err := s.paymentService.IsOrderSuccessful(userSubscriptionToUpdate.ProviderSubscriptionID)
	if err != nil {
		return nil, err
	}

	if !isSuccessful {
		return nil, errors.New("payment for subscription was not successful")
	}

	userSubscriptionToUpdate.Status = usersubscription.StatusACTIVE
	userSubscriptionToUpdate.StartDate = time.Now()
	userSubscriptionToUpdate.IsActive = true
	userSubscriptionToUpdate.EndDate = time.Now().AddDate(0, userSubscriptionToUpdate.Edges.Subscription.DurationInMonths, 0)

	go func() {
		bgCtx := context.Background()
		payment, err := s.StorePaymentForSubscription(bgCtx, transaction, userSubscriptionId, userId)
		if err != nil {
			log.Printf("could not store payment %v", err)
			return
		}

		log.Println("payment info stored successfully with id", payment.ID)
	}()

	err = s.userSubscriptionRepository.Update(ctx, userSubscriptionToUpdate)
	if err != nil {
		return nil, err
	}

	return userSubscriptionToUpdate, nil
}

func (s *SubscriptionService) StorePaymentForSubscription(ctx context.Context, transaction *cashfree_pg.PaymentEntity, userSubscriptionId int, userId string) (*ent.Payment, error) {

	const layout = time.RFC3339

	paymentDate, err := time.Parse(layout, *transaction.PaymentCompletionTime)
	if err != nil {
		return nil, err
	}

	paymentModel := commonRepositories.CreatePaymentModel{
		Status:             *transaction.PaymentStatus,
		PaymentMethod:      determinePaymentMethod(transaction.PaymentMethod),
		PaymentDate:        paymentDate,
		Amount:             float64(*transaction.PaymentAmount),
		UserSubscriptionId: userSubscriptionId,
		ProviderPaymentId:  *transaction.CfPaymentId,
	}

	return s.paymentrepository.Create(ctx, paymentModel, userId)

}

func determinePaymentMethod(transaction *cashfree_pg.PaymentEntityPaymentMethod) string {
	if transaction.PaymentMethodUPIInPaymentsEntity != nil {
		return "UPI"
	} else if transaction.PaymentMethodCardInPaymentsEntity != nil {
		return "Card"
	} else if transaction.PaymentMethodNetBankingInPaymentsEntity != nil {
		return "Netbanking"
	}
	// Add more checks for other payment methods as needed
	return "Unknown"
}

func (s *SubscriptionService) UserHasActiveSubscription(ctx context.Context, subscription *ent.Subscription, user *ent.User) (bool, error) {
	userSubscriptions, err := s.userSubscriptionRepository.GetByUserId(ctx, user.ID.String())
	if err != nil {
		return false, err
	}

	now := time.Now()

	for _, userSubscription := range userSubscriptions {
		if userSubscription.StartDate.Before(now) &&
			userSubscription.EndDate.After(now) &&
			userSubscription.Edges.Subscription.ID == subscription.ID {
			return true, nil
		}
	}

	return false, nil
}
