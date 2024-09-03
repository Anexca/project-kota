// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/examassesment"
	"common/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamAssesmentDelete is the builder for deleting a ExamAssesment entity.
type ExamAssesmentDelete struct {
	config
	hooks    []Hook
	mutation *ExamAssesmentMutation
}

// Where appends a list predicates to the ExamAssesmentDelete builder.
func (ead *ExamAssesmentDelete) Where(ps ...predicate.ExamAssesment) *ExamAssesmentDelete {
	ead.mutation.Where(ps...)
	return ead
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ead *ExamAssesmentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ead.sqlExec, ead.mutation, ead.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ead *ExamAssesmentDelete) ExecX(ctx context.Context) int {
	n, err := ead.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ead *ExamAssesmentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(examassesment.Table, sqlgraph.NewFieldSpec(examassesment.FieldID, field.TypeInt))
	if ps := ead.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ead.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ead.mutation.done = true
	return affected, err
}

// ExamAssesmentDeleteOne is the builder for deleting a single ExamAssesment entity.
type ExamAssesmentDeleteOne struct {
	ead *ExamAssesmentDelete
}

// Where appends a list predicates to the ExamAssesmentDelete builder.
func (eado *ExamAssesmentDeleteOne) Where(ps ...predicate.ExamAssesment) *ExamAssesmentDeleteOne {
	eado.ead.mutation.Where(ps...)
	return eado
}

// Exec executes the deletion query.
func (eado *ExamAssesmentDeleteOne) Exec(ctx context.Context) error {
	n, err := eado.ead.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{examassesment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eado *ExamAssesmentDeleteOne) ExecX(ctx context.Context) {
	if err := eado.Exec(ctx); err != nil {
		panic(err)
	}
}
