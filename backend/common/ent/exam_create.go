// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/cachedexam"
	"common/ent/exam"
	"common/ent/examcategory"
	"common/ent/examgroup"
	"common/ent/examsetting"
	"common/ent/generatedexam"
	"common/ent/subscriptionexam"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamCreate is the builder for creating a Exam entity.
type ExamCreate struct {
	config
	mutation *ExamMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ec *ExamCreate) SetName(s string) *ExamCreate {
	ec.mutation.SetName(s)
	return ec
}

// SetStage sets the "stage" field.
func (ec *ExamCreate) SetStage(s string) *ExamCreate {
	ec.mutation.SetStage(s)
	return ec
}

// SetNillableStage sets the "stage" field if the given value is not nil.
func (ec *ExamCreate) SetNillableStage(s *string) *ExamCreate {
	if s != nil {
		ec.SetStage(*s)
	}
	return ec
}

// SetIsSectional sets the "is_sectional" field.
func (ec *ExamCreate) SetIsSectional(b bool) *ExamCreate {
	ec.mutation.SetIsSectional(b)
	return ec
}

// SetNillableIsSectional sets the "is_sectional" field if the given value is not nil.
func (ec *ExamCreate) SetNillableIsSectional(b *bool) *ExamCreate {
	if b != nil {
		ec.SetIsSectional(*b)
	}
	return ec
}

// SetDescription sets the "description" field.
func (ec *ExamCreate) SetDescription(s string) *ExamCreate {
	ec.mutation.SetDescription(s)
	return ec
}

// SetType sets the "type" field.
func (ec *ExamCreate) SetType(e exam.Type) *ExamCreate {
	ec.mutation.SetType(e)
	return ec
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ec *ExamCreate) SetNillableType(e *exam.Type) *ExamCreate {
	if e != nil {
		ec.SetType(*e)
	}
	return ec
}

// SetIsActive sets the "is_active" field.
func (ec *ExamCreate) SetIsActive(b bool) *ExamCreate {
	ec.mutation.SetIsActive(b)
	return ec
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (ec *ExamCreate) SetNillableIsActive(b *bool) *ExamCreate {
	if b != nil {
		ec.SetIsActive(*b)
	}
	return ec
}

// SetLogoURL sets the "logo_url" field.
func (ec *ExamCreate) SetLogoURL(s string) *ExamCreate {
	ec.mutation.SetLogoURL(s)
	return ec
}

// SetNillableLogoURL sets the "logo_url" field if the given value is not nil.
func (ec *ExamCreate) SetNillableLogoURL(s *string) *ExamCreate {
	if s != nil {
		ec.SetLogoURL(*s)
	}
	return ec
}

// SetCreatedAt sets the "created_at" field.
func (ec *ExamCreate) SetCreatedAt(t time.Time) *ExamCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *ExamCreate) SetNillableCreatedAt(t *time.Time) *ExamCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// SetUpdatedAt sets the "updated_at" field.
func (ec *ExamCreate) SetUpdatedAt(t time.Time) *ExamCreate {
	ec.mutation.SetUpdatedAt(t)
	return ec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ec *ExamCreate) SetNillableUpdatedAt(t *time.Time) *ExamCreate {
	if t != nil {
		ec.SetUpdatedAt(*t)
	}
	return ec
}

// SetCategoryID sets the "category" edge to the ExamCategory entity by ID.
func (ec *ExamCreate) SetCategoryID(id int) *ExamCreate {
	ec.mutation.SetCategoryID(id)
	return ec
}

// SetNillableCategoryID sets the "category" edge to the ExamCategory entity by ID if the given value is not nil.
func (ec *ExamCreate) SetNillableCategoryID(id *int) *ExamCreate {
	if id != nil {
		ec = ec.SetCategoryID(*id)
	}
	return ec
}

// SetCategory sets the "category" edge to the ExamCategory entity.
func (ec *ExamCreate) SetCategory(e *ExamCategory) *ExamCreate {
	return ec.SetCategoryID(e.ID)
}

// SetGroupID sets the "group" edge to the ExamGroup entity by ID.
func (ec *ExamCreate) SetGroupID(id int) *ExamCreate {
	ec.mutation.SetGroupID(id)
	return ec
}

// SetNillableGroupID sets the "group" edge to the ExamGroup entity by ID if the given value is not nil.
func (ec *ExamCreate) SetNillableGroupID(id *int) *ExamCreate {
	if id != nil {
		ec = ec.SetGroupID(*id)
	}
	return ec
}

// SetGroup sets the "group" edge to the ExamGroup entity.
func (ec *ExamCreate) SetGroup(e *ExamGroup) *ExamCreate {
	return ec.SetGroupID(e.ID)
}

// AddSubscriptionIDs adds the "subscriptions" edge to the SubscriptionExam entity by IDs.
func (ec *ExamCreate) AddSubscriptionIDs(ids ...int) *ExamCreate {
	ec.mutation.AddSubscriptionIDs(ids...)
	return ec
}

// AddSubscriptions adds the "subscriptions" edges to the SubscriptionExam entity.
func (ec *ExamCreate) AddSubscriptions(s ...*SubscriptionExam) *ExamCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ec.AddSubscriptionIDs(ids...)
}

