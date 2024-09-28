// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/cachedexam"
	"common/ent/exam"
	"common/ent/examcategory"
	"common/ent/examgroup"
	"common/ent/examsetting"
	"common/ent/generatedexam"
	"common/ent/predicate"
	"common/ent/subscriptionexam"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamUpdate is the builder for updating Exam entities.
type ExamUpdate struct {
	config
	hooks    []Hook
	mutation *ExamMutation
}

// Where appends a list predicates to the ExamUpdate builder.
func (eu *ExamUpdate) Where(ps ...predicate.Exam) *ExamUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *ExamUpdate) SetName(s string) *ExamUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (eu *ExamUpdate) SetNillableName(s *string) *ExamUpdate {
	if s != nil {
		eu.SetName(*s)
	}
	return eu
}

// SetStage sets the "stage" field.
func (eu *ExamUpdate) SetStage(s string) *ExamUpdate {
	eu.mutation.SetStage(s)
	return eu
}

// SetNillableStage sets the "stage" field if the given value is not nil.
func (eu *ExamUpdate) SetNillableStage(s *string) *ExamUpdate {
	if s != nil {
		eu.SetStage(*s)
	}
	return eu
}

// ClearStage clears the value of the "stage" field.
func (eu *ExamUpdate) ClearStage() *ExamUpdate {
	eu.mutation.ClearStage()
	return eu
}

