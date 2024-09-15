// Code generated by ent, DO NOT EDIT.

package subscription

import (
	"common/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Subscription {
	return predicate.Subscription(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Subscription {
	return predicate.Subscription(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Subscription {
	return predicate.Subscription(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Subscription {
	return predicate.Subscription(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Subscription {
	return predicate.Subscription(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Subscription {
	return predicate.Subscription(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Subscription {
	return predicate.Subscription(sql.FieldLTE(FieldID, id))
}

// ProviderPlanID applies equality check predicate on the "provider_plan_id" field. It's identical to ProviderPlanIDEQ.
func ProviderPlanID(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldProviderPlanID, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldPrice, v))
}

// DurationInMonths applies equality check predicate on the "duration_in_months" field. It's identical to DurationInMonthsEQ.
func DurationInMonths(v int) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldDurationInMonths, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldIsActive, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldName, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldUpdatedAt, v))
}

// ProviderPlanIDEQ applies the EQ predicate on the "provider_plan_id" field.
func ProviderPlanIDEQ(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldProviderPlanID, v))
}

// ProviderPlanIDNEQ applies the NEQ predicate on the "provider_plan_id" field.
func ProviderPlanIDNEQ(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldNEQ(FieldProviderPlanID, v))
}

// ProviderPlanIDIn applies the In predicate on the "provider_plan_id" field.
func ProviderPlanIDIn(vs ...string) predicate.Subscription {
	return predicate.Subscription(sql.FieldIn(FieldProviderPlanID, vs...))
}

// ProviderPlanIDNotIn applies the NotIn predicate on the "provider_plan_id" field.
func ProviderPlanIDNotIn(vs ...string) predicate.Subscription {
	return predicate.Subscription(sql.FieldNotIn(FieldProviderPlanID, vs...))
}

// ProviderPlanIDGT applies the GT predicate on the "provider_plan_id" field.
func ProviderPlanIDGT(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldGT(FieldProviderPlanID, v))
}

// ProviderPlanIDGTE applies the GTE predicate on the "provider_plan_id" field.
func ProviderPlanIDGTE(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldGTE(FieldProviderPlanID, v))
}

// ProviderPlanIDLT applies the LT predicate on the "provider_plan_id" field.
func ProviderPlanIDLT(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldLT(FieldProviderPlanID, v))
}

// ProviderPlanIDLTE applies the LTE predicate on the "provider_plan_id" field.
func ProviderPlanIDLTE(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldLTE(FieldProviderPlanID, v))
}

// ProviderPlanIDContains applies the Contains predicate on the "provider_plan_id" field.
func ProviderPlanIDContains(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldContains(FieldProviderPlanID, v))
}

// ProviderPlanIDHasPrefix applies the HasPrefix predicate on the "provider_plan_id" field.
func ProviderPlanIDHasPrefix(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldHasPrefix(FieldProviderPlanID, v))
}

// ProviderPlanIDHasSuffix applies the HasSuffix predicate on the "provider_plan_id" field.
func ProviderPlanIDHasSuffix(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldHasSuffix(FieldProviderPlanID, v))
}

// ProviderPlanIDEqualFold applies the EqualFold predicate on the "provider_plan_id" field.
func ProviderPlanIDEqualFold(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldEqualFold(FieldProviderPlanID, v))
}

// ProviderPlanIDContainsFold applies the ContainsFold predicate on the "provider_plan_id" field.
func ProviderPlanIDContainsFold(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldContainsFold(FieldProviderPlanID, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int) predicate.Subscription {
	return predicate.Subscription(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int) predicate.Subscription {
	return predicate.Subscription(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...int) predicate.Subscription {
	return predicate.Subscription(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v int) predicate.Subscription {
	return predicate.Subscription(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int) predicate.Subscription {
	return predicate.Subscription(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int) predicate.Subscription {
	return predicate.Subscription(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int) predicate.Subscription {
	return predicate.Subscription(sql.FieldLTE(FieldPrice, v))
}

// DurationInMonthsEQ applies the EQ predicate on the "duration_in_months" field.
func DurationInMonthsEQ(v int) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldDurationInMonths, v))
}

// DurationInMonthsNEQ applies the NEQ predicate on the "duration_in_months" field.
func DurationInMonthsNEQ(v int) predicate.Subscription {
	return predicate.Subscription(sql.FieldNEQ(FieldDurationInMonths, v))
}

// DurationInMonthsIn applies the In predicate on the "duration_in_months" field.
func DurationInMonthsIn(vs ...int) predicate.Subscription {
	return predicate.Subscription(sql.FieldIn(FieldDurationInMonths, vs...))
}

// DurationInMonthsNotIn applies the NotIn predicate on the "duration_in_months" field.
func DurationInMonthsNotIn(vs ...int) predicate.Subscription {
	return predicate.Subscription(sql.FieldNotIn(FieldDurationInMonths, vs...))
}

// DurationInMonthsGT applies the GT predicate on the "duration_in_months" field.
func DurationInMonthsGT(v int) predicate.Subscription {
	return predicate.Subscription(sql.FieldGT(FieldDurationInMonths, v))
}

// DurationInMonthsGTE applies the GTE predicate on the "duration_in_months" field.
func DurationInMonthsGTE(v int) predicate.Subscription {
	return predicate.Subscription(sql.FieldGTE(FieldDurationInMonths, v))
}

// DurationInMonthsLT applies the LT predicate on the "duration_in_months" field.
func DurationInMonthsLT(v int) predicate.Subscription {
	return predicate.Subscription(sql.FieldLT(FieldDurationInMonths, v))
}

// DurationInMonthsLTE applies the LTE predicate on the "duration_in_months" field.
func DurationInMonthsLTE(v int) predicate.Subscription {
	return predicate.Subscription(sql.FieldLTE(FieldDurationInMonths, v))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.Subscription {
	return predicate.Subscription(sql.FieldNEQ(FieldIsActive, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Subscription {
	return predicate.Subscription(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Subscription {
	return predicate.Subscription(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Subscription {
	return predicate.Subscription(sql.FieldContainsFold(FieldName, v))
}

// RawSubscriptionDataIsNil applies the IsNil predicate on the "raw_subscription_data" field.
func RawSubscriptionDataIsNil() predicate.Subscription {
	return predicate.Subscription(sql.FieldIsNull(FieldRawSubscriptionData))
}

// RawSubscriptionDataNotNil applies the NotNil predicate on the "raw_subscription_data" field.
func RawSubscriptionDataNotNil() predicate.Subscription {
	return predicate.Subscription(sql.FieldNotNull(FieldRawSubscriptionData))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Subscription {
	return predicate.Subscription(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasExams applies the HasEdge predicate on the "exams" edge.
func HasExams() predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ExamsTable, ExamsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExamsWith applies the HasEdge predicate on the "exams" edge with a given conditions (other predicates).
func HasExamsWith(preds ...predicate.SubscriptionExam) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		step := newExamsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserSubscriptions applies the HasEdge predicate on the "user_subscriptions" edge.
func HasUserSubscriptions() predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserSubscriptionsTable, UserSubscriptionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserSubscriptionsWith applies the HasEdge predicate on the "user_subscriptions" edge with a given conditions (other predicates).
func HasUserSubscriptionsWith(preds ...predicate.UserSubscription) predicate.Subscription {
	return predicate.Subscription(func(s *sql.Selector) {
		step := newUserSubscriptionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Subscription) predicate.Subscription {
	return predicate.Subscription(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Subscription) predicate.Subscription {
	return predicate.Subscription(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Subscription) predicate.Subscription {
	return predicate.Subscription(sql.NotPredicates(p))
}