// SetSettingID sets the "setting" edge to the ExamSetting entity by ID.
func (ec *ExamCreate) SetSettingID(id int) *ExamCreate {
	ec.mutation.SetSettingID(id)
	return ec
}

// SetNillableSettingID sets the "setting" edge to the ExamSetting entity by ID if the given value is not nil.
func (ec *ExamCreate) SetNillableSettingID(id *int) *ExamCreate {
	if id != nil {
		ec = ec.SetSettingID(*id)
	}
	return ec
}

// SetSetting sets the "setting" edge to the ExamSetting entity.
func (ec *ExamCreate) SetSetting(e *ExamSetting) *ExamCreate {
	return ec.SetSettingID(e.ID)
}

// AddCachedExamIDs adds the "cached_exam" edge to the CachedExam entity by IDs.
func (ec *ExamCreate) AddCachedExamIDs(ids ...int) *ExamCreate {
	ec.mutation.AddCachedExamIDs(ids...)
	return ec
}

// AddCachedExam adds the "cached_exam" edges to the CachedExam entity.
func (ec *ExamCreate) AddCachedExam(c ...*CachedExam) *ExamCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ec.AddCachedExamIDs(ids...)
}

// AddGeneratedexamIDs adds the "generatedexams" edge to the GeneratedExam entity by IDs.
func (ec *ExamCreate) AddGeneratedexamIDs(ids ...int) *ExamCreate {
	ec.mutation.AddGeneratedexamIDs(ids...)
	return ec
}

// AddGeneratedexams adds the "generatedexams" edges to the GeneratedExam entity.
func (ec *ExamCreate) AddGeneratedexams(g ...*GeneratedExam) *ExamCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ec.AddGeneratedexamIDs(ids...)
}

// Mutation returns the ExamMutation object of the builder.
func (ec *ExamCreate) Mutation() *ExamMutation {
	return ec.mutation
}

// Save creates the Exam in the database.
func (ec *ExamCreate) Save(ctx context.Context) (*Exam, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *ExamCreate) SaveX(ctx context.Context) *Exam {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *ExamCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *ExamCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *ExamCreate) defaults() {
	if _, ok := ec.mutation.IsSectional(); !ok {
		v := exam.DefaultIsSectional
		ec.mutation.SetIsSectional(v)
	}
	if _, ok := ec.mutation.GetType(); !ok {
		v := exam.DefaultType
		ec.mutation.SetType(v)
	}
	if _, ok := ec.mutation.IsActive(); !ok {
		v := exam.DefaultIsActive
		ec.mutation.SetIsActive(v)
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		v := exam.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		v := exam.DefaultUpdatedAt()
		ec.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *ExamCreate) check() error {
	if _, ok := ec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Exam.name"`)}
	}
	if _, ok := ec.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Exam.description"`)}
	}
	if _, ok := ec.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Exam.type"`)}
	}
	if v, ok := ec.mutation.GetType(); ok {
		if err := exam.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Exam.type": %w`, err)}
		}
	}
	if _, ok := ec.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "Exam.is_active"`)}
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Exam.created_at"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Exam.updated_at"`)}
	}
	return nil
}

func (ec *ExamCreate) sqlSave(ctx context.Context) (*Exam, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *ExamCreate) createSpec() (*Exam, *sqlgraph.CreateSpec) {
	var (
		_node = &Exam{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(exam.Table, sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt))
	)
	if value, ok := ec.mutation.Name(); ok {
		_spec.SetField(exam.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ec.mutation.Stage(); ok {
		_spec.SetField(exam.FieldStage, field.TypeString, value)
		_node.Stage = value
	}
	if value, ok := ec.mutation.IsSectional(); ok {
		_spec.SetField(exam.FieldIsSectional, field.TypeBool, value)
		_node.IsSectional = value
	}
	if value, ok := ec.mutation.Description(); ok {
		_spec.SetField(exam.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ec.mutation.GetType(); ok {
		_spec.SetField(exam.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := ec.mutation.IsActive(); ok {
		_spec.SetField(exam.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := ec.mutation.LogoURL(); ok {
		_spec.SetField(exam.FieldLogoURL, field.TypeString, value)
		_node.LogoURL = value
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(exam.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.SetField(exam.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := ec.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   exam.CategoryTable,
			Columns: []string{exam.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examcategory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.exam_category_exams = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   exam.GroupTable,
			Columns: []string{exam.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.exam_group_exams = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.SubscriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exam.SubscriptionsTable,
			Columns: []string{exam.SubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subscriptionexam.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.SettingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   exam.SettingTable,
			Columns: []string{exam.SettingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examsetting.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.CachedExamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exam.CachedExamTable,
			Columns: []string{exam.CachedExamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cachedexam.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.GeneratedexamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exam.GeneratedexamsTable,
			Columns: []string{exam.GeneratedexamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(generatedexam.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ExamCreateBulk is the builder for creating many Exam entities in bulk.
type ExamCreateBulk struct {
	config
	err      error
	builders []*ExamCreate
}

// Save creates the Exam entities in the database.
func (ecb *ExamCreateBulk) Save(ctx context.Context) ([]*Exam, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Exam, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ExamMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *ExamCreateBulk) SaveX(ctx context.Context) []*Exam {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *ExamCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *ExamCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
