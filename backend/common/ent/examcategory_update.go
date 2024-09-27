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

// ExamCategoryUpdate is the builder for updating ExamCategory entities.
type ExamCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *ExamCategoryMutation
}

// Where appends a list predicates to the ExamCategoryUpdate builder.
func (ecu *ExamCategoryUpdate) Where(ps ...predicate.ExamCategory) *ExamCategoryUpdate {
	ecu.mutation.Where(ps...)
	return ecu
}

// SetName sets the "name" field.
func (ecu *ExamCategoryUpdate) SetName(e examcategory.Name) *ExamCategoryUpdate {
	ecu.mutation.SetName(e)
	return ecu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ecu *ExamCategoryUpdate) SetNillableName(e *examcategory.Name) *ExamCategoryUpdate {
	if e != nil {
		ecu.SetName(*e)
	}
	return ecu
}

// SetDescription sets the "description" field.
func (ecu *ExamCategoryUpdate) SetDescription(s string) *ExamCategoryUpdate {
	ecu.mutation.SetDescription(s)
	return ecu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ecu *ExamCategoryUpdate) SetNillableDescription(s *string) *ExamCategoryUpdate {
	if s != nil {
		ecu.SetDescription(*s)
	}
	return ecu
}

// SetIsActive sets the "is_active" field.
func (ecu *ExamCategoryUpdate) SetIsActive(b bool) *ExamCategoryUpdate {
	ecu.mutation.SetIsActive(b)
	return ecu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (ecu *ExamCategoryUpdate) SetNillableIsActive(b *bool) *ExamCategoryUpdate {
	if b != nil {
		ecu.SetIsActive(*b)
	}
	return ecu
}

// SetUpdatedAt sets the "updated_at" field.
func (ecu *ExamCategoryUpdate) SetUpdatedAt(t time.Time) *ExamCategoryUpdate {
	ecu.mutation.SetUpdatedAt(t)
	return ecu
}

// AddExamIDs adds the "exams" edge to the Exam entity by IDs.
func (ecu *ExamCategoryUpdate) AddExamIDs(ids ...int) *ExamCategoryUpdate {
	ecu.mutation.AddExamIDs(ids...)
	return ecu
}

// AddExams adds the "exams" edges to the Exam entity.
func (ecu *ExamCategoryUpdate) AddExams(e ...*Exam) *ExamCategoryUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecu.AddExamIDs(ids...)
}

// AddExamGroupIDs adds the "exam_groups" edge to the ExamGroup entity by IDs.
func (ecu *ExamCategoryUpdate) AddExamGroupIDs(ids ...int) *ExamCategoryUpdate {
	ecu.mutation.AddExamGroupIDs(ids...)
	return ecu
}

// AddExamGroups adds the "exam_groups" edges to the ExamGroup entity.
func (ecu *ExamCategoryUpdate) AddExamGroups(e ...*ExamGroup) *ExamCategoryUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecu.AddExamGroupIDs(ids...)
}

// Mutation returns the ExamCategoryMutation object of the builder.
func (ecu *ExamCategoryUpdate) Mutation() *ExamCategoryMutation {
	return ecu.mutation
}

// ClearExams clears all "exams" edges to the Exam entity.
func (ecu *ExamCategoryUpdate) ClearExams() *ExamCategoryUpdate {
	ecu.mutation.ClearExams()
	return ecu
}

// RemoveExamIDs removes the "exams" edge to Exam entities by IDs.
func (ecu *ExamCategoryUpdate) RemoveExamIDs(ids ...int) *ExamCategoryUpdate {
	ecu.mutation.RemoveExamIDs(ids...)
	return ecu
}

// RemoveExams removes "exams" edges to Exam entities.
func (ecu *ExamCategoryUpdate) RemoveExams(e ...*Exam) *ExamCategoryUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecu.RemoveExamIDs(ids...)
}

// ClearExamGroups clears all "exam_groups" edges to the ExamGroup entity.
func (ecu *ExamCategoryUpdate) ClearExamGroups() *ExamCategoryUpdate {
	ecu.mutation.ClearExamGroups()
	return ecu
}

// RemoveExamGroupIDs removes the "exam_groups" edge to ExamGroup entities by IDs.
func (ecu *ExamCategoryUpdate) RemoveExamGroupIDs(ids ...int) *ExamCategoryUpdate {
	ecu.mutation.RemoveExamGroupIDs(ids...)
	return ecu
}

