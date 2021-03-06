// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/shop-go/ent/customer"
	"github.com/a20070322/shop-go/ent/customeraddress"
	"github.com/a20070322/shop-go/ent/orderinfo"
)

// CustomerCreate is the builder for creating a Customer entity.
type CustomerCreate struct {
	config
	mutation *CustomerMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CustomerCreate) SetCreatedAt(t time.Time) *CustomerCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableCreatedAt(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CustomerCreate) SetUpdatedAt(t time.Time) *CustomerCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableUpdatedAt(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CustomerCreate) SetDeletedAt(t time.Time) *CustomerCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableDeletedAt(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetMiniOpenid sets the "mini_openid" field.
func (cc *CustomerCreate) SetMiniOpenid(s string) *CustomerCreate {
	cc.mutation.SetMiniOpenid(s)
	return cc
}

// SetNillableMiniOpenid sets the "mini_openid" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableMiniOpenid(s *string) *CustomerCreate {
	if s != nil {
		cc.SetMiniOpenid(*s)
	}
	return cc
}

// SetWechatOpenid sets the "wechat_openid" field.
func (cc *CustomerCreate) SetWechatOpenid(s string) *CustomerCreate {
	cc.mutation.SetWechatOpenid(s)
	return cc
}

// SetNillableWechatOpenid sets the "wechat_openid" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableWechatOpenid(s *string) *CustomerCreate {
	if s != nil {
		cc.SetWechatOpenid(*s)
	}
	return cc
}

// SetUnionID sets the "union_id" field.
func (cc *CustomerCreate) SetUnionID(s string) *CustomerCreate {
	cc.mutation.SetUnionID(s)
	return cc
}

// SetNillableUnionID sets the "union_id" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableUnionID(s *string) *CustomerCreate {
	if s != nil {
		cc.SetUnionID(*s)
	}
	return cc
}

// SetPhone sets the "phone" field.
func (cc *CustomerCreate) SetPhone(s string) *CustomerCreate {
	cc.mutation.SetPhone(s)
	return cc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (cc *CustomerCreate) SetNillablePhone(s *string) *CustomerCreate {
	if s != nil {
		cc.SetPhone(*s)
	}
	return cc
}

// SetAvatar sets the "avatar" field.
func (cc *CustomerCreate) SetAvatar(s string) *CustomerCreate {
	cc.mutation.SetAvatar(s)
	return cc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableAvatar(s *string) *CustomerCreate {
	if s != nil {
		cc.SetAvatar(*s)
	}
	return cc
}

// SetIsDisable sets the "is_disable" field.
func (cc *CustomerCreate) SetIsDisable(b bool) *CustomerCreate {
	cc.mutation.SetIsDisable(b)
	return cc
}

// SetNillableIsDisable sets the "is_disable" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableIsDisable(b *bool) *CustomerCreate {
	if b != nil {
		cc.SetIsDisable(*b)
	}
	return cc
}

// AddAddresIDs adds the "address" edge to the CustomerAddress entity by IDs.
func (cc *CustomerCreate) AddAddresIDs(ids ...int) *CustomerCreate {
	cc.mutation.AddAddresIDs(ids...)
	return cc
}

// AddAddress adds the "address" edges to the CustomerAddress entity.
func (cc *CustomerCreate) AddAddress(c ...*CustomerAddress) *CustomerCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddAddresIDs(ids...)
}

// AddOrderInfoIDs adds the "order_info" edge to the OrderInfo entity by IDs.
func (cc *CustomerCreate) AddOrderInfoIDs(ids ...int) *CustomerCreate {
	cc.mutation.AddOrderInfoIDs(ids...)
	return cc
}

// AddOrderInfo adds the "order_info" edges to the OrderInfo entity.
func (cc *CustomerCreate) AddOrderInfo(o ...*OrderInfo) *CustomerCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cc.AddOrderInfoIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cc *CustomerCreate) Mutation() *CustomerMutation {
	return cc.mutation
}

// Save creates the Customer in the database.
func (cc *CustomerCreate) Save(ctx context.Context) (*Customer, error) {
	var (
		err  error
		node *Customer
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CustomerCreate) SaveX(ctx context.Context) *Customer {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (cc *CustomerCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := customer.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.IsDisable(); !ok {
		v := customer.DefaultIsDisable
		cc.mutation.SetIsDisable(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CustomerCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	return nil
}

func (cc *CustomerCreate) sqlSave(ctx context.Context) (*Customer, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CustomerCreate) createSpec() (*Customer, *sqlgraph.CreateSpec) {
	var (
		_node = &Customer{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: customer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customer.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.MiniOpenid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldMiniOpenid,
		})
		_node.MiniOpenid = value
	}
	if value, ok := cc.mutation.WechatOpenid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldWechatOpenid,
		})
		_node.WechatOpenid = value
	}
	if value, ok := cc.mutation.UnionID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldUnionID,
		})
		_node.UnionID = value
	}
	if value, ok := cc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPhone,
		})
		_node.Phone = value
	}
	if value, ok := cc.mutation.Avatar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAvatar,
		})
		_node.Avatar = value
	}
	if value, ok := cc.mutation.IsDisable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: customer.FieldIsDisable,
		})
		_node.IsDisable = value
	}
	if nodes := cc.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.AddressTable,
			Columns: []string{customer.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customeraddress.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.OrderInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.OrderInfoTable,
			Columns: []string{customer.OrderInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CustomerCreateBulk is the builder for creating many Customer entities in bulk.
type CustomerCreateBulk struct {
	config
	builders []*CustomerCreate
}

// Save creates the Customer entities in the database.
func (ccb *CustomerCreateBulk) Save(ctx context.Context) ([]*Customer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Customer, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CustomerCreateBulk) SaveX(ctx context.Context) []*Customer {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
