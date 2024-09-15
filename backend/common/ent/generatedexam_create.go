// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/exam"
	"common/ent/examattempt"
	"common/ent/generatedexam"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GeneratedExamCreate is the builder for creating a GeneratedExam entity.
type GeneratedExamCreate struct {
	config
	mutation *GeneratedExamMutation
	hooks    []Hook
}

// SetIsActive sets the "is_active" field.
func (gec *GeneratedExamCreate) SetIsActive(b bool) *GeneratedExamCreate {
	gec.mutation.SetIsActive(b)
	return gec
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (gec *GeneratedExamCreate) SetNillableIsActive(b *bool) *GeneratedExamCreate {
	if b != nil {
		gec.SetIsActive(*b)
	}
	return gec
}

// SetRawExamData sets the "raw_exam_data" field.
func (gec *GeneratedExamCreate) SetRawExamData(m map[string]interface{}) *GeneratedExamCreate {
	gec.mutation.SetRawExamData(m)
	return gec
}

// SetIsOpen sets the "is_open" field.
func (gec *GeneratedExamCreate) SetIsOpen(b bool) *GeneratedExamCreate {
	gec.mutation.SetIsOpen(b)
	return gec
}

// SetNillableIsOpen sets the "is_open" field if the given value is not nil.
func (gec *GeneratedExamCreate) SetNillableIsOpen(b *bool) *GeneratedExamCreate {
	if b != nil {
		gec.SetIsOpen(*b)
	}
	return gec
}

// SetCreatedAt sets the "created_at" field.
func (gec *GeneratedExamCreate) SetCreatedAt(t time.Time) *GeneratedExamCreate {
	gec.mutation.SetCreatedAt(t)
	return gec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gec *GeneratedExamCreate) SetNillableCreatedAt(t *time.Time) *GeneratedExamCreate {
	if t != nil {
		gec.SetCreatedAt(*t)
	}
	return gec
}

// SetUpdatedAt sets the "updated_at" field.
func (gec *GeneratedExamCreate) SetUpdatedAt(t time.Time) *GeneratedExamCreate {
	gec.mutation.SetUpdatedAt(t)
	return gec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gec *GeneratedExamCreate) SetNillableUpdatedAt(t *time.Time) *GeneratedExamCreate {
	if t != nil {
		gec.SetUpdatedAt(*t)
	}
	return gec
}

// SetExamID sets the "exam" edge to the Exam entity by ID.
func (gec *GeneratedExamCreate) SetExamID(id int) *GeneratedExamCreate {
	gec.mutation.SetExamID(id)
	return gec
}

// SetNillableExamID sets the "exam" edge to the Exam entity by ID if the given value is not nil.
func (gec *GeneratedExamCreate) SetNillableExamID(id *int) *GeneratedExamCreate {
	if id != nil {
		gec = gec.SetExamID(*id)
	}
	return gec
}

// SetExam sets the "exam" edge to the Exam entity.
func (gec *GeneratedExamCreate) SetExam(e *Exam) *GeneratedExamCreate {
	return gec.SetExamID(e.ID)
}

// AddAttemptIDs adds the "attempts" edge to the ExamAttempt entity by IDs.
func (gec *GeneratedExamCreate) AddAttemptIDs(ids ...int) *GeneratedExamCreate {
	gec.mutation.AddAttemptIDs(ids...)
	return gec
}

// AddAttempts adds the "attempts" edges to the ExamAttempt entity.
func (gec *GeneratedExamCreate) AddAttempts(e ...*ExamAttempt) *GeneratedExamCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return gec.AddAttemptIDs(ids...)
}

// Mutation returns the GeneratedExamMutation object of the builder.
func (gec *GeneratedExamCreate) Mutation() *GeneratedExamMutation {
	return gec.mutation
}

