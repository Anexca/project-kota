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

// CachedQuestionMetadataDelete is the builder for deleting a CachedQuestionMetadata entity.
type CachedQuestionMetadataDelete struct {
	config
	hooks    []Hook
	mutation *CachedQuestionMetadataMutation
}

// Where appends a list predicates to the CachedQuestionMetadataDelete builder.
func (cqmd *CachedQuestionMetadataDelete) Where(ps ...predicate.CachedQuestionMetadata) *CachedQuestionMetadataDelete {
	cqmd.mutation.Where(ps...)
	return cqmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cqmd *CachedQuestionMetadataDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cqmd.sqlExec, cqmd.mutation, cqmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cqmd *CachedQuestionMetadataDelete) ExecX(ctx context.Context) int {
	n, err := cqmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cqmd *CachedQuestionMetadataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(cachedquestionmetadata.Table, sqlgraph.NewFieldSpec(cachedquestionmetadata.FieldID, field.TypeInt))
	if ps := cqmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cqmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cqmd.mutation.done = true
	return affected, err
}

// CachedQuestionMetadataDeleteOne is the builder for deleting a single CachedQuestionMetadata entity.
type CachedQuestionMetadataDeleteOne struct {
	cqmd *CachedQuestionMetadataDelete
}

// Where appends a list predicates to the CachedQuestionMetadataDelete builder.
func (cqmdo *CachedQuestionMetadataDeleteOne) Where(ps ...predicate.CachedQuestionMetadata) *CachedQuestionMetadataDeleteOne {
	cqmdo.cqmd.mutation.Where(ps...)
	return cqmdo
}

// Exec executes the deletion query.
func (cqmdo *CachedQuestionMetadataDeleteOne) Exec(ctx context.Context) error {
	n, err := cqmdo.cqmd.Exec(ctx)
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
func (cqmdo *CachedQuestionMetadataDeleteOne) ExecX(ctx context.Context) {
	if err := cqmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
