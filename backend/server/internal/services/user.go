package services

import (
	"context"
	"fmt"
	"time"

	"common/ent"
	"common/repositories"

	commonInterfaces "common/interfaces"
	"server/internal/interfaces"
	"server/pkg/models"
)

type UserService struct {
	paymentService             interfaces.PaymentServiceInterface
	userRepository             commonInterfaces.UserRepositoryInterface
	paymentRepository          commonInterfaces.PaymentRepositoryInterface
	userSubscriptionRepository commonInterfaces.UserSubscriptionRepositoryInterface
}

// NewUserService creates a new UserService with the provided dependencies.
func NewUserService(paymentService interfaces.PaymentServiceInterface, userRepo commonInterfaces.UserRepositoryInterface, paymentRepo commonInterfaces.PaymentRepositoryInterface, userSubscriptionRepo commonInterfaces.UserSubscriptionRepositoryInterface) *UserService {
	return &UserService{
		paymentService:             paymentService,
		userRepository:             userRepo,
		paymentRepository:          paymentRepo,
		userSubscriptionRepository: userSubscriptionRepo,
	}
}

// InitUserService initializes the UserService for production use.
func InitUserService(dbClient *ent.Client) *UserService {
	paymentService := NewPaymentService() // Assume NewPaymentService creates the service
	userRepository := repositories.NewUserRepository(dbClient)
	paymentRepository := repositories.NewPaymentRepository(dbClient)
	userSubscriptionRepository := repositories.NewUserSubscriptionRepository(dbClient)

	return NewUserService(paymentService, userRepository, paymentRepository, userSubscriptionRepository)
}

func (u *UserService) GetUserProfile(ctx context.Context, userId string) (models.UserProfileResponse, error) {
	user, err := u.userRepository.Get(ctx, userId)
	if err != nil {
		return models.UserProfileResponse{}, err
	}

	userSubscriptions, err := u.userSubscriptionRepository.GetByUserId(ctx, userId)
	if err != nil {
		return models.UserProfileResponse{}, err
	}

	var activeSubscriptions []models.UserSubscriptionDetails
	now := time.Now()
	for _, userSubscription := range userSubscriptions {
		if userSubscription.StartDate.Before(now) && userSubscription.EndDate.After(now) {
			subscriptionDetails := models.UserSubscriptionDetails{
				SubscriptionID:         userSubscription.Edges.Subscription.ID,
				ProviderPlanID:         userSubscription.Edges.Subscription.ProviderPlanID,
				ProviderSubscriptionID: userSubscription.ProviderSubscriptionID,
				DurationInMonths:       userSubscription.Edges.Subscription.DurationInMonths,
				StartDate:              userSubscription.StartDate,
				EndDate:                userSubscription.EndDate,
			}

			if len(userSubscription.Edges.Payments) > 0 {
				paymentDetails := models.SubscriptionPaymentDetails{
					Amount:        userSubscription.Edges.Payments[0].Amount, // Assuming the first payment holds the necessary details
					PaymentDate:   userSubscription.Edges.Payments[0].PaymentDate,
					PaymentStatus: string(userSubscription.Edges.Payments[0].Status),
					PaymentMethod: userSubscription.Edges.Payments[0].PaymentMethod,
				}

				subscriptionDetails.PaymentDetails = paymentDetails
			}

			activeSubscriptions = append(activeSubscriptions, subscriptionDetails)
		}
	}

	// Step 4: Create the response model
	responseModel := models.UserProfileResponse{
		Id:                  user.ID,
		Email:               user.Email,
		FirstName:           user.FirstName,
		LastName:            user.LastName,
		PhoneNumber:         user.PhoneNumber,
		ActiveSubscriptions: activeSubscriptions,
	}
	return responseModel, nil
}

func (u *UserService) UpdateUser(ctx context.Context, userId string, request models.UpdateUserRequest) (*ent.User, error) {
	userToUpdate, err := u.userRepository.Get(ctx, userId)
	if err != nil {
		return nil, err
	}

	if request.FirstName != "" {
		userToUpdate.FirstName = request.FirstName
	}

	if request.LastName != "" {
		userToUpdate.LastName = request.LastName
	}

	if request.PhoneNumber != "" {
		userToUpdate.PhoneNumber = request.PhoneNumber
	}

	if err := u.updatePaymentProviderCustomer(userToUpdate); err != nil {
		return nil, err
	}

	updatedUser, err := u.userRepository.Update(ctx, userToUpdate)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *UserService) GetUserTransactions(ctx context.Context, userId string) ([]models.SubscriptionPaymentDetails, error) {
	var subscriptionPayments []models.SubscriptionPaymentDetails

	payments, err := u.paymentRepository.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	for _, payment := range payments {
		subscriptionPayment := models.SubscriptionPaymentDetails{
			Amount:        payment.Amount,
			PaymentDate:   payment.PaymentDate,
			PaymentStatus: string(payment.Status),
			PaymentMethod: payment.PaymentMethod,
		}

		subscriptionPayments = append(subscriptionPayments, subscriptionPayment)
	}

	return subscriptionPayments, nil
}

func (u *UserService) updatePaymentProviderCustomer(user *ent.User) error {
	model := models.UpsertPaymentProviderCustomerModel{
		Name:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		Phone: user.PhoneNumber,
		Email: user.Email,
	}

	customer, err := u.paymentService.CreateCustomer(model)
	if err != nil {
		return err
	}

	customerId := customer.CustomerUid
	if *customerId == "" {
		return fmt.Errorf("could not extract customer id from response %v", customer)
	}

	user.PaymentProviderCustomerID = *customerId
	return nil
}
