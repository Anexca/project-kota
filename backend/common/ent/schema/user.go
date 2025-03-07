package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable(),
		field.String("email").
			Unique().
			NotEmpty(),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.String("phone_number").Optional(),
		field.String("payment_provider_customer_id").Optional().Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attempts", ExamAttempt.Type),           // One User can have Many Attempts
		edge.To("subscriptions", UserSubscription.Type), // One User can have many UserSubscriptions
		edge.To("payments", Payment.Type),               // One User can have many Payments
	}
}
