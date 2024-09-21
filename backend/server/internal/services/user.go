package services

import (
	"common/ent"
	"common/repositories"
	"context"
	"fmt"
	"server/pkg/models"
	"time"

	"github.com/google/uuid"
	"github.com/razorpay/razorpay-go"
)

type UserService struct {
	paymentService             *PaymentService
	userRepository             *repositories.UserRepository
	paymentRepositry           *repositories.PaymentRepository
	userSubscriptionRepository *repositories.UserSubscriptioRepository
}

type UpdateUserRequest struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone,omitempty"`
}

type UserProfileResponse struct {
	Id                  uuid.UUID                        `json:"id"`
	Email               string                           `json:"email"`
	FirstName           string                           `json:"first_name"`
	LastName            string                           `json:"last_name"`
	PhoneNumber         string                           `json:"phone_number"`
	ActiveSubscriptions []models.UserSubscriptionDetails `json:"active_subscriptions"`
}

func NewUserService(dbClient *ent.Client, paymentClient *razorpay.Client) *UserService {
	paymentService := NewPaymentService()
	paymentRepositry := repositories.NewPaymentRepository(dbClient)
	userRepository := repositories.NewUserRepository(dbClient)
	userSubscriptionRepository := repositories.NewUserSubscriptioRepository(dbClient)

	return &UserService{
		paymentService:             paymentService,
		userRepository:             userRepository,
		paymentRepositry:           paymentRepositry,
		userSubscriptionRepository: userSubscriptionRepository,
	}
}

func (u *UserService) GetUserProfile(ctx context.Context, userId string) (UserProfileResponse, error) {
	user, err := u.userRepository.Get(ctx, userId)
	if err != nil {
		return UserProfileResponse{}, err
	}

	userSubscriptions, err := u.userSubscriptionRepository.GetByUserId(ctx, userId)
	if err != nil {
		return UserProfileResponse{}, err
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
				PaymentDetails: models.SubscriptionPaymentDetails{
					Amount:        userSubscription.Edges.Payments[0].Amount, // Assuming the first payment holds the necessary details
					PaymentDate:   userSubscription.Edges.Payments[0].PaymentDate,
					PaymentStatus: string(userSubscription.Edges.Payments[0].Status),
					PaymentMethod: userSubscription.Edges.Payments[0].PaymentMethod,
				},
			}

			activeSubscriptions = append(activeSubscriptions, subscriptionDetails)
		}
	}

	// Step 4: Create the response model
	responseModel := UserProfileResponse{
		Id:                  user.ID,
		Email:               user.Email,
		FirstName:           user.FirstName,
		LastName:            user.LastName,
		PhoneNumber:         user.PhoneNumber,
		ActiveSubscriptions: activeSubscriptions,
	}
	return responseModel, nil
}

func (u *UserService) UpdateUser(ctx context.Context, userId string, request UpdateUserRequest) (*ent.User, error) {
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

	payments, err := u.paymentRepositry.GetByUserId(ctx, userId)
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
	model := UpsertPaymentProviderCustomerModel{
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
