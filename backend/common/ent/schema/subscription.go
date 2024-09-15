package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.String("provider_plan_id"),
		field.Int("price"),
		field.Int("duration_in_months"),
		field.Bool("is_active"),
		field.String("name"),
		field.JSON("raw_subscription_data", map[string]interface{}{}).
			SchemaType(map[string]string{
				dialect.Postgres: "jsonb",
			}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("exams", SubscriptionExam.Type),              // One Subscription can have many SubscriptionExams
		edge.To("user_subscriptions", UserSubscription.Type), // One Subscription can be linked to many UserSubscriptions
	}
}
