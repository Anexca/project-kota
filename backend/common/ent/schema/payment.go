package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// PaymentStatus represents the various statuses a payment can have.
type PaymentStatus string

const (
	PaymentStatusCreated           PaymentStatus = "CREATED"
	PaymentStatusAuthorized        PaymentStatus = "AUTHORIZED"
	PaymentStatusCaptured          PaymentStatus = "CAPTURED"
	PaymentStatusFailed            PaymentStatus = "FAILED"
	PaymentStatusRefunded          PaymentStatus = "REFUNDED"
	PaymentStatusPartiallyRefunded PaymentStatus = "PARTIALLY_REFUNDED"
	PaymentStatusPending           PaymentStatus = "PENDING"
	PaymentStatusProcessing        PaymentStatus = "PROCESSING"
	PaymentStatusCancelled         PaymentStatus = "CANCELLED"
	PaymentStatusDisputed          PaymentStatus = "DISPUTED"
)

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount"),
		field.Time("payment_date"),
		field.Enum("status").
			Values(
				string(PaymentStatusCreated),
				string(PaymentStatusAuthorized),
				string(PaymentStatusCaptured),
				string(PaymentStatusFailed),
				string(PaymentStatusRefunded),
				string(PaymentStatusPartiallyRefunded),
				string(PaymentStatusPending),
				string(PaymentStatusProcessing),
				string(PaymentStatusCancelled),
				string(PaymentStatusDisputed),
			),
		field.String("payment_method"),
		field.String("provider_payment_id").Unique(),
		field.String("receipt_id").Unique(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("payments").Unique(),                     // One User can have many Payments
		edge.From("subscription", UserSubscription.Type).Ref("payments").Unique(), // One Payment is linked to one UserSubscription
	}
}
