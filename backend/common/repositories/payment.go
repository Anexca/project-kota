package repositories

import (
	"common/ent"
	"common/ent/payment"
	"context"
	"time"

	"github.com/google/uuid"
)

type PaymentRepository struct {
	dbClient *ent.Client
}

type CreatePaymentModel struct {
	UserSubscriptionId int
	Amount             int
	PaymentDate        time.Time
	Status             string
	PaymentMethod      string
	ProviderPaymentId  string
}

func NewPaymentRepository(dbClient *ent.Client) *PaymentRepository {
	return &PaymentRepository{
		dbClient: dbClient,
	}
}

func (p *PaymentRepository) Create(ctx context.Context, model CreatePaymentModel, userId string) (*ent.Payment, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return p.dbClient.Payment.Create().
		SetAmount(model.Amount).
		SetSubscriptionID(model.UserSubscriptionId).
		SetUserID(userUid).
		SetStatus(payment.Status(model.Status)).
		SetPaymentMethod(model.PaymentMethod).
		SetProviderPaymentID(model.ProviderPaymentId).
		SetPaymentDate(model.PaymentDate).
		Save(ctx)
}