// SetDescription sets the "description" field.
func (eu *ExamUpdate) SetDescription(s string) *ExamUpdate {
	eu.mutation.SetDescription(s)
	return eu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (eu *ExamUpdate) SetNillableDescription(s *string) *ExamUpdate {
	if s != nil {
		eu.SetDescription(*s)
	}
	return eu
}

// SetType sets the "type" field.
func (eu *ExamUpdate) SetType(e exam.Type) *ExamUpdate {
	eu.mutation.SetType(e)
	return eu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (eu *ExamUpdate) SetNillableType(e *exam.Type) *ExamUpdate {
	if e != nil {
		eu.SetType(*e)
	}
	return eu
}

// SetIsActive sets the "is_active" field.
func (eu *ExamUpdate) SetIsActive(b bool) *ExamUpdate {
	eu.mutation.SetIsActive(b)
	return eu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (eu *ExamUpdate) SetNillableIsActive(b *bool) *ExamUpdate {
	if b != nil {
		eu.SetIsActive(*b)
	}
	return eu
}

// SetLogoURL sets the "logo_url" field.
func (eu *ExamUpdate) SetLogoURL(s string) *ExamUpdate {
	eu.mutation.SetLogoURL(s)
	return eu
}

// SetNillableLogoURL sets the "logo_url" field if the given value is not nil.
func (eu *ExamUpdate) SetNillableLogoURL(s *string) *ExamUpdate {
	if s != nil {
		eu.SetLogoURL(*s)
	}
	return eu
}

// ClearLogoURL clears the value of the "logo_url" field.
func (eu *ExamUpdate) ClearLogoURL() *ExamUpdate {
	eu.mutation.ClearLogoURL()
	return eu
}

// SetUpdatedAt sets the "updated_at" field.
func (eu *ExamUpdate) SetUpdatedAt(t time.Time) *ExamUpdate {
	eu.mutation.SetUpdatedAt(t)
	return eu
}

// SetCategoryID sets the "category" edge to the ExamCategory entity by ID.
func (eu *ExamUpdate) SetCategoryID(id int) *ExamUpdate {
	eu.mutation.SetCategoryID(id)
	return eu
}

// SetNillableCategoryID sets the "category" edge to the ExamCategory entity by ID if the given value is not nil.
func (eu *ExamUpdate) SetNillableCategoryID(id *int) *ExamUpdate {
	if id != nil {
		eu = eu.SetCategoryID(*id)
	}
	return eu
}

// SetCategory sets the "category" edge to the ExamCategory entity.
func (eu *ExamUpdate) SetCategory(e *ExamCategory) *ExamUpdate {
	return eu.SetCategoryID(e.ID)
}

// SetGroupID sets the "group" edge to the ExamGroup entity by ID.
func (eu *ExamUpdate) SetGroupID(id int) *ExamUpdate {
	eu.mutation.SetGroupID(id)
	return eu
}

// SetNillableGroupID sets the "group" edge to the ExamGroup entity by ID if the given value is not nil.
func (eu *ExamUpdate) SetNillableGroupID(id *int) *ExamUpdate {
	if id != nil {
		eu = eu.SetGroupID(*id)
	}
	return eu
}

// SetGroup sets the "group" edge to the ExamGroup entity.
func (eu *ExamUpdate) SetGroup(e *ExamGroup) *ExamUpdate {
	return eu.SetGroupID(e.ID)
}

// AddSubscriptionIDs adds the "subscriptions" edge to the SubscriptionExam entity by IDs.
func (eu *ExamUpdate) AddSubscriptionIDs(ids ...int) *ExamUpdate {
	eu.mutation.AddSubscriptionIDs(ids...)
	return eu
}

// AddSubscriptions adds the "subscriptions" edges to the SubscriptionExam entity.
func (eu *ExamUpdate) AddSubscriptions(s ...*SubscriptionExam) *ExamUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return eu.AddSubscriptionIDs(ids...)
}

// SetSettingID sets the "setting" edge to the ExamSetting entity by ID.
func (eu *ExamUpdate) SetSettingID(id int) *ExamUpdate {
	eu.mutation.SetSettingID(id)
	return eu
}

// SetNillableSettingID sets the "setting" edge to the ExamSetting entity by ID if the given value is not nil.
func (eu *ExamUpdate) SetNillableSettingID(id *int) *ExamUpdate {
	if id != nil {
		eu = eu.SetSettingID(*id)
	}
	return eu
}

// SetSetting sets the "setting" edge to the ExamSetting entity.
func (eu *ExamUpdate) SetSetting(e *ExamSetting) *ExamUpdate {
	return eu.SetSettingID(e.ID)
}

// AddCachedExamIDs adds the "cached_exam" edge to the CachedExam entity by IDs.
func (eu *ExamUpdate) AddCachedExamIDs(ids ...int) *ExamUpdate {
	eu.mutation.AddCachedExamIDs(ids...)
	return eu
}

// AddCachedExam adds the "cached_exam" edges to the CachedExam entity.
func (eu *ExamUpdate) AddCachedExam(c ...*CachedExam) *ExamUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return eu.AddCachedExamIDs(ids...)
}

// AddGeneratedexamIDs adds the "generatedexams" edge to the GeneratedExam entity by IDs.
func (eu *ExamUpdate) AddGeneratedexamIDs(ids ...int) *ExamUpdate {
	eu.mutation.AddGeneratedexamIDs(ids...)
	return eu
}

// AddGeneratedexams adds the "generatedexams" edges to the GeneratedExam entity.
func (eu *ExamUpdate) AddGeneratedexams(g ...*GeneratedExam) *ExamUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return eu.AddGeneratedexamIDs(ids...)
}

// Mutation returns the ExamMutation object of the builder.
func (eu *ExamUpdate) Mutation() *ExamMutation {
	return eu.mutation
}

// ClearCategory clears the "category" edge to the ExamCategory entity.
func (eu *ExamUpdate) ClearCategory() *ExamUpdate {
	eu.mutation.ClearCategory()
	return eu
}

// ClearGroup clears the "group" edge to the ExamGroup entity.
func (eu *ExamUpdate) ClearGroup() *ExamUpdate {
	eu.mutation.ClearGroup()
	return eu
}

// ClearSubscriptions clears all "subscriptions" edges to the SubscriptionExam entity.
func (eu *ExamUpdate) ClearSubscriptions() *ExamUpdate {
	eu.mutation.ClearSubscriptions()
	return eu
}

// RemoveSubscriptionIDs removes the "subscriptions" edge to SubscriptionExam entities by IDs.
func (eu *ExamUpdate) RemoveSubscriptionIDs(ids ...int) *ExamUpdate {
	eu.mutation.RemoveSubscriptionIDs(ids...)
	return eu
}

// RemoveSubscriptions removes "subscriptions" edges to SubscriptionExam entities.
func (eu *ExamUpdate) RemoveSubscriptions(s ...*SubscriptionExam) *ExamUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return eu.RemoveSubscriptionIDs(ids...)
}

