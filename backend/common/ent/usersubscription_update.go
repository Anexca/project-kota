// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/payment"
	"common/ent/predicate"
	"common/ent/subscription"
	"common/ent/user"
	"common/ent/usersubscription"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserSubscriptionUpdate is the builder for updating UserSubscription entities.
type UserSubscriptionUpdate struct {
	config
	hooks    []Hook
	mutation *UserSubscriptionMutation
}

// Where appends a list predicates to the UserSubscriptionUpdate builder.
func (usu *UserSubscriptionUpdate) Where(ps ...predicate.UserSubscription) *UserSubscriptionUpdate {
	usu.mutation.Where(ps...)
	return usu
}

// SetIsActive sets the "is_active" field.
func (usu *UserSubscriptionUpdate) SetIsActive(b bool) *UserSubscriptionUpdate {
	usu.mutation.SetIsActive(b)
	return usu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (usu *UserSubscriptionUpdate) SetNillableIsActive(b *bool) *UserSubscriptionUpdate {
	if b != nil {
		usu.SetIsActive(*b)
	}
	return usu
}

// SetStartDate sets the "start_date" field.
func (usu *UserSubscriptionUpdate) SetStartDate(t time.Time) *UserSubscriptionUpdate {
	usu.mutation.SetStartDate(t)
	return usu
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (usu *UserSubscriptionUpdate) SetNillableStartDate(t *time.Time) *UserSubscriptionUpdate {
	if t != nil {
		usu.SetStartDate(*t)
	}
	return usu
}

// ClearStartDate clears the value of the "start_date" field.
func (usu *UserSubscriptionUpdate) ClearStartDate() *UserSubscriptionUpdate {
	usu.mutation.ClearStartDate()
	return usu
}

// SetEndDate sets the "end_date" field.
func (usu *UserSubscriptionUpdate) SetEndDate(t time.Time) *UserSubscriptionUpdate {
	usu.mutation.SetEndDate(t)
	return usu
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (usu *UserSubscriptionUpdate) SetNillableEndDate(t *time.Time) *UserSubscriptionUpdate {
	if t != nil {
		usu.SetEndDate(*t)
	}
	return usu
}

// ClearEndDate clears the value of the "end_date" field.
func (usu *UserSubscriptionUpdate) ClearEndDate() *UserSubscriptionUpdate {
	usu.mutation.ClearEndDate()
	return usu
}

// SetProviderSubscriptionID sets the "provider_subscription_id" field.
func (usu *UserSubscriptionUpdate) SetProviderSubscriptionID(s string) *UserSubscriptionUpdate {
	usu.mutation.SetProviderSubscriptionID(s)
	return usu
}

// SetNillableProviderSubscriptionID sets the "provider_subscription_id" field if the given value is not nil.
func (usu *UserSubscriptionUpdate) SetNillableProviderSubscriptionID(s *string) *UserSubscriptionUpdate {
	if s != nil {
		usu.SetProviderSubscriptionID(*s)
	}
	return usu
}

// SetUpdatedAt sets the "updated_at" field.
func (usu *UserSubscriptionUpdate) SetUpdatedAt(t time.Time) *UserSubscriptionUpdate {
	usu.mutation.SetUpdatedAt(t)
	return usu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (usu *UserSubscriptionUpdate) SetUserID(id uuid.UUID) *UserSubscriptionUpdate {
	usu.mutation.SetUserID(id)
	return usu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (usu *UserSubscriptionUpdate) SetNillableUserID(id *uuid.UUID) *UserSubscriptionUpdate {
	if id != nil {
		usu = usu.SetUserID(*id)
	}
	return usu
}

// SetUser sets the "user" edge to the User entity.
func (usu *UserSubscriptionUpdate) SetUser(u *User) *UserSubscriptionUpdate {
	return usu.SetUserID(u.ID)
}

// SetSubscriptionID sets the "subscription" edge to the Subscription entity by ID.
func (usu *UserSubscriptionUpdate) SetSubscriptionID(id int) *UserSubscriptionUpdate {
	usu.mutation.SetSubscriptionID(id)
	return usu
}

// SetNillableSubscriptionID sets the "subscription" edge to the Subscription entity by ID if the given value is not nil.
func (usu *UserSubscriptionUpdate) SetNillableSubscriptionID(id *int) *UserSubscriptionUpdate {
	if id != nil {
		usu = usu.SetSubscriptionID(*id)
	}
	return usu
}

// SetSubscription sets the "subscription" edge to the Subscription entity.
func (usu *UserSubscriptionUpdate) SetSubscription(s *Subscription) *UserSubscriptionUpdate {
	return usu.SetSubscriptionID(s.ID)
}

// AddPaymentIDs adds the "payments" edge to the Payment entity by IDs.
func (usu *UserSubscriptionUpdate) AddPaymentIDs(ids ...int) *UserSubscriptionUpdate {
	usu.mutation.AddPaymentIDs(ids...)
	return usu
}

// AddPayments adds the "payments" edges to the Payment entity.
func (usu *UserSubscriptionUpdate) AddPayments(p ...*Payment) *UserSubscriptionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return usu.AddPaymentIDs(ids...)
}

// Mutation returns the UserSubscriptionMutation object of the builder.
func (usu *UserSubscriptionUpdate) Mutation() *UserSubscriptionMutation {
	return usu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (usu *UserSubscriptionUpdate) ClearUser() *UserSubscriptionUpdate {
	usu.mutation.ClearUser()
	return usu
}

// ClearSubscription clears the "subscription" edge to the Subscription entity.
func (usu *UserSubscriptionUpdate) ClearSubscription() *UserSubscriptionUpdate {
	usu.mutation.ClearSubscription()
	return usu
}

// ClearPayments clears all "payments" edges to the Payment entity.
func (usu *UserSubscriptionUpdate) ClearPayments() *UserSubscriptionUpdate {
	usu.mutation.ClearPayments()
	return usu
}

// RemovePaymentIDs removes the "payments" edge to Payment entities by IDs.
func (usu *UserSubscriptionUpdate) RemovePaymentIDs(ids ...int) *UserSubscriptionUpdate {
	usu.mutation.RemovePaymentIDs(ids...)
	return usu
}

// RemovePayments removes "payments" edges to Payment entities.
func (usu *UserSubscriptionUpdate) RemovePayments(p ...*Payment) *UserSubscriptionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return usu.RemovePaymentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (usu *UserSubscriptionUpdate) Save(ctx context.Context) (int, error) {
	usu.defaults()
	return withHooks(ctx, usu.sqlSave, usu.mutation, usu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (usu *UserSubscriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := usu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (usu *UserSubscriptionUpdate) Exec(ctx context.Context) error {
	_, err := usu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usu *UserSubscriptionUpdate) ExecX(ctx context.Context) {
	if err := usu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usu *UserSubscriptionUpdate) defaults() {
	if _, ok := usu.mutation.UpdatedAt(); !ok {
		v := usersubscription.UpdateDefaultUpdatedAt()
		usu.mutation.SetUpdatedAt(v)
	}
}

func (usu *UserSubscriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(usersubscription.Table, usersubscription.Columns, sqlgraph.NewFieldSpec(usersubscription.FieldID, field.TypeInt))
	if ps := usu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usu.mutation.IsActive(); ok {
		_spec.SetField(usersubscription.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := usu.mutation.StartDate(); ok {
		_spec.SetField(usersubscription.FieldStartDate, field.TypeTime, value)
	}
	if usu.mutation.StartDateCleared() {
		_spec.ClearField(usersubscription.FieldStartDate, field.TypeTime)
	}
	if value, ok := usu.mutation.EndDate(); ok {
		_spec.SetField(usersubscription.FieldEndDate, field.TypeTime, value)
	}
	if usu.mutation.EndDateCleared() {
		_spec.ClearField(usersubscription.FieldEndDate, field.TypeTime)
	}
	if value, ok := usu.mutation.ProviderSubscriptionID(); ok {
		_spec.SetField(usersubscription.FieldProviderSubscriptionID, field.TypeString, value)
	}
	if value, ok := usu.mutation.UpdatedAt(); ok {
		_spec.SetField(usersubscription.FieldUpdatedAt, field.TypeTime, value)
	}
	if usu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usersubscription.UserTable,
			Columns: []string{usersubscription.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usersubscription.UserTable,
			Columns: []string{usersubscription.UserColumn},
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
	if usu.mutation.SubscriptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usersubscription.SubscriptionTable,
			Columns: []string{usersubscription.SubscriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usu.mutation.SubscriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usersubscription.SubscriptionTable,
			Columns: []string{usersubscription.SubscriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if usu.mutation.PaymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usersubscription.PaymentsTable,
			Columns: []string{usersubscription.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usu.mutation.RemovedPaymentsIDs(); len(nodes) > 0 && !usu.mutation.PaymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usersubscription.PaymentsTable,
			Columns: []string{usersubscription.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usu.mutation.PaymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usersubscription.PaymentsTable,
			Columns: []string{usersubscription.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, usu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersubscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	usu.mutation.done = true
	return n, nil
}

// UserSubscriptionUpdateOne is the builder for updating a single UserSubscription entity.
type UserSubscriptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserSubscriptionMutation
}

// SetIsActive sets the "is_active" field.
func (usuo *UserSubscriptionUpdateOne) SetIsActive(b bool) *UserSubscriptionUpdateOne {
	usuo.mutation.SetIsActive(b)
	return usuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (usuo *UserSubscriptionUpdateOne) SetNillableIsActive(b *bool) *UserSubscriptionUpdateOne {
	if b != nil {
		usuo.SetIsActive(*b)
	}
	return usuo
}

// SetStartDate sets the "start_date" field.
func (usuo *UserSubscriptionUpdateOne) SetStartDate(t time.Time) *UserSubscriptionUpdateOne {
	usuo.mutation.SetStartDate(t)
	return usuo
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (usuo *UserSubscriptionUpdateOne) SetNillableStartDate(t *time.Time) *UserSubscriptionUpdateOne {
	if t != nil {
		usuo.SetStartDate(*t)
	}
	return usuo
}

// ClearStartDate clears the value of the "start_date" field.
func (usuo *UserSubscriptionUpdateOne) ClearStartDate() *UserSubscriptionUpdateOne {
	usuo.mutation.ClearStartDate()
	return usuo
}

// SetEndDate sets the "end_date" field.
func (usuo *UserSubscriptionUpdateOne) SetEndDate(t time.Time) *UserSubscriptionUpdateOne {
	usuo.mutation.SetEndDate(t)
	return usuo
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (usuo *UserSubscriptionUpdateOne) SetNillableEndDate(t *time.Time) *UserSubscriptionUpdateOne {
	if t != nil {
		usuo.SetEndDate(*t)
	}
	return usuo
}

// ClearEndDate clears the value of the "end_date" field.
func (usuo *UserSubscriptionUpdateOne) ClearEndDate() *UserSubscriptionUpdateOne {
	usuo.mutation.ClearEndDate()
	return usuo
}

// SetProviderSubscriptionID sets the "provider_subscription_id" field.
func (usuo *UserSubscriptionUpdateOne) SetProviderSubscriptionID(s string) *UserSubscriptionUpdateOne {
	usuo.mutation.SetProviderSubscriptionID(s)
	return usuo
}

// SetNillableProviderSubscriptionID sets the "provider_subscription_id" field if the given value is not nil.
func (usuo *UserSubscriptionUpdateOne) SetNillableProviderSubscriptionID(s *string) *UserSubscriptionUpdateOne {
	if s != nil {
		usuo.SetProviderSubscriptionID(*s)
	}
	return usuo
}

// SetUpdatedAt sets the "updated_at" field.
func (usuo *UserSubscriptionUpdateOne) SetUpdatedAt(t time.Time) *UserSubscriptionUpdateOne {
	usuo.mutation.SetUpdatedAt(t)
	return usuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (usuo *UserSubscriptionUpdateOne) SetUserID(id uuid.UUID) *UserSubscriptionUpdateOne {
	usuo.mutation.SetUserID(id)
	return usuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (usuo *UserSubscriptionUpdateOne) SetNillableUserID(id *uuid.UUID) *UserSubscriptionUpdateOne {
	if id != nil {
		usuo = usuo.SetUserID(*id)
	}
	return usuo
}

// SetUser sets the "user" edge to the User entity.
func (usuo *UserSubscriptionUpdateOne) SetUser(u *User) *UserSubscriptionUpdateOne {
	return usuo.SetUserID(u.ID)
}

// SetSubscriptionID sets the "subscription" edge to the Subscription entity by ID.
func (usuo *UserSubscriptionUpdateOne) SetSubscriptionID(id int) *UserSubscriptionUpdateOne {
	usuo.mutation.SetSubscriptionID(id)
	return usuo
}

// SetNillableSubscriptionID sets the "subscription" edge to the Subscription entity by ID if the given value is not nil.
func (usuo *UserSubscriptionUpdateOne) SetNillableSubscriptionID(id *int) *UserSubscriptionUpdateOne {
	if id != nil {
		usuo = usuo.SetSubscriptionID(*id)
	}
	return usuo
}

// SetSubscription sets the "subscription" edge to the Subscription entity.
func (usuo *UserSubscriptionUpdateOne) SetSubscription(s *Subscription) *UserSubscriptionUpdateOne {
	return usuo.SetSubscriptionID(s.ID)
}

// AddPaymentIDs adds the "payments" edge to the Payment entity by IDs.
func (usuo *UserSubscriptionUpdateOne) AddPaymentIDs(ids ...int) *UserSubscriptionUpdateOne {
	usuo.mutation.AddPaymentIDs(ids...)
	return usuo
}

// AddPayments adds the "payments" edges to the Payment entity.
func (usuo *UserSubscriptionUpdateOne) AddPayments(p ...*Payment) *UserSubscriptionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return usuo.AddPaymentIDs(ids...)
}

// Mutation returns the UserSubscriptionMutation object of the builder.
func (usuo *UserSubscriptionUpdateOne) Mutation() *UserSubscriptionMutation {
	return usuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (usuo *UserSubscriptionUpdateOne) ClearUser() *UserSubscriptionUpdateOne {
	usuo.mutation.ClearUser()
	return usuo
}

// ClearSubscription clears the "subscription" edge to the Subscription entity.
func (usuo *UserSubscriptionUpdateOne) ClearSubscription() *UserSubscriptionUpdateOne {
	usuo.mutation.ClearSubscription()
	return usuo
}

// ClearPayments clears all "payments" edges to the Payment entity.
func (usuo *UserSubscriptionUpdateOne) ClearPayments() *UserSubscriptionUpdateOne {
	usuo.mutation.ClearPayments()
	return usuo
}

// RemovePaymentIDs removes the "payments" edge to Payment entities by IDs.
func (usuo *UserSubscriptionUpdateOne) RemovePaymentIDs(ids ...int) *UserSubscriptionUpdateOne {
	usuo.mutation.RemovePaymentIDs(ids...)
	return usuo
}

// RemovePayments removes "payments" edges to Payment entities.
func (usuo *UserSubscriptionUpdateOne) RemovePayments(p ...*Payment) *UserSubscriptionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return usuo.RemovePaymentIDs(ids...)
}

// Where appends a list predicates to the UserSubscriptionUpdate builder.
func (usuo *UserSubscriptionUpdateOne) Where(ps ...predicate.UserSubscription) *UserSubscriptionUpdateOne {
	usuo.mutation.Where(ps...)
	return usuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (usuo *UserSubscriptionUpdateOne) Select(field string, fields ...string) *UserSubscriptionUpdateOne {
	usuo.fields = append([]string{field}, fields...)
	return usuo
}

// Save executes the query and returns the updated UserSubscription entity.
func (usuo *UserSubscriptionUpdateOne) Save(ctx context.Context) (*UserSubscription, error) {
	usuo.defaults()
	return withHooks(ctx, usuo.sqlSave, usuo.mutation, usuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (usuo *UserSubscriptionUpdateOne) SaveX(ctx context.Context) *UserSubscription {
	node, err := usuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (usuo *UserSubscriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := usuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usuo *UserSubscriptionUpdateOne) ExecX(ctx context.Context) {
	if err := usuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usuo *UserSubscriptionUpdateOne) defaults() {
	if _, ok := usuo.mutation.UpdatedAt(); !ok {
		v := usersubscription.UpdateDefaultUpdatedAt()
		usuo.mutation.SetUpdatedAt(v)
	}
}

func (usuo *UserSubscriptionUpdateOne) sqlSave(ctx context.Context) (_node *UserSubscription, err error) {
	_spec := sqlgraph.NewUpdateSpec(usersubscription.Table, usersubscription.Columns, sqlgraph.NewFieldSpec(usersubscription.FieldID, field.TypeInt))
	id, ok := usuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserSubscription.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := usuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usersubscription.FieldID)
		for _, f := range fields {
			if !usersubscription.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usersubscription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := usuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usuo.mutation.IsActive(); ok {
		_spec.SetField(usersubscription.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := usuo.mutation.StartDate(); ok {
		_spec.SetField(usersubscription.FieldStartDate, field.TypeTime, value)
	}
	if usuo.mutation.StartDateCleared() {
		_spec.ClearField(usersubscription.FieldStartDate, field.TypeTime)
	}
	if value, ok := usuo.mutation.EndDate(); ok {
		_spec.SetField(usersubscription.FieldEndDate, field.TypeTime, value)
	}
	if usuo.mutation.EndDateCleared() {
		_spec.ClearField(usersubscription.FieldEndDate, field.TypeTime)
	}
	if value, ok := usuo.mutation.ProviderSubscriptionID(); ok {
		_spec.SetField(usersubscription.FieldProviderSubscriptionID, field.TypeString, value)
	}
	if value, ok := usuo.mutation.UpdatedAt(); ok {
		_spec.SetField(usersubscription.FieldUpdatedAt, field.TypeTime, value)
	}
	if usuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usersubscription.UserTable,
			Columns: []string{usersubscription.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usersubscription.UserTable,
			Columns: []string{usersubscription.UserColumn},
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
	if usuo.mutation.SubscriptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usersubscription.SubscriptionTable,
			Columns: []string{usersubscription.SubscriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usuo.mutation.SubscriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usersubscription.SubscriptionTable,
			Columns: []string{usersubscription.SubscriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if usuo.mutation.PaymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usersubscription.PaymentsTable,
			Columns: []string{usersubscription.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usuo.mutation.RemovedPaymentsIDs(); len(nodes) > 0 && !usuo.mutation.PaymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usersubscription.PaymentsTable,
			Columns: []string{usersubscription.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usuo.mutation.PaymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usersubscription.PaymentsTable,
			Columns: []string{usersubscription.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserSubscription{config: usuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, usuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersubscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	usuo.mutation.done = true
	return _node, nil
}
