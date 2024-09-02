package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ExamAttempt holds the schema definition for the ExamAttempt entity.
type ExamAttempt struct {
	ent.Schema
}

// Fields of the ExamAttempt.
func (ExamAttempt) Fields() []ent.Field {
	return []ent.Field{
		field.Int("attempt_number"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ExamAttempt.
func (ExamAttempt) Edges() []ent.Edge {
	return nil
}
