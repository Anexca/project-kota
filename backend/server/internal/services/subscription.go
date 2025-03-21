package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"common/ent"
	"common/ent/usersubscription"
	commonInterfaces "common/interfaces"

	commonRepositories "common/repositories"

	cashfree_pg "github.com/cashfree/cashfree-pg/v4"

	"server/internal/interfaces"
	"server/pkg/config"
	"server/pkg/models"
)

// SubscriptionService defines the service for managing subscriptions.
type SubscriptionService struct {
	environment                *config.Environment
	paymentService             interfaces.PaymentServiceInterface // Use interface for better testability
	userRepository             commonInterfaces.UserRepositoryInterface
	subscriptionRepository     commonInterfaces.SubscriptionRepositoryInterface
	userSubscriptionRepository commonInterfaces.UserSubscriptionRepositoryInterface
	paymentrepository          commonInterfaces.PaymentRepositoryInterface
}

// NewSubscriptionService creates a new SubscriptionService with the provided dependencies.
func NewSubscriptionService(paymentService interfaces.PaymentServiceInterface, userRepo commonInterfaces.UserRepositoryInterface, subscriptionRepo commonInterfaces.SubscriptionRepositoryInterface, userSubscriptionRepo commonInterfaces.UserSubscriptionRepositoryInterface, paymentRepo commonInterfaces.PaymentRepositoryInterface) *SubscriptionService {
	return &SubscriptionService{
		environment:                nil, // To be set in Init method
		paymentService:             paymentService,
		userRepository:             userRepo,
		subscriptionRepository:     subscriptionRepo,
		userSubscriptionRepository: userSubscriptionRepo,
		paymentrepository:          paymentRepo,
	}
}

// InitSubscriptionService initializes the SubscriptionService for production use.
func InitSubscriptionService(dbClient *ent.Client) *SubscriptionService {
	environment, err := config.LoadEnvironment()
	if err != nil {
		log.Fatalf("error loading environment: %v", err)
	}

	paymentService := NewPaymentService()
	userRepository := commonRepositories.NewUserRepository(dbClient)
	subscriptionRepository := commonRepositories.NewSubscriptionRepository(dbClient)
	userSubscriptionRepository := commonRepositories.NewUserSubscriptionRepository(dbClient)
	paymentrepository := commonRepositories.NewPaymentRepository(dbClient)

	service := NewSubscriptionService(paymentService, userRepository, subscriptionRepository, userSubscriptionRepository, paymentrepository)
	service.environment = environment

	return service
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
			FinalPrice:          subscription.FinalPrice,
			BasePrice:           subscription.BasePrice,
		}

		subscriptionOverviews = append(subscriptionOverviews, subscriptionOverview)
	}

	return subscriptionOverviews, nil
}

func (s *SubscriptionService) StartUserSubscription(ctx context.Context, subscriptionId int, returnUrl *string, userId string) (*models.SubscriptionToActivate, error) {
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

	userHasSubscription, err := s.UserHasActiveSubscription(ctx, subscription, user)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if !errors.As(err, &notFoundError) {
			return nil, err
		}
	}

	if userHasSubscription {
		return nil, fmt.Errorf("user already has active subscription")
	}

	model := models.CreateOrderModel{
		Amount:              subscription.FinalPrice,
		CustomerId:          user.PaymentProviderCustomerID,
		CustomerPhoneNumber: user.PhoneNumber,
		CustomerEmail:       user.Email,
		UserId:              userId,
		CustomerName:        fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		ReturnUrl:           returnUrl,
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

func (s *SubscriptionService) ActivateUserSubscription(ctx context.Context, providerSubscriptionId string, userEmail string) (*models.ActivatedSubscription, error) {
	user, err := s.userRepository.GetByEmail(ctx, userEmail)
	if err != nil {
		return nil, err
	}

	userSubscriptionToUpdate, err := s.userSubscriptionRepository.GetByProviderSubscriptionId(ctx, providerSubscriptionId, user.ID.String())
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

	payment, err := s.StorePaymentForSubscription(ctx, transaction, userSubscriptionToUpdate.ID, user.ID.String())
	if err != nil {
		return nil, err
	}

	log.Println("payment info stored successfully with id", payment.ID)

	err = s.userSubscriptionRepository.Update(ctx, userSubscriptionToUpdate)
	if err != nil {
		return nil, err
	}

	activatedSubscription := models.ActivatedSubscription{
		Id:        userSubscriptionToUpdate.ID,
		Status:    string(userSubscriptionToUpdate.Status),
		StartDate: userSubscriptionToUpdate.StartDate,
		EndDate:   userSubscriptionToUpdate.EndDate,
	}

	return &activatedSubscription, nil
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
	switch {
	case transaction.PaymentMethodUPIInPaymentsEntity != nil:
		return "UPI"
	case transaction.PaymentMethodCardInPaymentsEntity != nil:
		return "Card"
	case transaction.PaymentMethodNetBankingInPaymentsEntity != nil:
		return "Netbanking"
	default:
		return "Unknown"
	}
}

func (s *SubscriptionService) UserHasActiveSubscription(ctx context.Context, subscription *ent.Subscription, user *ent.User) (bool, error) {
	userSubscriptions, err := s.userSubscriptionRepository.GetByUserId(ctx, user.ID.String())
	if err != nil {
		return false, err
	}

	now := time.Now()

	for _, userSubscription := range userSubscriptions {
		if userSubscription.StartDate.Before(now) && userSubscription.EndDate.After(now) && userSubscription.Edges.Subscription.ID == subscription.ID {
			return true, nil
		}
	}

	return false, nil
}
