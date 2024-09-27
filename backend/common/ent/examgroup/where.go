// Code generated by ent, DO NOT EDIT.

package examgroup

import (
	"common/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEQ(FieldDescription, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEQ(FieldIsActive, v))
}

// LogoURL applies equality check predicate on the "logo_url" field. It's identical to LogoURLEQ.
func LogoURL(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEQ(FieldLogoURL, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldContainsFold(FieldDescription, v))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldNEQ(FieldIsActive, v))
}

// LogoURLEQ applies the EQ predicate on the "logo_url" field.
func LogoURLEQ(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEQ(FieldLogoURL, v))
}

// LogoURLNEQ applies the NEQ predicate on the "logo_url" field.
func LogoURLNEQ(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldNEQ(FieldLogoURL, v))
}

// LogoURLIn applies the In predicate on the "logo_url" field.
func LogoURLIn(vs ...string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldIn(FieldLogoURL, vs...))
}

// LogoURLNotIn applies the NotIn predicate on the "logo_url" field.
func LogoURLNotIn(vs ...string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldNotIn(FieldLogoURL, vs...))
}

// LogoURLGT applies the GT predicate on the "logo_url" field.
func LogoURLGT(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldGT(FieldLogoURL, v))
}

// LogoURLGTE applies the GTE predicate on the "logo_url" field.
func LogoURLGTE(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldGTE(FieldLogoURL, v))
}

// LogoURLLT applies the LT predicate on the "logo_url" field.
func LogoURLLT(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldLT(FieldLogoURL, v))
}

// LogoURLLTE applies the LTE predicate on the "logo_url" field.
func LogoURLLTE(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldLTE(FieldLogoURL, v))
}

// LogoURLContains applies the Contains predicate on the "logo_url" field.
func LogoURLContains(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldContains(FieldLogoURL, v))
}

// LogoURLHasPrefix applies the HasPrefix predicate on the "logo_url" field.
func LogoURLHasPrefix(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldHasPrefix(FieldLogoURL, v))
}

// LogoURLHasSuffix applies the HasSuffix predicate on the "logo_url" field.
func LogoURLHasSuffix(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldHasSuffix(FieldLogoURL, v))
}

// LogoURLIsNil applies the IsNil predicate on the "logo_url" field.
func LogoURLIsNil() predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldIsNull(FieldLogoURL))
}

// LogoURLNotNil applies the NotNil predicate on the "logo_url" field.
func LogoURLNotNil() predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldNotNull(FieldLogoURL))
}

// LogoURLEqualFold applies the EqualFold predicate on the "logo_url" field.
func LogoURLEqualFold(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEqualFold(FieldLogoURL, v))
}

// LogoURLContainsFold applies the ContainsFold predicate on the "logo_url" field.
func LogoURLContainsFold(v string) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldContainsFold(FieldLogoURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ExamGroup {
	return predicate.ExamGroup(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.ExamGroup {
	return predicate.ExamGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.ExamCategory) predicate.ExamGroup {
	return predicate.ExamGroup(func(s *sql.Selector) {
		step := newCategoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ExamGroup) predicate.ExamGroup {
	return predicate.ExamGroup(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ExamGroup) predicate.ExamGroup {
	return predicate.ExamGroup(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ExamGroup) predicate.ExamGroup {
	return predicate.ExamGroup(sql.NotPredicates(p))
}