// Save creates the GeneratedExam in the database.
func (gec *GeneratedExamCreate) Save(ctx context.Context) (*GeneratedExam, error) {
	gec.defaults()
	return withHooks(ctx, gec.sqlSave, gec.mutation, gec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gec *GeneratedExamCreate) SaveX(ctx context.Context) *GeneratedExam {
	v, err := gec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gec *GeneratedExamCreate) Exec(ctx context.Context) error {
	_, err := gec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gec *GeneratedExamCreate) ExecX(ctx context.Context) {
	if err := gec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gec *GeneratedExamCreate) defaults() {
	if _, ok := gec.mutation.IsActive(); !ok {
		v := generatedexam.DefaultIsActive
		gec.mutation.SetIsActive(v)
	}
	if _, ok := gec.mutation.IsOpen(); !ok {
		v := generatedexam.DefaultIsOpen
		gec.mutation.SetIsOpen(v)
	}
	if _, ok := gec.mutation.CreatedAt(); !ok {
		v := generatedexam.DefaultCreatedAt()
		gec.mutation.SetCreatedAt(v)
	}
	if _, ok := gec.mutation.UpdatedAt(); !ok {
		v := generatedexam.DefaultUpdatedAt()
		gec.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gec *GeneratedExamCreate) check() error {
	if _, ok := gec.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "GeneratedExam.is_active"`)}
	}
	if _, ok := gec.mutation.IsOpen(); !ok {
		return &ValidationError{Name: "is_open", err: errors.New(`ent: missing required field "GeneratedExam.is_open"`)}
	}
	if _, ok := gec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "GeneratedExam.created_at"`)}
	}
	if _, ok := gec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "GeneratedExam.updated_at"`)}
	}
	return nil
}

func (gec *GeneratedExamCreate) sqlSave(ctx context.Context) (*GeneratedExam, error) {
	if err := gec.check(); err != nil {
		return nil, err
	}
	_node, _spec := gec.createSpec()
	if err := sqlgraph.CreateNode(ctx, gec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	gec.mutation.id = &_node.ID
	gec.mutation.done = true
	return _node, nil
}

func (gec *GeneratedExamCreate) createSpec() (*GeneratedExam, *sqlgraph.CreateSpec) {
	var (
		_node = &GeneratedExam{config: gec.config}
		_spec = sqlgraph.NewCreateSpec(generatedexam.Table, sqlgraph.NewFieldSpec(generatedexam.FieldID, field.TypeInt))
	)
	if value, ok := gec.mutation.IsActive(); ok {
		_spec.SetField(generatedexam.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := gec.mutation.RawExamData(); ok {
		_spec.SetField(generatedexam.FieldRawExamData, field.TypeJSON, value)
		_node.RawExamData = value
	}
	if value, ok := gec.mutation.IsOpen(); ok {
		_spec.SetField(generatedexam.FieldIsOpen, field.TypeBool, value)
		_node.IsOpen = value
	}
	if value, ok := gec.mutation.CreatedAt(); ok {
		_spec.SetField(generatedexam.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gec.mutation.UpdatedAt(); ok {
		_spec.SetField(generatedexam.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := gec.mutation.ExamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   generatedexam.ExamTable,
			Columns: []string{generatedexam.ExamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.exam_generatedexams = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gec.mutation.AttemptsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   generatedexam.AttemptsTable,
			Columns: []string{generatedexam.AttemptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examattempt.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GeneratedExamCreateBulk is the builder for creating many GeneratedExam entities in bulk.
type GeneratedExamCreateBulk struct {
	config
	err      error
	builders []*GeneratedExamCreate
}

// Save creates the GeneratedExam entities in the database.
func (gecb *GeneratedExamCreateBulk) Save(ctx context.Context) ([]*GeneratedExam, error) {
	if gecb.err != nil {
		return nil, gecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gecb.builders))
	nodes := make([]*GeneratedExam, len(gecb.builders))
	mutators := make([]Mutator, len(gecb.builders))
	for i := range gecb.builders {
		func(i int, root context.Context) {
			builder := gecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GeneratedExamMutation)
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
					_, err = mutators[i+1].Mutate(root, gecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gecb *GeneratedExamCreateBulk) SaveX(ctx context.Context) []*GeneratedExam {
	v, err := gecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gecb *GeneratedExamCreateBulk) Exec(ctx context.Context) error {
	_, err := gecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gecb *GeneratedExamCreateBulk) ExecX(ctx context.Context) {
	if err := gecb.Exec(ctx); err != nil {
		panic(err)
	}
}
