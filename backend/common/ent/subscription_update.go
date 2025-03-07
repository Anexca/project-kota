// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/predicate"
	"common/ent/subscription"
	"common/ent/subscriptionexam"
	"common/ent/usersubscription"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubscriptionUpdate is the builder for updating Subscription entities.
type SubscriptionUpdate struct {
	config
	hooks    []Hook
	mutation *SubscriptionMutation
}

// Where appends a list predicates to the SubscriptionUpdate builder.
func (su *SubscriptionUpdate) Where(ps ...predicate.Subscription) *SubscriptionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetProviderPlanID sets the "provider_plan_id" field.
func (su *SubscriptionUpdate) SetProviderPlanID(s string) *SubscriptionUpdate {
	su.mutation.SetProviderPlanID(s)
	return su
}

// SetNillableProviderPlanID sets the "provider_plan_id" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableProviderPlanID(s *string) *SubscriptionUpdate {
	if s != nil {
		su.SetProviderPlanID(*s)
	}
	return su
}

// SetBasePrice sets the "base_price" field.
func (su *SubscriptionUpdate) SetBasePrice(f float64) *SubscriptionUpdate {
	su.mutation.ResetBasePrice()
	su.mutation.SetBasePrice(f)
	return su
}

// SetNillableBasePrice sets the "base_price" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableBasePrice(f *float64) *SubscriptionUpdate {
	if f != nil {
		su.SetBasePrice(*f)
	}
	return su
}

// AddBasePrice adds f to the "base_price" field.
func (su *SubscriptionUpdate) AddBasePrice(f float64) *SubscriptionUpdate {
	su.mutation.AddBasePrice(f)
	return su
}

// ClearBasePrice clears the value of the "base_price" field.
func (su *SubscriptionUpdate) ClearBasePrice() *SubscriptionUpdate {
	su.mutation.ClearBasePrice()
	return su
}

// SetFinalPrice sets the "final_price" field.
func (su *SubscriptionUpdate) SetFinalPrice(f float64) *SubscriptionUpdate {
	su.mutation.ResetFinalPrice()
	su.mutation.SetFinalPrice(f)
	return su
}

// SetNillableFinalPrice sets the "final_price" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableFinalPrice(f *float64) *SubscriptionUpdate {
	if f != nil {
		su.SetFinalPrice(*f)
	}
	return su
}

// AddFinalPrice adds f to the "final_price" field.
func (su *SubscriptionUpdate) AddFinalPrice(f float64) *SubscriptionUpdate {
	su.mutation.AddFinalPrice(f)
	return su
}

// ClearFinalPrice clears the value of the "final_price" field.
func (su *SubscriptionUpdate) ClearFinalPrice() *SubscriptionUpdate {
	su.mutation.ClearFinalPrice()
	return su
}

// SetPrice sets the "price" field.
func (su *SubscriptionUpdate) SetPrice(f float64) *SubscriptionUpdate {
	su.mutation.ResetPrice()
	su.mutation.SetPrice(f)
	return su
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillablePrice(f *float64) *SubscriptionUpdate {
	if f != nil {
		su.SetPrice(*f)
	}
	return su
}

// AddPrice adds f to the "price" field.
func (su *SubscriptionUpdate) AddPrice(f float64) *SubscriptionUpdate {
	su.mutation.AddPrice(f)
	return su
}

// ClearPrice clears the value of the "price" field.
func (su *SubscriptionUpdate) ClearPrice() *SubscriptionUpdate {
	su.mutation.ClearPrice()
	return su
}

// SetDurationInMonths sets the "duration_in_months" field.
func (su *SubscriptionUpdate) SetDurationInMonths(i int) *SubscriptionUpdate {
	su.mutation.ResetDurationInMonths()
	su.mutation.SetDurationInMonths(i)
	return su
}

// SetNillableDurationInMonths sets the "duration_in_months" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableDurationInMonths(i *int) *SubscriptionUpdate {
	if i != nil {
		su.SetDurationInMonths(*i)
	}
	return su
}

