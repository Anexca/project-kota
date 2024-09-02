package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Exam holds the schema definition for the Exam entity.
type Exam struct {
	ent.Schema
}

// Fields of the Exam.
func (Exam) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.Bool("is_active").Default(true),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Exam.
func (Exam) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", ExamCategory.Type).
			Ref("exams").
			Unique(), // Many Exams belong to one ExamCategory

		edge.To("setting", ExamSetting.Type).
			Unique(), // Each Exam has one ExamSetting

		edge.To("attempts", ExamAttempt.Type), // One Exam can have Many Attempts

		edge.To("cached_question_metadata", CachedQuestionMetaData.Type),

		edge.To("generatedexams", GeneratedExam.Type), // One Exam can have many Questions
	}
}
