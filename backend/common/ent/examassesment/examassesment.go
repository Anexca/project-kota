// Code generated by ent, DO NOT EDIT.

package examassesment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the examassesment type in the database.
	Label = "exam_assesment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCompletedSeconds holds the string denoting the completed_seconds field in the database.
	FieldCompletedSeconds = "completed_seconds"
	// FieldRawAssesmentData holds the string denoting the raw_assesment_data field in the database.
	FieldRawAssesmentData = "raw_assesment_data"
	// FieldIsReady holds the string denoting the is_ready field in the database.
	FieldIsReady = "is_ready"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeAttempt holds the string denoting the attempt edge name in mutations.
	EdgeAttempt = "attempt"
	// Table holds the table name of the examassesment in the database.
	Table = "exam_assesments"
	// AttemptTable is the table that holds the attempt relation/edge.
	AttemptTable = "exam_assesments"
	// AttemptInverseTable is the table name for the ExamAttempt entity.
	// It exists in this package in order to avoid circular dependency with the "examattempt" package.
	AttemptInverseTable = "exam_attempts"
	// AttemptColumn is the table column denoting the attempt relation/edge.
	AttemptColumn = "exam_attempt_assesment"
)

// Columns holds all SQL columns for examassesment fields.
var Columns = []string{
	FieldID,
	FieldCompletedSeconds,
	FieldRawAssesmentData,
	FieldIsReady,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "exam_assesments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"exam_attempt_assesment",
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
	// DefaultIsReady holds the default value on creation for the "is_ready" field.
	DefaultIsReady bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the ExamAssesment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCompletedSeconds orders the results by the completed_seconds field.
func ByCompletedSeconds(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompletedSeconds, opts...).ToFunc()
}

// ByIsReady orders the results by the is_ready field.
func ByIsReady(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsReady, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByAttemptField orders the results by attempt field.
func ByAttemptField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAttemptStep(), sql.OrderByField(field, opts...))
	}
}
func newAttemptStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AttemptInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, AttemptTable, AttemptColumn),
	)
}
