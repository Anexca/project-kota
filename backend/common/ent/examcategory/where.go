// Code generated by ent, DO NOT EDIT.

package examcategory

import (
	"common/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldLTE(FieldID, id))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldEQ(FieldDescription, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldEQ(FieldIsActive, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v Name) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v Name) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...Name) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...Name) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldNotIn(FieldName, vs...))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldContainsFold(FieldDescription, v))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldNEQ(FieldIsActive, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ExamCategory {
	return predicate.ExamCategory(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasExams applies the HasEdge predicate on the "exams" edge.
func HasExams() predicate.ExamCategory {
	return predicate.ExamCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ExamsTable, ExamsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExamsWith applies the HasEdge predicate on the "exams" edge with a given conditions (other predicates).
func HasExamsWith(preds ...predicate.Exam) predicate.ExamCategory {
	return predicate.ExamCategory(func(s *sql.Selector) {
		step := newExamsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() predicate.ExamCategory {
	return predicate.ExamCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GroupsTable, GroupsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.ExamGroup) predicate.ExamCategory {
	return predicate.ExamCategory(func(s *sql.Selector) {
		step := newGroupsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ExamCategory) predicate.ExamCategory {
	return predicate.ExamCategory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ExamCategory) predicate.ExamCategory {
	return predicate.ExamCategory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ExamCategory) predicate.ExamCategory {
	return predicate.ExamCategory(sql.NotPredicates(p))
}
