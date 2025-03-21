// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/exam"
	"common/ent/examcategory"
	"common/ent/examgroup"
	"common/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamGroupUpdate is the builder for updating ExamGroup entities.
type ExamGroupUpdate struct {
	config
	hooks    []Hook
	mutation *ExamGroupMutation
}

// Where appends a list predicates to the ExamGroupUpdate builder.
func (egu *ExamGroupUpdate) Where(ps ...predicate.ExamGroup) *ExamGroupUpdate {
	egu.mutation.Where(ps...)
	return egu
}

// SetName sets the "name" field.
func (egu *ExamGroupUpdate) SetName(s string) *ExamGroupUpdate {
	egu.mutation.SetName(s)
	return egu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (egu *ExamGroupUpdate) SetNillableName(s *string) *ExamGroupUpdate {
	if s != nil {
		egu.SetName(*s)
	}
	return egu
}

// SetDescription sets the "description" field.
func (egu *ExamGroupUpdate) SetDescription(s string) *ExamGroupUpdate {
	egu.mutation.SetDescription(s)
	return egu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (egu *ExamGroupUpdate) SetNillableDescription(s *string) *ExamGroupUpdate {
	if s != nil {
		egu.SetDescription(*s)
	}
	return egu
}

// SetIsActive sets the "is_active" field.
func (egu *ExamGroupUpdate) SetIsActive(b bool) *ExamGroupUpdate {
	egu.mutation.SetIsActive(b)
	return egu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (egu *ExamGroupUpdate) SetNillableIsActive(b *bool) *ExamGroupUpdate {
	if b != nil {
		egu.SetIsActive(*b)
	}
	return egu
}

// SetLogoURL sets the "logo_url" field.
func (egu *ExamGroupUpdate) SetLogoURL(s string) *ExamGroupUpdate {
	egu.mutation.SetLogoURL(s)
	return egu
}

// SetNillableLogoURL sets the "logo_url" field if the given value is not nil.
func (egu *ExamGroupUpdate) SetNillableLogoURL(s *string) *ExamGroupUpdate {
	if s != nil {
		egu.SetLogoURL(*s)
	}
	return egu
}

// ClearLogoURL clears the value of the "logo_url" field.
func (egu *ExamGroupUpdate) ClearLogoURL() *ExamGroupUpdate {
	egu.mutation.ClearLogoURL()
	return egu
}

// SetUpdatedAt sets the "updated_at" field.
func (egu *ExamGroupUpdate) SetUpdatedAt(t time.Time) *ExamGroupUpdate {
	egu.mutation.SetUpdatedAt(t)
	return egu
}

// SetCategoryID sets the "category" edge to the ExamCategory entity by ID.
func (egu *ExamGroupUpdate) SetCategoryID(id int) *ExamGroupUpdate {
	egu.mutation.SetCategoryID(id)
	return egu
}

// SetNillableCategoryID sets the "category" edge to the ExamCategory entity by ID if the given value is not nil.
func (egu *ExamGroupUpdate) SetNillableCategoryID(id *int) *ExamGroupUpdate {
	if id != nil {
		egu = egu.SetCategoryID(*id)
	}
	return egu
}

// SetCategory sets the "category" edge to the ExamCategory entity.
func (egu *ExamGroupUpdate) SetCategory(e *ExamCategory) *ExamGroupUpdate {
	return egu.SetCategoryID(e.ID)
}

// AddExamIDs adds the "exams" edge to the Exam entity by IDs.
func (egu *ExamGroupUpdate) AddExamIDs(ids ...int) *ExamGroupUpdate {
	egu.mutation.AddExamIDs(ids...)
	return egu
}

// AddExams adds the "exams" edges to the Exam entity.
func (egu *ExamGroupUpdate) AddExams(e ...*Exam) *ExamGroupUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return egu.AddExamIDs(ids...)
}

// Mutation returns the ExamGroupMutation object of the builder.
func (egu *ExamGroupUpdate) Mutation() *ExamGroupMutation {
	return egu.mutation
}

// ClearCategory clears the "category" edge to the ExamCategory entity.
func (egu *ExamGroupUpdate) ClearCategory() *ExamGroupUpdate {
	egu.mutation.ClearCategory()
	return egu
}

// ClearExams clears all "exams" edges to the Exam entity.
func (egu *ExamGroupUpdate) ClearExams() *ExamGroupUpdate {
	egu.mutation.ClearExams()
	return egu
}

// RemoveExamIDs removes the "exams" edge to Exam entities by IDs.
func (egu *ExamGroupUpdate) RemoveExamIDs(ids ...int) *ExamGroupUpdate {
	egu.mutation.RemoveExamIDs(ids...)
	return egu
}

// RemoveExams removes "exams" edges to Exam entities.
func (egu *ExamGroupUpdate) RemoveExams(e ...*Exam) *ExamGroupUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return egu.RemoveExamIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (egu *ExamGroupUpdate) Save(ctx context.Context) (int, error) {
	egu.defaults()
	return withHooks(ctx, egu.sqlSave, egu.mutation, egu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (egu *ExamGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := egu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (egu *ExamGroupUpdate) Exec(ctx context.Context) error {
	_, err := egu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (egu *ExamGroupUpdate) ExecX(ctx context.Context) {
	if err := egu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (egu *ExamGroupUpdate) defaults() {
	if _, ok := egu.mutation.UpdatedAt(); !ok {
		v := examgroup.UpdateDefaultUpdatedAt()
		egu.mutation.SetUpdatedAt(v)
	}
}

func (egu *ExamGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(examgroup.Table, examgroup.Columns, sqlgraph.NewFieldSpec(examgroup.FieldID, field.TypeInt))
	if ps := egu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := egu.mutation.Name(); ok {
		_spec.SetField(examgroup.FieldName, field.TypeString, value)
	}
	if value, ok := egu.mutation.Description(); ok {
		_spec.SetField(examgroup.FieldDescription, field.TypeString, value)
	}
	if value, ok := egu.mutation.IsActive(); ok {
		_spec.SetField(examgroup.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := egu.mutation.LogoURL(); ok {
		_spec.SetField(examgroup.FieldLogoURL, field.TypeString, value)
	}
	if egu.mutation.LogoURLCleared() {
		_spec.ClearField(examgroup.FieldLogoURL, field.TypeString)
	}
	if value, ok := egu.mutation.UpdatedAt(); ok {
		_spec.SetField(examgroup.FieldUpdatedAt, field.TypeTime, value)
	}
	if egu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   examgroup.CategoryTable,
			Columns: []string{examgroup.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examcategory.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := egu.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   examgroup.CategoryTable,
			Columns: []string{examgroup.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examcategory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if egu.mutation.ExamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examgroup.ExamsTable,
			Columns: []string{examgroup.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := egu.mutation.RemovedExamsIDs(); len(nodes) > 0 && !egu.mutation.ExamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examgroup.ExamsTable,
			Columns: []string{examgroup.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := egu.mutation.ExamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examgroup.ExamsTable,
			Columns: []string{examgroup.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, egu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{examgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	egu.mutation.done = true
	return n, nil
}

// ExamGroupUpdateOne is the builder for updating a single ExamGroup entity.
type ExamGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ExamGroupMutation
}

// SetName sets the "name" field.
func (eguo *ExamGroupUpdateOne) SetName(s string) *ExamGroupUpdateOne {
	eguo.mutation.SetName(s)
	return eguo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (eguo *ExamGroupUpdateOne) SetNillableName(s *string) *ExamGroupUpdateOne {
	if s != nil {
		eguo.SetName(*s)
	}
	return eguo
}

// SetDescription sets the "description" field.
func (eguo *ExamGroupUpdateOne) SetDescription(s string) *ExamGroupUpdateOne {
	eguo.mutation.SetDescription(s)
	return eguo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (eguo *ExamGroupUpdateOne) SetNillableDescription(s *string) *ExamGroupUpdateOne {
	if s != nil {
		eguo.SetDescription(*s)
	}
	return eguo
}

// SetIsActive sets the "is_active" field.
func (eguo *ExamGroupUpdateOne) SetIsActive(b bool) *ExamGroupUpdateOne {
	eguo.mutation.SetIsActive(b)
	return eguo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (eguo *ExamGroupUpdateOne) SetNillableIsActive(b *bool) *ExamGroupUpdateOne {
	if b != nil {
		eguo.SetIsActive(*b)
	}
	return eguo
}

// SetLogoURL sets the "logo_url" field.
func (eguo *ExamGroupUpdateOne) SetLogoURL(s string) *ExamGroupUpdateOne {
	eguo.mutation.SetLogoURL(s)
	return eguo
}

// SetNillableLogoURL sets the "logo_url" field if the given value is not nil.
func (eguo *ExamGroupUpdateOne) SetNillableLogoURL(s *string) *ExamGroupUpdateOne {
	if s != nil {
		eguo.SetLogoURL(*s)
	}
	return eguo
}

// ClearLogoURL clears the value of the "logo_url" field.
func (eguo *ExamGroupUpdateOne) ClearLogoURL() *ExamGroupUpdateOne {
	eguo.mutation.ClearLogoURL()
	return eguo
}

// SetUpdatedAt sets the "updated_at" field.
func (eguo *ExamGroupUpdateOne) SetUpdatedAt(t time.Time) *ExamGroupUpdateOne {
	eguo.mutation.SetUpdatedAt(t)
	return eguo
}

// SetCategoryID sets the "category" edge to the ExamCategory entity by ID.
func (eguo *ExamGroupUpdateOne) SetCategoryID(id int) *ExamGroupUpdateOne {
	eguo.mutation.SetCategoryID(id)
	return eguo
}

// SetNillableCategoryID sets the "category" edge to the ExamCategory entity by ID if the given value is not nil.
func (eguo *ExamGroupUpdateOne) SetNillableCategoryID(id *int) *ExamGroupUpdateOne {
	if id != nil {
		eguo = eguo.SetCategoryID(*id)
	}
	return eguo
}

// SetCategory sets the "category" edge to the ExamCategory entity.
func (eguo *ExamGroupUpdateOne) SetCategory(e *ExamCategory) *ExamGroupUpdateOne {
	return eguo.SetCategoryID(e.ID)
}

// AddExamIDs adds the "exams" edge to the Exam entity by IDs.
func (eguo *ExamGroupUpdateOne) AddExamIDs(ids ...int) *ExamGroupUpdateOne {
	eguo.mutation.AddExamIDs(ids...)
	return eguo
}

// AddExams adds the "exams" edges to the Exam entity.
func (eguo *ExamGroupUpdateOne) AddExams(e ...*Exam) *ExamGroupUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return eguo.AddExamIDs(ids...)
}

// Mutation returns the ExamGroupMutation object of the builder.
func (eguo *ExamGroupUpdateOne) Mutation() *ExamGroupMutation {
	return eguo.mutation
}

// ClearCategory clears the "category" edge to the ExamCategory entity.
func (eguo *ExamGroupUpdateOne) ClearCategory() *ExamGroupUpdateOne {
	eguo.mutation.ClearCategory()
	return eguo
}

// ClearExams clears all "exams" edges to the Exam entity.
func (eguo *ExamGroupUpdateOne) ClearExams() *ExamGroupUpdateOne {
	eguo.mutation.ClearExams()
	return eguo
}

// RemoveExamIDs removes the "exams" edge to Exam entities by IDs.
func (eguo *ExamGroupUpdateOne) RemoveExamIDs(ids ...int) *ExamGroupUpdateOne {
	eguo.mutation.RemoveExamIDs(ids...)
	return eguo
}

// RemoveExams removes "exams" edges to Exam entities.
func (eguo *ExamGroupUpdateOne) RemoveExams(e ...*Exam) *ExamGroupUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return eguo.RemoveExamIDs(ids...)
}

// Where appends a list predicates to the ExamGroupUpdate builder.
func (eguo *ExamGroupUpdateOne) Where(ps ...predicate.ExamGroup) *ExamGroupUpdateOne {
	eguo.mutation.Where(ps...)
	return eguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eguo *ExamGroupUpdateOne) Select(field string, fields ...string) *ExamGroupUpdateOne {
	eguo.fields = append([]string{field}, fields...)
	return eguo
}

// Save executes the query and returns the updated ExamGroup entity.
func (eguo *ExamGroupUpdateOne) Save(ctx context.Context) (*ExamGroup, error) {
	eguo.defaults()
	return withHooks(ctx, eguo.sqlSave, eguo.mutation, eguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eguo *ExamGroupUpdateOne) SaveX(ctx context.Context) *ExamGroup {
	node, err := eguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eguo *ExamGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := eguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eguo *ExamGroupUpdateOne) ExecX(ctx context.Context) {
	if err := eguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eguo *ExamGroupUpdateOne) defaults() {
	if _, ok := eguo.mutation.UpdatedAt(); !ok {
		v := examgroup.UpdateDefaultUpdatedAt()
		eguo.mutation.SetUpdatedAt(v)
	}
}

func (eguo *ExamGroupUpdateOne) sqlSave(ctx context.Context) (_node *ExamGroup, err error) {
	_spec := sqlgraph.NewUpdateSpec(examgroup.Table, examgroup.Columns, sqlgraph.NewFieldSpec(examgroup.FieldID, field.TypeInt))
	id, ok := eguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ExamGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, examgroup.FieldID)
		for _, f := range fields {
			if !examgroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != examgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eguo.mutation.Name(); ok {
		_spec.SetField(examgroup.FieldName, field.TypeString, value)
	}
	if value, ok := eguo.mutation.Description(); ok {
		_spec.SetField(examgroup.FieldDescription, field.TypeString, value)
	}
	if value, ok := eguo.mutation.IsActive(); ok {
		_spec.SetField(examgroup.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := eguo.mutation.LogoURL(); ok {
		_spec.SetField(examgroup.FieldLogoURL, field.TypeString, value)
	}
	if eguo.mutation.LogoURLCleared() {
		_spec.ClearField(examgroup.FieldLogoURL, field.TypeString)
	}
	if value, ok := eguo.mutation.UpdatedAt(); ok {
		_spec.SetField(examgroup.FieldUpdatedAt, field.TypeTime, value)
	}
	if eguo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   examgroup.CategoryTable,
			Columns: []string{examgroup.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examcategory.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eguo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   examgroup.CategoryTable,
			Columns: []string{examgroup.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examcategory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eguo.mutation.ExamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examgroup.ExamsTable,
			Columns: []string{examgroup.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eguo.mutation.RemovedExamsIDs(); len(nodes) > 0 && !eguo.mutation.ExamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examgroup.ExamsTable,
			Columns: []string{examgroup.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eguo.mutation.ExamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examgroup.ExamsTable,
			Columns: []string{examgroup.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ExamGroup{config: eguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{examgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	eguo.mutation.done = true
	return _node, nil
}
