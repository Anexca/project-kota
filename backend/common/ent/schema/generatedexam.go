package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GeneratedExam holds the schema definition for the GeneratedExam entity.
type GeneratedExam struct {
	ent.Schema
}

// Fields of the GeneratedExam.
func (GeneratedExam) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_active").Default(true),
		field.JSON("raw_exam_data", map[string]interface{}{}).
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "jsonb",
			}),
		field.Bool("is_open").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GeneratedExam.
func (GeneratedExam) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("exam", Exam.Type).
			Ref("generatedexams").
			Unique(), // Many GeneratedExams have one Exam

		edge.To("attempts", ExamAttempt.Type), // One GenratedExam can have Many Attempts
	}
}
