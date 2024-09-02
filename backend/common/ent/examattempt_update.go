// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/exam"
	"common/ent/examattempt"
	"common/ent/predicate"
	"common/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ExamAttemptUpdate is the builder for updating ExamAttempt entities.
type ExamAttemptUpdate struct {
	config
	hooks    []Hook
	mutation *ExamAttemptMutation
}

// Where appends a list predicates to the ExamAttemptUpdate builder.
func (eau *ExamAttemptUpdate) Where(ps ...predicate.ExamAttempt) *ExamAttemptUpdate {
	eau.mutation.Where(ps...)
	return eau
}

// SetAttemptNumber sets the "attempt_number" field.
func (eau *ExamAttemptUpdate) SetAttemptNumber(i int) *ExamAttemptUpdate {
	eau.mutation.ResetAttemptNumber()
	eau.mutation.SetAttemptNumber(i)
	return eau
}

// SetNillableAttemptNumber sets the "attempt_number" field if the given value is not nil.
func (eau *ExamAttemptUpdate) SetNillableAttemptNumber(i *int) *ExamAttemptUpdate {
	if i != nil {
		eau.SetAttemptNumber(*i)
	}
	return eau
}

// AddAttemptNumber adds i to the "attempt_number" field.
func (eau *ExamAttemptUpdate) AddAttemptNumber(i int) *ExamAttemptUpdate {
	eau.mutation.AddAttemptNumber(i)
	return eau
}

// SetUpdatedAt sets the "updated_at" field.
func (eau *ExamAttemptUpdate) SetUpdatedAt(t time.Time) *ExamAttemptUpdate {
	eau.mutation.SetUpdatedAt(t)
	return eau
}

// SetExamID sets the "exam" edge to the Exam entity by ID.
func (eau *ExamAttemptUpdate) SetExamID(id int) *ExamAttemptUpdate {
	eau.mutation.SetExamID(id)
	return eau
}

// SetNillableExamID sets the "exam" edge to the Exam entity by ID if the given value is not nil.
func (eau *ExamAttemptUpdate) SetNillableExamID(id *int) *ExamAttemptUpdate {
	if id != nil {
		eau = eau.SetExamID(*id)
	}
	return eau
}

