package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserSubscription holds the schema definition for the UserSubscription entity.
type UserSubscription struct {
	ent.Schema
}

type UserSubscriptionStatus string

const (
	UserSubscriptionStatusActive   UserSubscriptionStatus = "ACTIVE"
	UserSubscriptionStatusCanceled UserSubscriptionStatus = "CANCELED"
	UserSubscriptionStatusExpired  UserSubscriptionStatus = "EXPIRED"
	UserSubscriptionStatusPending  UserSubscriptionStatus = "PENDING"
	UserSubscriptionStatusPaused   UserSubscriptionStatus = "PAUSED"
)

// Fields of the UserSubscription.
func (UserSubscription) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_active"),
		field.Enum("status"). // Define the status field as an enum
					Values(
				string(UserSubscriptionStatusActive),
				string(UserSubscriptionStatusCanceled),
				string(UserSubscriptionStatusExpired),
				string(UserSubscriptionStatusPending),
				string(UserSubscriptionStatusPaused),
			).Default(string(UserSubscriptionStatusPending)),
		field.Time("start_date").Optional(),
		field.Time("end_date").Optional(),
		field.String("provider_subscription_id").Unique(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserSubscription.
func (UserSubscription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("subscriptions").Unique(),                      // One User can have many UserSubscriptions
		edge.From("subscription", Subscription.Type).Ref("user_subscriptions").Unique(), // One Subscription can be linked to many UserSubscriptions
		edge.To("payments", Payment.Type),                                               // One UserSubscription can have multiple payments
	}
}
