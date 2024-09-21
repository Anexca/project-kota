package services

import (
	"common/util"
	"log"
	"server/pkg/config"

	cashfree_pg "github.com/cashfree/cashfree-pg/v4"
)

type PaymentService struct {
	environment *config.Environment
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
}

var xAPIVersion = "2023-08-01"

func NewPaymentService() *PaymentService {
	environment, err := config.LoadEnvironment()
	if err != nil {
		log.Fatalf("error initiating payment service, %v", err)
	}

	cashfree_pg.XClientId = &environment.PaymentProviderKey
	cashfree_pg.XClientSecret = &environment.PaymentProviderSecret
	cashfree_pg.XEnvironment = cashfree_pg.SANDBOX

	if environment.IsProduction {
		cashfree_pg.XEnvironment = cashfree_pg.PRODUCTION
	}

	return &PaymentService{
		environment: environment,
	}
}

func (p *PaymentService) CreateCustomer(model UpsertPaymentProviderCustomerModel) (*cashfree_pg.CustomerEntity, error) {

	createCustomerRequest := cashfree_pg.CreateCustomerRequest{
		CustomerPhone: model.Phone,
		CustomerEmail: &model.Email,
		CustomerName:  &model.Name,
	}

	xRequestId := util.GenerateUUID()
	xIdempotencyKey := util.GenerateUUID()

	resp, _, err := cashfree_pg.PGCreateCustomer(&xAPIVersion, &createCustomerRequest, &xRequestId, &xIdempotencyKey, nil)
	return resp, err
}

func (p *PaymentService) CreateOrder(model CreateOrderModel) (*cashfree_pg.OrderEntity, error) {
	request := cashfree_pg.CreateOrderRequest{
		OrderAmount: model.Amount,
		CustomerDetails: cashfree_pg.CustomerDetails{
			// CustomerUid:   &model.CustomerId,
			CustomerPhone: model.CustomerPhoneNumber,
			CustomerId:    model.UserId,
			CustomerEmail: &model.CustomerEmail,
			CustomerName:  &model.CustomerName,
		},
		OrderCurrency: "INR",
	}

	response, _, err := cashfree_pg.PGCreateOrder(&xAPIVersion, &request, nil, nil, nil)
	return response, err
}

func (p *PaymentService) IsOrderSuccessful(orderId string) (bool, error) {

	response, _, err := cashfree_pg.PGOrderFetchPayments(&xAPIVersion, orderId, nil, nil, nil)
	if err != nil {
		return false, err
	}

	for _, transaction := range response {
		if *transaction.PaymentStatus == "SUCCESS" {
			return true, nil
		}
	}

	return false, nil
}
