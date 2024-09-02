package schema

import "entgo.io/ent"

// ExamResult holds the schema definition for the ExamResult entity.
type ExamResult struct {
	ent.Schema
}

// Fields of the ExamResult.
func (ExamResult) Fields() []ent.Field {
	return nil
}

// Edges of the ExamResult.
func (ExamResult) Edges() []ent.Edge {
	return nil
}
