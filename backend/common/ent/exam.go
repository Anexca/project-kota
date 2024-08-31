// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/exam"
	"common/ent/examcategory"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Exam is the model entity for the Exam schema.
type Exam struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExamQuery when eager-loading is set.
	Edges               ExamEdges `json:"edges"`
	exam_category_exams *int
	selectValues        sql.SelectValues
}

// ExamEdges holds the relations/edges for other nodes in the graph.
type ExamEdges struct {
	// Category holds the value of the category edge.
	Category *ExamCategory `json:"category,omitempty"`
	// Settings holds the value of the settings edge.
	Settings []*ExamSetting `json:"settings,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ExamEdges) CategoryOrErr() (*ExamCategory, error) {
	if e.Category != nil {
		return e.Category, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: examcategory.Label}
	}
	return nil, &NotLoadedError{edge: "category"}
}

// SettingsOrErr returns the Settings value or an error if the edge
// was not loaded in eager-loading.
func (e ExamEdges) SettingsOrErr() ([]*ExamSetting, error) {
	if e.loadedTypes[1] {
		return e.Settings, nil
	}
	return nil, &NotLoadedError{edge: "settings"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Exam) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case exam.FieldIsActive:
			values[i] = new(sql.NullBool)
		case exam.FieldID:
			values[i] = new(sql.NullInt64)
		case exam.FieldName, exam.FieldDescription:
			values[i] = new(sql.NullString)
		case exam.FieldCreatedAt, exam.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case exam.ForeignKeys[0]: // exam_category_exams
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Exam fields.
func (e *Exam) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case exam.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case exam.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case exam.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				e.Description = value.String
			}
		case exam.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				e.IsActive = value.Bool
			}
		case exam.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case exam.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				e.UpdatedAt = value.Time
			}
		case exam.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_category_exams", value)
			} else if value.Valid {
				e.exam_category_exams = new(int)
				*e.exam_category_exams = int(value.Int64)
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Exam.
// This includes values selected through modifiers, order, etc.
func (e *Exam) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QueryCategory queries the "category" edge of the Exam entity.
func (e *Exam) QueryCategory() *ExamCategoryQuery {
	return NewExamClient(e.config).QueryCategory(e)
}

// QuerySettings queries the "settings" edge of the Exam entity.
func (e *Exam) QuerySettings() *ExamSettingQuery {
	return NewExamClient(e.config).QuerySettings(e)
}

// Update returns a builder for updating this Exam.
// Note that you need to call Exam.Unwrap() before calling this method if this Exam
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Exam) Update() *ExamUpdateOne {
	return NewExamClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Exam entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Exam) Unwrap() *Exam {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Exam is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Exam) String() string {
	var builder strings.Builder
	builder.WriteString("Exam(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(e.Description)
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", e.IsActive))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(e.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Exams is a parsable slice of Exam.
type Exams []*Exam
