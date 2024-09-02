// Code generated by ent, DO NOT EDIT.

package examattempt

import (
	"common/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldLTE(FieldID, id))
}

// AttemptNumber applies equality check predicate on the "attempt_number" field. It's identical to AttemptNumberEQ.
func AttemptNumber(v int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldEQ(FieldAttemptNumber, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldEQ(FieldUpdatedAt, v))
}

// AttemptNumberEQ applies the EQ predicate on the "attempt_number" field.
func AttemptNumberEQ(v int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldEQ(FieldAttemptNumber, v))
}

// AttemptNumberNEQ applies the NEQ predicate on the "attempt_number" field.
func AttemptNumberNEQ(v int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldNEQ(FieldAttemptNumber, v))
}

// AttemptNumberIn applies the In predicate on the "attempt_number" field.
func AttemptNumberIn(vs ...int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldIn(FieldAttemptNumber, vs...))
}

// AttemptNumberNotIn applies the NotIn predicate on the "attempt_number" field.
func AttemptNumberNotIn(vs ...int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldNotIn(FieldAttemptNumber, vs...))
}

// AttemptNumberGT applies the GT predicate on the "attempt_number" field.
func AttemptNumberGT(v int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldGT(FieldAttemptNumber, v))
}

// AttemptNumberGTE applies the GTE predicate on the "attempt_number" field.
func AttemptNumberGTE(v int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldGTE(FieldAttemptNumber, v))
}

// AttemptNumberLT applies the LT predicate on the "attempt_number" field.
func AttemptNumberLT(v int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldLT(FieldAttemptNumber, v))
}

// AttemptNumberLTE applies the LTE predicate on the "attempt_number" field.
func AttemptNumberLTE(v int) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldLTE(FieldAttemptNumber, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasGeneratedexam applies the HasEdge predicate on the "generatedexam" edge.
func HasGeneratedexam() predicate.ExamAttempt {
	return predicate.ExamAttempt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GeneratedexamTable, GeneratedexamColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGeneratedexamWith applies the HasEdge predicate on the "generatedexam" edge with a given conditions (other predicates).
func HasGeneratedexamWith(preds ...predicate.GeneratedExam) predicate.ExamAttempt {
	return predicate.ExamAttempt(func(s *sql.Selector) {
		step := newGeneratedexamStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.ExamAttempt {
	return predicate.ExamAttempt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.ExamAttempt {
	return predicate.ExamAttempt(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ExamAttempt) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ExamAttempt) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ExamAttempt) predicate.ExamAttempt {
	return predicate.ExamAttempt(sql.NotPredicates(p))
}
