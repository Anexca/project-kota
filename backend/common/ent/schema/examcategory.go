package schema

import (
	"common/constants"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ExamCategory holds the schema definition for the ExamCategory entity.
type ExamCategory struct {
	ent.Schema
}

// Fields of the ExamCategory.
func (ExamCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("name").Values(
			string(constants.ExamCategoryNameBanking),
		),
		field.String("description"),
		field.Bool("is_active").Default(true),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ExamCategory.
func (ExamCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("exams", Exam.Type), // One ExamCategory can have many Exams

		edge.To("groups", ExamGroup.Type), // One ExamCategory can have many Exams
	}
}
