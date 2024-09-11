package services

import "github.com/razorpay/razorpay-go"

type PaymentService struct {
	paymentClient *razorpay.Client
}

func NewPaymentService(paymentClient *razorpay.Client) *PaymentService {
	return &PaymentService{
		paymentClient: paymentClient,
	}
}

func (p *PaymentService) CreateOrder() (map[string]interface{}, error) {
	data := map[string]interface{}{
		"amount":          50000,
		"currency":        "INR",
		"receipt":         "some_receipt_id",
		"partial_payment": false,
		"notes": map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	return p.paymentClient.Order.Create(data, nil)
}
