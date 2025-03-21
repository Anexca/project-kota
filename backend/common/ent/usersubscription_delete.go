// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/predicate"
	"common/ent/usersubscription"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserSubscriptionDelete is the builder for deleting a UserSubscription entity.
type UserSubscriptionDelete struct {
	config
	hooks    []Hook
	mutation *UserSubscriptionMutation
}

// Where appends a list predicates to the UserSubscriptionDelete builder.
func (usd *UserSubscriptionDelete) Where(ps ...predicate.UserSubscription) *UserSubscriptionDelete {
	usd.mutation.Where(ps...)
	return usd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (usd *UserSubscriptionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, usd.sqlExec, usd.mutation, usd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (usd *UserSubscriptionDelete) ExecX(ctx context.Context) int {
	n, err := usd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (usd *UserSubscriptionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(usersubscription.Table, sqlgraph.NewFieldSpec(usersubscription.FieldID, field.TypeInt))
	if ps := usd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, usd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	usd.mutation.done = true
	return affected, err
}

// UserSubscriptionDeleteOne is the builder for deleting a single UserSubscription entity.
type UserSubscriptionDeleteOne struct {
	usd *UserSubscriptionDelete
}

// Where appends a list predicates to the UserSubscriptionDelete builder.
func (usdo *UserSubscriptionDeleteOne) Where(ps ...predicate.UserSubscription) *UserSubscriptionDeleteOne {
	usdo.usd.mutation.Where(ps...)
	return usdo
}

// Exec executes the deletion query.
func (usdo *UserSubscriptionDeleteOne) Exec(ctx context.Context) error {
	n, err := usdo.usd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usersubscription.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (usdo *UserSubscriptionDeleteOne) ExecX(ctx context.Context) {
	if err := usdo.Exec(ctx); err != nil {
		panic(err)
	}
}
