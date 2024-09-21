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

	resp, _, err := cashfree_pg.PGCreateCustomer(nil, &createCustomerRequest, &xRequestId, &xIdempotencyKey, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
