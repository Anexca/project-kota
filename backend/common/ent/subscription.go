// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/subscription"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Subscription is the model entity for the Subscription schema.
type Subscription struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ProviderPlanID holds the value of the "provider_plan_id" field.
	ProviderPlanID string `json:"provider_plan_id,omitempty"`
	// BasePrice holds the value of the "base_price" field.
	BasePrice float64 `json:"base_price,omitempty"`
	// FinalPrice holds the value of the "final_price" field.
	FinalPrice float64 `json:"final_price,omitempty"`
	// Price holds the value of the "price" field.
	//
	// Deprecated: Field "price" was marked as deprecated in the schema.
	Price float64 `json:"price,omitempty"`
	// DurationInMonths holds the value of the "duration_in_months" field.
	DurationInMonths int `json:"duration_in_months,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// RawSubscriptionData holds the value of the "raw_subscription_data" field.
	RawSubscriptionData map[string]interface{} `json:"raw_subscription_data,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubscriptionQuery when eager-loading is set.
	Edges        SubscriptionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SubscriptionEdges holds the relations/edges for other nodes in the graph.
type SubscriptionEdges struct {
	// Exams holds the value of the exams edge.
	Exams []*SubscriptionExam `json:"exams,omitempty"`
	// UserSubscriptions holds the value of the user_subscriptions edge.
	UserSubscriptions []*UserSubscription `json:"user_subscriptions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ExamsOrErr returns the Exams value or an error if the edge
// was not loaded in eager-loading.
func (e SubscriptionEdges) ExamsOrErr() ([]*SubscriptionExam, error) {
	if e.loadedTypes[0] {
		return e.Exams, nil
	}
	return nil, &NotLoadedError{edge: "exams"}
}

// UserSubscriptionsOrErr returns the UserSubscriptions value or an error if the edge
// was not loaded in eager-loading.
func (e SubscriptionEdges) UserSubscriptionsOrErr() ([]*UserSubscription, error) {
	if e.loadedTypes[1] {
		return e.UserSubscriptions, nil
	}
	return nil, &NotLoadedError{edge: "user_subscriptions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subscription) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subscription.FieldRawSubscriptionData:
			values[i] = new([]byte)
		case subscription.FieldIsActive:
			values[i] = new(sql.NullBool)
		case subscription.FieldBasePrice, subscription.FieldFinalPrice, subscription.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case subscription.FieldID, subscription.FieldDurationInMonths:
			values[i] = new(sql.NullInt64)
		case subscription.FieldProviderPlanID, subscription.FieldName:
			values[i] = new(sql.NullString)
		case subscription.FieldCreatedAt, subscription.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subscription fields.
func (s *Subscription) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subscription.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case subscription.FieldProviderPlanID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider_plan_id", values[i])
			} else if value.Valid {
				s.ProviderPlanID = value.String
			}
		case subscription.FieldBasePrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field base_price", values[i])
			} else if value.Valid {
				s.BasePrice = value.Float64
			}
		case subscription.FieldFinalPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field final_price", values[i])
			} else if value.Valid {
				s.FinalPrice = value.Float64
			}
		case subscription.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				s.Price = value.Float64
			}
		case subscription.FieldDurationInMonths:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration_in_months", values[i])
			} else if value.Valid {
				s.DurationInMonths = int(value.Int64)
			}
		case subscription.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				s.IsActive = value.Bool
			}
		case subscription.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case subscription.FieldRawSubscriptionData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field raw_subscription_data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.RawSubscriptionData); err != nil {
					return fmt.Errorf("unmarshal field raw_subscription_data: %w", err)
				}
			}
		case subscription.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case subscription.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Subscription.
// This includes values selected through modifiers, order, etc.
func (s *Subscription) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryExams queries the "exams" edge of the Subscription entity.
func (s *Subscription) QueryExams() *SubscriptionExamQuery {
	return NewSubscriptionClient(s.config).QueryExams(s)
}

// QueryUserSubscriptions queries the "user_subscriptions" edge of the Subscription entity.
func (s *Subscription) QueryUserSubscriptions() *UserSubscriptionQuery {
	return NewSubscriptionClient(s.config).QueryUserSubscriptions(s)
}

// Update returns a builder for updating this Subscription.
// Note that you need to call Subscription.Unwrap() before calling this method if this Subscription
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subscription) Update() *SubscriptionUpdateOne {
	return NewSubscriptionClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Subscription entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subscription) Unwrap() *Subscription {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subscription is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subscription) String() string {
	var builder strings.Builder
	builder.WriteString("Subscription(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("provider_plan_id=")
	builder.WriteString(s.ProviderPlanID)
	builder.WriteString(", ")
	builder.WriteString("base_price=")
	builder.WriteString(fmt.Sprintf("%v", s.BasePrice))
	builder.WriteString(", ")
	builder.WriteString("final_price=")
	builder.WriteString(fmt.Sprintf("%v", s.FinalPrice))
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", s.Price))
	builder.WriteString(", ")
	builder.WriteString("duration_in_months=")
	builder.WriteString(fmt.Sprintf("%v", s.DurationInMonths))
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", s.IsActive))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("raw_subscription_data=")
	builder.WriteString(fmt.Sprintf("%v", s.RawSubscriptionData))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Subscriptions is a parsable slice of Subscription.
type Subscriptions []*Subscription