// AddDurationInMonths adds i to the "duration_in_months" field.
func (su *SubscriptionUpdate) AddDurationInMonths(i int) *SubscriptionUpdate {
	su.mutation.AddDurationInMonths(i)
	return su
}

// SetIsActive sets the "is_active" field.
func (su *SubscriptionUpdate) SetIsActive(b bool) *SubscriptionUpdate {
	su.mutation.SetIsActive(b)
	return su
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableIsActive(b *bool) *SubscriptionUpdate {
	if b != nil {
		su.SetIsActive(*b)
	}
	return su
}

// SetName sets the "name" field.
func (su *SubscriptionUpdate) SetName(s string) *SubscriptionUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableName(s *string) *SubscriptionUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// SetRawSubscriptionData sets the "raw_subscription_data" field.
func (su *SubscriptionUpdate) SetRawSubscriptionData(m map[string]interface{}) *SubscriptionUpdate {
	su.mutation.SetRawSubscriptionData(m)
	return su
}

// ClearRawSubscriptionData clears the value of the "raw_subscription_data" field.
func (su *SubscriptionUpdate) ClearRawSubscriptionData() *SubscriptionUpdate {
	su.mutation.ClearRawSubscriptionData()
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SubscriptionUpdate) SetUpdatedAt(t time.Time) *SubscriptionUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// AddExamIDs adds the "exams" edge to the SubscriptionExam entity by IDs.
func (su *SubscriptionUpdate) AddExamIDs(ids ...int) *SubscriptionUpdate {
	su.mutation.AddExamIDs(ids...)
	return su
}

// AddExams adds the "exams" edges to the SubscriptionExam entity.
func (su *SubscriptionUpdate) AddExams(s ...*SubscriptionExam) *SubscriptionUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddExamIDs(ids...)
}

// AddUserSubscriptionIDs adds the "user_subscriptions" edge to the UserSubscription entity by IDs.
func (su *SubscriptionUpdate) AddUserSubscriptionIDs(ids ...int) *SubscriptionUpdate {
	su.mutation.AddUserSubscriptionIDs(ids...)
	return su
}

// AddUserSubscriptions adds the "user_subscriptions" edges to the UserSubscription entity.
func (su *SubscriptionUpdate) AddUserSubscriptions(u ...*UserSubscription) *SubscriptionUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.AddUserSubscriptionIDs(ids...)
}

// Mutation returns the SubscriptionMutation object of the builder.
func (su *SubscriptionUpdate) Mutation() *SubscriptionMutation {
	return su.mutation
}

// ClearExams clears all "exams" edges to the SubscriptionExam entity.
func (su *SubscriptionUpdate) ClearExams() *SubscriptionUpdate {
	su.mutation.ClearExams()
	return su
}

// RemoveExamIDs removes the "exams" edge to SubscriptionExam entities by IDs.
func (su *SubscriptionUpdate) RemoveExamIDs(ids ...int) *SubscriptionUpdate {
	su.mutation.RemoveExamIDs(ids...)
	return su
}

// RemoveExams removes "exams" edges to SubscriptionExam entities.
func (su *SubscriptionUpdate) RemoveExams(s ...*SubscriptionExam) *SubscriptionUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveExamIDs(ids...)
}

// ClearUserSubscriptions clears all "user_subscriptions" edges to the UserSubscription entity.
func (su *SubscriptionUpdate) ClearUserSubscriptions() *SubscriptionUpdate {
	su.mutation.ClearUserSubscriptions()
	return su
}

// RemoveUserSubscriptionIDs removes the "user_subscriptions" edge to UserSubscription entities by IDs.
func (su *SubscriptionUpdate) RemoveUserSubscriptionIDs(ids ...int) *SubscriptionUpdate {
	su.mutation.RemoveUserSubscriptionIDs(ids...)
	return su
}