// ClearSetting clears the "setting" edge to the ExamSetting entity.
func (eu *ExamUpdate) ClearSetting() *ExamUpdate {
	eu.mutation.ClearSetting()
	return eu
}

// ClearCachedExam clears all "cached_exam" edges to the CachedExam entity.
func (eu *ExamUpdate) ClearCachedExam() *ExamUpdate {
	eu.mutation.ClearCachedExam()
	return eu
}

// RemoveCachedExamIDs removes the "cached_exam" edge to CachedExam entities by IDs.
func (eu *ExamUpdate) RemoveCachedExamIDs(ids ...int) *ExamUpdate {
	eu.mutation.RemoveCachedExamIDs(ids...)
	return eu
}

// RemoveCachedExam removes "cached_exam" edges to CachedExam entities.
func (eu *ExamUpdate) RemoveCachedExam(c ...*CachedExam) *ExamUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return eu.RemoveCachedExamIDs(ids...)
}

// ClearGeneratedexams clears all "generatedexams" edges to the GeneratedExam entity.
func (eu *ExamUpdate) ClearGeneratedexams() *ExamUpdate {
	eu.mutation.ClearGeneratedexams()
	return eu
}

// RemoveGeneratedexamIDs removes the "generatedexams" edge to GeneratedExam entities by IDs.
func (eu *ExamUpdate) RemoveGeneratedexamIDs(ids ...int) *ExamUpdate {
	eu.mutation.RemoveGeneratedexamIDs(ids...)
	return eu
}

// RemoveGeneratedexams removes "generatedexams" edges to GeneratedExam entities.
func (eu *ExamUpdate) RemoveGeneratedexams(g ...*GeneratedExam) *ExamUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return eu.RemoveGeneratedexamIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *ExamUpdate) Save(ctx context.Context) (int, error) {
	eu.defaults()
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *ExamUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *ExamUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *ExamUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *ExamUpdate) defaults() {
	if _, ok := eu.mutation.UpdatedAt(); !ok {
		v := exam.UpdateDefaultUpdatedAt()
		eu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *ExamUpdate) check() error {
	if v, ok := eu.mutation.GetType(); ok {
		if err := exam.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Exam.type": %w`, err)}
		}
	}
	return nil
}

