// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/exam"
	"common/ent/examcategory"
	"common/ent/examgroup"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamCategoryCreate is the builder for creating a ExamCategory entity.
type ExamCategoryCreate struct {
	config
	mutation *ExamCategoryMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ecc *ExamCategoryCreate) SetName(e examcategory.Name) *ExamCategoryCreate {
	ecc.mutation.SetName(e)
	return ecc
}

// SetDescription sets the "description" field.
func (ecc *ExamCategoryCreate) SetDescription(s string) *ExamCategoryCreate {
	ecc.mutation.SetDescription(s)
	return ecc
}

// SetIsActive sets the "is_active" field.
func (ecc *ExamCategoryCreate) SetIsActive(b bool) *ExamCategoryCreate {
	ecc.mutation.SetIsActive(b)
	return ecc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (ecc *ExamCategoryCreate) SetNillableIsActive(b *bool) *ExamCategoryCreate {
	if b != nil {
		ecc.SetIsActive(*b)
	}
	return ecc
}

// SetCreatedAt sets the "created_at" field.
func (ecc *ExamCategoryCreate) SetCreatedAt(t time.Time) *ExamCategoryCreate {
	ecc.mutation.SetCreatedAt(t)
	return ecc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ecc *ExamCategoryCreate) SetNillableCreatedAt(t *time.Time) *ExamCategoryCreate {
	if t != nil {
		ecc.SetCreatedAt(*t)
	}
	return ecc
}

// SetUpdatedAt sets the "updated_at" field.
func (ecc *ExamCategoryCreate) SetUpdatedAt(t time.Time) *ExamCategoryCreate {
	ecc.mutation.SetUpdatedAt(t)
	return ecc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ecc *ExamCategoryCreate) SetNillableUpdatedAt(t *time.Time) *ExamCategoryCreate {
	if t != nil {
		ecc.SetUpdatedAt(*t)
	}
	return ecc
}

// AddExamIDs adds the "exams" edge to the Exam entity by IDs.
func (ecc *ExamCategoryCreate) AddExamIDs(ids ...int) *ExamCategoryCreate {
	ecc.mutation.AddExamIDs(ids...)
	return ecc
}

// AddExams adds the "exams" edges to the Exam entity.
func (ecc *ExamCategoryCreate) AddExams(e ...*Exam) *ExamCategoryCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecc.AddExamIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the ExamGroup entity by IDs.
func (ecc *ExamCategoryCreate) AddGroupIDs(ids ...int) *ExamCategoryCreate {
	ecc.mutation.AddGroupIDs(ids...)
	return ecc
}

// AddGroups adds the "groups" edges to the ExamGroup entity.
func (ecc *ExamCategoryCreate) AddGroups(e ...*ExamGroup) *ExamCategoryCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecc.AddGroupIDs(ids...)
}

// Mutation returns the ExamCategoryMutation object of the builder.
func (ecc *ExamCategoryCreate) Mutation() *ExamCategoryMutation {
	return ecc.mutation
}

// Save creates the ExamCategory in the database.
func (ecc *ExamCategoryCreate) Save(ctx context.Context) (*ExamCategory, error) {
	ecc.defaults()
	return withHooks(ctx, ecc.sqlSave, ecc.mutation, ecc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ecc *ExamCategoryCreate) SaveX(ctx context.Context) *ExamCategory {
	v, err := ecc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecc *ExamCategoryCreate) Exec(ctx context.Context) error {
	_, err := ecc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecc *ExamCategoryCreate) ExecX(ctx context.Context) {
	if err := ecc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ecc *ExamCategoryCreate) defaults() {
	if _, ok := ecc.mutation.IsActive(); !ok {
		v := examcategory.DefaultIsActive
		ecc.mutation.SetIsActive(v)
	}
	if _, ok := ecc.mutation.CreatedAt(); !ok {
		v := examcategory.DefaultCreatedAt()
		ecc.mutation.SetCreatedAt(v)
	}
	if _, ok := ecc.mutation.UpdatedAt(); !ok {
		v := examcategory.DefaultUpdatedAt()
		ecc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecc *ExamCategoryCreate) check() error {
	if _, ok := ecc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ExamCategory.name"`)}
	}
	if v, ok := ecc.mutation.Name(); ok {
		if err := examcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ExamCategory.name": %w`, err)}
		}
	}
	if _, ok := ecc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "ExamCategory.description"`)}
	}
	if _, ok := ecc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "ExamCategory.is_active"`)}
	}
	if _, ok := ecc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ExamCategory.created_at"`)}
	}
	if _, ok := ecc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ExamCategory.updated_at"`)}
	}
	return nil
}

func (ecc *ExamCategoryCreate) sqlSave(ctx context.Context) (*ExamCategory, error) {
	if err := ecc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ecc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ecc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ecc.mutation.id = &_node.ID
	ecc.mutation.done = true
	return _node, nil
}

func (ecc *ExamCategoryCreate) createSpec() (*ExamCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &ExamCategory{config: ecc.config}
		_spec = sqlgraph.NewCreateSpec(examcategory.Table, sqlgraph.NewFieldSpec(examcategory.FieldID, field.TypeInt))
	)
	if value, ok := ecc.mutation.Name(); ok {
		_spec.SetField(examcategory.FieldName, field.TypeEnum, value)
		_node.Name = value
	}
	if value, ok := ecc.mutation.Description(); ok {
		_spec.SetField(examcategory.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ecc.mutation.IsActive(); ok {
		_spec.SetField(examcategory.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := ecc.mutation.CreatedAt(); ok {
		_spec.SetField(examcategory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ecc.mutation.UpdatedAt(); ok {
		_spec.SetField(examcategory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := ecc.mutation.ExamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examcategory.ExamsTable,
			Columns: []string{examcategory.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ecc.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examcategory.GroupsTable,
			Columns: []string{examcategory.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ExamCategoryCreateBulk is the builder for creating many ExamCategory entities in bulk.
type ExamCategoryCreateBulk struct {
	config
	err      error
	builders []*ExamCategoryCreate
}

// Save creates the ExamCategory entities in the database.
func (eccb *ExamCategoryCreateBulk) Save(ctx context.Context) ([]*ExamCategory, error) {
	if eccb.err != nil {
		return nil, eccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(eccb.builders))
	nodes := make([]*ExamCategory, len(eccb.builders))
	mutators := make([]Mutator, len(eccb.builders))
	for i := range eccb.builders {
		func(i int, root context.Context) {
			builder := eccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ExamCategoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, eccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, eccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eccb *ExamCategoryCreateBulk) SaveX(ctx context.Context) []*ExamCategory {
	v, err := eccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eccb *ExamCategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := eccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eccb *ExamCategoryCreateBulk) ExecX(ctx context.Context) {
	if err := eccb.Exec(ctx); err != nil {
		panic(err)
	}
}
