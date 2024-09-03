package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CachedExam holds the schema definition for the CachedExam entity.
type CachedExam struct {
	ent.Schema
}

// Fields of the CachedExam.
func (CachedExam) Fields() []ent.Field {
	return []ent.Field{
		field.String("cache_uid").Unique(),
		field.Bool("is_used").Default(false),
		field.Time("expires_at"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CachedExam.
func (CachedExam) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("exam", Exam.Type).
			Ref("cached_exam").
			Unique().
			Required(),
	}
}
