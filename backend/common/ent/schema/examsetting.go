package schema

import (
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
		field.Time("duration_minutes"),
		field.Float("negative_marking"),
		field.JSON("other_details", map[string]interface{}{}).
			SchemaType(map[string]string{
				dialect.Postgres: "json",
			}),
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
