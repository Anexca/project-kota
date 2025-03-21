// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/cachedexam"
	"common/ent/exam"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CachedExamCreate is the builder for creating a CachedExam entity.
type CachedExamCreate struct {
	config
	mutation *CachedExamMutation
	hooks    []Hook
}

// SetCacheUID sets the "cache_uid" field.
func (cec *CachedExamCreate) SetCacheUID(s string) *CachedExamCreate {
	cec.mutation.SetCacheUID(s)
	return cec
}

// SetIsUsed sets the "is_used" field.
func (cec *CachedExamCreate) SetIsUsed(b bool) *CachedExamCreate {
	cec.mutation.SetIsUsed(b)
	return cec
}

// SetNillableIsUsed sets the "is_used" field if the given value is not nil.
func (cec *CachedExamCreate) SetNillableIsUsed(b *bool) *CachedExamCreate {
	if b != nil {
		cec.SetIsUsed(*b)
	}
	return cec
}

// SetExpiresAt sets the "expires_at" field.
func (cec *CachedExamCreate) SetExpiresAt(t time.Time) *CachedExamCreate {
	cec.mutation.SetExpiresAt(t)
	return cec
}

// SetCreatedAt sets the "created_at" field.
func (cec *CachedExamCreate) SetCreatedAt(t time.Time) *CachedExamCreate {
	cec.mutation.SetCreatedAt(t)
	return cec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cec *CachedExamCreate) SetNillableCreatedAt(t *time.Time) *CachedExamCreate {
	if t != nil {
		cec.SetCreatedAt(*t)
	}
	return cec
}

// SetUpdatedAt sets the "updated_at" field.
func (cec *CachedExamCreate) SetUpdatedAt(t time.Time) *CachedExamCreate {
	cec.mutation.SetUpdatedAt(t)
	return cec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cec *CachedExamCreate) SetNillableUpdatedAt(t *time.Time) *CachedExamCreate {
	if t != nil {
		cec.SetUpdatedAt(*t)
	}
	return cec
}

// SetExamID sets the "exam" edge to the Exam entity by ID.
func (cec *CachedExamCreate) SetExamID(id int) *CachedExamCreate {
	cec.mutation.SetExamID(id)
	return cec
}

// SetExam sets the "exam" edge to the Exam entity.
func (cec *CachedExamCreate) SetExam(e *Exam) *CachedExamCreate {
	return cec.SetExamID(e.ID)
}

// Mutation returns the CachedExamMutation object of the builder.
func (cec *CachedExamCreate) Mutation() *CachedExamMutation {
	return cec.mutation
}

// Save creates the CachedExam in the database.
func (cec *CachedExamCreate) Save(ctx context.Context) (*CachedExam, error) {
	cec.defaults()
	return withHooks(ctx, cec.sqlSave, cec.mutation, cec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cec *CachedExamCreate) SaveX(ctx context.Context) *CachedExam {
	v, err := cec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cec *CachedExamCreate) Exec(ctx context.Context) error {
	_, err := cec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cec *CachedExamCreate) ExecX(ctx context.Context) {
	if err := cec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cec *CachedExamCreate) defaults() {
	if _, ok := cec.mutation.IsUsed(); !ok {
		v := cachedexam.DefaultIsUsed
		cec.mutation.SetIsUsed(v)
	}
	if _, ok := cec.mutation.CreatedAt(); !ok {
		v := cachedexam.DefaultCreatedAt()
		cec.mutation.SetCreatedAt(v)
	}
	if _, ok := cec.mutation.UpdatedAt(); !ok {
		v := cachedexam.DefaultUpdatedAt()
		cec.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cec *CachedExamCreate) check() error {
	if _, ok := cec.mutation.CacheUID(); !ok {
		return &ValidationError{Name: "cache_uid", err: errors.New(`ent: missing required field "CachedExam.cache_uid"`)}
	}
	if _, ok := cec.mutation.IsUsed(); !ok {
		return &ValidationError{Name: "is_used", err: errors.New(`ent: missing required field "CachedExam.is_used"`)}
	}
	if _, ok := cec.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`ent: missing required field "CachedExam.expires_at"`)}
	}
	if _, ok := cec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CachedExam.created_at"`)}
	}
	if _, ok := cec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CachedExam.updated_at"`)}
	}
	if len(cec.mutation.ExamIDs()) == 0 {
		return &ValidationError{Name: "exam", err: errors.New(`ent: missing required edge "CachedExam.exam"`)}
	}
	return nil
}

func (cec *CachedExamCreate) sqlSave(ctx context.Context) (*CachedExam, error) {
	if err := cec.check(); err != nil {
		return nil, err
	}
	_node, _spec := cec.createSpec()
	if err := sqlgraph.CreateNode(ctx, cec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cec.mutation.id = &_node.ID
	cec.mutation.done = true
	return _node, nil
}

func (cec *CachedExamCreate) createSpec() (*CachedExam, *sqlgraph.CreateSpec) {
	var (
		_node = &CachedExam{config: cec.config}
		_spec = sqlgraph.NewCreateSpec(cachedexam.Table, sqlgraph.NewFieldSpec(cachedexam.FieldID, field.TypeInt))
	)
	if value, ok := cec.mutation.CacheUID(); ok {
		_spec.SetField(cachedexam.FieldCacheUID, field.TypeString, value)
		_node.CacheUID = value
	}
	if value, ok := cec.mutation.IsUsed(); ok {
		_spec.SetField(cachedexam.FieldIsUsed, field.TypeBool, value)
		_node.IsUsed = value
	}
	if value, ok := cec.mutation.ExpiresAt(); ok {
		_spec.SetField(cachedexam.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if value, ok := cec.mutation.CreatedAt(); ok {
		_spec.SetField(cachedexam.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cec.mutation.UpdatedAt(); ok {
		_spec.SetField(cachedexam.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := cec.mutation.ExamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cachedexam.ExamTable,
			Columns: []string{cachedexam.ExamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.exam_cached_exam = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CachedExamCreateBulk is the builder for creating many CachedExam entities in bulk.
type CachedExamCreateBulk struct {
	config
	err      error
	builders []*CachedExamCreate
}

// Save creates the CachedExam entities in the database.
func (cecb *CachedExamCreateBulk) Save(ctx context.Context) ([]*CachedExam, error) {
	if cecb.err != nil {
		return nil, cecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cecb.builders))
	nodes := make([]*CachedExam, len(cecb.builders))
	mutators := make([]Mutator, len(cecb.builders))
	for i := range cecb.builders {
		func(i int, root context.Context) {
			builder := cecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CachedExamMutation)
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
					_, err = mutators[i+1].Mutate(root, cecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cecb *CachedExamCreateBulk) SaveX(ctx context.Context) []*CachedExam {
	v, err := cecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cecb *CachedExamCreateBulk) Exec(ctx context.Context) error {
	_, err := cecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cecb *CachedExamCreateBulk) ExecX(ctx context.Context) {
	if err := cecb.Exec(ctx); err != nil {
		panic(err)
	}
}
