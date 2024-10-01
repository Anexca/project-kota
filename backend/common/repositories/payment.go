package repositories

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"

	"common/ent"
	"common/ent/payment"
	"common/ent/user"
)

// CreatePaymentModel represents the data needed to create a payment record.
type CreatePaymentModel struct {
	UserSubscriptionId int
	Amount             float64
	PaymentDate        time.Time
	Status             string
	PaymentMethod      string
	ProviderPaymentId  string
}

// PaymentRepository is a concrete implementation of PaymentRepositoryInterface.
type PaymentRepository struct {
	dbClient *ent.Client
}

// NewPaymentRepository creates a new instance of PaymentRepository.
func NewPaymentRepository(dbClient *ent.Client) *PaymentRepository {
	return &PaymentRepository{
		dbClient: dbClient,
	}
}

// GetByUserId retrieves all payments for a specific user by their ID.
func (p *PaymentRepository) GetByUserId(ctx context.Context, userId string) ([]*ent.Payment, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return p.dbClient.Payment.Query().
		Where(payment.HasUserWith(user.IDEQ(userUid))).
		WithSubscription().
		Order(ent.Desc(payment.FieldUpdatedAt)).
		All(ctx)
}

// Create adds a new payment record to the database.
func (p *PaymentRepository) Create(ctx context.Context, model CreatePaymentModel, userId string) (*ent.Payment, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return p.dbClient.Payment.Create().
		SetAmount(model.Amount).
		SetSubscriptionID(model.UserSubscriptionId).
		SetUserID(userUid).
		SetStatus(payment.Status(strings.ToUpper(model.Status))).
		SetPaymentMethod(model.PaymentMethod).
		SetProviderPaymentID(model.ProviderPaymentId).
		SetPaymentDate(model.PaymentDate).
		Save(ctx)
}

// GetByProviderPaymentId retrieves a payment by the provider's payment ID.
func (p *PaymentRepository) GetByProviderPaymentId(ctx context.Context, paymentProviderId string) (*ent.Payment, error) {
	return p.dbClient.Payment.Query().
		Where(payment.ProviderPaymentIDEQ(paymentProviderId)).
		Only(ctx)
}
