package services

import (
	"server/pkg/config"

	"github.com/razorpay/razorpay-go"
	utils "github.com/razorpay/razorpay-go/utils"
)

type PaymentServiceV2 struct {
	paymentClient *razorpay.Client
	environment   *config.Environment
}

type CreateSubscriptionModelV2 struct {
	ProviderPlanId     string
	TotalBillingCycles int
	CustomerId         string
}

type UpsertPaymentProviderCustomerModelV2 struct {
	Name  string
	Email string
	Phone string
}

func NewPaymentServiceV2(paymentClient *razorpay.Client) *PaymentServiceV2 {
	environment, _ := config.LoadEnvironment()

	return &PaymentServiceV2{
		environment:   environment,
		paymentClient: paymentClient,
	}
}

func (p *PaymentServiceV2) CreateCustomer(model UpsertPaymentProviderCustomerModelV2) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"name":          model.Name,
		"contact":       model.Phone,
		"email":         model.Email,
		"fail_existing": 0,
	}

	return p.paymentClient.Customer.Create(data, nil)
}

func (p *PaymentServiceV2) UpdateCustomer(customerID string, model UpsertPaymentProviderCustomerModelV2) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"name":    model.Name,
		"contact": model.Phone,
		"email":   model.Email,
	}

	return p.paymentClient.Customer.Edit(customerID, data, nil)
}

func (p *PaymentServiceV2) CreateSubscription(model CreateSubscriptionModelV2) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"plan_id":         model.ProviderPlanId,
		"total_count":     model.TotalBillingCycles,
		"customer_notify": 1,
		"customer_id":     model.CustomerId,
	}

	return p.paymentClient.Subscription.Create(data, nil)
}

func (p *PaymentServiceV2) CancelUserSubscription(subscriptionId string) (map[string]interface{}, error) {
	return p.paymentClient.Subscription.Cancel(subscriptionId, nil, nil)
}

func (p *PaymentServiceV2) GetPayment(paymentId string) (map[string]interface{}, error) {
	return p.paymentClient.Payment.Fetch(paymentId, nil, nil)
}

func (p *PaymentServiceV2) IsSubscriptionPaymentSignatureValid(paymentId, subscriptionId, signatureToVerify string) bool {
	params := map[string]interface{}{
		"razorpay_subscription_id": subscriptionId,
		"razorpay_payment_id":      paymentId,
	}

	signature := signatureToVerify
	secret := p.environment.RazorpaySecret
	return utils.VerifySubscriptionSignature(params, signature, secret)
}
