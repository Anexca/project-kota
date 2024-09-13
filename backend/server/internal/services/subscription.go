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
	"time"

	"github.com/razorpay/razorpay-go"
)

type SubscriptionService struct {
	environment                *config.Environment
	paymentService             *PaymentService
	userService                *UserService
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
	paymentService := NewPaymentService(paymentClient)
	userService := NewUserService(dbClient, paymentClient)
	subscriptionRepository := commonRepositories.NewSubscriptionRepository(dbClient)
	userSubscriptionRepository := commonRepositories.NewUserSubscriptioRepository(dbClient)
	paymentrepository := commonRepositories.NewPaymentRepository(dbClient)

	return &SubscriptionService{
		environment:                environment,
		paymentService:             paymentService,
		userService:                userService,
		subscriptionRepository:     subscriptionRepository,
		userSubscriptionRepository: userSubscriptionRepository,
		paymentrepository:          paymentrepository,
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
		return nil, fmt.Errorf("could not get id from object %v", createdSubscription)
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

	if !s.paymentService.IsSubscriptionPaymentSignatureValid(request.PaymentId, userSubscriptionToUpdate.ProviderSubscriptionID, request.Signature) {
		return nil, errors.New("payment verification failed")
	}

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
	paymentInfo, err := s.paymentService.GetPayment(providerPaymentId)
	if err != nil {
		return nil, err
	}

	amount, ok := paymentInfo["amount"].(float64)
	if !ok {
		return nil, fmt.Errorf("could not get amount from object %v", paymentInfo)
	}

	createdAt, ok := paymentInfo["created_at"].(float64)
	if !ok {
		return nil, fmt.Errorf("could not get created_at from object %v", paymentInfo)
	}

	method, ok := paymentInfo["method"].(string)
	if !ok {
		return nil, fmt.Errorf("could not get method from object %v", paymentInfo)
	}

	status, ok := paymentInfo["status"].(string)
	if !ok {
		return nil, fmt.Errorf("could not get status from object %v", paymentInfo)
	}

	invoiceId, ok := paymentInfo["invoice_id"].(string)
	if !ok {
		return nil, fmt.Errorf("could not get invoice_id from object %v", paymentInfo)
	}

	paymentModel := commonRepositories.CreatePaymentModel{
		Status:             status,
		PaymentMethod:      method,
		PaymentDate:        time.Unix(int64(createdAt), 0),
		Amount:             int(amount),
		UserSubscriptionId: userSubscriptionId,
		ProviderPaymentId:  providerPaymentId,
		ProviderInvoiceId:  invoiceId,
	}

	return s.paymentrepository.Create(ctx, paymentModel, userId)
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

func (s *SubscriptionService) UserHasActiveSubscription(ctx context.Context, subscription *ent.Subscription, user *ent.User) (bool, error) {
	userSubscription, err := s.userSubscriptionRepository.GetByUserId(ctx, subscription.ID, user.ID.String())
	if err != nil {
		return false, err
	}

	now := time.Now()
	return userSubscription.StartDate.Before(now) && userSubscription.EndDate.After(now), nil
}
