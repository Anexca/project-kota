// Code generated by ent, DO NOT EDIT.

package examattempt

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the examattempt type in the database.
	Label = "exam_attempt"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAttemptNumber holds the string denoting the attempt_number field in the database.
	FieldAttemptNumber = "attempt_number"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeGeneratedexam holds the string denoting the generatedexam edge name in mutations.
	EdgeGeneratedexam = "generatedexam"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeAssesment holds the string denoting the assesment edge name in mutations.
	EdgeAssesment = "assesment"
	// Table holds the table name of the examattempt in the database.
	Table = "exam_attempts"
	// GeneratedexamTable is the table that holds the generatedexam relation/edge.
	GeneratedexamTable = "exam_attempts"
	// GeneratedexamInverseTable is the table name for the GeneratedExam entity.
	// It exists in this package in order to avoid circular dependency with the "generatedexam" package.
	GeneratedexamInverseTable = "generated_exams"
	// GeneratedexamColumn is the table column denoting the generatedexam relation/edge.
	GeneratedexamColumn = "generated_exam_attempts"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "exam_attempts"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_attempts"
	// AssesmentTable is the table that holds the assesment relation/edge.
	AssesmentTable = "exam_assesments"
	// AssesmentInverseTable is the table name for the ExamAssesment entity.
	// It exists in this package in order to avoid circular dependency with the "examassesment" package.
	AssesmentInverseTable = "exam_assesments"
	// AssesmentColumn is the table column denoting the assesment relation/edge.
	AssesmentColumn = "exam_attempt_assesment"
)

// Columns holds all SQL columns for examattempt fields.
var Columns = []string{
	FieldID,
	FieldAttemptNumber,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "exam_attempts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"generated_exam_attempts",
	"user_attempts",
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the ExamAttempt queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAttemptNumber orders the results by the attempt_number field.
func ByAttemptNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAttemptNumber, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByGeneratedexamField orders the results by generatedexam field.
func ByGeneratedexamField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGeneratedexamStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByAssesmentField orders the results by assesment field.
func ByAssesmentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAssesmentStep(), sql.OrderByField(field, opts...))
	}
}
func newGeneratedexamStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GeneratedexamInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, GeneratedexamTable, GeneratedexamColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newAssesmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AssesmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, AssesmentTable, AssesmentColumn),
	)
}