func (eu *ExamUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(exam.Table, exam.Columns, sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.SetField(exam.FieldName, field.TypeString, value)
	}
	if value, ok := eu.mutation.Stage(); ok {
		_spec.SetField(exam.FieldStage, field.TypeString, value)
	}
	if eu.mutation.StageCleared() {
		_spec.ClearField(exam.FieldStage, field.TypeString)
	}
	if value, ok := eu.mutation.Description(); ok {
		_spec.SetField(exam.FieldDescription, field.TypeString, value)
	}
	if value, ok := eu.mutation.GetType(); ok {
		_spec.SetField(exam.FieldType, field.TypeEnum, value)
	}
	if value, ok := eu.mutation.IsActive(); ok {
		_spec.SetField(exam.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := eu.mutation.LogoURL(); ok {
		_spec.SetField(exam.FieldLogoURL, field.TypeString, value)
	}
	if eu.mutation.LogoURLCleared() {
		_spec.ClearField(exam.FieldLogoURL, field.TypeString)
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.SetField(exam.FieldUpdatedAt, field.TypeTime, value)
	}
	if eu.mutation.CategoryCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.CategoryIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.GroupCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.GroupIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.SubscriptionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedSubscriptionsIDs(); len(nodes) > 0 && !eu.mutation.SubscriptionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.SubscriptionsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.SettingCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.SettingIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.CachedExamCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedCachedExamIDs(); len(nodes) > 0 && !eu.mutation.CachedExamCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.CachedExamIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.GeneratedexamsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedGeneratedexamsIDs(); len(nodes) > 0 && !eu.mutation.GeneratedexamsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.GeneratedexamsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{exam.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// ExamUpdateOne is the builder for updating a single Exam entity.
type ExamUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ExamMutation
}

// SetName sets the "name" field.
func (euo *ExamUpdateOne) SetName(s string) *ExamUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (euo *ExamUpdateOne) SetNillableName(s *string) *ExamUpdateOne {
	if s != nil {
		euo.SetName(*s)
	}
	return euo
}

// SetStage sets the "stage" field.
func (euo *ExamUpdateOne) SetStage(s string) *ExamUpdateOne {
	euo.mutation.SetStage(s)
	return euo
}

// SetNillableStage sets the "stage" field if the given value is not nil.
func (euo *ExamUpdateOne) SetNillableStage(s *string) *ExamUpdateOne {
	if s != nil {
		euo.SetStage(*s)
	}
	return euo
}

// ClearStage clears the value of the "stage" field.
func (euo *ExamUpdateOne) ClearStage() *ExamUpdateOne {
	euo.mutation.ClearStage()
	return euo
}

// SetDescription sets the "description" field.
func (euo *ExamUpdateOne) SetDescription(s string) *ExamUpdateOne {
	euo.mutation.SetDescription(s)
	return euo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (euo *ExamUpdateOne) SetNillableDescription(s *string) *ExamUpdateOne {
	if s != nil {
		euo.SetDescription(*s)
	}
	return euo
}

// SetType sets the "type" field.
func (euo *ExamUpdateOne) SetType(e exam.Type) *ExamUpdateOne {
	euo.mutation.SetType(e)
	return euo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (euo *ExamUpdateOne) SetNillableType(e *exam.Type) *ExamUpdateOne {
	if e != nil {
		euo.SetType(*e)
	}
	return euo
}

// SetIsActive sets the "is_active" field.
func (euo *ExamUpdateOne) SetIsActive(b bool) *ExamUpdateOne {
	euo.mutation.SetIsActive(b)
	return euo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (euo *ExamUpdateOne) SetNillableIsActive(b *bool) *ExamUpdateOne {
	if b != nil {
		euo.SetIsActive(*b)
	}
	return euo
}

// SetLogoURL sets the "logo_url" field.
func (euo *ExamUpdateOne) SetLogoURL(s string) *ExamUpdateOne {
	euo.mutation.SetLogoURL(s)
	return euo
}

// SetNillableLogoURL sets the "logo_url" field if the given value is not nil.
func (euo *ExamUpdateOne) SetNillableLogoURL(s *string) *ExamUpdateOne {
	if s != nil {
		euo.SetLogoURL(*s)
	}
	return euo
}

// ClearLogoURL clears the value of the "logo_url" field.
func (euo *ExamUpdateOne) ClearLogoURL() *ExamUpdateOne {
	euo.mutation.ClearLogoURL()
	return euo
}

// SetUpdatedAt sets the "updated_at" field.
func (euo *ExamUpdateOne) SetUpdatedAt(t time.Time) *ExamUpdateOne {
	euo.mutation.SetUpdatedAt(t)
	return euo
}

// SetCategoryID sets the "category" edge to the ExamCategory entity by ID.
func (euo *ExamUpdateOne) SetCategoryID(id int) *ExamUpdateOne {
	euo.mutation.SetCategoryID(id)
	return euo
}

// SetNillableCategoryID sets the "category" edge to the ExamCategory entity by ID if the given value is not nil.
func (euo *ExamUpdateOne) SetNillableCategoryID(id *int) *ExamUpdateOne {
	if id != nil {
		euo = euo.SetCategoryID(*id)
	}
	return euo
}

// SetCategory sets the "category" edge to the ExamCategory entity.
func (euo *ExamUpdateOne) SetCategory(e *ExamCategory) *ExamUpdateOne {
	return euo.SetCategoryID(e.ID)
}

// SetGroupID sets the "group" edge to the ExamGroup entity by ID.
func (euo *ExamUpdateOne) SetGroupID(id int) *ExamUpdateOne {
	euo.mutation.SetGroupID(id)
	return euo
}

// SetNillableGroupID sets the "group" edge to the ExamGroup entity by ID if the given value is not nil.
func (euo *ExamUpdateOne) SetNillableGroupID(id *int) *ExamUpdateOne {
	if id != nil {
		euo = euo.SetGroupID(*id)
	}
	return euo
}

// SetGroup sets the "group" edge to the ExamGroup entity.
func (euo *ExamUpdateOne) SetGroup(e *ExamGroup) *ExamUpdateOne {
	return euo.SetGroupID(e.ID)
}

// AddSubscriptionIDs adds the "subscriptions" edge to the SubscriptionExam entity by IDs.
func (euo *ExamUpdateOne) AddSubscriptionIDs(ids ...int) *ExamUpdateOne {
	euo.mutation.AddSubscriptionIDs(ids...)
	return euo
}

// AddSubscriptions adds the "subscriptions" edges to the SubscriptionExam entity.
func (euo *ExamUpdateOne) AddSubscriptions(s ...*SubscriptionExam) *ExamUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return euo.AddSubscriptionIDs(ids...)
}

// SetSettingID sets the "setting" edge to the ExamSetting entity by ID.
func (euo *ExamUpdateOne) SetSettingID(id int) *ExamUpdateOne {
	euo.mutation.SetSettingID(id)
	return euo
}

// SetNillableSettingID sets the "setting" edge to the ExamSetting entity by ID if the given value is not nil.
func (euo *ExamUpdateOne) SetNillableSettingID(id *int) *ExamUpdateOne {
	if id != nil {
		euo = euo.SetSettingID(*id)
	}
	return euo
}

// SetSetting sets the "setting" edge to the ExamSetting entity.
func (euo *ExamUpdateOne) SetSetting(e *ExamSetting) *ExamUpdateOne {
	return euo.SetSettingID(e.ID)
}

// AddCachedExamIDs adds the "cached_exam" edge to the CachedExam entity by IDs.
func (euo *ExamUpdateOne) AddCachedExamIDs(ids ...int) *ExamUpdateOne {
	euo.mutation.AddCachedExamIDs(ids...)
	return euo
}

// AddCachedExam adds the "cached_exam" edges to the CachedExam entity.
func (euo *ExamUpdateOne) AddCachedExam(c ...*CachedExam) *ExamUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return euo.AddCachedExamIDs(ids...)
}

// AddGeneratedexamIDs adds the "generatedexams" edge to the GeneratedExam entity by IDs.
func (euo *ExamUpdateOne) AddGeneratedexamIDs(ids ...int) *ExamUpdateOne {
	euo.mutation.AddGeneratedexamIDs(ids...)
	return euo
}

// AddGeneratedexams adds the "generatedexams" edges to the GeneratedExam entity.
func (euo *ExamUpdateOne) AddGeneratedexams(g ...*GeneratedExam) *ExamUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return euo.AddGeneratedexamIDs(ids...)
}

// Mutation returns the ExamMutation object of the builder.
func (euo *ExamUpdateOne) Mutation() *ExamMutation {
	return euo.mutation
}

// ClearCategory clears the "category" edge to the ExamCategory entity.
func (euo *ExamUpdateOne) ClearCategory() *ExamUpdateOne {
	euo.mutation.ClearCategory()
	return euo
}

// ClearGroup clears the "group" edge to the ExamGroup entity.
func (euo *ExamUpdateOne) ClearGroup() *ExamUpdateOne {
	euo.mutation.ClearGroup()
	return euo
}

// ClearSubscriptions clears all "subscriptions" edges to the SubscriptionExam entity.
func (euo *ExamUpdateOne) ClearSubscriptions() *ExamUpdateOne {
	euo.mutation.ClearSubscriptions()
	return euo
}

// RemoveSubscriptionIDs removes the "subscriptions" edge to SubscriptionExam entities by IDs.
func (euo *ExamUpdateOne) RemoveSubscriptionIDs(ids ...int) *ExamUpdateOne {
	euo.mutation.RemoveSubscriptionIDs(ids...)
	return euo
}

// RemoveSubscriptions removes "subscriptions" edges to SubscriptionExam entities.
func (euo *ExamUpdateOne) RemoveSubscriptions(s ...*SubscriptionExam) *ExamUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return euo.RemoveSubscriptionIDs(ids...)
}

// ClearSetting clears the "setting" edge to the ExamSetting entity.
func (euo *ExamUpdateOne) ClearSetting() *ExamUpdateOne {
	euo.mutation.ClearSetting()
	return euo
}

// ClearCachedExam clears all "cached_exam" edges to the CachedExam entity.
func (euo *ExamUpdateOne) ClearCachedExam() *ExamUpdateOne {
	euo.mutation.ClearCachedExam()
	return euo
}

// RemoveCachedExamIDs removes the "cached_exam" edge to CachedExam entities by IDs.
func (euo *ExamUpdateOne) RemoveCachedExamIDs(ids ...int) *ExamUpdateOne {
	euo.mutation.RemoveCachedExamIDs(ids...)
	return euo
}

// RemoveCachedExam removes "cached_exam" edges to CachedExam entities.
func (euo *ExamUpdateOne) RemoveCachedExam(c ...*CachedExam) *ExamUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return euo.RemoveCachedExamIDs(ids...)
}

// ClearGeneratedexams clears all "generatedexams" edges to the GeneratedExam entity.
func (euo *ExamUpdateOne) ClearGeneratedexams() *ExamUpdateOne {
	euo.mutation.ClearGeneratedexams()
	return euo
}

// RemoveGeneratedexamIDs removes the "generatedexams" edge to GeneratedExam entities by IDs.
func (euo *ExamUpdateOne) RemoveGeneratedexamIDs(ids ...int) *ExamUpdateOne {
	euo.mutation.RemoveGeneratedexamIDs(ids...)
	return euo
}

// RemoveGeneratedexams removes "generatedexams" edges to GeneratedExam entities.
func (euo *ExamUpdateOne) RemoveGeneratedexams(g ...*GeneratedExam) *ExamUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return euo.RemoveGeneratedexamIDs(ids...)
}

