package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ExamAssesment holds the schema definition for the ExamAssesment entity.
type ExamAssesment struct {
	ent.Schema
}

// Fields of the ExamAssesment.
func (ExamAssesment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("completed_seconds"),
		field.JSON("raw_assesment_data", map[string]interface{}{}).
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "jsonb",
			}),
		field.Bool("is_ready").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ExamAssesment.
func (ExamAssesment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("attempt", ExamAttempt.Type).
			Ref("assesment").
			Unique(), // Each Assesment belong to one attempt
	}
}
