// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/cachedquestionmetadata"
	"common/ent/exam"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// CachedQuestionMetaData is the model entity for the CachedQuestionMetaData schema.
type CachedQuestionMetaData struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CacheUID holds the value of the "cache_uid" field.
	CacheUID string `json:"cache_uid,omitempty"`
	// IsUsed holds the value of the "is_used" field.
	IsUsed bool `json:"is_used,omitempty"`
	// ExpiresAt holds the value of the "expires_at" field.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CachedQuestionMetaDataQuery when eager-loading is set.
	Edges                         CachedQuestionMetaDataEdges `json:"edges"`
	exam_cached_question_metadata *int
	selectValues                  sql.SelectValues
}

// CachedQuestionMetaDataEdges holds the relations/edges for other nodes in the graph.
type CachedQuestionMetaDataEdges struct {
	// Exam holds the value of the exam edge.
	Exam *Exam `json:"exam,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ExamOrErr returns the Exam value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CachedQuestionMetaDataEdges) ExamOrErr() (*Exam, error) {
	if e.Exam != nil {
		return e.Exam, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: exam.Label}
	}
	return nil, &NotLoadedError{edge: "exam"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CachedQuestionMetaData) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cachedquestionmetadata.FieldIsUsed:
			values[i] = new(sql.NullBool)
		case cachedquestionmetadata.FieldID:
			values[i] = new(sql.NullInt64)
		case cachedquestionmetadata.FieldCacheUID:
			values[i] = new(sql.NullString)
		case cachedquestionmetadata.FieldExpiresAt, cachedquestionmetadata.FieldCreatedAt, cachedquestionmetadata.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case cachedquestionmetadata.ForeignKeys[0]: // exam_cached_question_metadata
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CachedQuestionMetaData fields.
func (cqmd *CachedQuestionMetaData) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cachedquestionmetadata.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cqmd.ID = int(value.Int64)
		case cachedquestionmetadata.FieldCacheUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cache_uid", values[i])
			} else if value.Valid {
				cqmd.CacheUID = value.String
			}
		case cachedquestionmetadata.FieldIsUsed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_used", values[i])
			} else if value.Valid {
				cqmd.IsUsed = value.Bool
			}
		case cachedquestionmetadata.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				cqmd.ExpiresAt = value.Time
			}
		case cachedquestionmetadata.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cqmd.CreatedAt = value.Time
			}
		case cachedquestionmetadata.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cqmd.UpdatedAt = value.Time
			}
		case cachedquestionmetadata.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exam_cached_question_metadata", value)
			} else if value.Valid {
				cqmd.exam_cached_question_metadata = new(int)
				*cqmd.exam_cached_question_metadata = int(value.Int64)
			}
		default:
			cqmd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CachedQuestionMetaData.
// This includes values selected through modifiers, order, etc.
func (cqmd *CachedQuestionMetaData) Value(name string) (ent.Value, error) {
	return cqmd.selectValues.Get(name)
}

// QueryExam queries the "exam" edge of the CachedQuestionMetaData entity.
func (cqmd *CachedQuestionMetaData) QueryExam() *ExamQuery {
	return NewCachedQuestionMetaDataClient(cqmd.config).QueryExam(cqmd)
}

// Update returns a builder for updating this CachedQuestionMetaData.
// Note that you need to call CachedQuestionMetaData.Unwrap() before calling this method if this CachedQuestionMetaData
// was returned from a transaction, and the transaction was committed or rolled back.
func (cqmd *CachedQuestionMetaData) Update() *CachedQuestionMetaDataUpdateOne {
	return NewCachedQuestionMetaDataClient(cqmd.config).UpdateOne(cqmd)
}

// Unwrap unwraps the CachedQuestionMetaData entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cqmd *CachedQuestionMetaData) Unwrap() *CachedQuestionMetaData {
	_tx, ok := cqmd.config.driver.(*txDriver)
	if !ok {
		panic("ent: CachedQuestionMetaData is not a transactional entity")
	}
	cqmd.config.driver = _tx.drv
	return cqmd
}

// String implements the fmt.Stringer.
func (cqmd *CachedQuestionMetaData) String() string {
	var builder strings.Builder
	builder.WriteString("CachedQuestionMetaData(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cqmd.ID))
	builder.WriteString("cache_uid=")
	builder.WriteString(cqmd.CacheUID)
	builder.WriteString(", ")
	builder.WriteString("is_used=")
	builder.WriteString(fmt.Sprintf("%v", cqmd.IsUsed))
	builder.WriteString(", ")
	builder.WriteString("expires_at=")
	builder.WriteString(cqmd.ExpiresAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(cqmd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(cqmd.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CachedQuestionMetaDataSlice is a parsable slice of CachedQuestionMetaData.
type CachedQuestionMetaDataSlice []*CachedQuestionMetaData