// RemoveExamGroups removes "exam_groups" edges to ExamGroup entities.
func (ecu *ExamCategoryUpdate) RemoveExamGroups(e ...*ExamGroup) *ExamCategoryUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecu.RemoveExamGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ecu *ExamCategoryUpdate) Save(ctx context.Context) (int, error) {
	ecu.defaults()
	return withHooks(ctx, ecu.sqlSave, ecu.mutation, ecu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ecu *ExamCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ecu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ecu *ExamCategoryUpdate) Exec(ctx context.Context) error {
	_, err := ecu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecu *ExamCategoryUpdate) ExecX(ctx context.Context) {
	if err := ecu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ecu *ExamCategoryUpdate) defaults() {
	if _, ok := ecu.mutation.UpdatedAt(); !ok {
		v := examcategory.UpdateDefaultUpdatedAt()
		ecu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecu *ExamCategoryUpdate) check() error {
	if v, ok := ecu.mutation.Name(); ok {
		if err := examcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ExamCategory.name": %w`, err)}
		}
	}
	return nil
}

func (ecu *ExamCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ecu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(examcategory.Table, examcategory.Columns, sqlgraph.NewFieldSpec(examcategory.FieldID, field.TypeInt))
	if ps := ecu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecu.mutation.Name(); ok {
		_spec.SetField(examcategory.FieldName, field.TypeEnum, value)
	}
	if value, ok := ecu.mutation.Description(); ok {
		_spec.SetField(examcategory.FieldDescription, field.TypeString, value)
	}
	if value, ok := ecu.mutation.IsActive(); ok {
		_spec.SetField(examcategory.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := ecu.mutation.UpdatedAt(); ok {
		_spec.SetField(examcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ecu.mutation.ExamsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.RemovedExamsIDs(); len(nodes) > 0 && !ecu.mutation.ExamsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.ExamsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ecu.mutation.ExamGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examcategory.ExamGroupsTable,
			Columns: []string{examcategory.ExamGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examgroup.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.RemovedExamGroupsIDs(); len(nodes) > 0 && !ecu.mutation.ExamGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examcategory.ExamGroupsTable,
			Columns: []string{examcategory.ExamGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.ExamGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examcategory.ExamGroupsTable,
			Columns: []string{examcategory.ExamGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ecu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{examcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ecu.mutation.done = true
	return n, nil
}

// ExamCategoryUpdateOne is the builder for updating a single ExamCategory entity.
type ExamCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ExamCategoryMutation
}

// SetName sets the "name" field.
func (ecuo *ExamCategoryUpdateOne) SetName(e examcategory.Name) *ExamCategoryUpdateOne {
	ecuo.mutation.SetName(e)
	return ecuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ecuo *ExamCategoryUpdateOne) SetNillableName(e *examcategory.Name) *ExamCategoryUpdateOne {
	if e != nil {
		ecuo.SetName(*e)
	}
	return ecuo
}

// SetDescription sets the "description" field.
func (ecuo *ExamCategoryUpdateOne) SetDescription(s string) *ExamCategoryUpdateOne {
	ecuo.mutation.SetDescription(s)
	return ecuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ecuo *ExamCategoryUpdateOne) SetNillableDescription(s *string) *ExamCategoryUpdateOne {
	if s != nil {
		ecuo.SetDescription(*s)
	}
	return ecuo
}

// SetIsActive sets the "is_active" field.
func (ecuo *ExamCategoryUpdateOne) SetIsActive(b bool) *ExamCategoryUpdateOne {
	ecuo.mutation.SetIsActive(b)
	return ecuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (ecuo *ExamCategoryUpdateOne) SetNillableIsActive(b *bool) *ExamCategoryUpdateOne {
	if b != nil {
		ecuo.SetIsActive(*b)
	}
	return ecuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ecuo *ExamCategoryUpdateOne) SetUpdatedAt(t time.Time) *ExamCategoryUpdateOne {
	ecuo.mutation.SetUpdatedAt(t)
	return ecuo
}

// AddExamIDs adds the "exams" edge to the Exam entity by IDs.
func (ecuo *ExamCategoryUpdateOne) AddExamIDs(ids ...int) *ExamCategoryUpdateOne {
	ecuo.mutation.AddExamIDs(ids...)
	return ecuo
}

// AddExams adds the "exams" edges to the Exam entity.
func (ecuo *ExamCategoryUpdateOne) AddExams(e ...*Exam) *ExamCategoryUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecuo.AddExamIDs(ids...)
}

// AddExamGroupIDs adds the "exam_groups" edge to the ExamGroup entity by IDs.
func (ecuo *ExamCategoryUpdateOne) AddExamGroupIDs(ids ...int) *ExamCategoryUpdateOne {
	ecuo.mutation.AddExamGroupIDs(ids...)
	return ecuo
}

// AddExamGroups adds the "exam_groups" edges to the ExamGroup entity.
func (ecuo *ExamCategoryUpdateOne) AddExamGroups(e ...*ExamGroup) *ExamCategoryUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecuo.AddExamGroupIDs(ids...)
}

// Mutation returns the ExamCategoryMutation object of the builder.
func (ecuo *ExamCategoryUpdateOne) Mutation() *ExamCategoryMutation {
	return ecuo.mutation
}

// ClearExams clears all "exams" edges to the Exam entity.
func (ecuo *ExamCategoryUpdateOne) ClearExams() *ExamCategoryUpdateOne {
	ecuo.mutation.ClearExams()
	return ecuo
}

// RemoveExamIDs removes the "exams" edge to Exam entities by IDs.
func (ecuo *ExamCategoryUpdateOne) RemoveExamIDs(ids ...int) *ExamCategoryUpdateOne {
	ecuo.mutation.RemoveExamIDs(ids...)
	return ecuo
}

// RemoveExams removes "exams" edges to Exam entities.
func (ecuo *ExamCategoryUpdateOne) RemoveExams(e ...*Exam) *ExamCategoryUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecuo.RemoveExamIDs(ids...)
}

// ClearExamGroups clears all "exam_groups" edges to the ExamGroup entity.
func (ecuo *ExamCategoryUpdateOne) ClearExamGroups() *ExamCategoryUpdateOne {
	ecuo.mutation.ClearExamGroups()
	return ecuo
}

// RemoveExamGroupIDs removes the "exam_groups" edge to ExamGroup entities by IDs.
func (ecuo *ExamCategoryUpdateOne) RemoveExamGroupIDs(ids ...int) *ExamCategoryUpdateOne {
	ecuo.mutation.RemoveExamGroupIDs(ids...)
	return ecuo
}

// RemoveExamGroups removes "exam_groups" edges to ExamGroup entities.
func (ecuo *ExamCategoryUpdateOne) RemoveExamGroups(e ...*ExamGroup) *ExamCategoryUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecuo.RemoveExamGroupIDs(ids...)
}

// Where appends a list predicates to the ExamCategoryUpdate builder.
func (ecuo *ExamCategoryUpdateOne) Where(ps ...predicate.ExamCategory) *ExamCategoryUpdateOne {
	ecuo.mutation.Where(ps...)
	return ecuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ecuo *ExamCategoryUpdateOne) Select(field string, fields ...string) *ExamCategoryUpdateOne {
	ecuo.fields = append([]string{field}, fields...)
	return ecuo
}

// Save executes the query and returns the updated ExamCategory entity.
func (ecuo *ExamCategoryUpdateOne) Save(ctx context.Context) (*ExamCategory, error) {
	ecuo.defaults()
	return withHooks(ctx, ecuo.sqlSave, ecuo.mutation, ecuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ecuo *ExamCategoryUpdateOne) SaveX(ctx context.Context) *ExamCategory {
	node, err := ecuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ecuo *ExamCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ecuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecuo *ExamCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := ecuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ecuo *ExamCategoryUpdateOne) defaults() {
	if _, ok := ecuo.mutation.UpdatedAt(); !ok {
		v := examcategory.UpdateDefaultUpdatedAt()
		ecuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecuo *ExamCategoryUpdateOne) check() error {
	if v, ok := ecuo.mutation.Name(); ok {
		if err := examcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ExamCategory.name": %w`, err)}
		}
	}
	return nil
}

func (ecuo *ExamCategoryUpdateOne) sqlSave(ctx context.Context) (_node *ExamCategory, err error) {
	if err := ecuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(examcategory.Table, examcategory.Columns, sqlgraph.NewFieldSpec(examcategory.FieldID, field.TypeInt))
	id, ok := ecuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ExamCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ecuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, examcategory.FieldID)
		for _, f := range fields {
			if !examcategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != examcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ecuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecuo.mutation.Name(); ok {
		_spec.SetField(examcategory.FieldName, field.TypeEnum, value)
	}
	if value, ok := ecuo.mutation.Description(); ok {
		_spec.SetField(examcategory.FieldDescription, field.TypeString, value)
	}
	if value, ok := ecuo.mutation.IsActive(); ok {
		_spec.SetField(examcategory.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := ecuo.mutation.UpdatedAt(); ok {
		_spec.SetField(examcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ecuo.mutation.ExamsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.RemovedExamsIDs(); len(nodes) > 0 && !ecuo.mutation.ExamsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.ExamsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ecuo.mutation.ExamGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examcategory.ExamGroupsTable,
			Columns: []string{examcategory.ExamGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examgroup.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.RemovedExamGroupsIDs(); len(nodes) > 0 && !ecuo.mutation.ExamGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examcategory.ExamGroupsTable,
			Columns: []string{examcategory.ExamGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.ExamGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examcategory.ExamGroupsTable,
			Columns: []string{examcategory.ExamGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(examgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ExamCategory{config: ecuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ecuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{examcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ecuo.mutation.done = true
	return _node, nil
}
