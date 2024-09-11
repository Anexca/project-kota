package services

import (
	"common/ent"
	"common/repositories"
	"context"
	"fmt"

	"github.com/razorpay/razorpay-go"
)

type UserService struct {
	paymentService *PaymentService
	userRepository *repositories.UserRepository
}

type UpdateUserRequest struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone,omitempty"`
}

func NewUserService(dbClient *ent.Client, paymentClient *razorpay.Client) *UserService {
	paymentService := NewPaymentService(paymentClient)
	userRepository := repositories.NewUserRepository(dbClient)

	return &UserService{
		paymentService: paymentService,
		userRepository: userRepository,
	}
}

func (u *UserService) GetUser(ctx context.Context, userId string) (*ent.User, error) {
	return u.userRepository.Get(ctx, userId)
}

func (u *UserService) UpdateUser(ctx context.Context, userId string, request UpdateUserRequest) (*ent.User, error) {
	userToUpdate, err := u.GetUser(ctx, userId)
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

func (u *UserService) updatePaymentProviderCustomer(user *ent.User) error {
	model := UpsertPaymentProviderCustomerModel{
		Name:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		Phone: user.PhoneNumber,
		Email: user.Email,
	}

	if user.PaymentProviderCustomerID != "" {
		_, err := u.paymentService.UpdateCustomer(user.PaymentProviderCustomerID, model)
		return err
	}

	customer, err := u.paymentService.CreateCustomer(model)
	if err != nil {
		return err
	}

	customerId, ok := customer["id"].(string)
	if !ok || customerId == "" {
		return fmt.Errorf("could not extract customer id from response %v", customer)
	}

	user.PaymentProviderCustomerID = customerId

	return nil
}