// Where appends a list predicates to the ExamUpdate builder.
func (euo *ExamUpdateOne) Where(ps ...predicate.Exam) *ExamUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *ExamUpdateOne) Select(field string, fields ...string) *ExamUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Exam entity.
func (euo *ExamUpdateOne) Save(ctx context.Context) (*Exam, error) {
	euo.defaults()
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *ExamUpdateOne) SaveX(ctx context.Context) *Exam {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *ExamUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *ExamUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *ExamUpdateOne) defaults() {
	if _, ok := euo.mutation.UpdatedAt(); !ok {
		v := exam.UpdateDefaultUpdatedAt()
		euo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *ExamUpdateOne) check() error {
	if v, ok := euo.mutation.GetType(); ok {
		if err := exam.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Exam.type": %w`, err)}
		}
	}
	return nil
}

func (euo *ExamUpdateOne) sqlSave(ctx context.Context) (_node *Exam, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(exam.Table, exam.Columns, sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Exam.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, exam.FieldID)
		for _, f := range fields {
			if !exam.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != exam.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.SetField(exam.FieldName, field.TypeString, value)
	}
	if value, ok := euo.mutation.Stage(); ok {
		_spec.SetField(exam.FieldStage, field.TypeString, value)
	}
	if euo.mutation.StageCleared() {
		_spec.ClearField(exam.FieldStage, field.TypeString)
	}
	if value, ok := euo.mutation.Description(); ok {
		_spec.SetField(exam.FieldDescription, field.TypeString, value)
	}
	if value, ok := euo.mutation.GetType(); ok {
		_spec.SetField(exam.FieldType, field.TypeEnum, value)
	}
	if value, ok := euo.mutation.IsActive(); ok {
		_spec.SetField(exam.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := euo.mutation.LogoURL(); ok {
		_spec.SetField(exam.FieldLogoURL, field.TypeString, value)
	}
	if euo.mutation.LogoURLCleared() {
		_spec.ClearField(exam.FieldLogoURL, field.TypeString)
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.SetField(exam.FieldUpdatedAt, field.TypeTime, value)
	}
	if euo.mutation.CategoryCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.CategoryIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.GroupCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.GroupIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.SubscriptionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedSubscriptionsIDs(); len(nodes) > 0 && !euo.mutation.SubscriptionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.SubscriptionsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.SettingCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.SettingIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.CachedExamCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedCachedExamIDs(); len(nodes) > 0 && !euo.mutation.CachedExamCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.CachedExamIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.GeneratedexamsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedGeneratedexamsIDs(); len(nodes) > 0 && !euo.mutation.GeneratedexamsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.GeneratedexamsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Exam{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{exam.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