// RemoveUserSubscriptions removes "user_subscriptions" edges to UserSubscription entities.
func (su *SubscriptionUpdate) RemoveUserSubscriptions(u ...*UserSubscription) *SubscriptionUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.RemoveUserSubscriptionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SubscriptionUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SubscriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SubscriptionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SubscriptionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SubscriptionUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := subscription.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

func (su *SubscriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(subscription.Table, subscription.Columns, sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.ProviderPlanID(); ok {
		_spec.SetField(subscription.FieldProviderPlanID, field.TypeString, value)
	}
	if value, ok := su.mutation.BasePrice(); ok {
		_spec.SetField(subscription.FieldBasePrice, field.TypeFloat64, value)
	}
	if value, ok := su.mutation.AddedBasePrice(); ok {
		_spec.AddField(subscription.FieldBasePrice, field.TypeFloat64, value)
	}
	if su.mutation.BasePriceCleared() {
		_spec.ClearField(subscription.FieldBasePrice, field.TypeFloat64)
	}
	if value, ok := su.mutation.FinalPrice(); ok {
		_spec.SetField(subscription.FieldFinalPrice, field.TypeFloat64, value)
	}
	if value, ok := su.mutation.AddedFinalPrice(); ok {
		_spec.AddField(subscription.FieldFinalPrice, field.TypeFloat64, value)
	}
	if su.mutation.FinalPriceCleared() {
		_spec.ClearField(subscription.FieldFinalPrice, field.TypeFloat64)
	}
	if value, ok := su.mutation.Price(); ok {
		_spec.SetField(subscription.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := su.mutation.AddedPrice(); ok {
		_spec.AddField(subscription.FieldPrice, field.TypeFloat64, value)
	}
	if su.mutation.PriceCleared() {
		_spec.ClearField(subscription.FieldPrice, field.TypeFloat64)
	}
	if value, ok := su.mutation.DurationInMonths(); ok {
		_spec.SetField(subscription.FieldDurationInMonths, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedDurationInMonths(); ok {
		_spec.AddField(subscription.FieldDurationInMonths, field.TypeInt, value)
	}
	if value, ok := su.mutation.IsActive(); ok {
		_spec.SetField(subscription.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(subscription.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.RawSubscriptionData(); ok {
		_spec.SetField(subscription.FieldRawSubscriptionData, field.TypeJSON, value)
	}
	if su.mutation.RawSubscriptionDataCleared() {
		_spec.ClearField(subscription.FieldRawSubscriptionData, field.TypeJSON)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(subscription.FieldUpdatedAt, field.TypeTime, value)
	}
	if su.mutation.ExamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.ExamsTable,
			Columns: []string{subscription.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subscriptionexam.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedExamsIDs(); len(nodes) > 0 && !su.mutation.ExamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.ExamsTable,
			Columns: []string{subscription.ExamsColumn},
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
	if nodes := su.mutation.ExamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.ExamsTable,
			Columns: []string{subscription.ExamsColumn},
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
	if su.mutation.UserSubscriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.UserSubscriptionsTable,
			Columns: []string{subscription.UserSubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersubscription.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedUserSubscriptionsIDs(); len(nodes) > 0 && !su.mutation.UserSubscriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.UserSubscriptionsTable,
			Columns: []string{subscription.UserSubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersubscription.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserSubscriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.UserSubscriptionsTable,
			Columns: []string{subscription.UserSubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersubscription.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SubscriptionUpdateOne is the builder for updating a single Subscription entity.
type SubscriptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubscriptionMutation
}

// SetProviderPlanID sets the "provider_plan_id" field.
func (suo *SubscriptionUpdateOne) SetProviderPlanID(s string) *SubscriptionUpdateOne {
	suo.mutation.SetProviderPlanID(s)
	return suo
}

// SetNillableProviderPlanID sets the "provider_plan_id" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableProviderPlanID(s *string) *SubscriptionUpdateOne {
	if s != nil {
		suo.SetProviderPlanID(*s)
	}
	return suo
}

// SetBasePrice sets the "base_price" field.
func (suo *SubscriptionUpdateOne) SetBasePrice(f float64) *SubscriptionUpdateOne {
	suo.mutation.ResetBasePrice()
	suo.mutation.SetBasePrice(f)
	return suo
}

// SetNillableBasePrice sets the "base_price" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableBasePrice(f *float64) *SubscriptionUpdateOne {
	if f != nil {
		suo.SetBasePrice(*f)
	}
	return suo
}

// AddBasePrice adds f to the "base_price" field.
func (suo *SubscriptionUpdateOne) AddBasePrice(f float64) *SubscriptionUpdateOne {
	suo.mutation.AddBasePrice(f)
	return suo
}

// ClearBasePrice clears the value of the "base_price" field.
func (suo *SubscriptionUpdateOne) ClearBasePrice() *SubscriptionUpdateOne {
	suo.mutation.ClearBasePrice()
	return suo
}

// SetFinalPrice sets the "final_price" field.
func (suo *SubscriptionUpdateOne) SetFinalPrice(f float64) *SubscriptionUpdateOne {
	suo.mutation.ResetFinalPrice()
	suo.mutation.SetFinalPrice(f)
	return suo
}

// SetNillableFinalPrice sets the "final_price" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableFinalPrice(f *float64) *SubscriptionUpdateOne {
	if f != nil {
		suo.SetFinalPrice(*f)
	}
	return suo
}

// AddFinalPrice adds f to the "final_price" field.
func (suo *SubscriptionUpdateOne) AddFinalPrice(f float64) *SubscriptionUpdateOne {
	suo.mutation.AddFinalPrice(f)
	return suo
}

// ClearFinalPrice clears the value of the "final_price" field.
func (suo *SubscriptionUpdateOne) ClearFinalPrice() *SubscriptionUpdateOne {
	suo.mutation.ClearFinalPrice()
	return suo
}

// SetPrice sets the "price" field.
func (suo *SubscriptionUpdateOne) SetPrice(f float64) *SubscriptionUpdateOne {
	suo.mutation.ResetPrice()
	suo.mutation.SetPrice(f)
	return suo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillablePrice(f *float64) *SubscriptionUpdateOne {
	if f != nil {
		suo.SetPrice(*f)
	}
	return suo
}

// AddPrice adds f to the "price" field.
func (suo *SubscriptionUpdateOne) AddPrice(f float64) *SubscriptionUpdateOne {
	suo.mutation.AddPrice(f)
	return suo
}

// ClearPrice clears the value of the "price" field.
func (suo *SubscriptionUpdateOne) ClearPrice() *SubscriptionUpdateOne {
	suo.mutation.ClearPrice()
	return suo
}

// SetDurationInMonths sets the "duration_in_months" field.
func (suo *SubscriptionUpdateOne) SetDurationInMonths(i int) *SubscriptionUpdateOne {
	suo.mutation.ResetDurationInMonths()
	suo.mutation.SetDurationInMonths(i)
	return suo
}

// SetNillableDurationInMonths sets the "duration_in_months" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableDurationInMonths(i *int) *SubscriptionUpdateOne {
	if i != nil {
		suo.SetDurationInMonths(*i)
	}
	return suo
}

// AddDurationInMonths adds i to the "duration_in_months" field.
func (suo *SubscriptionUpdateOne) AddDurationInMonths(i int) *SubscriptionUpdateOne {
	suo.mutation.AddDurationInMonths(i)
	return suo
}

// SetIsActive sets the "is_active" field.
func (suo *SubscriptionUpdateOne) SetIsActive(b bool) *SubscriptionUpdateOne {
	suo.mutation.SetIsActive(b)
	return suo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableIsActive(b *bool) *SubscriptionUpdateOne {
	if b != nil {
		suo.SetIsActive(*b)
	}
	return suo
}

// SetName sets the "name" field.
func (suo *SubscriptionUpdateOne) SetName(s string) *SubscriptionUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableName(s *string) *SubscriptionUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// SetRawSubscriptionData sets the "raw_subscription_data" field.
func (suo *SubscriptionUpdateOne) SetRawSubscriptionData(m map[string]interface{}) *SubscriptionUpdateOne {
	suo.mutation.SetRawSubscriptionData(m)
	return suo
}

// ClearRawSubscriptionData clears the value of the "raw_subscription_data" field.
func (suo *SubscriptionUpdateOne) ClearRawSubscriptionData() *SubscriptionUpdateOne {
	suo.mutation.ClearRawSubscriptionData()
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SubscriptionUpdateOne) SetUpdatedAt(t time.Time) *SubscriptionUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// AddExamIDs adds the "exams" edge to the SubscriptionExam entity by IDs.
func (suo *SubscriptionUpdateOne) AddExamIDs(ids ...int) *SubscriptionUpdateOne {
	suo.mutation.AddExamIDs(ids...)
	return suo
}

// AddExams adds the "exams" edges to the SubscriptionExam entity.
func (suo *SubscriptionUpdateOne) AddExams(s ...*SubscriptionExam) *SubscriptionUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddExamIDs(ids...)
}

// AddUserSubscriptionIDs adds the "user_subscriptions" edge to the UserSubscription entity by IDs.
func (suo *SubscriptionUpdateOne) AddUserSubscriptionIDs(ids ...int) *SubscriptionUpdateOne {
	suo.mutation.AddUserSubscriptionIDs(ids...)
	return suo
}

// AddUserSubscriptions adds the "user_subscriptions" edges to the UserSubscription entity.
func (suo *SubscriptionUpdateOne) AddUserSubscriptions(u ...*UserSubscription) *SubscriptionUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.AddUserSubscriptionIDs(ids...)
}

// Mutation returns the SubscriptionMutation object of the builder.
func (suo *SubscriptionUpdateOne) Mutation() *SubscriptionMutation {
	return suo.mutation
}

// ClearExams clears all "exams" edges to the SubscriptionExam entity.
func (suo *SubscriptionUpdateOne) ClearExams() *SubscriptionUpdateOne {
	suo.mutation.ClearExams()
	return suo
}

// RemoveExamIDs removes the "exams" edge to SubscriptionExam entities by IDs.
func (suo *SubscriptionUpdateOne) RemoveExamIDs(ids ...int) *SubscriptionUpdateOne {
	suo.mutation.RemoveExamIDs(ids...)
	return suo
}

// RemoveExams removes "exams" edges to SubscriptionExam entities.
func (suo *SubscriptionUpdateOne) RemoveExams(s ...*SubscriptionExam) *SubscriptionUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveExamIDs(ids...)
}

// ClearUserSubscriptions clears all "user_subscriptions" edges to the UserSubscription entity.
func (suo *SubscriptionUpdateOne) ClearUserSubscriptions() *SubscriptionUpdateOne {
	suo.mutation.ClearUserSubscriptions()
	return suo
}

// RemoveUserSubscriptionIDs removes the "user_subscriptions" edge to UserSubscription entities by IDs.
func (suo *SubscriptionUpdateOne) RemoveUserSubscriptionIDs(ids ...int) *SubscriptionUpdateOne {
	suo.mutation.RemoveUserSubscriptionIDs(ids...)
	return suo
}

// RemoveUserSubscriptions removes "user_subscriptions" edges to UserSubscription entities.
func (suo *SubscriptionUpdateOne) RemoveUserSubscriptions(u ...*UserSubscription) *SubscriptionUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.RemoveUserSubscriptionIDs(ids...)
}

// Where appends a list predicates to the SubscriptionUpdate builder.
func (suo *SubscriptionUpdateOne) Where(ps ...predicate.Subscription) *SubscriptionUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SubscriptionUpdateOne) Select(field string, fields ...string) *SubscriptionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Subscription entity.
func (suo *SubscriptionUpdateOne) Save(ctx context.Context) (*Subscription, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SubscriptionUpdateOne) SaveX(ctx context.Context) *Subscription {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SubscriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubscriptionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SubscriptionUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := subscription.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

func (suo *SubscriptionUpdateOne) sqlSave(ctx context.Context) (_node *Subscription, err error) {
	_spec := sqlgraph.NewUpdateSpec(subscription.Table, subscription.Columns, sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Subscription.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subscription.FieldID)
		for _, f := range fields {
			if !subscription.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subscription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.ProviderPlanID(); ok {
		_spec.SetField(subscription.FieldProviderPlanID, field.TypeString, value)
	}
	if value, ok := suo.mutation.BasePrice(); ok {
		_spec.SetField(subscription.FieldBasePrice, field.TypeFloat64, value)
	}
	if value, ok := suo.mutation.AddedBasePrice(); ok {
		_spec.AddField(subscription.FieldBasePrice, field.TypeFloat64, value)
	}
	if suo.mutation.BasePriceCleared() {
		_spec.ClearField(subscription.FieldBasePrice, field.TypeFloat64)
	}
	if value, ok := suo.mutation.FinalPrice(); ok {
		_spec.SetField(subscription.FieldFinalPrice, field.TypeFloat64, value)
	}
	if value, ok := suo.mutation.AddedFinalPrice(); ok {
		_spec.AddField(subscription.FieldFinalPrice, field.TypeFloat64, value)
	}
	if suo.mutation.FinalPriceCleared() {
		_spec.ClearField(subscription.FieldFinalPrice, field.TypeFloat64)
	}
	if value, ok := suo.mutation.Price(); ok {
		_spec.SetField(subscription.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := suo.mutation.AddedPrice(); ok {
		_spec.AddField(subscription.FieldPrice, field.TypeFloat64, value)
	}
	if suo.mutation.PriceCleared() {
		_spec.ClearField(subscription.FieldPrice, field.TypeFloat64)
	}
	if value, ok := suo.mutation.DurationInMonths(); ok {
		_spec.SetField(subscription.FieldDurationInMonths, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedDurationInMonths(); ok {
		_spec.AddField(subscription.FieldDurationInMonths, field.TypeInt, value)
	}
	if value, ok := suo.mutation.IsActive(); ok {
		_spec.SetField(subscription.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(subscription.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.RawSubscriptionData(); ok {
		_spec.SetField(subscription.FieldRawSubscriptionData, field.TypeJSON, value)
	}
	if suo.mutation.RawSubscriptionDataCleared() {
		_spec.ClearField(subscription.FieldRawSubscriptionData, field.TypeJSON)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(subscription.FieldUpdatedAt, field.TypeTime, value)
	}
	if suo.mutation.ExamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.ExamsTable,
			Columns: []string{subscription.ExamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subscriptionexam.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedExamsIDs(); len(nodes) > 0 && !suo.mutation.ExamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.ExamsTable,
			Columns: []string{subscription.ExamsColumn},
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
	if nodes := suo.mutation.ExamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.ExamsTable,
			Columns: []string{subscription.ExamsColumn},
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
	if suo.mutation.UserSubscriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.UserSubscriptionsTable,
			Columns: []string{subscription.UserSubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersubscription.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedUserSubscriptionsIDs(); len(nodes) > 0 && !suo.mutation.UserSubscriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.UserSubscriptionsTable,
			Columns: []string{subscription.UserSubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersubscription.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserSubscriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.UserSubscriptionsTable,
			Columns: []string{subscription.UserSubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersubscription.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Subscription{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
