package services

import "github.com/razorpay/razorpay-go"

type PaymentService struct {
	paymentClient *razorpay.Client
}

type CreateSubscriptionModel struct {
	ProviderPlanId     string
	TotalBillingCycles int
}

type CreatePaymentProviderCustomerModel struct {
	Name  string
	Email string
	Phone string
}

func NewPaymentService(paymentClient *razorpay.Client) *PaymentService {
	return &PaymentService{
		paymentClient: paymentClient,
	}
}

func (p *PaymentService) CreateCustomer(model CreatePaymentProviderCustomerModel) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"name":          model.Name,
		"contact":       model.Phone,
		"email":         model.Email,
		"fail_existing": 0,
	}

	return p.paymentClient.Customer.Create(data, nil)
}

func (p *PaymentService) CreateSubscription(model CreateSubscriptionModel) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"plan_id":         model.ProviderPlanId,
		"total_count":     model.TotalBillingCycles,
		"customer_notify": 1,
	}

	return p.paymentClient.Subscription.Create(data, nil)
}
