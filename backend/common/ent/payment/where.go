// Code generated by ent, DO NOT EDIT.

package payment

import (
	"common/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldID, id))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldAmount, v))
}

// PaymentDate applies equality check predicate on the "payment_date" field. It's identical to PaymentDateEQ.
func PaymentDate(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldPaymentDate, v))
}

// PaymentMethod applies equality check predicate on the "payment_method" field. It's identical to PaymentMethodEQ.
func PaymentMethod(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldPaymentMethod, v))
}

// PaymentPaymentID applies equality check predicate on the "payment_payment_id" field. It's identical to PaymentPaymentIDEQ.
func PaymentPaymentID(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldPaymentPaymentID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldUpdatedAt, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v int) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...int) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...int) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v int) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v int) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v int) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v int) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldAmount, v))
}

// PaymentDateEQ applies the EQ predicate on the "payment_date" field.
func PaymentDateEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldPaymentDate, v))
}

// PaymentDateNEQ applies the NEQ predicate on the "payment_date" field.
func PaymentDateNEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldPaymentDate, v))
}

// PaymentDateIn applies the In predicate on the "payment_date" field.
func PaymentDateIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldPaymentDate, vs...))
}

// PaymentDateNotIn applies the NotIn predicate on the "payment_date" field.
func PaymentDateNotIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldPaymentDate, vs...))
}

// PaymentDateGT applies the GT predicate on the "payment_date" field.
func PaymentDateGT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldPaymentDate, v))
}

// PaymentDateGTE applies the GTE predicate on the "payment_date" field.
func PaymentDateGTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldPaymentDate, v))
}

// PaymentDateLT applies the LT predicate on the "payment_date" field.
func PaymentDateLT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldPaymentDate, v))
}

// PaymentDateLTE applies the LTE predicate on the "payment_date" field.
func PaymentDateLTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldPaymentDate, v))
}

// PaymentStatusEQ applies the EQ predicate on the "payment_status" field.
func PaymentStatusEQ(v PaymentStatus) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldPaymentStatus, v))
}

// PaymentStatusNEQ applies the NEQ predicate on the "payment_status" field.
func PaymentStatusNEQ(v PaymentStatus) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldPaymentStatus, v))
}

// PaymentStatusIn applies the In predicate on the "payment_status" field.
func PaymentStatusIn(vs ...PaymentStatus) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldPaymentStatus, vs...))
}

// PaymentStatusNotIn applies the NotIn predicate on the "payment_status" field.
func PaymentStatusNotIn(vs ...PaymentStatus) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldPaymentStatus, vs...))
}

// PaymentMethodEQ applies the EQ predicate on the "payment_method" field.
func PaymentMethodEQ(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldPaymentMethod, v))
}

// PaymentMethodNEQ applies the NEQ predicate on the "payment_method" field.
func PaymentMethodNEQ(v string) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldPaymentMethod, v))
}

// PaymentMethodIn applies the In predicate on the "payment_method" field.
func PaymentMethodIn(vs ...string) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldPaymentMethod, vs...))
}

// PaymentMethodNotIn applies the NotIn predicate on the "payment_method" field.
func PaymentMethodNotIn(vs ...string) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldPaymentMethod, vs...))
}

// PaymentMethodGT applies the GT predicate on the "payment_method" field.
func PaymentMethodGT(v string) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldPaymentMethod, v))
}

// PaymentMethodGTE applies the GTE predicate on the "payment_method" field.
func PaymentMethodGTE(v string) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldPaymentMethod, v))
}

// PaymentMethodLT applies the LT predicate on the "payment_method" field.
func PaymentMethodLT(v string) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldPaymentMethod, v))
}

// PaymentMethodLTE applies the LTE predicate on the "payment_method" field.
func PaymentMethodLTE(v string) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldPaymentMethod, v))
}

// PaymentMethodContains applies the Contains predicate on the "payment_method" field.
func PaymentMethodContains(v string) predicate.Payment {
	return predicate.Payment(sql.FieldContains(FieldPaymentMethod, v))
}

