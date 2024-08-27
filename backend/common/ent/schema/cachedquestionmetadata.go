package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CachedQuestionMetadata holds the schema definition for the CachedQuestionMetadata entity.
type CachedQuestionMetadata struct {
	ent.Schema
}

// Fields of the CachedQuestionMetadata.
func (CachedQuestionMetadata) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Unique(),
		field.String("type"),
		field.String("subject"),
		field.String("exam"),
		field.Bool("is_processed").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the CachedQuestionMetadata.
func (CachedQuestionMetadata) Edges() []ent.Edge {
	return nil
}
