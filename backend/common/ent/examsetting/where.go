// Code generated by ent, DO NOT EDIT.

package examsetting

import (
	"common/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldID, id))
}

// NumberOfQuestions applies equality check predicate on the "number_of_questions" field. It's identical to NumberOfQuestionsEQ.
func NumberOfQuestions(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldNumberOfQuestions, v))
}

// DurationMinutes applies equality check predicate on the "duration_minutes" field. It's identical to DurationMinutesEQ.
func DurationMinutes(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldDurationMinutes, v))
}

// NegativeMarking applies equality check predicate on the "negative_marking" field. It's identical to NegativeMarkingEQ.
func NegativeMarking(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldNegativeMarking, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldUpdatedAt, v))
}

// NumberOfQuestionsEQ applies the EQ predicate on the "number_of_questions" field.
func NumberOfQuestionsEQ(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldNumberOfQuestions, v))
}

// NumberOfQuestionsNEQ applies the NEQ predicate on the "number_of_questions" field.
func NumberOfQuestionsNEQ(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldNumberOfQuestions, v))
}

// NumberOfQuestionsIn applies the In predicate on the "number_of_questions" field.
func NumberOfQuestionsIn(vs ...int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldNumberOfQuestions, vs...))
}

// NumberOfQuestionsNotIn applies the NotIn predicate on the "number_of_questions" field.
func NumberOfQuestionsNotIn(vs ...int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldNumberOfQuestions, vs...))
}

// NumberOfQuestionsGT applies the GT predicate on the "number_of_questions" field.
func NumberOfQuestionsGT(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldNumberOfQuestions, v))
}

// NumberOfQuestionsGTE applies the GTE predicate on the "number_of_questions" field.
func NumberOfQuestionsGTE(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldNumberOfQuestions, v))
}

// NumberOfQuestionsLT applies the LT predicate on the "number_of_questions" field.
func NumberOfQuestionsLT(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldNumberOfQuestions, v))
}

// NumberOfQuestionsLTE applies the LTE predicate on the "number_of_questions" field.
func NumberOfQuestionsLTE(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldNumberOfQuestions, v))
}

// DurationMinutesEQ applies the EQ predicate on the "duration_minutes" field.
func DurationMinutesEQ(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldDurationMinutes, v))
}

// DurationMinutesNEQ applies the NEQ predicate on the "duration_minutes" field.
func DurationMinutesNEQ(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldDurationMinutes, v))
}

// DurationMinutesIn applies the In predicate on the "duration_minutes" field.
func DurationMinutesIn(vs ...int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldDurationMinutes, vs...))
}

// DurationMinutesNotIn applies the NotIn predicate on the "duration_minutes" field.
func DurationMinutesNotIn(vs ...int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldDurationMinutes, vs...))
}

// DurationMinutesGT applies the GT predicate on the "duration_minutes" field.
func DurationMinutesGT(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldDurationMinutes, v))
}

// DurationMinutesGTE applies the GTE predicate on the "duration_minutes" field.
func DurationMinutesGTE(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldDurationMinutes, v))
}

// DurationMinutesLT applies the LT predicate on the "duration_minutes" field.
func DurationMinutesLT(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldDurationMinutes, v))
}

// DurationMinutesLTE applies the LTE predicate on the "duration_minutes" field.
func DurationMinutesLTE(v int) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldDurationMinutes, v))
}

// NegativeMarkingEQ applies the EQ predicate on the "negative_marking" field.
func NegativeMarkingEQ(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldNegativeMarking, v))
}

// NegativeMarkingNEQ applies the NEQ predicate on the "negative_marking" field.
func NegativeMarkingNEQ(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldNegativeMarking, v))
}

// NegativeMarkingIn applies the In predicate on the "negative_marking" field.
func NegativeMarkingIn(vs ...float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldNegativeMarking, vs...))
}

// NegativeMarkingNotIn applies the NotIn predicate on the "negative_marking" field.
func NegativeMarkingNotIn(vs ...float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldNegativeMarking, vs...))
}

// NegativeMarkingGT applies the GT predicate on the "negative_marking" field.
func NegativeMarkingGT(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldNegativeMarking, v))
}

// NegativeMarkingGTE applies the GTE predicate on the "negative_marking" field.
func NegativeMarkingGTE(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldNegativeMarking, v))
}

// NegativeMarkingLT applies the LT predicate on the "negative_marking" field.
func NegativeMarkingLT(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldNegativeMarking, v))
}

// NegativeMarkingLTE applies the LTE predicate on the "negative_marking" field.
func NegativeMarkingLTE(v float64) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldNegativeMarking, v))
}

// NegativeMarkingIsNil applies the IsNil predicate on the "negative_marking" field.
func NegativeMarkingIsNil() predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIsNull(FieldNegativeMarking))
}

// NegativeMarkingNotNil applies the NotNil predicate on the "negative_marking" field.
func NegativeMarkingNotNil() predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotNull(FieldNegativeMarking))
}

// OtherDetailsIsNil applies the IsNil predicate on the "other_details" field.
func OtherDetailsIsNil() predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIsNull(FieldOtherDetails))
}

// OtherDetailsNotNil applies the NotNil predicate on the "other_details" field.
func OtherDetailsNotNil() predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotNull(FieldOtherDetails))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ExamSetting {
	return predicate.ExamSetting(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasExam applies the HasEdge predicate on the "exam" edge.
func HasExam() predicate.ExamSetting {
	return predicate.ExamSetting(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ExamTable, ExamColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExamWith applies the HasEdge predicate on the "exam" edge with a given conditions (other predicates).
func HasExamWith(preds ...predicate.Exam) predicate.ExamSetting {
	return predicate.ExamSetting(func(s *sql.Selector) {
		step := newExamStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ExamSetting) predicate.ExamSetting {
	return predicate.ExamSetting(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ExamSetting) predicate.ExamSetting {
	return predicate.ExamSetting(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ExamSetting) predicate.ExamSetting {
	return predicate.ExamSetting(sql.NotPredicates(p))
}
