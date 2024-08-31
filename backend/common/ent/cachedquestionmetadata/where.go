// Code generated by ent, DO NOT EDIT.

package cachedquestionmetadata

import (
	"common/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldLTE(FieldID, id))
}

// CacheUID applies equality check predicate on the "cache_uid" field. It's identical to CacheUIDEQ.
func CacheUID(v string) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldEQ(FieldCacheUID, v))
}

// IsUsed applies equality check predicate on the "is_used" field. It's identical to IsUsedEQ.
func IsUsed(v bool) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldEQ(FieldIsUsed, v))
}

// ExpiresAt applies equality check predicate on the "expires_at" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldEQ(FieldExpiresAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldEQ(FieldUpdatedAt, v))
}

// CacheUIDEQ applies the EQ predicate on the "cache_uid" field.
func CacheUIDEQ(v string) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldEQ(FieldCacheUID, v))
}

// CacheUIDNEQ applies the NEQ predicate on the "cache_uid" field.
func CacheUIDNEQ(v string) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldNEQ(FieldCacheUID, v))
}

// CacheUIDIn applies the In predicate on the "cache_uid" field.
func CacheUIDIn(vs ...string) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldIn(FieldCacheUID, vs...))
}

// CacheUIDNotIn applies the NotIn predicate on the "cache_uid" field.
func CacheUIDNotIn(vs ...string) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldNotIn(FieldCacheUID, vs...))
}

// CacheUIDGT applies the GT predicate on the "cache_uid" field.
func CacheUIDGT(v string) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldGT(FieldCacheUID, v))
}

// CacheUIDGTE applies the GTE predicate on the "cache_uid" field.
func CacheUIDGTE(v string) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldGTE(FieldCacheUID, v))
}

// CacheUIDLT applies the LT predicate on the "cache_uid" field.
func CacheUIDLT(v string) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldLT(FieldCacheUID, v))
}

// CacheUIDLTE applies the LTE predicate on the "cache_uid" field.
func CacheUIDLTE(v string) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldLTE(FieldCacheUID, v))
}

// CacheUIDContains applies the Contains predicate on the "cache_uid" field.
func CacheUIDContains(v string) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldContains(FieldCacheUID, v))
}

// CacheUIDHasPrefix applies the HasPrefix predicate on the "cache_uid" field.
func CacheUIDHasPrefix(v string) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldHasPrefix(FieldCacheUID, v))
}

// CacheUIDHasSuffix applies the HasSuffix predicate on the "cache_uid" field.
func CacheUIDHasSuffix(v string) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldHasSuffix(FieldCacheUID, v))
}

// CacheUIDEqualFold applies the EqualFold predicate on the "cache_uid" field.
func CacheUIDEqualFold(v string) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldEqualFold(FieldCacheUID, v))
}

// CacheUIDContainsFold applies the ContainsFold predicate on the "cache_uid" field.
func CacheUIDContainsFold(v string) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldContainsFold(FieldCacheUID, v))
}

// IsUsedEQ applies the EQ predicate on the "is_used" field.
func IsUsedEQ(v bool) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldEQ(FieldIsUsed, v))
}

// IsUsedNEQ applies the NEQ predicate on the "is_used" field.
func IsUsedNEQ(v bool) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldNEQ(FieldIsUsed, v))
}

// ExpiresAtEQ applies the EQ predicate on the "expires_at" field.
func ExpiresAtEQ(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldEQ(FieldExpiresAt, v))
}

// ExpiresAtNEQ applies the NEQ predicate on the "expires_at" field.
func ExpiresAtNEQ(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldNEQ(FieldExpiresAt, v))
}

// ExpiresAtIn applies the In predicate on the "expires_at" field.
func ExpiresAtIn(vs ...time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldIn(FieldExpiresAt, vs...))
}

// ExpiresAtNotIn applies the NotIn predicate on the "expires_at" field.
func ExpiresAtNotIn(vs ...time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldNotIn(FieldExpiresAt, vs...))
}

// ExpiresAtGT applies the GT predicate on the "expires_at" field.
func ExpiresAtGT(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldGT(FieldExpiresAt, v))
}

// ExpiresAtGTE applies the GTE predicate on the "expires_at" field.
func ExpiresAtGTE(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldGTE(FieldExpiresAt, v))
}

// ExpiresAtLT applies the LT predicate on the "expires_at" field.
func ExpiresAtLT(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldLT(FieldExpiresAt, v))
}

// ExpiresAtLTE applies the LTE predicate on the "expires_at" field.
func ExpiresAtLTE(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldLTE(FieldExpiresAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasExam applies the HasEdge predicate on the "exam" edge.
func HasExam() predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ExamTable, ExamColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExamWith applies the HasEdge predicate on the "exam" edge with a given conditions (other predicates).
func HasExamWith(preds ...predicate.Exam) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(func(s *sql.Selector) {
		step := newExamStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CachedQuestionMetaData) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CachedQuestionMetaData) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CachedQuestionMetaData) predicate.CachedQuestionMetaData {
	return predicate.CachedQuestionMetaData(sql.NotPredicates(p))
}
