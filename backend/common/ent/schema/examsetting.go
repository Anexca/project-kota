package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ExamSetting holds the schema definition for the ExamSetting entity.
type ExamSetting struct {
	ent.Schema
}

// Fields of the ExamSetting.
func (ExamSetting) Fields() []ent.Field {
	return []ent.Field{
		field.Int("number_of_questions"),
		field.Int("duration_seconds"),
		field.Float("negative_marking").Optional(),
		field.String("ai_prompt").Optional(),
		field.JSON("other_details", map[string]interface{}{}).
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "json",
			}),
		field.Int("max_attempts").Default(2),
		field.String("evaluation_ai_prompt").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ExamSetting.
func (ExamSetting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("exam", Exam.Type).
			Ref("setting").
			Unique(), // Each ExamSetting belongs to one Exam
	}
}
