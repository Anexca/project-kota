// Code generated by ent, DO NOT EDIT.

package examcategory

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the examcategory type in the database.
	Label = "exam_category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeExams holds the string denoting the exams edge name in mutations.
	EdgeExams = "exams"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// Table holds the table name of the examcategory in the database.
	Table = "exam_categories"
	// ExamsTable is the table that holds the exams relation/edge.
	ExamsTable = "exams"
	// ExamsInverseTable is the table name for the Exam entity.
	// It exists in this package in order to avoid circular dependency with the "exam" package.
	ExamsInverseTable = "exams"
	// ExamsColumn is the table column denoting the exams relation/edge.
	ExamsColumn = "exam_category_exams"
	// GroupsTable is the table that holds the groups relation/edge.
	GroupsTable = "exam_groups"
	// GroupsInverseTable is the table name for the ExamGroup entity.
	// It exists in this package in order to avoid circular dependency with the "examgroup" package.
	GroupsInverseTable = "exam_groups"
	// GroupsColumn is the table column denoting the groups relation/edge.
	GroupsColumn = "exam_category_groups"
)

// Columns holds all SQL columns for examcategory fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldIsActive,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Name defines the type for the "name" enum field.
type Name string

// Name values.
const (
	NameBANKING Name = "BANKING"
)

func (n Name) String() string {
	return string(n)
}

// NameValidator is a validator for the "name" field enum values. It is called by the builders before save.
func NameValidator(n Name) error {
	switch n {
	case NameBANKING:
		return nil
	default:
		return fmt.Errorf("examcategory: invalid enum value for name field: %q", n)
	}
}

// OrderOption defines the ordering options for the ExamCategory queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByIsActive orders the results by the is_active field.
func ByIsActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsActive, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByExamsCount orders the results by exams count.
func ByExamsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExamsStep(), opts...)
	}
}

// ByExams orders the results by exams terms.
func ByExams(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExamsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByGroupsCount orders the results by groups count.
func ByGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGroupsStep(), opts...)
	}
}

// ByGroups orders the results by groups terms.
func ByGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newExamsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExamsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExamsTable, ExamsColumn),
	)
}
func newGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, GroupsTable, GroupsColumn),
	)
}
