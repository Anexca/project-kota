package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CachedQuestionMetaData holds the schema definition for the CachedQuestionMetaData entity.
type CachedQuestionMetaData struct {
	ent.Schema
}

// Fields of the CachedQuestionMetaData.
func (CachedQuestionMetaData) Fields() []ent.Field {
	return []ent.Field{
		field.String("cache_uid").Unique(),
		field.Bool("is_used").Default(false),
		field.Time("expires_at"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CachedQuestionMetaData.
func (CachedQuestionMetaData) Edges() []ent.Edge {
	return nil
}
