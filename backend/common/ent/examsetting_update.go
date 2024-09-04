// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/exam"
	"common/ent/examsetting"
	"common/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamSettingUpdate is the builder for updating ExamSetting entities.
type ExamSettingUpdate struct {
	config
	hooks    []Hook
	mutation *ExamSettingMutation
}

// Where appends a list predicates to the ExamSettingUpdate builder.
func (esu *ExamSettingUpdate) Where(ps ...predicate.ExamSetting) *ExamSettingUpdate {
	esu.mutation.Where(ps...)
	return esu
}

// SetNumberOfQuestions sets the "number_of_questions" field.
func (esu *ExamSettingUpdate) SetNumberOfQuestions(i int) *ExamSettingUpdate {
	esu.mutation.ResetNumberOfQuestions()
	esu.mutation.SetNumberOfQuestions(i)
	return esu
}

// SetNillableNumberOfQuestions sets the "number_of_questions" field if the given value is not nil.
func (esu *ExamSettingUpdate) SetNillableNumberOfQuestions(i *int) *ExamSettingUpdate {
	if i != nil {
		esu.SetNumberOfQuestions(*i)
	}
	return esu
}

// AddNumberOfQuestions adds i to the "number_of_questions" field.
func (esu *ExamSettingUpdate) AddNumberOfQuestions(i int) *ExamSettingUpdate {
	esu.mutation.AddNumberOfQuestions(i)
	return esu
}

// SetDurationSeconds sets the "duration_seconds" field.
func (esu *ExamSettingUpdate) SetDurationSeconds(i int) *ExamSettingUpdate {
	esu.mutation.ResetDurationSeconds()
	esu.mutation.SetDurationSeconds(i)
	return esu
}

// SetNillableDurationSeconds sets the "duration_seconds" field if the given value is not nil.
func (esu *ExamSettingUpdate) SetNillableDurationSeconds(i *int) *ExamSettingUpdate {
	if i != nil {
		esu.SetDurationSeconds(*i)
	}
	return esu
}

// AddDurationSeconds adds i to the "duration_seconds" field.
func (esu *ExamSettingUpdate) AddDurationSeconds(i int) *ExamSettingUpdate {
	esu.mutation.AddDurationSeconds(i)
	return esu
}

// SetNegativeMarking sets the "negative_marking" field.
func (esu *ExamSettingUpdate) SetNegativeMarking(f float64) *ExamSettingUpdate {
	esu.mutation.ResetNegativeMarking()
	esu.mutation.SetNegativeMarking(f)
	return esu
}

// SetNillableNegativeMarking sets the "negative_marking" field if the given value is not nil.
func (esu *ExamSettingUpdate) SetNillableNegativeMarking(f *float64) *ExamSettingUpdate {
	if f != nil {
		esu.SetNegativeMarking(*f)
	}
	return esu
}

// AddNegativeMarking adds f to the "negative_marking" field.
func (esu *ExamSettingUpdate) AddNegativeMarking(f float64) *ExamSettingUpdate {
	esu.mutation.AddNegativeMarking(f)
	return esu
}

// ClearNegativeMarking clears the value of the "negative_marking" field.
func (esu *ExamSettingUpdate) ClearNegativeMarking() *ExamSettingUpdate {
	esu.mutation.ClearNegativeMarking()
	return esu
}

// SetAiPrompt sets the "ai_prompt" field.
func (esu *ExamSettingUpdate) SetAiPrompt(s string) *ExamSettingUpdate {
	esu.mutation.SetAiPrompt(s)
	return esu
}

// SetNillableAiPrompt sets the "ai_prompt" field if the given value is not nil.
func (esu *ExamSettingUpdate) SetNillableAiPrompt(s *string) *ExamSettingUpdate {
	if s != nil {
		esu.SetAiPrompt(*s)
	}
	return esu
}

// ClearAiPrompt clears the value of the "ai_prompt" field.
func (esu *ExamSettingUpdate) ClearAiPrompt() *ExamSettingUpdate {
	esu.mutation.ClearAiPrompt()
	return esu
}

