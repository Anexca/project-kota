package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SubscriptionExam holds the schema definition for the SubscriptionExam entity.
type SubscriptionExam struct {
	ent.Schema
}

// Fields of the SubscriptionExam.
func (SubscriptionExam) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the SubscriptionExam.
func (SubscriptionExam) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("subscription", Subscription.Type).Ref("exams").Unique(), // Many exams belong to one Subscription
		edge.From("exam", Exam.Type).Ref("subscriptions").Unique(),         // One Exam can be linked to many Subscriptions
	}
}
