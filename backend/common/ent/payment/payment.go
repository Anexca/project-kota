// Code generated by ent, DO NOT EDIT.

package payment

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the payment type in the database.
	Label = "payment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldPaymentDate holds the string denoting the payment_date field in the database.
	FieldPaymentDate = "payment_date"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPaymentMethod holds the string denoting the payment_method field in the database.
	FieldPaymentMethod = "payment_method"
	// FieldProviderPaymentID holds the string denoting the provider_payment_id field in the database.
	FieldProviderPaymentID = "provider_payment_id"
	// FieldProviderInvoiceID holds the string denoting the provider_invoice_id field in the database.
	FieldProviderInvoiceID = "provider_invoice_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeSubscription holds the string denoting the subscription edge name in mutations.
	EdgeSubscription = "subscription"
	// Table holds the table name of the payment in the database.
	Table = "payments"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "payments"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_payments"
	// SubscriptionTable is the table that holds the subscription relation/edge.
	SubscriptionTable = "payments"
	// SubscriptionInverseTable is the table name for the UserSubscription entity.
	// It exists in this package in order to avoid circular dependency with the "usersubscription" package.
	SubscriptionInverseTable = "user_subscriptions"
	// SubscriptionColumn is the table column denoting the subscription relation/edge.
	SubscriptionColumn = "user_subscription_payments"
)

// Columns holds all SQL columns for payment fields.
var Columns = []string{
	FieldID,
	FieldAmount,
	FieldPaymentDate,
	FieldStatus,
	FieldPaymentMethod,
	FieldProviderPaymentID,
	FieldProviderInvoiceID,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "payments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_payments",
	"user_subscription_payments",
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

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusSUCCESS       Status = "SUCCESS"
	StatusNOT_ATTEMPTED Status = "NOT_ATTEMPTED"
	StatusUSER_DROPPED  Status = "USER_DROPPED"
	StatusFAILED        Status = "FAILED"
	StatusPENDING       Status = "PENDING"
	StatusCANCELLED     Status = "CANCELLED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusSUCCESS, StatusNOT_ATTEMPTED, StatusUSER_DROPPED, StatusFAILED, StatusPENDING, StatusCANCELLED:
		return nil
	default:
		return fmt.Errorf("payment: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Payment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByPaymentDate orders the results by the payment_date field.
func ByPaymentDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentDate, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPaymentMethod orders the results by the payment_method field.
func ByPaymentMethod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentMethod, opts...).ToFunc()
}

// ByProviderPaymentID orders the results by the provider_payment_id field.
func ByProviderPaymentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProviderPaymentID, opts...).ToFunc()
}

// ByProviderInvoiceID orders the results by the provider_invoice_id field.
func ByProviderInvoiceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProviderInvoiceID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// BySubscriptionField orders the results by subscription field.
func BySubscriptionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubscriptionStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newSubscriptionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubscriptionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SubscriptionTable, SubscriptionColumn),
	)
}
