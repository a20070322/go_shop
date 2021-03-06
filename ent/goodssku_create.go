// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/shop-go/ent/goodssku"
	"github.com/a20070322/shop-go/ent/goodsspecsoption"
	"github.com/a20070322/shop-go/ent/goodsspu"
)

// GoodsSkuCreate is the builder for creating a GoodsSku entity.
type GoodsSkuCreate struct {
	config
	mutation *GoodsSkuMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (gsc *GoodsSkuCreate) SetCreatedAt(t time.Time) *GoodsSkuCreate {
	gsc.mutation.SetCreatedAt(t)
	return gsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gsc *GoodsSkuCreate) SetNillableCreatedAt(t *time.Time) *GoodsSkuCreate {
	if t != nil {
		gsc.SetCreatedAt(*t)
	}
	return gsc
}

// SetUpdatedAt sets the "updated_at" field.
func (gsc *GoodsSkuCreate) SetUpdatedAt(t time.Time) *GoodsSkuCreate {
	gsc.mutation.SetUpdatedAt(t)
	return gsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gsc *GoodsSkuCreate) SetNillableUpdatedAt(t *time.Time) *GoodsSkuCreate {
	if t != nil {
		gsc.SetUpdatedAt(*t)
	}
	return gsc
}

// SetDeletedAt sets the "deleted_at" field.
func (gsc *GoodsSkuCreate) SetDeletedAt(t time.Time) *GoodsSkuCreate {
	gsc.mutation.SetDeletedAt(t)
	return gsc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gsc *GoodsSkuCreate) SetNillableDeletedAt(t *time.Time) *GoodsSkuCreate {
	if t != nil {
		gsc.SetDeletedAt(*t)
	}
	return gsc
}

// SetSkuName sets the "sku_name" field.
func (gsc *GoodsSkuCreate) SetSkuName(s string) *GoodsSkuCreate {
	gsc.mutation.SetSkuName(s)
	return gsc
}

// SetSkuCode sets the "sku_code" field.
func (gsc *GoodsSkuCreate) SetSkuCode(s string) *GoodsSkuCreate {
	gsc.mutation.SetSkuCode(s)
	return gsc
}

// SetStockNum sets the "stock_num" field.
func (gsc *GoodsSkuCreate) SetStockNum(i int) *GoodsSkuCreate {
	gsc.mutation.SetStockNum(i)
	return gsc
}

// SetNillableStockNum sets the "stock_num" field if the given value is not nil.
func (gsc *GoodsSkuCreate) SetNillableStockNum(i *int) *GoodsSkuCreate {
	if i != nil {
		gsc.SetStockNum(*i)
	}
	return gsc
}

// SetSalesNum sets the "sales_num" field.
func (gsc *GoodsSkuCreate) SetSalesNum(i int) *GoodsSkuCreate {
	gsc.mutation.SetSalesNum(i)
	return gsc
}

// SetNillableSalesNum sets the "sales_num" field if the given value is not nil.
func (gsc *GoodsSkuCreate) SetNillableSalesNum(i *int) *GoodsSkuCreate {
	if i != nil {
		gsc.SetSalesNum(*i)
	}
	return gsc
}

// SetPrice sets the "price" field.
func (gsc *GoodsSkuCreate) SetPrice(i int) *GoodsSkuCreate {
	gsc.mutation.SetPrice(i)
	return gsc
}

// SetGoodsSpuID sets the "goods_spu" edge to the GoodsSpu entity by ID.
func (gsc *GoodsSkuCreate) SetGoodsSpuID(id int) *GoodsSkuCreate {
	gsc.mutation.SetGoodsSpuID(id)
	return gsc
}

// SetNillableGoodsSpuID sets the "goods_spu" edge to the GoodsSpu entity by ID if the given value is not nil.
func (gsc *GoodsSkuCreate) SetNillableGoodsSpuID(id *int) *GoodsSkuCreate {
	if id != nil {
		gsc = gsc.SetGoodsSpuID(*id)
	}
	return gsc
}

// SetGoodsSpu sets the "goods_spu" edge to the GoodsSpu entity.
func (gsc *GoodsSkuCreate) SetGoodsSpu(g *GoodsSpu) *GoodsSkuCreate {
	return gsc.SetGoodsSpuID(g.ID)
}

// AddGoodsSpecsOptionIDs adds the "goods_specs_option" edge to the GoodsSpecsOption entity by IDs.
func (gsc *GoodsSkuCreate) AddGoodsSpecsOptionIDs(ids ...int) *GoodsSkuCreate {
	gsc.mutation.AddGoodsSpecsOptionIDs(ids...)
	return gsc
}

// AddGoodsSpecsOption adds the "goods_specs_option" edges to the GoodsSpecsOption entity.
func (gsc *GoodsSkuCreate) AddGoodsSpecsOption(g ...*GoodsSpecsOption) *GoodsSkuCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gsc.AddGoodsSpecsOptionIDs(ids...)
}

