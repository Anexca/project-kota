// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/payment"
	"common/ent/user"
	"common/ent/usersubscription"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PaymentCreate is the builder for creating a Payment entity.
type PaymentCreate struct {
	config
	mutation *PaymentMutation
	hooks    []Hook
}

// SetAmount sets the "amount" field.
func (pc *PaymentCreate) SetAmount(f float64) *PaymentCreate {
	pc.mutation.SetAmount(f)
	return pc
}

// SetPaymentDate sets the "payment_date" field.
func (pc *PaymentCreate) SetPaymentDate(t time.Time) *PaymentCreate {
	pc.mutation.SetPaymentDate(t)
	return pc
}

// SetStatus sets the "status" field.
func (pc *PaymentCreate) SetStatus(pa payment.Status) *PaymentCreate {
	pc.mutation.SetStatus(pa)
	return pc
}

// SetPaymentMethod sets the "payment_method" field.
func (pc *PaymentCreate) SetPaymentMethod(s string) *PaymentCreate {
	pc.mutation.SetPaymentMethod(s)
	return pc
}

// SetProviderPaymentID sets the "provider_payment_id" field.
func (pc *PaymentCreate) SetProviderPaymentID(s string) *PaymentCreate {
	pc.mutation.SetProviderPaymentID(s)
	return pc
}

// SetProviderInvoiceID sets the "provider_invoice_id" field.
func (pc *PaymentCreate) SetProviderInvoiceID(s string) *PaymentCreate {
	pc.mutation.SetProviderInvoiceID(s)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PaymentCreate) SetCreatedAt(t time.Time) *PaymentCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PaymentCreate) SetNillableCreatedAt(t *time.Time) *PaymentCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PaymentCreate) SetUpdatedAt(t time.Time) *PaymentCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PaymentCreate) SetNillableUpdatedAt(t *time.Time) *PaymentCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pc *PaymentCreate) SetUserID(id uuid.UUID) *PaymentCreate {
	pc.mutation.SetUserID(id)
	return pc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pc *PaymentCreate) SetNillableUserID(id *uuid.UUID) *PaymentCreate {
	if id != nil {
		pc = pc.SetUserID(*id)
	}
	return pc
}

// SetUser sets the "user" edge to the User entity.
func (pc *PaymentCreate) SetUser(u *User) *PaymentCreate {
	return pc.SetUserID(u.ID)
}

// SetSubscriptionID sets the "subscription" edge to the UserSubscription entity by ID.
func (pc *PaymentCreate) SetSubscriptionID(id int) *PaymentCreate {
	pc.mutation.SetSubscriptionID(id)
	return pc
}

// SetNillableSubscriptionID sets the "subscription" edge to the UserSubscription entity by ID if the given value is not nil.
func (pc *PaymentCreate) SetNillableSubscriptionID(id *int) *PaymentCreate {
	if id != nil {
		pc = pc.SetSubscriptionID(*id)
	}
	return pc
}

// SetSubscription sets the "subscription" edge to the UserSubscription entity.
func (pc *PaymentCreate) SetSubscription(u *UserSubscription) *PaymentCreate {
	return pc.SetSubscriptionID(u.ID)
}

// Mutation returns the PaymentMutation object of the builder.
func (pc *PaymentCreate) Mutation() *PaymentMutation {
	return pc.mutation
}

// Save creates the Payment in the database.
func (pc *PaymentCreate) Save(ctx context.Context) (*Payment, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PaymentCreate) SaveX(ctx context.Context) *Payment {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PaymentCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PaymentCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PaymentCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := payment.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := payment.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PaymentCreate) check() error {
	if _, ok := pc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Payment.amount"`)}
	}
	if _, ok := pc.mutation.PaymentDate(); !ok {
		return &ValidationError{Name: "payment_date", err: errors.New(`ent: missing required field "Payment.payment_date"`)}
	}
	if _, ok := pc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Payment.status"`)}
	}
	if v, ok := pc.mutation.Status(); ok {
		if err := payment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Payment.status": %w`, err)}
		}
	}
	if _, ok := pc.mutation.PaymentMethod(); !ok {
		return &ValidationError{Name: "payment_method", err: errors.New(`ent: missing required field "Payment.payment_method"`)}
	}
	if _, ok := pc.mutation.ProviderPaymentID(); !ok {
		return &ValidationError{Name: "provider_payment_id", err: errors.New(`ent: missing required field "Payment.provider_payment_id"`)}
	}
	if _, ok := pc.mutation.ProviderInvoiceID(); !ok {
		return &ValidationError{Name: "provider_invoice_id", err: errors.New(`ent: missing required field "Payment.provider_invoice_id"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Payment.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Payment.updated_at"`)}
	}
	return nil
}

func (pc *PaymentCreate) sqlSave(ctx context.Context) (*Payment, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PaymentCreate) createSpec() (*Payment, *sqlgraph.CreateSpec) {
	var (
		_node = &Payment{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(payment.Table, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if value, ok := pc.mutation.PaymentDate(); ok {
		_spec.SetField(payment.FieldPaymentDate, field.TypeTime, value)
		_node.PaymentDate = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(payment.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := pc.mutation.PaymentMethod(); ok {
		_spec.SetField(payment.FieldPaymentMethod, field.TypeString, value)
		_node.PaymentMethod = value
	}
	if value, ok := pc.mutation.ProviderPaymentID(); ok {
		_spec.SetField(payment.FieldProviderPaymentID, field.TypeString, value)
		_node.ProviderPaymentID = value
	}
	if value, ok := pc.mutation.ProviderInvoiceID(); ok {
		_spec.SetField(payment.FieldProviderInvoiceID, field.TypeString, value)
		_node.ProviderInvoiceID = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(payment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(payment.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.user_payments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.SubscriptionIDs(); len(nodes) > 0 {
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
		_node.user_subscription_payments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PaymentCreateBulk is the builder for creating many Payment entities in bulk.
type PaymentCreateBulk struct {
	config
	err      error
	builders []*PaymentCreate
}

// Save creates the Payment entities in the database.
func (pcb *PaymentCreateBulk) Save(ctx context.Context) ([]*Payment, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Payment, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PaymentMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PaymentCreateBulk) SaveX(ctx context.Context) []*Payment {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PaymentCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PaymentCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
