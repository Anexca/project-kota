// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/payment"
	"common/ent/subscription"
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

// UserSubscriptionCreate is the builder for creating a UserSubscription entity.
type UserSubscriptionCreate struct {
	config
	mutation *UserSubscriptionMutation
	hooks    []Hook
}

// SetIsActive sets the "is_active" field.
func (usc *UserSubscriptionCreate) SetIsActive(b bool) *UserSubscriptionCreate {
	usc.mutation.SetIsActive(b)
	return usc
}

// SetStatus sets the "status" field.
func (usc *UserSubscriptionCreate) SetStatus(u usersubscription.Status) *UserSubscriptionCreate {
	usc.mutation.SetStatus(u)
	return usc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (usc *UserSubscriptionCreate) SetNillableStatus(u *usersubscription.Status) *UserSubscriptionCreate {
	if u != nil {
		usc.SetStatus(*u)
	}
	return usc
}

// SetStartDate sets the "start_date" field.
func (usc *UserSubscriptionCreate) SetStartDate(t time.Time) *UserSubscriptionCreate {
	usc.mutation.SetStartDate(t)
	return usc
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (usc *UserSubscriptionCreate) SetNillableStartDate(t *time.Time) *UserSubscriptionCreate {
	if t != nil {
		usc.SetStartDate(*t)
	}
	return usc
}

// SetEndDate sets the "end_date" field.
func (usc *UserSubscriptionCreate) SetEndDate(t time.Time) *UserSubscriptionCreate {
	usc.mutation.SetEndDate(t)
	return usc
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (usc *UserSubscriptionCreate) SetNillableEndDate(t *time.Time) *UserSubscriptionCreate {
	if t != nil {
		usc.SetEndDate(*t)
	}
	return usc
}

// SetProviderSubscriptionID sets the "provider_subscription_id" field.
func (usc *UserSubscriptionCreate) SetProviderSubscriptionID(s string) *UserSubscriptionCreate {
	usc.mutation.SetProviderSubscriptionID(s)
	return usc
}

// SetCreatedAt sets the "created_at" field.
func (usc *UserSubscriptionCreate) SetCreatedAt(t time.Time) *UserSubscriptionCreate {
	usc.mutation.SetCreatedAt(t)
	return usc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (usc *UserSubscriptionCreate) SetNillableCreatedAt(t *time.Time) *UserSubscriptionCreate {
	if t != nil {
		usc.SetCreatedAt(*t)
	}
	return usc
}

// SetUpdatedAt sets the "updated_at" field.
func (usc *UserSubscriptionCreate) SetUpdatedAt(t time.Time) *UserSubscriptionCreate {
	usc.mutation.SetUpdatedAt(t)
	return usc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (usc *UserSubscriptionCreate) SetNillableUpdatedAt(t *time.Time) *UserSubscriptionCreate {
	if t != nil {
		usc.SetUpdatedAt(*t)
	}
	return usc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (usc *UserSubscriptionCreate) SetUserID(id uuid.UUID) *UserSubscriptionCreate {
	usc.mutation.SetUserID(id)
	return usc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (usc *UserSubscriptionCreate) SetNillableUserID(id *uuid.UUID) *UserSubscriptionCreate {
	if id != nil {
		usc = usc.SetUserID(*id)
	}
	return usc
}

// SetUser sets the "user" edge to the User entity.
func (usc *UserSubscriptionCreate) SetUser(u *User) *UserSubscriptionCreate {
	return usc.SetUserID(u.ID)
}

// SetSubscriptionID sets the "subscription" edge to the Subscription entity by ID.
func (usc *UserSubscriptionCreate) SetSubscriptionID(id int) *UserSubscriptionCreate {
	usc.mutation.SetSubscriptionID(id)
	return usc
}

// SetNillableSubscriptionID sets the "subscription" edge to the Subscription entity by ID if the given value is not nil.
func (usc *UserSubscriptionCreate) SetNillableSubscriptionID(id *int) *UserSubscriptionCreate {
	if id != nil {
		usc = usc.SetSubscriptionID(*id)
	}
	return usc
}

// SetSubscription sets the "subscription" edge to the Subscription entity.
func (usc *UserSubscriptionCreate) SetSubscription(s *Subscription) *UserSubscriptionCreate {
	return usc.SetSubscriptionID(s.ID)
}

// AddPaymentIDs adds the "payments" edge to the Payment entity by IDs.
func (usc *UserSubscriptionCreate) AddPaymentIDs(ids ...int) *UserSubscriptionCreate {
	usc.mutation.AddPaymentIDs(ids...)
	return usc
}

// AddPayments adds the "payments" edges to the Payment entity.
func (usc *UserSubscriptionCreate) AddPayments(p ...*Payment) *UserSubscriptionCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return usc.AddPaymentIDs(ids...)
}

// Mutation returns the UserSubscriptionMutation object of the builder.
func (usc *UserSubscriptionCreate) Mutation() *UserSubscriptionMutation {
	return usc.mutation
}

// Save creates the UserSubscription in the database.
func (usc *UserSubscriptionCreate) Save(ctx context.Context) (*UserSubscription, error) {
	usc.defaults()
	return withHooks(ctx, usc.sqlSave, usc.mutation, usc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (usc *UserSubscriptionCreate) SaveX(ctx context.Context) *UserSubscription {
	v, err := usc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (usc *UserSubscriptionCreate) Exec(ctx context.Context) error {
	_, err := usc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usc *UserSubscriptionCreate) ExecX(ctx context.Context) {
	if err := usc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usc *UserSubscriptionCreate) defaults() {
	if _, ok := usc.mutation.Status(); !ok {
		v := usersubscription.DefaultStatus
		usc.mutation.SetStatus(v)
	}
	if _, ok := usc.mutation.CreatedAt(); !ok {
		v := usersubscription.DefaultCreatedAt()
		usc.mutation.SetCreatedAt(v)
	}
	if _, ok := usc.mutation.UpdatedAt(); !ok {
		v := usersubscription.DefaultUpdatedAt()
		usc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usc *UserSubscriptionCreate) check() error {
	if _, ok := usc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "UserSubscription.is_active"`)}
	}
	if _, ok := usc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "UserSubscription.status"`)}
	}
	if v, ok := usc.mutation.Status(); ok {
		if err := usersubscription.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "UserSubscription.status": %w`, err)}
		}
	}
	if _, ok := usc.mutation.ProviderSubscriptionID(); !ok {
		return &ValidationError{Name: "provider_subscription_id", err: errors.New(`ent: missing required field "UserSubscription.provider_subscription_id"`)}
	}
	if _, ok := usc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserSubscription.created_at"`)}
	}
	if _, ok := usc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UserSubscription.updated_at"`)}
	}
	return nil
}

func (usc *UserSubscriptionCreate) sqlSave(ctx context.Context) (*UserSubscription, error) {
	if err := usc.check(); err != nil {
		return nil, err
	}
	_node, _spec := usc.createSpec()
	if err := sqlgraph.CreateNode(ctx, usc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	usc.mutation.id = &_node.ID
	usc.mutation.done = true
	return _node, nil
}

func (usc *UserSubscriptionCreate) createSpec() (*UserSubscription, *sqlgraph.CreateSpec) {
	var (
		_node = &UserSubscription{config: usc.config}
		_spec = sqlgraph.NewCreateSpec(usersubscription.Table, sqlgraph.NewFieldSpec(usersubscription.FieldID, field.TypeInt))
	)
	if value, ok := usc.mutation.IsActive(); ok {
		_spec.SetField(usersubscription.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := usc.mutation.Status(); ok {
		_spec.SetField(usersubscription.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := usc.mutation.StartDate(); ok {
		_spec.SetField(usersubscription.FieldStartDate, field.TypeTime, value)
		_node.StartDate = value
	}
	if value, ok := usc.mutation.EndDate(); ok {
		_spec.SetField(usersubscription.FieldEndDate, field.TypeTime, value)
		_node.EndDate = value
	}
	if value, ok := usc.mutation.ProviderSubscriptionID(); ok {
		_spec.SetField(usersubscription.FieldProviderSubscriptionID, field.TypeString, value)
		_node.ProviderSubscriptionID = value
	}
	if value, ok := usc.mutation.CreatedAt(); ok {
		_spec.SetField(usersubscription.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := usc.mutation.UpdatedAt(); ok {
		_spec.SetField(usersubscription.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := usc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.user_subscriptions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := usc.mutation.SubscriptionIDs(); len(nodes) > 0 {
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
		_node.subscription_user_subscriptions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := usc.mutation.PaymentsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserSubscriptionCreateBulk is the builder for creating many UserSubscription entities in bulk.
type UserSubscriptionCreateBulk struct {
	config
	err      error
	builders []*UserSubscriptionCreate
}

// Save creates the UserSubscription entities in the database.
func (uscb *UserSubscriptionCreateBulk) Save(ctx context.Context) ([]*UserSubscription, error) {
	if uscb.err != nil {
		return nil, uscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uscb.builders))
	nodes := make([]*UserSubscription, len(uscb.builders))
	mutators := make([]Mutator, len(uscb.builders))
	for i := range uscb.builders {
		func(i int, root context.Context) {
			builder := uscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserSubscriptionMutation)
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
					_, err = mutators[i+1].Mutate(root, uscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uscb *UserSubscriptionCreateBulk) SaveX(ctx context.Context) []*UserSubscription {
	v, err := uscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uscb *UserSubscriptionCreateBulk) Exec(ctx context.Context) error {
	_, err := uscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uscb *UserSubscriptionCreateBulk) ExecX(ctx context.Context) {
	if err := uscb.Exec(ctx); err != nil {
		panic(err)
	}
}
