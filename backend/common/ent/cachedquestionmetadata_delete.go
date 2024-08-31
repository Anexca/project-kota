// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/cachedquestionmetadata"
	"common/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CachedQuestionMetaDataDelete is the builder for deleting a CachedQuestionMetaData entity.
type CachedQuestionMetaDataDelete struct {
	config
	hooks    []Hook
	mutation *CachedQuestionMetaDataMutation
}

// Where appends a list predicates to the CachedQuestionMetaDataDelete builder.
func (cqmdd *CachedQuestionMetaDataDelete) Where(ps ...predicate.CachedQuestionMetaData) *CachedQuestionMetaDataDelete {
	cqmdd.mutation.Where(ps...)
	return cqmdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cqmdd *CachedQuestionMetaDataDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cqmdd.sqlExec, cqmdd.mutation, cqmdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cqmdd *CachedQuestionMetaDataDelete) ExecX(ctx context.Context) int {
	n, err := cqmdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cqmdd *CachedQuestionMetaDataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(cachedquestionmetadata.Table, sqlgraph.NewFieldSpec(cachedquestionmetadata.FieldID, field.TypeInt))
	if ps := cqmdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cqmdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cqmdd.mutation.done = true
	return affected, err
}

// CachedQuestionMetaDataDeleteOne is the builder for deleting a single CachedQuestionMetaData entity.
type CachedQuestionMetaDataDeleteOne struct {
	cqmdd *CachedQuestionMetaDataDelete
}

// Where appends a list predicates to the CachedQuestionMetaDataDelete builder.
func (cqmddo *CachedQuestionMetaDataDeleteOne) Where(ps ...predicate.CachedQuestionMetaData) *CachedQuestionMetaDataDeleteOne {
	cqmddo.cqmdd.mutation.Where(ps...)
	return cqmddo
}

// Exec executes the deletion query.
func (cqmddo *CachedQuestionMetaDataDeleteOne) Exec(ctx context.Context) error {
	n, err := cqmddo.cqmdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cachedquestionmetadata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cqmddo *CachedQuestionMetaDataDeleteOne) ExecX(ctx context.Context) {
	if err := cqmddo.Exec(ctx); err != nil {
		panic(err)
	}
}
