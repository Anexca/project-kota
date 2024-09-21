package services

import cashfree_pg "github.com/cashfree/cashfree-pg/v4"

type PaymentService struct {

}

func NewPaymentService() *PaymentService{
	cashfree_pg.XClientId
	return &PaymentService{}
}

func (p *PaymentService) New