package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"common/constants"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.Float("amount"),
		field.Time("payment_date"),
		field.Enum("status").
			Values(
				string(constants.PaymentStatusSuccess),
				string(constants.PaymentStatusNotAttempted),
				string(constants.PaymentStatusUserDropped),
				string(constants.PaymentStatusFailed),
				string(constants.PaymentStatusPending),
				string(constants.PaymentStatusCancelled),
			),
		field.String("payment_method"),
		field.String("provider_payment_id").Unique(),
		field.String("provider_invoice_id").Unique().Optional(),
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
