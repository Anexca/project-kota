package models

import "time"

type UserSubscriptionDetails struct {
	SubscriptionID         int                        `json:"subscription_id"`          // Subscription defined in the DB
	ProviderSubscriptionID string                     `json:"provider_subscription_id"` // Subscription purchased by user
	ProviderPlanID         string                     `json:"provider_plan_id"`         // Provider plan ID
	DurationInMonths       int                        `json:"duration_in_months"`       // Subscription duration in months
	StartDate              time.Time                  `json:"start_date"`               // Subscription start date
	EndDate                time.Time                  `json:"end_date"`                 // Subscription end date
	PaymentDetails         SubscriptionPaymentDetails `json:"payment_details"`          // Payment details
}

type SubscriptionPaymentDetails struct {
	Amount        float64   `json:"amount"`         // Payment amount
	PaymentDate   time.Time `json:"payment_date"`   // Payment date
	PaymentStatus string    `json:"payment_status"` // Status of the payment
	PaymentMethod string    `json:"payment_method"` // Payment method (e.g., card, upi)
}

type SubscriptionOverview struct {
	Id                  int                    `json:"id"`
	ProviderPlanID      string                 `json:"provider_plan_id"`
	BasePrice           float64                `json:"base_price"`
	FinalPrice          float64                `json:"final_price"`
	DurationInMonths    int                    `json:"duration_in_months"`
	IsActive            bool                   `json:"is_active"`
	Name                string                 `json:"name"`
	RawSubscriptionData map[string]interface{} `json:"raw_subscription_data,omitempty"`
}

type SubscriptionToActivate struct {
	Id               int    `json:"id"`
	Status           string `json:"status"`
	SubscriptionId   string `json:"subscription_id"`
	PaymentSessionId string `json:"payment_session_id"`
}

type ActivatedSubscription struct {
	Id        int       `json:"id"`
	Status    string    `json:"status"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
