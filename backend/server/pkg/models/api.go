package models

import "github.com/google/uuid"

type DescriptiveExamAssesmentRequest struct {
	CompletedSeconds int    `json:"completed_seconds" validate:"required"`
	Content          string `json:"content" validate:"required"`
}

type UpsertPaymentProviderCustomerModel struct {
	Name  string
	Email string
	Phone string
}

type CreateOrderModel struct {
	Amount              float64
	UserId              string
	CustomerId          string
	CustomerPhoneNumber string
	CustomerName        string
	CustomerEmail       string
	ReturnUrl           *string
}

type UpdateUserRequest struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone,omitempty"`
}

type UserProfileResponse struct {
	Id                  uuid.UUID                 `json:"id"`
	Email               string                    `json:"email"`
	FirstName           string                    `json:"first_name"`
	LastName            string                    `json:"last_name"`
	PhoneNumber         string                    `json:"phone_number"`
	ActiveSubscriptions []UserSubscriptionDetails `json:"active_subscriptions"`
}
