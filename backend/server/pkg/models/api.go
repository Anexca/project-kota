package models

import "github.com/google/uuid"

type DescriptiveExamAssesmentRequest struct {
	CompletedSeconds int    `json:"completed_seconds" validate:"gte=0"`
	Content          string `json:"content" validate:"required"`
}

type MCQExamAssessmentRequestModel struct {
	QuestionNumber          int   `json:"question_number" validate:"required"`
	UserSelectedOptionIndex []int `json:"user_selected_option_index" validate:"required"`
}

type MCQExamAssessmentRequest struct {
	AttemptedQuestions []MCQExamAssessmentRequestModel `json:"attempted_questions" validate:"required"`
	CompletedSeconds   int                             `json:"completed_seconds" validate:"gte=0"`
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
