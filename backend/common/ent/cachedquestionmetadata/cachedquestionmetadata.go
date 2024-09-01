// Code generated by ent, DO NOT EDIT.

package cachedquestionmetadata

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the cachedquestionmetadata type in the database.
	Label = "cached_question_meta_data"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCacheUID holds the string denoting the cache_uid field in the database.
	FieldCacheUID = "cache_uid"
	// FieldIsUsed holds the string denoting the is_used field in the database.
	FieldIsUsed = "is_used"
	// FieldExpiresAt holds the string denoting the expires_at field in the database.
	FieldExpiresAt = "expires_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeExam holds the string denoting the exam edge name in mutations.
	EdgeExam = "exam"
	// Table holds the table name of the cachedquestionmetadata in the database.
	Table = "cached_question_meta_data"
	// ExamTable is the table that holds the exam relation/edge.
	ExamTable = "cached_question_meta_data"
	// ExamInverseTable is the table name for the Exam entity.
	// It exists in this package in order to avoid circular dependency with the "exam" package.
	ExamInverseTable = "exams"
	// ExamColumn is the table column denoting the exam relation/edge.
	ExamColumn = "exam_cached_question_metadata"
)

// Columns holds all SQL columns for cachedquestionmetadata fields.
var Columns = []string{
	FieldID,
	FieldCacheUID,
	FieldIsUsed,
	FieldExpiresAt,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "cached_question_meta_data"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"exam_cached_question_metadata",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsUsed holds the default value on creation for the "is_used" field.
	DefaultIsUsed bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the CachedQuestionMetaData queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCacheUID orders the results by the cache_uid field.
func ByCacheUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCacheUID, opts...).ToFunc()
}

// ByIsUsed orders the results by the is_used field.
func ByIsUsed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsUsed, opts...).ToFunc()
}

// ByExpiresAt orders the results by the expires_at field.
func ByExpiresAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpiresAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByExamField orders the results by exam field.
func ByExamField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExamStep(), sql.OrderByField(field, opts...))
	}
}
func newExamStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExamInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ExamTable, ExamColumn),
	)
}