// SetOtherDetails sets the "other_details" field.
func (esu *ExamSettingUpdate) SetOtherDetails(m map[string]interface{}) *ExamSettingUpdate {
	esu.mutation.SetOtherDetails(m)
	return esu
}

// ClearOtherDetails clears the value of the "other_details" field.
func (esu *ExamSettingUpdate) ClearOtherDetails() *ExamSettingUpdate {
	esu.mutation.ClearOtherDetails()
	return esu
}

// SetMaxAttempts sets the "max_attempts" field.
func (esu *ExamSettingUpdate) SetMaxAttempts(i int) *ExamSettingUpdate {
	esu.mutation.ResetMaxAttempts()
	esu.mutation.SetMaxAttempts(i)
	return esu
}

// SetNillableMaxAttempts sets the "max_attempts" field if the given value is not nil.
func (esu *ExamSettingUpdate) SetNillableMaxAttempts(i *int) *ExamSettingUpdate {
	if i != nil {
		esu.SetMaxAttempts(*i)
	}
	return esu
}

// AddMaxAttempts adds i to the "max_attempts" field.
func (esu *ExamSettingUpdate) AddMaxAttempts(i int) *ExamSettingUpdate {
	esu.mutation.AddMaxAttempts(i)
	return esu
}

// SetEvaluationAiPrompt sets the "evaluation_ai_prompt" field.
func (esu *ExamSettingUpdate) SetEvaluationAiPrompt(s string) *ExamSettingUpdate {
	esu.mutation.SetEvaluationAiPrompt(s)
	return esu
}

// SetNillableEvaluationAiPrompt sets the "evaluation_ai_prompt" field if the given value is not nil.
func (esu *ExamSettingUpdate) SetNillableEvaluationAiPrompt(s *string) *ExamSettingUpdate {
	if s != nil {
		esu.SetEvaluationAiPrompt(*s)
	}
	return esu
}

// ClearEvaluationAiPrompt clears the value of the "evaluation_ai_prompt" field.
func (esu *ExamSettingUpdate) ClearEvaluationAiPrompt() *ExamSettingUpdate {
	esu.mutation.ClearEvaluationAiPrompt()
	return esu
}

// SetUpdatedAt sets the "updated_at" field.
func (esu *ExamSettingUpdate) SetUpdatedAt(t time.Time) *ExamSettingUpdate {
	esu.mutation.SetUpdatedAt(t)
	return esu
}

// SetExamID sets the "exam" edge to the Exam entity by ID.
func (esu *ExamSettingUpdate) SetExamID(id int) *ExamSettingUpdate {
	esu.mutation.SetExamID(id)
	return esu
}

// SetNillableExamID sets the "exam" edge to the Exam entity by ID if the given value is not nil.
func (esu *ExamSettingUpdate) SetNillableExamID(id *int) *ExamSettingUpdate {
	if id != nil {
		esu = esu.SetExamID(*id)
	}
	return esu
}

// SetExam sets the "exam" edge to the Exam entity.
func (esu *ExamSettingUpdate) SetExam(e *Exam) *ExamSettingUpdate {
	return esu.SetExamID(e.ID)
}

// Mutation returns the ExamSettingMutation object of the builder.
func (esu *ExamSettingUpdate) Mutation() *ExamSettingMutation {
	return esu.mutation
}