// PaymentMethodHasPrefix applies the HasPrefix predicate on the "payment_method" field.
func PaymentMethodHasPrefix(v string) predicate.Payment {
	return predicate.Payment(sql.FieldHasPrefix(FieldPaymentMethod, v))
}

// PaymentMethodHasSuffix applies the HasSuffix predicate on the "payment_method" field.
func PaymentMethodHasSuffix(v string) predicate.Payment {
	return predicate.Payment(sql.FieldHasSuffix(FieldPaymentMethod, v))
}

// PaymentMethodEqualFold applies the EqualFold predicate on the "payment_method" field.
func PaymentMethodEqualFold(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEqualFold(FieldPaymentMethod, v))
}

// PaymentMethodContainsFold applies the ContainsFold predicate on the "payment_method" field.
func PaymentMethodContainsFold(v string) predicate.Payment {
	return predicate.Payment(sql.FieldContainsFold(FieldPaymentMethod, v))
}

// PaymentPaymentIDEQ applies the EQ predicate on the "payment_payment_id" field.
func PaymentPaymentIDEQ(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldPaymentPaymentID, v))
}

// PaymentPaymentIDNEQ applies the NEQ predicate on the "payment_payment_id" field.
func PaymentPaymentIDNEQ(v string) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldPaymentPaymentID, v))
}

// PaymentPaymentIDIn applies the In predicate on the "payment_payment_id" field.
func PaymentPaymentIDIn(vs ...string) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldPaymentPaymentID, vs...))
}

// PaymentPaymentIDNotIn applies the NotIn predicate on the "payment_payment_id" field.
func PaymentPaymentIDNotIn(vs ...string) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldPaymentPaymentID, vs...))
}

// PaymentPaymentIDGT applies the GT predicate on the "payment_payment_id" field.
func PaymentPaymentIDGT(v string) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldPaymentPaymentID, v))
}

// PaymentPaymentIDGTE applies the GTE predicate on the "payment_payment_id" field.
func PaymentPaymentIDGTE(v string) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldPaymentPaymentID, v))
}

// PaymentPaymentIDLT applies the LT predicate on the "payment_payment_id" field.
func PaymentPaymentIDLT(v string) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldPaymentPaymentID, v))
}

// PaymentPaymentIDLTE applies the LTE predicate on the "payment_payment_id" field.
func PaymentPaymentIDLTE(v string) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldPaymentPaymentID, v))
}

// PaymentPaymentIDContains applies the Contains predicate on the "payment_payment_id" field.
func PaymentPaymentIDContains(v string) predicate.Payment {
	return predicate.Payment(sql.FieldContains(FieldPaymentPaymentID, v))
}

// PaymentPaymentIDHasPrefix applies the HasPrefix predicate on the "payment_payment_id" field.
func PaymentPaymentIDHasPrefix(v string) predicate.Payment {
	return predicate.Payment(sql.FieldHasPrefix(FieldPaymentPaymentID, v))
}

// PaymentPaymentIDHasSuffix applies the HasSuffix predicate on the "payment_payment_id" field.
func PaymentPaymentIDHasSuffix(v string) predicate.Payment {
	return predicate.Payment(sql.FieldHasSuffix(FieldPaymentPaymentID, v))
}

// PaymentPaymentIDEqualFold applies the EqualFold predicate on the "payment_payment_id" field.
func PaymentPaymentIDEqualFold(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEqualFold(FieldPaymentPaymentID, v))
}

// PaymentPaymentIDContainsFold applies the ContainsFold predicate on the "payment_payment_id" field.
func PaymentPaymentIDContainsFold(v string) predicate.Payment {
	return predicate.Payment(sql.FieldContainsFold(FieldPaymentPaymentID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubscription applies the HasEdge predicate on the "subscription" edge.
func HasSubscription() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SubscriptionTable, SubscriptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubscriptionWith applies the HasEdge predicate on the "subscription" edge with a given conditions (other predicates).
func HasSubscriptionWith(preds ...predicate.UserSubscription) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := newSubscriptionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Payment) predicate.Payment {
	return predicate.Payment(sql.NotPredicates(p))
}
