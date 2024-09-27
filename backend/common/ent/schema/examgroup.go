package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ExamGroup holds the schema definition for the ExamGroup entity.
type ExamGroup struct {
	ent.Schema
}

// Fields of the ExamGroup.
func (ExamGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.Bool("is_active").Default(true),
		field.String("logo_url").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ExamGroup.
func (ExamGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", ExamCategory.Type).
			Ref("groups").
			Unique(), // Many Exams Group belong to one ExamCategory

		edge.To("exams", Exam.Type), // One ExamGroup can have many Exams

	}
}
