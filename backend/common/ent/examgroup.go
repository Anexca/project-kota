// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/examcategory"
	"common/ent/examgroup"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ExamGroup is the model entity for the ExamGroup schema.
type ExamGroup struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// LogoURL holds the value of the "logo_url" field.
	LogoURL string `json:"logo_url,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExamGroupQuery when eager-loading is set.
	Edges                ExamGroupEdges `json:"edges"`
	exam_category_groups *int
	selectValues         sql.SelectValues
}

// ExamGroupEdges holds the relations/edges for other nodes in the graph.
type ExamGroupEdges struct {
	// Category holds the value of the category edge.
	Category *ExamCategory `json:"category,omitempty"`
	// Exams holds the value of the exams edge.
	Exams []*Exam `json:"exams,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ExamGroupEdges) CategoryOrErr() (*ExamCategory, error) {
	if e.Category != nil {
		return e.Category, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: examcategory.Label}
	}
	return nil, &NotLoadedError{edge: "category"}
}

// ExamsOrErr returns the Exams value or an error if the edge
// was not loaded in eager-loading.
func (e ExamGroupEdges) ExamsOrErr() ([]*Exam, error) {
	if e.loadedTypes[1] {
		return e.Exams, nil
	}
	return nil, &NotLoadedError{edge: "exams"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ExamGroup) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case examgroup.FieldIsActive:
			values[i] = new(sql.NullBool)
		case examgroup.FieldID:
			values[i] = new(sql.NullInt64)
		case examgroup.FieldName, examgroup.FieldDescription, examgroup.FieldLogoURL:
			values[i] = new(sql.NullString)
		case examgroup.FieldCreatedAt, examgroup.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case examgroup.ForeignKeys[0]: // exam_category_groups
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ExamGroup fields.
func (eg *ExamGroup) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case examgroup.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			eg.ID = int(value.Int64)
		case examgroup.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				eg.Name = value.String
			}
		case examgroup.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				eg.Description = value.String
			}
		case examgroup.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				eg.IsActive = value.Bool
			}
		case examgroup.FieldLogoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo_url", values[i])
			} else if value.Valid {
				eg.LogoURL = value.String
			}
		case examgroup.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				eg.CreatedAt = value.Time
			}
		case examgroup.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				eg.UpdatedAt = value.Time
			}
		case examgroup.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_category_groups", value)
			} else if value.Valid {
				eg.exam_category_groups = new(int)
				*eg.exam_category_groups = int(value.Int64)
			}
		default:
			eg.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ExamGroup.
// This includes values selected through modifiers, order, etc.
func (eg *ExamGroup) Value(name string) (ent.Value, error) {
	return eg.selectValues.Get(name)
}

// QueryCategory queries the "category" edge of the ExamGroup entity.
func (eg *ExamGroup) QueryCategory() *ExamCategoryQuery {
	return NewExamGroupClient(eg.config).QueryCategory(eg)
}

// QueryExams queries the "exams" edge of the ExamGroup entity.
func (eg *ExamGroup) QueryExams() *ExamQuery {
	return NewExamGroupClient(eg.config).QueryExams(eg)
}

// Update returns a builder for updating this ExamGroup.
// Note that you need to call ExamGroup.Unwrap() before calling this method if this ExamGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (eg *ExamGroup) Update() *ExamGroupUpdateOne {
	return NewExamGroupClient(eg.config).UpdateOne(eg)
}

// Unwrap unwraps the ExamGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (eg *ExamGroup) Unwrap() *ExamGroup {
	_tx, ok := eg.config.driver.(*txDriver)
	if !ok {
		panic("ent: ExamGroup is not a transactional entity")
	}
	eg.config.driver = _tx.drv
	return eg
}

// String implements the fmt.Stringer.
func (eg *ExamGroup) String() string {
	var builder strings.Builder
	builder.WriteString("ExamGroup(")
	builder.WriteString(fmt.Sprintf("id=%v, ", eg.ID))
	builder.WriteString("name=")
	builder.WriteString(eg.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(eg.Description)
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", eg.IsActive))
	builder.WriteString(", ")
	builder.WriteString("logo_url=")
	builder.WriteString(eg.LogoURL)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(eg.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(eg.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ExamGroups is a parsable slice of ExamGroup.
type ExamGroups []*ExamGroup
