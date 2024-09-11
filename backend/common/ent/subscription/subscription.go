// Code generated by ent, DO NOT EDIT.

package subscription

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the subscription type in the database.
	Label = "subscription"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProviderSubscriptionID holds the string denoting the provider_subscription_id field in the database.
	FieldProviderSubscriptionID = "provider_subscription_id"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldDurationInMonths holds the string denoting the duration_in_months field in the database.
	FieldDurationInMonths = "duration_in_months"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldRawSubscriptionData holds the string denoting the raw_subscription_data field in the database.
	FieldRawSubscriptionData = "raw_subscription_data"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeExams holds the string denoting the exams edge name in mutations.
	EdgeExams = "exams"
	// EdgeUserSubscriptions holds the string denoting the user_subscriptions edge name in mutations.
	EdgeUserSubscriptions = "user_subscriptions"
	// Table holds the table name of the subscription in the database.
	Table = "subscriptions"
	// ExamsTable is the table that holds the exams relation/edge.
	ExamsTable = "subscription_exams"
	// ExamsInverseTable is the table name for the SubscriptionExam entity.
	// It exists in this package in order to avoid circular dependency with the "subscriptionexam" package.
	ExamsInverseTable = "subscription_exams"
	// ExamsColumn is the table column denoting the exams relation/edge.
	ExamsColumn = "subscription_exams"
	// UserSubscriptionsTable is the table that holds the user_subscriptions relation/edge.
	UserSubscriptionsTable = "user_subscriptions"
	// UserSubscriptionsInverseTable is the table name for the UserSubscription entity.
	// It exists in this package in order to avoid circular dependency with the "usersubscription" package.
	UserSubscriptionsInverseTable = "user_subscriptions"
	// UserSubscriptionsColumn is the table column denoting the user_subscriptions relation/edge.
	UserSubscriptionsColumn = "subscription_user_subscriptions"
)

// Columns holds all SQL columns for subscription fields.
var Columns = []string{
	FieldID,
	FieldProviderSubscriptionID,
	FieldPrice,
	FieldDurationInMonths,
	FieldIsActive,
	FieldName,
	FieldRawSubscriptionData,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Subscription queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByProviderSubscriptionID orders the results by the provider_subscription_id field.
func ByProviderSubscriptionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProviderSubscriptionID, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByDurationInMonths orders the results by the duration_in_months field.
func ByDurationInMonths(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDurationInMonths, opts...).ToFunc()
}

// ByIsActive orders the results by the is_active field.
func ByIsActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsActive, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
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

// ByUserSubscriptionsCount orders the results by user_subscriptions count.
func ByUserSubscriptionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserSubscriptionsStep(), opts...)
	}
}

// ByUserSubscriptions orders the results by user_subscriptions terms.
func ByUserSubscriptions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserSubscriptionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newExamsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExamsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExamsTable, ExamsColumn),
	)
}
func newUserSubscriptionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserSubscriptionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserSubscriptionsTable, UserSubscriptionsColumn),
	)
}
