// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/payment"
	"common/ent/predicate"
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

// PaymentUpdate is the builder for updating Payment entities.
type PaymentUpdate struct {
	config
	hooks    []Hook
	mutation *PaymentMutation
}

// Where appends a list predicates to the PaymentUpdate builder.
func (pu *PaymentUpdate) Where(ps ...predicate.Payment) *PaymentUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetAmount sets the "amount" field.
func (pu *PaymentUpdate) SetAmount(i int) *PaymentUpdate {
	pu.mutation.ResetAmount()
	pu.mutation.SetAmount(i)
	return pu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableAmount(i *int) *PaymentUpdate {
	if i != nil {
		pu.SetAmount(*i)
	}
	return pu
}

// AddAmount adds i to the "amount" field.
func (pu *PaymentUpdate) AddAmount(i int) *PaymentUpdate {
	pu.mutation.AddAmount(i)
	return pu
}

// SetPaymentDate sets the "payment_date" field.
func (pu *PaymentUpdate) SetPaymentDate(t time.Time) *PaymentUpdate {
	pu.mutation.SetPaymentDate(t)
	return pu
}

// SetNillablePaymentDate sets the "payment_date" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillablePaymentDate(t *time.Time) *PaymentUpdate {
	if t != nil {
		pu.SetPaymentDate(*t)
	}
	return pu
}

// SetStatus sets the "status" field.
func (pu *PaymentUpdate) SetStatus(pa payment.Status) *PaymentUpdate {
	pu.mutation.SetStatus(pa)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableStatus(pa *payment.Status) *PaymentUpdate {
	if pa != nil {
		pu.SetStatus(*pa)
	}
	return pu
}

// SetPaymentMethod sets the "payment_method" field.
func (pu *PaymentUpdate) SetPaymentMethod(s string) *PaymentUpdate {
	pu.mutation.SetPaymentMethod(s)
	return pu
}

// SetNillablePaymentMethod sets the "payment_method" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillablePaymentMethod(s *string) *PaymentUpdate {
	if s != nil {
		pu.SetPaymentMethod(*s)
	}
	return pu
}

// SetProviderPaymentID sets the "provider_payment_id" field.
func (pu *PaymentUpdate) SetProviderPaymentID(s string) *PaymentUpdate {
	pu.mutation.SetProviderPaymentID(s)
	return pu
}

// SetNillableProviderPaymentID sets the "provider_payment_id" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableProviderPaymentID(s *string) *PaymentUpdate {
	if s != nil {
		pu.SetProviderPaymentID(*s)
	}
	return pu
}

// SetProviderInvoiceID sets the "provider_invoice_id" field.
func (pu *PaymentUpdate) SetProviderInvoiceID(s string) *PaymentUpdate {
	pu.mutation.SetProviderInvoiceID(s)
	return pu
}

// SetNillableProviderInvoiceID sets the "provider_invoice_id" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableProviderInvoiceID(s *string) *PaymentUpdate {
	if s != nil {
		pu.SetProviderInvoiceID(*s)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PaymentUpdate) SetUpdatedAt(t time.Time) *PaymentUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pu *PaymentUpdate) SetUserID(id uuid.UUID) *PaymentUpdate {
	pu.mutation.SetUserID(id)
	return pu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pu *PaymentUpdate) SetNillableUserID(id *uuid.UUID) *PaymentUpdate {
	if id != nil {
		pu = pu.SetUserID(*id)
	}
	return pu
}

// SetUser sets the "user" edge to the User entity.
func (pu *PaymentUpdate) SetUser(u *User) *PaymentUpdate {
	return pu.SetUserID(u.ID)
}

// SetSubscriptionID sets the "subscription" edge to the UserSubscription entity by ID.
func (pu *PaymentUpdate) SetSubscriptionID(id int) *PaymentUpdate {
	pu.mutation.SetSubscriptionID(id)
	return pu
}

// SetNillableSubscriptionID sets the "subscription" edge to the UserSubscription entity by ID if the given value is not nil.
func (pu *PaymentUpdate) SetNillableSubscriptionID(id *int) *PaymentUpdate {
	if id != nil {
		pu = pu.SetSubscriptionID(*id)
	}
	return pu
}