// Mutation returns the GoodsSkuMutation object of the builder.
func (gsc *GoodsSkuCreate) Mutation() *GoodsSkuMutation {
	return gsc.mutation
}

// Save creates the GoodsSku in the database.
func (gsc *GoodsSkuCreate) Save(ctx context.Context) (*GoodsSku, error) {
	var (
		err  error
		node *GoodsSku
	)
	gsc.defaults()
	if len(gsc.hooks) == 0 {
		if err = gsc.check(); err != nil {
			return nil, err
		}
		node, err = gsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodsSkuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gsc.check(); err != nil {
				return nil, err
			}
			gsc.mutation = mutation
			node, err = gsc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gsc.hooks) - 1; i >= 0; i-- {
			mut = gsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gsc *GoodsSkuCreate) SaveX(ctx context.Context) *GoodsSku {
	v, err := gsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (gsc *GoodsSkuCreate) defaults() {
	if _, ok := gsc.mutation.CreatedAt(); !ok {
		v := goodssku.DefaultCreatedAt()
		gsc.mutation.SetCreatedAt(v)
	}
	if _, ok := gsc.mutation.StockNum(); !ok {
		v := goodssku.DefaultStockNum
		gsc.mutation.SetStockNum(v)
	}
	if _, ok := gsc.mutation.SalesNum(); !ok {
		v := goodssku.DefaultSalesNum
		gsc.mutation.SetSalesNum(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gsc *GoodsSkuCreate) check() error {
	if _, ok := gsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := gsc.mutation.SkuName(); !ok {
		return &ValidationError{Name: "sku_name", err: errors.New("ent: missing required field \"sku_name\"")}
	}
	if _, ok := gsc.mutation.SkuCode(); !ok {
		return &ValidationError{Name: "sku_code", err: errors.New("ent: missing required field \"sku_code\"")}
	}
	if _, ok := gsc.mutation.StockNum(); !ok {
		return &ValidationError{Name: "stock_num", err: errors.New("ent: missing required field \"stock_num\"")}
	}
	if _, ok := gsc.mutation.SalesNum(); !ok {
		return &ValidationError{Name: "sales_num", err: errors.New("ent: missing required field \"sales_num\"")}
	}
	if _, ok := gsc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	return nil
}

func (gsc *GoodsSkuCreate) sqlSave(ctx context.Context) (*GoodsSku, error) {
	_node, _spec := gsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gsc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gsc *GoodsSkuCreate) createSpec() (*GoodsSku, *sqlgraph.CreateSpec) {
	var (
		_node = &GoodsSku{config: gsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: goodssku.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: goodssku.FieldID,
			},
		}
	)
	if value, ok := gsc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goodssku.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := gsc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goodssku.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := gsc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goodssku.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := gsc.mutation.SkuName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodssku.FieldSkuName,
		})
		_node.SkuName = value
	}
	if value, ok := gsc.mutation.SkuCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodssku.FieldSkuCode,
		})
		_node.SkuCode = value
	}
	if value, ok := gsc.mutation.StockNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodssku.FieldStockNum,
		})
		_node.StockNum = value
	}
	if value, ok := gsc.mutation.SalesNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodssku.FieldSalesNum,
		})
		_node.SalesNum = value
	}
	if value, ok := gsc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodssku.FieldPrice,
		})
		_node.Price = value
	}
	if nodes := gsc.mutation.GoodsSpuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   goodssku.GoodsSpuTable,
			Columns: []string{goodssku.GoodsSpuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodsspu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.goods_spu_goods_sku = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gsc.mutation.GoodsSpecsOptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   goodssku.GoodsSpecsOptionTable,
			Columns: goodssku.GoodsSpecsOptionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodsspecsoption.FieldID,
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

// GoodsSkuCreateBulk is the builder for creating many GoodsSku entities in bulk.
type GoodsSkuCreateBulk struct {
	config
	builders []*GoodsSkuCreate
}

// Save creates the GoodsSku entities in the database.
func (gscb *GoodsSkuCreateBulk) Save(ctx context.Context) ([]*GoodsSku, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gscb.builders))
	nodes := make([]*GoodsSku, len(gscb.builders))
	mutators := make([]Mutator, len(gscb.builders))
	for i := range gscb.builders {
		func(i int, root context.Context) {
			builder := gscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GoodsSkuMutation)
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
					_, err = mutators[i+1].Mutate(root, gscb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gscb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gscb *GoodsSkuCreateBulk) SaveX(ctx context.Context) []*GoodsSku {
	v, err := gscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
