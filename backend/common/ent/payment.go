// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/payment"
	"common/ent/user"
	"common/ent/usersubscription"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Payment is the model entity for the Payment schema.
type Payment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount int `json:"amount,omitempty"`
	// PaymentDate holds the value of the "payment_date" field.
	PaymentDate time.Time `json:"payment_date,omitempty"`
	// Status holds the value of the "status" field.
	Status payment.Status `json:"status,omitempty"`
	// PaymentMethod holds the value of the "payment_method" field.
	PaymentMethod string `json:"payment_method,omitempty"`
	// ProviderPaymentID holds the value of the "provider_payment_id" field.
	ProviderPaymentID string `json:"provider_payment_id,omitempty"`
	// ReceiptID holds the value of the "receipt_id" field.
	ReceiptID string `json:"receipt_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PaymentQuery when eager-loading is set.
	Edges                      PaymentEdges `json:"edges"`
	user_payments              *uuid.UUID
	user_subscription_payments *int
	selectValues               sql.SelectValues
}

// PaymentEdges holds the relations/edges for other nodes in the graph.
type PaymentEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Subscription holds the value of the subscription edge.
	Subscription *UserSubscription `json:"subscription,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// SubscriptionOrErr returns the Subscription value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentEdges) SubscriptionOrErr() (*UserSubscription, error) {
	if e.Subscription != nil {
		return e.Subscription, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: usersubscription.Label}
	}
	return nil, &NotLoadedError{edge: "subscription"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Payment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case payment.FieldID, payment.FieldAmount:
			values[i] = new(sql.NullInt64)
		case payment.FieldStatus, payment.FieldPaymentMethod, payment.FieldProviderPaymentID, payment.FieldReceiptID:
			values[i] = new(sql.NullString)
		case payment.FieldPaymentDate, payment.FieldCreatedAt, payment.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case payment.ForeignKeys[0]: // user_payments
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case payment.ForeignKeys[1]: // user_subscription_payments
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Payment fields.
func (pa *Payment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case payment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case payment.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				pa.Amount = int(value.Int64)
			}
		case payment.FieldPaymentDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field payment_date", values[i])
			} else if value.Valid {
				pa.PaymentDate = value.Time
			}
		case payment.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pa.Status = payment.Status(value.String)
			}
		case payment.FieldPaymentMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payment_method", values[i])
			} else if value.Valid {
				pa.PaymentMethod = value.String
			}
		case payment.FieldProviderPaymentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider_payment_id", values[i])
			} else if value.Valid {
				pa.ProviderPaymentID = value.String
			}
		case payment.FieldReceiptID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field receipt_id", values[i])
			} else if value.Valid {
				pa.ReceiptID = value.String
			}
		case payment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pa.CreatedAt = value.Time
			}
		case payment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pa.UpdatedAt = value.Time
			}
		case payment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_payments", values[i])
			} else if value.Valid {
				pa.user_payments = new(uuid.UUID)
				*pa.user_payments = *value.S.(*uuid.UUID)
			}
		case payment.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_subscription_payments", value)
			} else if value.Valid {
				pa.user_subscription_payments = new(int)
				*pa.user_subscription_payments = int(value.Int64)
			}
		default:
			pa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Payment.
// This includes values selected through modifiers, order, etc.
func (pa *Payment) Value(name string) (ent.Value, error) {
	return pa.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Payment entity.
func (pa *Payment) QueryUser() *UserQuery {
	return NewPaymentClient(pa.config).QueryUser(pa)
}

// QuerySubscription queries the "subscription" edge of the Payment entity.
func (pa *Payment) QuerySubscription() *UserSubscriptionQuery {
	return NewPaymentClient(pa.config).QuerySubscription(pa)
}

// Update returns a builder for updating this Payment.
// Note that you need to call Payment.Unwrap() before calling this method if this Payment
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Payment) Update() *PaymentUpdateOne {
	return NewPaymentClient(pa.config).UpdateOne(pa)
}

// Unwrap unwraps the Payment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Payment) Unwrap() *Payment {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Payment is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Payment) String() string {
	var builder strings.Builder
	builder.WriteString("Payment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", pa.Amount))
	builder.WriteString(", ")
	builder.WriteString("payment_date=")
	builder.WriteString(pa.PaymentDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pa.Status))
	builder.WriteString(", ")
	builder.WriteString("payment_method=")
	builder.WriteString(pa.PaymentMethod)
	builder.WriteString(", ")
	builder.WriteString("provider_payment_id=")
	builder.WriteString(pa.ProviderPaymentID)
	builder.WriteString(", ")
	builder.WriteString("receipt_id=")
	builder.WriteString(pa.ReceiptID)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pa.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pa.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Payments is a parsable slice of Payment.
type Payments []*Payment
