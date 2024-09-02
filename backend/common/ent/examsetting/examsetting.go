// Code generated by ent, DO NOT EDIT.

package examsetting

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the examsetting type in the database.
	Label = "exam_setting"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNumberOfQuestions holds the string denoting the number_of_questions field in the database.
	FieldNumberOfQuestions = "number_of_questions"
	// FieldDurationMinutes holds the string denoting the duration_minutes field in the database.
	FieldDurationMinutes = "duration_minutes"
	// FieldNegativeMarking holds the string denoting the negative_marking field in the database.
	FieldNegativeMarking = "negative_marking"
	// FieldAiPrompt holds the string denoting the ai_prompt field in the database.
	FieldAiPrompt = "ai_prompt"
	// FieldOtherDetails holds the string denoting the other_details field in the database.
	FieldOtherDetails = "other_details"
	// FieldMaxAttempts holds the string denoting the max_attempts field in the database.
	FieldMaxAttempts = "max_attempts"
	// FieldEvaluationAiPrompt holds the string denoting the evaluation_ai_prompt field in the database.
	FieldEvaluationAiPrompt = "evaluation_ai_prompt"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeExam holds the string denoting the exam edge name in mutations.
	EdgeExam = "exam"
	// Table holds the table name of the examsetting in the database.
	Table = "exam_settings"
	// ExamTable is the table that holds the exam relation/edge.
	ExamTable = "exam_settings"
	// ExamInverseTable is the table name for the Exam entity.
	// It exists in this package in order to avoid circular dependency with the "exam" package.
	ExamInverseTable = "exams"
	// ExamColumn is the table column denoting the exam relation/edge.
	ExamColumn = "exam_setting"
)

// Columns holds all SQL columns for examsetting fields.
var Columns = []string{
	FieldID,
	FieldNumberOfQuestions,
	FieldDurationMinutes,
	FieldNegativeMarking,
	FieldAiPrompt,
	FieldOtherDetails,
	FieldMaxAttempts,
	FieldEvaluationAiPrompt,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "exam_settings"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"exam_setting",
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
	// DefaultMaxAttempts holds the default value on creation for the "max_attempts" field.
	DefaultMaxAttempts int
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the ExamSetting queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNumberOfQuestions orders the results by the number_of_questions field.
func ByNumberOfQuestions(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumberOfQuestions, opts...).ToFunc()
}

// ByDurationMinutes orders the results by the duration_minutes field.
func ByDurationMinutes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDurationMinutes, opts...).ToFunc()
}

// ByNegativeMarking orders the results by the negative_marking field.
func ByNegativeMarking(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNegativeMarking, opts...).ToFunc()
}

// ByAiPrompt orders the results by the ai_prompt field.
func ByAiPrompt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAiPrompt, opts...).ToFunc()
}

// ByMaxAttempts orders the results by the max_attempts field.
func ByMaxAttempts(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMaxAttempts, opts...).ToFunc()
}

// ByEvaluationAiPrompt orders the results by the evaluation_ai_prompt field.
func ByEvaluationAiPrompt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEvaluationAiPrompt, opts...).ToFunc()
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
		sqlgraph.Edge(sqlgraph.O2O, true, ExamTable, ExamColumn),
	)
}
