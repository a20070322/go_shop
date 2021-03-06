// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/shop-go/ent/orderaddress"
	"github.com/a20070322/shop-go/ent/orderinfo"
)

// OrderAddressCreate is the builder for creating a OrderAddress entity.
type OrderAddressCreate struct {
	config
	mutation *OrderAddressMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (oac *OrderAddressCreate) SetName(s string) *OrderAddressCreate {
	oac.mutation.SetName(s)
	return oac
}

// SetPhone sets the "phone" field.
func (oac *OrderAddressCreate) SetPhone(s string) *OrderAddressCreate {
	oac.mutation.SetPhone(s)
	return oac
}

// SetProvince sets the "province" field.
func (oac *OrderAddressCreate) SetProvince(s string) *OrderAddressCreate {
	oac.mutation.SetProvince(s)
	return oac
}

// SetCity sets the "city" field.
func (oac *OrderAddressCreate) SetCity(s string) *OrderAddressCreate {
	oac.mutation.SetCity(s)
	return oac
}

// SetArea sets the "area" field.
func (oac *OrderAddressCreate) SetArea(s string) *OrderAddressCreate {
	oac.mutation.SetArea(s)
	return oac
}

// SetDetailed sets the "detailed" field.
func (oac *OrderAddressCreate) SetDetailed(s string) *OrderAddressCreate {
	oac.mutation.SetDetailed(s)
	return oac
}

// SetRemark sets the "remark" field.
func (oac *OrderAddressCreate) SetRemark(s string) *OrderAddressCreate {
	oac.mutation.SetRemark(s)
	return oac
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (oac *OrderAddressCreate) SetNillableRemark(s *string) *OrderAddressCreate {
	if s != nil {
		oac.SetRemark(*s)
	}
	return oac
}

// SetOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID.
func (oac *OrderAddressCreate) SetOrderInfoID(id int) *OrderAddressCreate {
	oac.mutation.SetOrderInfoID(id)
	return oac
}

// SetNillableOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID if the given value is not nil.
func (oac *OrderAddressCreate) SetNillableOrderInfoID(id *int) *OrderAddressCreate {
	if id != nil {
		oac = oac.SetOrderInfoID(*id)
	}
	return oac
}

// SetOrderInfo sets the "order_info" edge to the OrderInfo entity.
func (oac *OrderAddressCreate) SetOrderInfo(o *OrderInfo) *OrderAddressCreate {
	return oac.SetOrderInfoID(o.ID)
}

// Mutation returns the OrderAddressMutation object of the builder.
func (oac *OrderAddressCreate) Mutation() *OrderAddressMutation {
	return oac.mutation
}

// Save creates the OrderAddress in the database.
func (oac *OrderAddressCreate) Save(ctx context.Context) (*OrderAddress, error) {
	var (
		err  error
		node *OrderAddress
	)
	if len(oac.hooks) == 0 {
		if err = oac.check(); err != nil {
			return nil, err
		}
		node, err = oac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderAddressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oac.check(); err != nil {
				return nil, err
			}
			oac.mutation = mutation
			node, err = oac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oac.hooks) - 1; i >= 0; i-- {
			mut = oac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oac *OrderAddressCreate) SaveX(ctx context.Context) *OrderAddress {
	v, err := oac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (oac *OrderAddressCreate) check() error {
	if _, ok := oac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := oac.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New("ent: missing required field \"phone\"")}
	}
	if _, ok := oac.mutation.Province(); !ok {
		return &ValidationError{Name: "province", err: errors.New("ent: missing required field \"province\"")}
	}
	if _, ok := oac.mutation.City(); !ok {
		return &ValidationError{Name: "city", err: errors.New("ent: missing required field \"city\"")}
	}
	if _, ok := oac.mutation.Area(); !ok {
		return &ValidationError{Name: "area", err: errors.New("ent: missing required field \"area\"")}
	}
	if _, ok := oac.mutation.Detailed(); !ok {
		return &ValidationError{Name: "detailed", err: errors.New("ent: missing required field \"detailed\"")}
	}
	return nil
}

func (oac *OrderAddressCreate) sqlSave(ctx context.Context) (*OrderAddress, error) {
	_node, _spec := oac.createSpec()
	if err := sqlgraph.CreateNode(ctx, oac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oac *OrderAddressCreate) createSpec() (*OrderAddress, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderAddress{config: oac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: orderaddress.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderaddress.FieldID,
			},
		}
	)
	if value, ok := oac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderaddress.FieldName,
		})
		_node.Name = value
	}
	if value, ok := oac.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderaddress.FieldPhone,
		})
		_node.Phone = value
	}
	if value, ok := oac.mutation.Province(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderaddress.FieldProvince,
		})
		_node.Province = value
	}
	if value, ok := oac.mutation.City(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderaddress.FieldCity,
		})
		_node.City = value
	}
	if value, ok := oac.mutation.Area(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderaddress.FieldArea,
		})
		_node.Area = value
	}
	if value, ok := oac.mutation.Detailed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderaddress.FieldDetailed,
		})
		_node.Detailed = value
	}
	if value, ok := oac.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderaddress.FieldRemark,
		})
		_node.Remark = value
	}
	if nodes := oac.mutation.OrderInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderaddress.OrderInfoTable,
			Columns: []string{orderaddress.OrderInfoColumn},
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
		_node.order_info_order_address = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderAddressCreateBulk is the builder for creating many OrderAddress entities in bulk.
type OrderAddressCreateBulk struct {
	config
	builders []*OrderAddressCreate
}

// Save creates the OrderAddress entities in the database.
func (oacb *OrderAddressCreateBulk) Save(ctx context.Context) ([]*OrderAddress, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oacb.builders))
	nodes := make([]*OrderAddress, len(oacb.builders))
	mutators := make([]Mutator, len(oacb.builders))
	for i := range oacb.builders {
		func(i int, root context.Context) {
			builder := oacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderAddressMutation)
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
					_, err = mutators[i+1].Mutate(root, oacb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oacb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oacb *OrderAddressCreateBulk) SaveX(ctx context.Context) []*OrderAddress {
	v, err := oacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
