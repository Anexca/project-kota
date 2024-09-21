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

	"github.com/razorpay/razorpay-go"
)

type SubscriptionService struct {
	environment                *config.Environment
	paymentService             *PaymentService
	userRepository             *commonRepositories.UserRepository
	subscriptionRepository     *commonRepositories.SubscriptionRepository
	userSubscriptionRepository *commonRepositories.UserSubscriptioRepository
	paymentrepository          *commonRepositories.PaymentRepository
}

type ActivateUserSubscriptionRequest struct {
	PaymentId string `json:"payment_id" validate:"required"`
	Signature string `json:"signature" validate:"required"`
}

func NewSubscriptionService(dbClient *ent.Client, paymentClient *razorpay.Client) *SubscriptionService {
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

func (s *SubscriptionService) ActivateUserSubscription(ctx context.Context, request ActivateUserSubscriptionRequest, userSubscriptionId int, userId string) (*ent.UserSubscription, error) {
	userSubscriptionToUpdate, err := s.userSubscriptionRepository.GetById(ctx, userSubscriptionId, userId)
	if err != nil {
		return nil, err
	}

	existingPayment, err := s.paymentrepository.GetByProviderPaymentId(ctx, request.PaymentId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if !errors.As(err, &notFoundError) {
			return nil, err
		}
	}

	if existingPayment != nil {
		return nil, fmt.Errorf("record of payment with %s id already exists", request.PaymentId)
	}

	// if !s.paymentService.IsSubscriptionPaymentSignatureValid(request.PaymentId, userSubscriptionToUpdate.ProviderSubscriptionID, request.Signature) {
	// 	return nil, errors.New("payment verification failed")
	// }

	userSubscriptionToUpdate.Status = usersubscription.StatusACTIVE
	userSubscriptionToUpdate.StartDate = time.Now()
	userSubscriptionToUpdate.IsActive = true
	userSubscriptionToUpdate.EndDate = time.Now().AddDate(0, userSubscriptionToUpdate.Edges.Subscription.DurationInMonths, 0)

	go func() {
		bgCtx := context.Background()
		payment, err := s.StorePaymentForSubscription(bgCtx, request.PaymentId, userSubscriptionId, userId)
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

func (s *SubscriptionService) StorePaymentForSubscription(ctx context.Context, providerPaymentId string, userSubscriptionId int, userId string) (*ent.Payment, error) {
	// paymentInfo, err := s.paymentService.GetPayment(providerPaymentId)
	// if err != nil {
	// 	return nil, err
	// }

	// amount, ok := paymentInfo["amount"].(float64)
	// if !ok {
	// 	return nil, fmt.Errorf("could not get amount from object %v", paymentInfo)
	// }

	// createdAt, ok := paymentInfo["created_at"].(float64)
	// if !ok {
	// 	return nil, fmt.Errorf("could not get created_at from object %v", paymentInfo)
	// }

	// method, ok := paymentInfo["method"].(string)
	// if !ok {
	// 	return nil, fmt.Errorf("could not get method from object %v", paymentInfo)
	// }

	// status, ok := paymentInfo["status"].(string)
	// if !ok {
	// 	return nil, fmt.Errorf("could not get status from object %v", paymentInfo)
	// }

	// invoiceId, ok := paymentInfo["invoice_id"].(string)
	// if !ok {
	// 	return nil, fmt.Errorf("could not get invoice_id from object %v", paymentInfo)
	// }

	// paymentModel := commonRepositories.CreatePaymentModel{
	// 	Status:             status,
	// 	PaymentMethod:      method,
	// 	PaymentDate:        time.Unix(int64(createdAt), 0),
	// 	Amount:             int(amount),
	// 	UserSubscriptionId: userSubscriptionId,
	// 	ProviderPaymentId:  providerPaymentId,
	// 	ProviderInvoiceId:  invoiceId,
	// }

	// return s.paymentrepository.Create(ctx, paymentModel, userId)
	return nil, nil
}

func (s *SubscriptionService) CancelUserSubscription(ctx context.Context, userSubscriptionId int, userId string) (*ent.UserSubscription, error) {
	userSubscriptionToCancel, err := s.userSubscriptionRepository.GetById(ctx, userSubscriptionId, userId)
	if err != nil {
		return nil, err
	}

	if userSubscriptionToCancel.Status == usersubscription.StatusCANCELED {
		return nil, errors.New("user subscrption is already cancelled")
	}

	// _, err = s.paymentService.CancelUserSubscription(userSubscriptionToCancel.ProviderSubscriptionID)
	// if err != nil {
	// 	return nil, fmt.Errorf("error canceling subscription with payment provider: %v", err)
	// }

	userSubscriptionToCancel.Status = usersubscription.StatusCANCELED

	err = s.userSubscriptionRepository.Update(ctx, userSubscriptionToCancel)
	if err != nil {
		return nil, err
	}

	return userSubscriptionToCancel, nil
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
