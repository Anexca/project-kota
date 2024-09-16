// Code generated by ent, DO NOT EDIT.

package exam

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the exam type in the database.
	Label = "exam"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// EdgeSubscriptions holds the string denoting the subscriptions edge name in mutations.
	EdgeSubscriptions = "subscriptions"
	// EdgeSetting holds the string denoting the setting edge name in mutations.
	EdgeSetting = "setting"
	// EdgeCachedExam holds the string denoting the cached_exam edge name in mutations.
	EdgeCachedExam = "cached_exam"
	// EdgeGeneratedexams holds the string denoting the generatedexams edge name in mutations.
	EdgeGeneratedexams = "generatedexams"
	// Table holds the table name of the exam in the database.
	Table = "exams"
	// CategoryTable is the table that holds the category relation/edge.
	CategoryTable = "exams"
	// CategoryInverseTable is the table name for the ExamCategory entity.
	// It exists in this package in order to avoid circular dependency with the "examcategory" package.
	CategoryInverseTable = "exam_categories"
	// CategoryColumn is the table column denoting the category relation/edge.
	CategoryColumn = "exam_category_exams"
	// SubscriptionsTable is the table that holds the subscriptions relation/edge.
	SubscriptionsTable = "subscription_exams"
	// SubscriptionsInverseTable is the table name for the SubscriptionExam entity.
	// It exists in this package in order to avoid circular dependency with the "subscriptionexam" package.
	SubscriptionsInverseTable = "subscription_exams"
	// SubscriptionsColumn is the table column denoting the subscriptions relation/edge.
	SubscriptionsColumn = "exam_subscriptions"
	// SettingTable is the table that holds the setting relation/edge.
	SettingTable = "exam_settings"
	// SettingInverseTable is the table name for the ExamSetting entity.
	// It exists in this package in order to avoid circular dependency with the "examsetting" package.
	SettingInverseTable = "exam_settings"
	// SettingColumn is the table column denoting the setting relation/edge.
	SettingColumn = "exam_setting"
	// CachedExamTable is the table that holds the cached_exam relation/edge.
	CachedExamTable = "cached_exams"
	// CachedExamInverseTable is the table name for the CachedExam entity.
	// It exists in this package in order to avoid circular dependency with the "cachedexam" package.
	CachedExamInverseTable = "cached_exams"
	// CachedExamColumn is the table column denoting the cached_exam relation/edge.
	CachedExamColumn = "exam_cached_exam"
	// GeneratedexamsTable is the table that holds the generatedexams relation/edge.
	GeneratedexamsTable = "generated_exams"
	// GeneratedexamsInverseTable is the table name for the GeneratedExam entity.
	// It exists in this package in order to avoid circular dependency with the "generatedexam" package.
	GeneratedexamsInverseTable = "generated_exams"
	// GeneratedexamsColumn is the table column denoting the generatedexams relation/edge.
	GeneratedexamsColumn = "exam_generatedexams"
)

// Columns holds all SQL columns for exam fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldType,
	FieldIsActive,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "exams"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"exam_category_exams",
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
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Type defines the type for the "type" enum field.
type Type string

// TypeDESCRIPTIVE is the default value of the Type enum.
const DefaultType = TypeDESCRIPTIVE

// Type values.
const (
	TypeMCQ         Type = "MCQ"
	TypeDESCRIPTIVE Type = "DESCRIPTIVE"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeMCQ, TypeDESCRIPTIVE:
		return nil
	default:
		return fmt.Errorf("exam: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the Exam queries.
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

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
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

// ByCategoryField orders the results by category field.
func ByCategoryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoryStep(), sql.OrderByField(field, opts...))
	}
}

// BySubscriptionsCount orders the results by subscriptions count.
func BySubscriptionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubscriptionsStep(), opts...)
	}
}

// BySubscriptions orders the results by subscriptions terms.
func BySubscriptions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubscriptionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySettingField orders the results by setting field.
func BySettingField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSettingStep(), sql.OrderByField(field, opts...))
	}
}

// ByCachedExamCount orders the results by cached_exam count.
func ByCachedExamCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCachedExamStep(), opts...)
	}
}

// ByCachedExam orders the results by cached_exam terms.
func ByCachedExam(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCachedExamStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByGeneratedexamsCount orders the results by generatedexams count.
func ByGeneratedexamsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGeneratedexamsStep(), opts...)
	}
}

// ByGeneratedexams orders the results by generatedexams terms.
func ByGeneratedexams(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGeneratedexamsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
	)
}
func newSubscriptionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubscriptionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubscriptionsTable, SubscriptionsColumn),
	)
}
func newSettingStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SettingInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, SettingTable, SettingColumn),
	)
}
func newCachedExamStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CachedExamInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CachedExamTable, CachedExamColumn),
	)
}
func newGeneratedexamsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GeneratedexamsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, GeneratedexamsTable, GeneratedexamsColumn),
	)
}