// SetSubscription sets the "subscription" edge to the UserSubscription entity.
func (pu *PaymentUpdate) SetSubscription(u *UserSubscription) *PaymentUpdate {
	return pu.SetSubscriptionID(u.ID)
}

// Mutation returns the PaymentMutation object of the builder.
func (pu *PaymentUpdate) Mutation() *PaymentMutation {
	return pu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pu *PaymentUpdate) ClearUser() *PaymentUpdate {
	pu.mutation.ClearUser()
	return pu
}

// ClearSubscription clears the "subscription" edge to the UserSubscription entity.
func (pu *PaymentUpdate) ClearSubscription() *PaymentUpdate {
	pu.mutation.ClearSubscription()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PaymentUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PaymentUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PaymentUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PaymentUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PaymentUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := payment.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PaymentUpdate) check() error {
	if v, ok := pu.mutation.Status(); ok {
		if err := payment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Payment.status": %w`, err)}
		}
	}
	return nil
}

func (pu *PaymentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedAmount(); ok {
		_spec.AddField(payment.FieldAmount, field.TypeInt, value)
	}
	if value, ok := pu.mutation.PaymentDate(); ok {
		_spec.SetField(payment.FieldPaymentDate, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(payment.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.PaymentMethod(); ok {
		_spec.SetField(payment.FieldPaymentMethod, field.TypeString, value)
	}
	if value, ok := pu.mutation.ProviderPaymentID(); ok {
		_spec.SetField(payment.FieldProviderPaymentID, field.TypeString, value)
	}
	if value, ok := pu.mutation.ProviderInvoiceID(); ok {
		_spec.SetField(payment.FieldProviderInvoiceID, field.TypeString, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(payment.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.UserTable,
			Columns: []string{payment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.UserTable,
			Columns: []string{payment.UserColumn},
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
	if pu.mutation.SubscriptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.SubscriptionTable,
			Columns: []string{payment.SubscriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersubscription.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SubscriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.SubscriptionTable,
			Columns: []string{payment.SubscriptionColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PaymentUpdateOne is the builder for updating a single Payment entity.
type PaymentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PaymentMutation
}

// SetAmount sets the "amount" field.
func (puo *PaymentUpdateOne) SetAmount(i int) *PaymentUpdateOne {
	puo.mutation.ResetAmount()
	puo.mutation.SetAmount(i)
	return puo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableAmount(i *int) *PaymentUpdateOne {
	if i != nil {
		puo.SetAmount(*i)
	}
	return puo
}

// AddAmount adds i to the "amount" field.
func (puo *PaymentUpdateOne) AddAmount(i int) *PaymentUpdateOne {
	puo.mutation.AddAmount(i)
	return puo
}

// SetPaymentDate sets the "payment_date" field.
func (puo *PaymentUpdateOne) SetPaymentDate(t time.Time) *PaymentUpdateOne {
	puo.mutation.SetPaymentDate(t)
	return puo
}

// SetNillablePaymentDate sets the "payment_date" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillablePaymentDate(t *time.Time) *PaymentUpdateOne {
	if t != nil {
		puo.SetPaymentDate(*t)
	}
	return puo
}

// SetStatus sets the "status" field.
func (puo *PaymentUpdateOne) SetStatus(pa payment.Status) *PaymentUpdateOne {
	puo.mutation.SetStatus(pa)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableStatus(pa *payment.Status) *PaymentUpdateOne {
	if pa != nil {
		puo.SetStatus(*pa)
	}
	return puo
}

// SetPaymentMethod sets the "payment_method" field.
func (puo *PaymentUpdateOne) SetPaymentMethod(s string) *PaymentUpdateOne {
	puo.mutation.SetPaymentMethod(s)
	return puo
}

// SetNillablePaymentMethod sets the "payment_method" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillablePaymentMethod(s *string) *PaymentUpdateOne {
	if s != nil {
		puo.SetPaymentMethod(*s)
	}
	return puo
}

// SetProviderPaymentID sets the "provider_payment_id" field.
func (puo *PaymentUpdateOne) SetProviderPaymentID(s string) *PaymentUpdateOne {
	puo.mutation.SetProviderPaymentID(s)
	return puo
}

// SetNillableProviderPaymentID sets the "provider_payment_id" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableProviderPaymentID(s *string) *PaymentUpdateOne {
	if s != nil {
		puo.SetProviderPaymentID(*s)
	}
	return puo
}

// SetProviderInvoiceID sets the "provider_invoice_id" field.
func (puo *PaymentUpdateOne) SetProviderInvoiceID(s string) *PaymentUpdateOne {
	puo.mutation.SetProviderInvoiceID(s)
	return puo
}

// SetNillableProviderInvoiceID sets the "provider_invoice_id" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableProviderInvoiceID(s *string) *PaymentUpdateOne {
	if s != nil {
		puo.SetProviderInvoiceID(*s)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PaymentUpdateOne) SetUpdatedAt(t time.Time) *PaymentUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (puo *PaymentUpdateOne) SetUserID(id uuid.UUID) *PaymentUpdateOne {
	puo.mutation.SetUserID(id)
	return puo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableUserID(id *uuid.UUID) *PaymentUpdateOne {
	if id != nil {
		puo = puo.SetUserID(*id)
	}
	return puo
}

// SetUser sets the "user" edge to the User entity.
func (puo *PaymentUpdateOne) SetUser(u *User) *PaymentUpdateOne {
	return puo.SetUserID(u.ID)
}

// SetSubscriptionID sets the "subscription" edge to the UserSubscription entity by ID.
func (puo *PaymentUpdateOne) SetSubscriptionID(id int) *PaymentUpdateOne {
	puo.mutation.SetSubscriptionID(id)
	return puo
}

// SetNillableSubscriptionID sets the "subscription" edge to the UserSubscription entity by ID if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableSubscriptionID(id *int) *PaymentUpdateOne {
	if id != nil {
		puo = puo.SetSubscriptionID(*id)
	}
	return puo
}

// SetSubscription sets the "subscription" edge to the UserSubscription entity.
func (puo *PaymentUpdateOne) SetSubscription(u *UserSubscription) *PaymentUpdateOne {
	return puo.SetSubscriptionID(u.ID)
}

// Mutation returns the PaymentMutation object of the builder.
func (puo *PaymentUpdateOne) Mutation() *PaymentMutation {
	return puo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (puo *PaymentUpdateOne) ClearUser() *PaymentUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// ClearSubscription clears the "subscription" edge to the UserSubscription entity.
func (puo *PaymentUpdateOne) ClearSubscription() *PaymentUpdateOne {
	puo.mutation.ClearSubscription()
	return puo
}

// Where appends a list predicates to the PaymentUpdate builder.
func (puo *PaymentUpdateOne) Where(ps ...predicate.Payment) *PaymentUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PaymentUpdateOne) Select(field string, fields ...string) *PaymentUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Payment entity.
func (puo *PaymentUpdateOne) Save(ctx context.Context) (*Payment, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PaymentUpdateOne) SaveX(ctx context.Context) *Payment {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PaymentUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PaymentUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PaymentUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := payment.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PaymentUpdateOne) check() error {
	if v, ok := puo.mutation.Status(); ok {
		if err := payment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Payment.status": %w`, err)}
		}
	}
	return nil
}

func (puo *PaymentUpdateOne) sqlSave(ctx context.Context) (_node *Payment, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Payment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payment.FieldID)
		for _, f := range fields {
			if !payment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != payment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedAmount(); ok {
		_spec.AddField(payment.FieldAmount, field.TypeInt, value)
	}
	if value, ok := puo.mutation.PaymentDate(); ok {
		_spec.SetField(payment.FieldPaymentDate, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(payment.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.PaymentMethod(); ok {
		_spec.SetField(payment.FieldPaymentMethod, field.TypeString, value)
	}
	if value, ok := puo.mutation.ProviderPaymentID(); ok {
		_spec.SetField(payment.FieldProviderPaymentID, field.TypeString, value)
	}
	if value, ok := puo.mutation.ProviderInvoiceID(); ok {
		_spec.SetField(payment.FieldProviderInvoiceID, field.TypeString, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(payment.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.UserTable,
			Columns: []string{payment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.UserTable,
			Columns: []string{payment.UserColumn},
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
	if puo.mutation.SubscriptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.SubscriptionTable,
			Columns: []string{payment.SubscriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersubscription.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SubscriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.SubscriptionTable,
			Columns: []string{payment.SubscriptionColumn},
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
	_node = &Payment{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
