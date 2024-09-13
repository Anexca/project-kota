package services

import (
	"server/pkg/config"

	"github.com/razorpay/razorpay-go"
	utils "github.com/razorpay/razorpay-go/utils"
)

type PaymentService struct {
	paymentClient *razorpay.Client
	environment   *config.Environment
}

type CreateSubscriptionModel struct {
	ProviderPlanId     string
	TotalBillingCycles int
	CustomerId         string
}

type UpsertPaymentProviderCustomerModel struct {
	Name  string
	Email string
	Phone string
}

func NewPaymentService(paymentClient *razorpay.Client) *PaymentService {
	environment, _ := config.LoadEnvironment()

	return &PaymentService{
		environment:   environment,
		paymentClient: paymentClient,
	}
}

func (p *PaymentService) CreateCustomer(model UpsertPaymentProviderCustomerModel) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"name":          model.Name,
		"contact":       model.Phone,
		"email":         model.Email,
		"fail_existing": 0,
	}

	return p.paymentClient.Customer.Create(data, nil)
}

func (p *PaymentService) UpdateCustomer(customerID string, model UpsertPaymentProviderCustomerModel) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"name":    model.Name,
		"contact": model.Phone,
		"email":   model.Email,
	}

	return p.paymentClient.Customer.Edit(customerID, data, nil)
}

func (p *PaymentService) CreateSubscription(model CreateSubscriptionModel) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"plan_id":         model.ProviderPlanId,
		"total_count":     model.TotalBillingCycles,
		"customer_notify": 1,
		"customer_id":     model.CustomerId,
	}

	return p.paymentClient.Subscription.Create(data, nil)
}

func (p *PaymentService) CancelUserSubscription(subscriptionId string) (map[string]interface{}, error) {
	return p.paymentClient.Subscription.Cancel(subscriptionId, nil, nil)
}

func (p *PaymentService) GetPayment(paymentId string) (map[string]interface{}, error) {
	return p.paymentClient.Payment.Fetch(paymentId, nil, nil)
}

func (p *PaymentService) IsSubscriptionPaymentSignatureValid(paymentId, subscriptionId, signatureToVerify string) bool {
	params := map[string]interface{}{
		"razorpay_subscription_id": subscriptionId,
		"razorpay_payment_id":      paymentId,
	}

	signature := signatureToVerify
	secret := p.environment.RazorpaySecret
	return utils.VerifySubscriptionSignature(params, signature, secret)
}
