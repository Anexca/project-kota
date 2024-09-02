// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/examattempt"
	"common/ent/generatedexam"
	"common/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ExamAttempt is the model entity for the ExamAttempt schema.
type ExamAttempt struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AttemptNumber holds the value of the "attempt_number" field.
	AttemptNumber int `json:"attempt_number,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExamAttemptQuery when eager-loading is set.
	Edges                   ExamAttemptEdges `json:"edges"`
	generated_exam_attempts *int
	user_attempts           *uuid.UUID
	selectValues            sql.SelectValues
}

// ExamAttemptEdges holds the relations/edges for other nodes in the graph.
type ExamAttemptEdges struct {
	// Generatedexam holds the value of the generatedexam edge.
	Generatedexam *GeneratedExam `json:"generatedexam,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GeneratedexamOrErr returns the Generatedexam value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ExamAttemptEdges) GeneratedexamOrErr() (*GeneratedExam, error) {
	if e.Generatedexam != nil {
		return e.Generatedexam, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: generatedexam.Label}
	}
	return nil, &NotLoadedError{edge: "generatedexam"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ExamAttemptEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ExamAttempt) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case examattempt.FieldID, examattempt.FieldAttemptNumber:
			values[i] = new(sql.NullInt64)
		case examattempt.FieldCreatedAt, examattempt.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case examattempt.ForeignKeys[0]: // generated_exam_attempts
			values[i] = new(sql.NullInt64)
		case examattempt.ForeignKeys[1]: // user_attempts
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ExamAttempt fields.
func (ea *ExamAttempt) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case examattempt.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ea.ID = int(value.Int64)
		case examattempt.FieldAttemptNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field attempt_number", values[i])
			} else if value.Valid {
				ea.AttemptNumber = int(value.Int64)
			}
		case examattempt.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ea.CreatedAt = value.Time
			}
		case examattempt.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ea.UpdatedAt = value.Time
			}
		case examattempt.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field generated_exam_attempts", value)
			} else if value.Valid {
				ea.generated_exam_attempts = new(int)
				*ea.generated_exam_attempts = int(value.Int64)
			}
		case examattempt.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_attempts", values[i])
			} else if value.Valid {
				ea.user_attempts = new(uuid.UUID)
				*ea.user_attempts = *value.S.(*uuid.UUID)
			}
		default:
			ea.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ExamAttempt.
// This includes values selected through modifiers, order, etc.
func (ea *ExamAttempt) Value(name string) (ent.Value, error) {
	return ea.selectValues.Get(name)
}

// QueryGeneratedexam queries the "generatedexam" edge of the ExamAttempt entity.
func (ea *ExamAttempt) QueryGeneratedexam() *GeneratedExamQuery {
	return NewExamAttemptClient(ea.config).QueryGeneratedexam(ea)
}

// QueryUser queries the "user" edge of the ExamAttempt entity.
func (ea *ExamAttempt) QueryUser() *UserQuery {
	return NewExamAttemptClient(ea.config).QueryUser(ea)
}

// Update returns a builder for updating this ExamAttempt.
// Note that you need to call ExamAttempt.Unwrap() before calling this method if this ExamAttempt
// was returned from a transaction, and the transaction was committed or rolled back.
func (ea *ExamAttempt) Update() *ExamAttemptUpdateOne {
	return NewExamAttemptClient(ea.config).UpdateOne(ea)
}

// Unwrap unwraps the ExamAttempt entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ea *ExamAttempt) Unwrap() *ExamAttempt {
	_tx, ok := ea.config.driver.(*txDriver)
	if !ok {
		panic("ent: ExamAttempt is not a transactional entity")
	}
	ea.config.driver = _tx.drv
	return ea
}

// String implements the fmt.Stringer.
func (ea *ExamAttempt) String() string {
	var builder strings.Builder
	builder.WriteString("ExamAttempt(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ea.ID))
	builder.WriteString("attempt_number=")
	builder.WriteString(fmt.Sprintf("%v", ea.AttemptNumber))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ea.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ea.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ExamAttempts is a parsable slice of ExamAttempt.
type ExamAttempts []*ExamAttempt
