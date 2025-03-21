// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/subscription"
	"common/ent/user"
	"common/ent/usersubscription"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// UserSubscription is the model entity for the UserSubscription schema.
type UserSubscription struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// Status holds the value of the "status" field.
	Status usersubscription.Status `json:"status,omitempty"`
	// StartDate holds the value of the "start_date" field.
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate holds the value of the "end_date" field.
	EndDate time.Time `json:"end_date,omitempty"`
	// ProviderSubscriptionID holds the value of the "provider_subscription_id" field.
	ProviderSubscriptionID string `json:"provider_subscription_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserSubscriptionQuery when eager-loading is set.
	Edges                           UserSubscriptionEdges `json:"edges"`
	subscription_user_subscriptions *int
	user_subscriptions              *uuid.UUID
	selectValues                    sql.SelectValues
}

// UserSubscriptionEdges holds the relations/edges for other nodes in the graph.
type UserSubscriptionEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Subscription holds the value of the subscription edge.
	Subscription *Subscription `json:"subscription,omitempty"`
	// Payments holds the value of the payments edge.
	Payments []*Payment `json:"payments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserSubscriptionEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// SubscriptionOrErr returns the Subscription value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserSubscriptionEdges) SubscriptionOrErr() (*Subscription, error) {
	if e.Subscription != nil {
		return e.Subscription, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: subscription.Label}
	}
	return nil, &NotLoadedError{edge: "subscription"}
}

// PaymentsOrErr returns the Payments value or an error if the edge
// was not loaded in eager-loading.
func (e UserSubscriptionEdges) PaymentsOrErr() ([]*Payment, error) {
	if e.loadedTypes[2] {
		return e.Payments, nil
	}
	return nil, &NotLoadedError{edge: "payments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserSubscription) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usersubscription.FieldIsActive:
			values[i] = new(sql.NullBool)
		case usersubscription.FieldID:
			values[i] = new(sql.NullInt64)
		case usersubscription.FieldStatus, usersubscription.FieldProviderSubscriptionID:
			values[i] = new(sql.NullString)
		case usersubscription.FieldStartDate, usersubscription.FieldEndDate, usersubscription.FieldCreatedAt, usersubscription.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case usersubscription.ForeignKeys[0]: // subscription_user_subscriptions
			values[i] = new(sql.NullInt64)
		case usersubscription.ForeignKeys[1]: // user_subscriptions
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserSubscription fields.
func (us *UserSubscription) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usersubscription.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			us.ID = int(value.Int64)
		case usersubscription.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				us.IsActive = value.Bool
			}
		case usersubscription.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				us.Status = usersubscription.Status(value.String)
			}
		case usersubscription.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				us.StartDate = value.Time
			}
		case usersubscription.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_date", values[i])
			} else if value.Valid {
				us.EndDate = value.Time
			}
		case usersubscription.FieldProviderSubscriptionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider_subscription_id", values[i])
			} else if value.Valid {
				us.ProviderSubscriptionID = value.String
			}
		case usersubscription.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				us.CreatedAt = value.Time
			}
		case usersubscription.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				us.UpdatedAt = value.Time
			}
		case usersubscription.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field subscription_user_subscriptions", value)
			} else if value.Valid {
				us.subscription_user_subscriptions = new(int)
				*us.subscription_user_subscriptions = int(value.Int64)
			}
		case usersubscription.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_subscriptions", values[i])
			} else if value.Valid {
				us.user_subscriptions = new(uuid.UUID)
				*us.user_subscriptions = *value.S.(*uuid.UUID)
			}
		default:
			us.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserSubscription.
// This includes values selected through modifiers, order, etc.
func (us *UserSubscription) Value(name string) (ent.Value, error) {
	return us.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserSubscription entity.
func (us *UserSubscription) QueryUser() *UserQuery {
	return NewUserSubscriptionClient(us.config).QueryUser(us)
}

// QuerySubscription queries the "subscription" edge of the UserSubscription entity.
func (us *UserSubscription) QuerySubscription() *SubscriptionQuery {
	return NewUserSubscriptionClient(us.config).QuerySubscription(us)
}

// QueryPayments queries the "payments" edge of the UserSubscription entity.
func (us *UserSubscription) QueryPayments() *PaymentQuery {
	return NewUserSubscriptionClient(us.config).QueryPayments(us)
}

// Update returns a builder for updating this UserSubscription.
// Note that you need to call UserSubscription.Unwrap() before calling this method if this UserSubscription
// was returned from a transaction, and the transaction was committed or rolled back.
func (us *UserSubscription) Update() *UserSubscriptionUpdateOne {
	return NewUserSubscriptionClient(us.config).UpdateOne(us)
}

// Unwrap unwraps the UserSubscription entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (us *UserSubscription) Unwrap() *UserSubscription {
	_tx, ok := us.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserSubscription is not a transactional entity")
	}
	us.config.driver = _tx.drv
	return us
}

// String implements the fmt.Stringer.
func (us *UserSubscription) String() string {
	var builder strings.Builder
	builder.WriteString("UserSubscription(")
	builder.WriteString(fmt.Sprintf("id=%v, ", us.ID))
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", us.IsActive))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", us.Status))
	builder.WriteString(", ")
	builder.WriteString("start_date=")
	builder.WriteString(us.StartDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_date=")
	builder.WriteString(us.EndDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("provider_subscription_id=")
	builder.WriteString(us.ProviderSubscriptionID)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(us.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(us.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserSubscriptions is a parsable slice of UserSubscription.
type UserSubscriptions []*UserSubscription