// ClearExam clears the "exam" edge to the Exam entity.
func (esu *ExamSettingUpdate) ClearExam() *ExamSettingUpdate {
	esu.mutation.ClearExam()
	return esu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (esu *ExamSettingUpdate) Save(ctx context.Context) (int, error) {
	esu.defaults()
	return withHooks(ctx, esu.sqlSave, esu.mutation, esu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (esu *ExamSettingUpdate) SaveX(ctx context.Context) int {
	affected, err := esu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (esu *ExamSettingUpdate) Exec(ctx context.Context) error {
	_, err := esu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (esu *ExamSettingUpdate) ExecX(ctx context.Context) {
	if err := esu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (esu *ExamSettingUpdate) defaults() {
	if _, ok := esu.mutation.UpdatedAt(); !ok {
		v := examsetting.UpdateDefaultUpdatedAt()
		esu.mutation.SetUpdatedAt(v)
	}
}

func (esu *ExamSettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(examsetting.Table, examsetting.Columns, sqlgraph.NewFieldSpec(examsetting.FieldID, field.TypeInt))
	if ps := esu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := esu.mutation.NumberOfQuestions(); ok {
		_spec.SetField(examsetting.FieldNumberOfQuestions, field.TypeInt, value)
	}
	if value, ok := esu.mutation.AddedNumberOfQuestions(); ok {
		_spec.AddField(examsetting.FieldNumberOfQuestions, field.TypeInt, value)
	}
	if value, ok := esu.mutation.DurationSeconds(); ok {
		_spec.SetField(examsetting.FieldDurationSeconds, field.TypeInt, value)
	}
	if value, ok := esu.mutation.AddedDurationSeconds(); ok {
		_spec.AddField(examsetting.FieldDurationSeconds, field.TypeInt, value)
	}
	if value, ok := esu.mutation.NegativeMarking(); ok {
		_spec.SetField(examsetting.FieldNegativeMarking, field.TypeFloat64, value)
	}
	if value, ok := esu.mutation.AddedNegativeMarking(); ok {
		_spec.AddField(examsetting.FieldNegativeMarking, field.TypeFloat64, value)
	}
	if esu.mutation.NegativeMarkingCleared() {
		_spec.ClearField(examsetting.FieldNegativeMarking, field.TypeFloat64)
	}
	if value, ok := esu.mutation.AiPrompt(); ok {
		_spec.SetField(examsetting.FieldAiPrompt, field.TypeString, value)
	}
	if esu.mutation.AiPromptCleared() {
		_spec.ClearField(examsetting.FieldAiPrompt, field.TypeString)
	}
	if value, ok := esu.mutation.OtherDetails(); ok {
		_spec.SetField(examsetting.FieldOtherDetails, field.TypeJSON, value)
	}
	if esu.mutation.OtherDetailsCleared() {
		_spec.ClearField(examsetting.FieldOtherDetails, field.TypeJSON)
	}
	if value, ok := esu.mutation.MaxAttempts(); ok {
		_spec.SetField(examsetting.FieldMaxAttempts, field.TypeInt, value)
	}
	if value, ok := esu.mutation.AddedMaxAttempts(); ok {
		_spec.AddField(examsetting.FieldMaxAttempts, field.TypeInt, value)
	}
	if value, ok := esu.mutation.EvaluationAiPrompt(); ok {
		_spec.SetField(examsetting.FieldEvaluationAiPrompt, field.TypeString, value)
	}
	if esu.mutation.EvaluationAiPromptCleared() {
		_spec.ClearField(examsetting.FieldEvaluationAiPrompt, field.TypeString)
	}
	if value, ok := esu.mutation.UpdatedAt(); ok {
		_spec.SetField(examsetting.FieldUpdatedAt, field.TypeTime, value)
	}
	if esu.mutation.ExamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   examsetting.ExamTable,
			Columns: []string{examsetting.ExamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esu.mutation.ExamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   examsetting.ExamTable,
			Columns: []string{examsetting.ExamColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, esu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{examsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	esu.mutation.done = true
	return n, nil
}

// ExamSettingUpdateOne is the builder for updating a single ExamSetting entity.
type ExamSettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ExamSettingMutation
}

// SetNumberOfQuestions sets the "number_of_questions" field.
func (esuo *ExamSettingUpdateOne) SetNumberOfQuestions(i int) *ExamSettingUpdateOne {
	esuo.mutation.ResetNumberOfQuestions()
	esuo.mutation.SetNumberOfQuestions(i)
	return esuo
}

// SetNillableNumberOfQuestions sets the "number_of_questions" field if the given value is not nil.
func (esuo *ExamSettingUpdateOne) SetNillableNumberOfQuestions(i *int) *ExamSettingUpdateOne {
	if i != nil {
		esuo.SetNumberOfQuestions(*i)
	}
	return esuo
}

// AddNumberOfQuestions adds i to the "number_of_questions" field.
func (esuo *ExamSettingUpdateOne) AddNumberOfQuestions(i int) *ExamSettingUpdateOne {
	esuo.mutation.AddNumberOfQuestions(i)
	return esuo
}

// SetDurationSeconds sets the "duration_seconds" field.
func (esuo *ExamSettingUpdateOne) SetDurationSeconds(i int) *ExamSettingUpdateOne {
	esuo.mutation.ResetDurationSeconds()
	esuo.mutation.SetDurationSeconds(i)
	return esuo
}

// SetNillableDurationSeconds sets the "duration_seconds" field if the given value is not nil.
func (esuo *ExamSettingUpdateOne) SetNillableDurationSeconds(i *int) *ExamSettingUpdateOne {
	if i != nil {
		esuo.SetDurationSeconds(*i)
	}
	return esuo
}

// AddDurationSeconds adds i to the "duration_seconds" field.
func (esuo *ExamSettingUpdateOne) AddDurationSeconds(i int) *ExamSettingUpdateOne {
	esuo.mutation.AddDurationSeconds(i)
	return esuo
}

// SetNegativeMarking sets the "negative_marking" field.
func (esuo *ExamSettingUpdateOne) SetNegativeMarking(f float64) *ExamSettingUpdateOne {
	esuo.mutation.ResetNegativeMarking()
	esuo.mutation.SetNegativeMarking(f)
	return esuo
}

// SetNillableNegativeMarking sets the "negative_marking" field if the given value is not nil.
func (esuo *ExamSettingUpdateOne) SetNillableNegativeMarking(f *float64) *ExamSettingUpdateOne {
	if f != nil {
		esuo.SetNegativeMarking(*f)
	}
	return esuo
}

// AddNegativeMarking adds f to the "negative_marking" field.
func (esuo *ExamSettingUpdateOne) AddNegativeMarking(f float64) *ExamSettingUpdateOne {
	esuo.mutation.AddNegativeMarking(f)
	return esuo
}

// ClearNegativeMarking clears the value of the "negative_marking" field.
func (esuo *ExamSettingUpdateOne) ClearNegativeMarking() *ExamSettingUpdateOne {
	esuo.mutation.ClearNegativeMarking()
	return esuo
}

// SetAiPrompt sets the "ai_prompt" field.
func (esuo *ExamSettingUpdateOne) SetAiPrompt(s string) *ExamSettingUpdateOne {
	esuo.mutation.SetAiPrompt(s)
	return esuo
}

// SetNillableAiPrompt sets the "ai_prompt" field if the given value is not nil.
func (esuo *ExamSettingUpdateOne) SetNillableAiPrompt(s *string) *ExamSettingUpdateOne {
	if s != nil {
		esuo.SetAiPrompt(*s)
	}
	return esuo
}

// ClearAiPrompt clears the value of the "ai_prompt" field.
func (esuo *ExamSettingUpdateOne) ClearAiPrompt() *ExamSettingUpdateOne {
	esuo.mutation.ClearAiPrompt()
	return esuo
}

// SetOtherDetails sets the "other_details" field.
func (esuo *ExamSettingUpdateOne) SetOtherDetails(m map[string]interface{}) *ExamSettingUpdateOne {
	esuo.mutation.SetOtherDetails(m)
	return esuo
}

// ClearOtherDetails clears the value of the "other_details" field.
func (esuo *ExamSettingUpdateOne) ClearOtherDetails() *ExamSettingUpdateOne {
	esuo.mutation.ClearOtherDetails()
	return esuo
}

// SetMaxAttempts sets the "max_attempts" field.
func (esuo *ExamSettingUpdateOne) SetMaxAttempts(i int) *ExamSettingUpdateOne {
	esuo.mutation.ResetMaxAttempts()
	esuo.mutation.SetMaxAttempts(i)
	return esuo
}

// SetNillableMaxAttempts sets the "max_attempts" field if the given value is not nil.
func (esuo *ExamSettingUpdateOne) SetNillableMaxAttempts(i *int) *ExamSettingUpdateOne {
	if i != nil {
		esuo.SetMaxAttempts(*i)
	}
	return esuo
}

// AddMaxAttempts adds i to the "max_attempts" field.
func (esuo *ExamSettingUpdateOne) AddMaxAttempts(i int) *ExamSettingUpdateOne {
	esuo.mutation.AddMaxAttempts(i)
	return esuo
}

// SetEvaluationAiPrompt sets the "evaluation_ai_prompt" field.
func (esuo *ExamSettingUpdateOne) SetEvaluationAiPrompt(s string) *ExamSettingUpdateOne {
	esuo.mutation.SetEvaluationAiPrompt(s)
	return esuo
}

// SetNillableEvaluationAiPrompt sets the "evaluation_ai_prompt" field if the given value is not nil.
func (esuo *ExamSettingUpdateOne) SetNillableEvaluationAiPrompt(s *string) *ExamSettingUpdateOne {
	if s != nil {
		esuo.SetEvaluationAiPrompt(*s)
	}
	return esuo
}

// ClearEvaluationAiPrompt clears the value of the "evaluation_ai_prompt" field.
func (esuo *ExamSettingUpdateOne) ClearEvaluationAiPrompt() *ExamSettingUpdateOne {
	esuo.mutation.ClearEvaluationAiPrompt()
	return esuo
}

// SetUpdatedAt sets the "updated_at" field.
func (esuo *ExamSettingUpdateOne) SetUpdatedAt(t time.Time) *ExamSettingUpdateOne {
	esuo.mutation.SetUpdatedAt(t)
	return esuo
}

// SetExamID sets the "exam" edge to the Exam entity by ID.
func (esuo *ExamSettingUpdateOne) SetExamID(id int) *ExamSettingUpdateOne {
	esuo.mutation.SetExamID(id)
	return esuo
}

// SetNillableExamID sets the "exam" edge to the Exam entity by ID if the given value is not nil.
func (esuo *ExamSettingUpdateOne) SetNillableExamID(id *int) *ExamSettingUpdateOne {
	if id != nil {
		esuo = esuo.SetExamID(*id)
	}
	return esuo
}

// SetExam sets the "exam" edge to the Exam entity.
func (esuo *ExamSettingUpdateOne) SetExam(e *Exam) *ExamSettingUpdateOne {
	return esuo.SetExamID(e.ID)
}

// Mutation returns the ExamSettingMutation object of the builder.
func (esuo *ExamSettingUpdateOne) Mutation() *ExamSettingMutation {
	return esuo.mutation
}

// ClearExam clears the "exam" edge to the Exam entity.
func (esuo *ExamSettingUpdateOne) ClearExam() *ExamSettingUpdateOne {
	esuo.mutation.ClearExam()
	return esuo
}

// Where appends a list predicates to the ExamSettingUpdate builder.
func (esuo *ExamSettingUpdateOne) Where(ps ...predicate.ExamSetting) *ExamSettingUpdateOne {
	esuo.mutation.Where(ps...)
	return esuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (esuo *ExamSettingUpdateOne) Select(field string, fields ...string) *ExamSettingUpdateOne {
	esuo.fields = append([]string{field}, fields...)
	return esuo
}

// Save executes the query and returns the updated ExamSetting entity.
func (esuo *ExamSettingUpdateOne) Save(ctx context.Context) (*ExamSetting, error) {
	esuo.defaults()
	return withHooks(ctx, esuo.sqlSave, esuo.mutation, esuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (esuo *ExamSettingUpdateOne) SaveX(ctx context.Context) *ExamSetting {
	node, err := esuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (esuo *ExamSettingUpdateOne) Exec(ctx context.Context) error {
	_, err := esuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (esuo *ExamSettingUpdateOne) ExecX(ctx context.Context) {
	if err := esuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (esuo *ExamSettingUpdateOne) defaults() {
	if _, ok := esuo.mutation.UpdatedAt(); !ok {
		v := examsetting.UpdateDefaultUpdatedAt()
		esuo.mutation.SetUpdatedAt(v)
	}
}

func (esuo *ExamSettingUpdateOne) sqlSave(ctx context.Context) (_node *ExamSetting, err error) {
	_spec := sqlgraph.NewUpdateSpec(examsetting.Table, examsetting.Columns, sqlgraph.NewFieldSpec(examsetting.FieldID, field.TypeInt))
	id, ok := esuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ExamSetting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := esuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, examsetting.FieldID)
		for _, f := range fields {
			if !examsetting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != examsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := esuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := esuo.mutation.NumberOfQuestions(); ok {
		_spec.SetField(examsetting.FieldNumberOfQuestions, field.TypeInt, value)
	}
	if value, ok := esuo.mutation.AddedNumberOfQuestions(); ok {
		_spec.AddField(examsetting.FieldNumberOfQuestions, field.TypeInt, value)
	}
	if value, ok := esuo.mutation.DurationSeconds(); ok {
		_spec.SetField(examsetting.FieldDurationSeconds, field.TypeInt, value)
	}
	if value, ok := esuo.mutation.AddedDurationSeconds(); ok {
		_spec.AddField(examsetting.FieldDurationSeconds, field.TypeInt, value)
	}
	if value, ok := esuo.mutation.NegativeMarking(); ok {
		_spec.SetField(examsetting.FieldNegativeMarking, field.TypeFloat64, value)
	}
	if value, ok := esuo.mutation.AddedNegativeMarking(); ok {
		_spec.AddField(examsetting.FieldNegativeMarking, field.TypeFloat64, value)
	}
	if esuo.mutation.NegativeMarkingCleared() {
		_spec.ClearField(examsetting.FieldNegativeMarking, field.TypeFloat64)
	}
	if value, ok := esuo.mutation.AiPrompt(); ok {
		_spec.SetField(examsetting.FieldAiPrompt, field.TypeString, value)
	}
	if esuo.mutation.AiPromptCleared() {
		_spec.ClearField(examsetting.FieldAiPrompt, field.TypeString)
	}
	if value, ok := esuo.mutation.OtherDetails(); ok {
		_spec.SetField(examsetting.FieldOtherDetails, field.TypeJSON, value)
	}
	if esuo.mutation.OtherDetailsCleared() {
		_spec.ClearField(examsetting.FieldOtherDetails, field.TypeJSON)
	}
	if value, ok := esuo.mutation.MaxAttempts(); ok {
		_spec.SetField(examsetting.FieldMaxAttempts, field.TypeInt, value)
	}
	if value, ok := esuo.mutation.AddedMaxAttempts(); ok {
		_spec.AddField(examsetting.FieldMaxAttempts, field.TypeInt, value)
	}
	if value, ok := esuo.mutation.EvaluationAiPrompt(); ok {
		_spec.SetField(examsetting.FieldEvaluationAiPrompt, field.TypeString, value)
	}
	if esuo.mutation.EvaluationAiPromptCleared() {
		_spec.ClearField(examsetting.FieldEvaluationAiPrompt, field.TypeString)
	}
	if value, ok := esuo.mutation.UpdatedAt(); ok {
		_spec.SetField(examsetting.FieldUpdatedAt, field.TypeTime, value)
	}
	if esuo.mutation.ExamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   examsetting.ExamTable,
			Columns: []string{examsetting.ExamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esuo.mutation.ExamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   examsetting.ExamTable,
			Columns: []string{examsetting.ExamColumn},
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
	_node = &ExamSetting{config: esuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, esuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{examsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	esuo.mutation.done = true
	return _node, nil
}