// SetExam sets the "exam" edge to the Exam entity.
func (eau *ExamAttemptUpdate) SetExam(e *Exam) *ExamAttemptUpdate {
	return eau.SetExamID(e.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (eau *ExamAttemptUpdate) SetUserID(id uuid.UUID) *ExamAttemptUpdate {
	eau.mutation.SetUserID(id)
	return eau
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (eau *ExamAttemptUpdate) SetNillableUserID(id *uuid.UUID) *ExamAttemptUpdate {
	if id != nil {
		eau = eau.SetUserID(*id)
	}
	return eau
}

// SetUser sets the "user" edge to the User entity.
func (eau *ExamAttemptUpdate) SetUser(u *User) *ExamAttemptUpdate {
	return eau.SetUserID(u.ID)
}

// Mutation returns the ExamAttemptMutation object of the builder.
func (eau *ExamAttemptUpdate) Mutation() *ExamAttemptMutation {
	return eau.mutation
}

// ClearExam clears the "exam" edge to the Exam entity.
func (eau *ExamAttemptUpdate) ClearExam() *ExamAttemptUpdate {
	eau.mutation.ClearExam()
	return eau
}

// ClearUser clears the "user" edge to the User entity.
func (eau *ExamAttemptUpdate) ClearUser() *ExamAttemptUpdate {
	eau.mutation.ClearUser()
	return eau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eau *ExamAttemptUpdate) Save(ctx context.Context) (int, error) {
	eau.defaults()
	return withHooks(ctx, eau.sqlSave, eau.mutation, eau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eau *ExamAttemptUpdate) SaveX(ctx context.Context) int {
	affected, err := eau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eau *ExamAttemptUpdate) Exec(ctx context.Context) error {
	_, err := eau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eau *ExamAttemptUpdate) ExecX(ctx context.Context) {
	if err := eau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eau *ExamAttemptUpdate) defaults() {
	if _, ok := eau.mutation.UpdatedAt(); !ok {
		v := examattempt.UpdateDefaultUpdatedAt()
		eau.mutation.SetUpdatedAt(v)
	}
}

func (eau *ExamAttemptUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(examattempt.Table, examattempt.Columns, sqlgraph.NewFieldSpec(examattempt.FieldID, field.TypeInt))
	if ps := eau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eau.mutation.AttemptNumber(); ok {
		_spec.SetField(examattempt.FieldAttemptNumber, field.TypeInt, value)
	}
	if value, ok := eau.mutation.AddedAttemptNumber(); ok {
		_spec.AddField(examattempt.FieldAttemptNumber, field.TypeInt, value)
	}
	if value, ok := eau.mutation.UpdatedAt(); ok {
		_spec.SetField(examattempt.FieldUpdatedAt, field.TypeTime, value)
	}
	if eau.mutation.ExamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   examattempt.ExamTable,
			Columns: []string{examattempt.ExamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eau.mutation.ExamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   examattempt.ExamTable,
			Columns: []string{examattempt.ExamColumn},
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
	if eau.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   examattempt.UserTable,
			Columns: []string{examattempt.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eau.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   examattempt.UserTable,
			Columns: []string{examattempt.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{examattempt.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eau.mutation.done = true
	return n, nil
}

// ExamAttemptUpdateOne is the builder for updating a single ExamAttempt entity.
type ExamAttemptUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ExamAttemptMutation
}

// SetAttemptNumber sets the "attempt_number" field.
func (eauo *ExamAttemptUpdateOne) SetAttemptNumber(i int) *ExamAttemptUpdateOne {
	eauo.mutation.ResetAttemptNumber()
	eauo.mutation.SetAttemptNumber(i)
	return eauo
}

// SetNillableAttemptNumber sets the "attempt_number" field if the given value is not nil.
func (eauo *ExamAttemptUpdateOne) SetNillableAttemptNumber(i *int) *ExamAttemptUpdateOne {
	if i != nil {
		eauo.SetAttemptNumber(*i)
	}
	return eauo
}

// AddAttemptNumber adds i to the "attempt_number" field.
func (eauo *ExamAttemptUpdateOne) AddAttemptNumber(i int) *ExamAttemptUpdateOne {
	eauo.mutation.AddAttemptNumber(i)
	return eauo
}

// SetUpdatedAt sets the "updated_at" field.
func (eauo *ExamAttemptUpdateOne) SetUpdatedAt(t time.Time) *ExamAttemptUpdateOne {
	eauo.mutation.SetUpdatedAt(t)
	return eauo
}

// SetExamID sets the "exam" edge to the Exam entity by ID.
func (eauo *ExamAttemptUpdateOne) SetExamID(id int) *ExamAttemptUpdateOne {
	eauo.mutation.SetExamID(id)
	return eauo
}

// SetNillableExamID sets the "exam" edge to the Exam entity by ID if the given value is not nil.
func (eauo *ExamAttemptUpdateOne) SetNillableExamID(id *int) *ExamAttemptUpdateOne {
	if id != nil {
		eauo = eauo.SetExamID(*id)
	}
	return eauo
}

// SetExam sets the "exam" edge to the Exam entity.
func (eauo *ExamAttemptUpdateOne) SetExam(e *Exam) *ExamAttemptUpdateOne {
	return eauo.SetExamID(e.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (eauo *ExamAttemptUpdateOne) SetUserID(id uuid.UUID) *ExamAttemptUpdateOne {
	eauo.mutation.SetUserID(id)
	return eauo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (eauo *ExamAttemptUpdateOne) SetNillableUserID(id *uuid.UUID) *ExamAttemptUpdateOne {
	if id != nil {
		eauo = eauo.SetUserID(*id)
	}
	return eauo
}

// SetUser sets the "user" edge to the User entity.
func (eauo *ExamAttemptUpdateOne) SetUser(u *User) *ExamAttemptUpdateOne {
	return eauo.SetUserID(u.ID)
}

// Mutation returns the ExamAttemptMutation object of the builder.
func (eauo *ExamAttemptUpdateOne) Mutation() *ExamAttemptMutation {
	return eauo.mutation
}

// ClearExam clears the "exam" edge to the Exam entity.
func (eauo *ExamAttemptUpdateOne) ClearExam() *ExamAttemptUpdateOne {
	eauo.mutation.ClearExam()
	return eauo
}

// ClearUser clears the "user" edge to the User entity.
func (eauo *ExamAttemptUpdateOne) ClearUser() *ExamAttemptUpdateOne {
	eauo.mutation.ClearUser()
	return eauo
}

// Where appends a list predicates to the ExamAttemptUpdate builder.
func (eauo *ExamAttemptUpdateOne) Where(ps ...predicate.ExamAttempt) *ExamAttemptUpdateOne {
	eauo.mutation.Where(ps...)
	return eauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eauo *ExamAttemptUpdateOne) Select(field string, fields ...string) *ExamAttemptUpdateOne {
	eauo.fields = append([]string{field}, fields...)
	return eauo
}

// Save executes the query and returns the updated ExamAttempt entity.
func (eauo *ExamAttemptUpdateOne) Save(ctx context.Context) (*ExamAttempt, error) {
	eauo.defaults()
	return withHooks(ctx, eauo.sqlSave, eauo.mutation, eauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eauo *ExamAttemptUpdateOne) SaveX(ctx context.Context) *ExamAttempt {
	node, err := eauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eauo *ExamAttemptUpdateOne) Exec(ctx context.Context) error {
	_, err := eauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eauo *ExamAttemptUpdateOne) ExecX(ctx context.Context) {
	if err := eauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eauo *ExamAttemptUpdateOne) defaults() {
	if _, ok := eauo.mutation.UpdatedAt(); !ok {
		v := examattempt.UpdateDefaultUpdatedAt()
		eauo.mutation.SetUpdatedAt(v)
	}
}

func (eauo *ExamAttemptUpdateOne) sqlSave(ctx context.Context) (_node *ExamAttempt, err error) {
	_spec := sqlgraph.NewUpdateSpec(examattempt.Table, examattempt.Columns, sqlgraph.NewFieldSpec(examattempt.FieldID, field.TypeInt))
	id, ok := eauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ExamAttempt.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, examattempt.FieldID)
		for _, f := range fields {
			if !examattempt.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != examattempt.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eauo.mutation.AttemptNumber(); ok {
		_spec.SetField(examattempt.FieldAttemptNumber, field.TypeInt, value)
	}
	if value, ok := eauo.mutation.AddedAttemptNumber(); ok {
		_spec.AddField(examattempt.FieldAttemptNumber, field.TypeInt, value)
	}
	if value, ok := eauo.mutation.UpdatedAt(); ok {
		_spec.SetField(examattempt.FieldUpdatedAt, field.TypeTime, value)
	}
	if eauo.mutation.ExamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   examattempt.ExamTable,
			Columns: []string{examattempt.ExamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eauo.mutation.ExamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   examattempt.ExamTable,
			Columns: []string{examattempt.ExamColumn},
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
	if eauo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   examattempt.UserTable,
			Columns: []string{examattempt.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eauo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   examattempt.UserTable,
			Columns: []string{examattempt.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ExamAttempt{config: eauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{examattempt.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	eauo.mutation.done = true
	return _node, nil
}
