// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/shop-go/ent/goodssku"
	"github.com/a20070322/shop-go/ent/goodsspecsoption"
	"github.com/a20070322/shop-go/ent/goodsspu"
	"github.com/a20070322/shop-go/ent/predicate"
)

// GoodsSkuUpdate is the builder for updating GoodsSku entities.
type GoodsSkuUpdate struct {
	config
	hooks    []Hook
	mutation *GoodsSkuMutation
}

// Where adds a new predicate for the GoodsSkuUpdate builder.
func (gsu *GoodsSkuUpdate) Where(ps ...predicate.GoodsSku) *GoodsSkuUpdate {
	gsu.mutation.predicates = append(gsu.mutation.predicates, ps...)
	return gsu
}

// SetUpdatedAt sets the "updated_at" field.
func (gsu *GoodsSkuUpdate) SetUpdatedAt(t time.Time) *GoodsSkuUpdate {
	gsu.mutation.SetUpdatedAt(t)
	return gsu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (gsu *GoodsSkuUpdate) ClearUpdatedAt() *GoodsSkuUpdate {
	gsu.mutation.ClearUpdatedAt()
	return gsu
}

// SetDeletedAt sets the "deleted_at" field.
func (gsu *GoodsSkuUpdate) SetDeletedAt(t time.Time) *GoodsSkuUpdate {
	gsu.mutation.SetDeletedAt(t)
	return gsu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gsu *GoodsSkuUpdate) SetNillableDeletedAt(t *time.Time) *GoodsSkuUpdate {
	if t != nil {
		gsu.SetDeletedAt(*t)
	}
	return gsu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gsu *GoodsSkuUpdate) ClearDeletedAt() *GoodsSkuUpdate {
	gsu.mutation.ClearDeletedAt()
	return gsu
}

// SetSkuName sets the "sku_name" field.
func (gsu *GoodsSkuUpdate) SetSkuName(s string) *GoodsSkuUpdate {
	gsu.mutation.SetSkuName(s)
	return gsu
}

// SetSkuCode sets the "sku_code" field.
func (gsu *GoodsSkuUpdate) SetSkuCode(s string) *GoodsSkuUpdate {
	gsu.mutation.SetSkuCode(s)
	return gsu
}

// SetStockNum sets the "stock_num" field.
func (gsu *GoodsSkuUpdate) SetStockNum(i int) *GoodsSkuUpdate {
	gsu.mutation.ResetStockNum()
	gsu.mutation.SetStockNum(i)
	return gsu
}

// SetNillableStockNum sets the "stock_num" field if the given value is not nil.
func (gsu *GoodsSkuUpdate) SetNillableStockNum(i *int) *GoodsSkuUpdate {
	if i != nil {
		gsu.SetStockNum(*i)
	}
	return gsu
}

// AddStockNum adds i to the "stock_num" field.
func (gsu *GoodsSkuUpdate) AddStockNum(i int) *GoodsSkuUpdate {
	gsu.mutation.AddStockNum(i)
	return gsu
}

// SetSalesNum sets the "sales_num" field.
func (gsu *GoodsSkuUpdate) SetSalesNum(i int) *GoodsSkuUpdate {
	gsu.mutation.ResetSalesNum()
	gsu.mutation.SetSalesNum(i)
	return gsu
}

// SetNillableSalesNum sets the "sales_num" field if the given value is not nil.
func (gsu *GoodsSkuUpdate) SetNillableSalesNum(i *int) *GoodsSkuUpdate {
	if i != nil {
		gsu.SetSalesNum(*i)
	}
	return gsu
}

// AddSalesNum adds i to the "sales_num" field.
func (gsu *GoodsSkuUpdate) AddSalesNum(i int) *GoodsSkuUpdate {
	gsu.mutation.AddSalesNum(i)
	return gsu
}

// SetPrice sets the "price" field.
func (gsu *GoodsSkuUpdate) SetPrice(i int) *GoodsSkuUpdate {
	gsu.mutation.ResetPrice()
	gsu.mutation.SetPrice(i)
	return gsu
}

// AddPrice adds i to the "price" field.
func (gsu *GoodsSkuUpdate) AddPrice(i int) *GoodsSkuUpdate {
	gsu.mutation.AddPrice(i)
	return gsu
}

// SetGoodsSpuID sets the "goods_spu" edge to the GoodsSpu entity by ID.
func (gsu *GoodsSkuUpdate) SetGoodsSpuID(id int) *GoodsSkuUpdate {
	gsu.mutation.SetGoodsSpuID(id)
	return gsu
}

// SetNillableGoodsSpuID sets the "goods_spu" edge to the GoodsSpu entity by ID if the given value is not nil.
func (gsu *GoodsSkuUpdate) SetNillableGoodsSpuID(id *int) *GoodsSkuUpdate {
	if id != nil {
		gsu = gsu.SetGoodsSpuID(*id)
	}
	return gsu
}

// SetGoodsSpu sets the "goods_spu" edge to the GoodsSpu entity.
func (gsu *GoodsSkuUpdate) SetGoodsSpu(g *GoodsSpu) *GoodsSkuUpdate {
	return gsu.SetGoodsSpuID(g.ID)
}

// AddGoodsSpecsOptionIDs adds the "goods_specs_option" edge to the GoodsSpecsOption entity by IDs.
func (gsu *GoodsSkuUpdate) AddGoodsSpecsOptionIDs(ids ...int) *GoodsSkuUpdate {
	gsu.mutation.AddGoodsSpecsOptionIDs(ids...)
	return gsu
}

// AddGoodsSpecsOption adds the "goods_specs_option" edges to the GoodsSpecsOption entity.
func (gsu *GoodsSkuUpdate) AddGoodsSpecsOption(g ...*GoodsSpecsOption) *GoodsSkuUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gsu.AddGoodsSpecsOptionIDs(ids...)
}

// Mutation returns the GoodsSkuMutation object of the builder.
func (gsu *GoodsSkuUpdate) Mutation() *GoodsSkuMutation {
	return gsu.mutation
}

// ClearGoodsSpu clears the "goods_spu" edge to the GoodsSpu entity.
func (gsu *GoodsSkuUpdate) ClearGoodsSpu() *GoodsSkuUpdate {
	gsu.mutation.ClearGoodsSpu()
	return gsu
}

// ClearGoodsSpecsOption clears all "goods_specs_option" edges to the GoodsSpecsOption entity.
func (gsu *GoodsSkuUpdate) ClearGoodsSpecsOption() *GoodsSkuUpdate {
	gsu.mutation.ClearGoodsSpecsOption()
	return gsu
}

// RemoveGoodsSpecsOptionIDs removes the "goods_specs_option" edge to GoodsSpecsOption entities by IDs.
func (gsu *GoodsSkuUpdate) RemoveGoodsSpecsOptionIDs(ids ...int) *GoodsSkuUpdate {
	gsu.mutation.RemoveGoodsSpecsOptionIDs(ids...)
	return gsu
}

// RemoveGoodsSpecsOption removes "goods_specs_option" edges to GoodsSpecsOption entities.
func (gsu *GoodsSkuUpdate) RemoveGoodsSpecsOption(g ...*GoodsSpecsOption) *GoodsSkuUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gsu.RemoveGoodsSpecsOptionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gsu *GoodsSkuUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	gsu.defaults()
	if len(gsu.hooks) == 0 {
		affected, err = gsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodsSkuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gsu.mutation = mutation
			affected, err = gsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gsu.hooks) - 1; i >= 0; i-- {
			mut = gsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gsu *GoodsSkuUpdate) SaveX(ctx context.Context) int {
	affected, err := gsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gsu *GoodsSkuUpdate) Exec(ctx context.Context) error {
	_, err := gsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gsu *GoodsSkuUpdate) ExecX(ctx context.Context) {
	if err := gsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gsu *GoodsSkuUpdate) defaults() {
	if _, ok := gsu.mutation.UpdatedAt(); !ok && !gsu.mutation.UpdatedAtCleared() {
		v := goodssku.UpdateDefaultUpdatedAt()
		gsu.mutation.SetUpdatedAt(v)
	}
}

func (gsu *GoodsSkuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodssku.Table,
			Columns: goodssku.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: goodssku.FieldID,
			},
		},
	}
	if ps := gsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gsu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goodssku.FieldUpdatedAt,
		})
	}
	if gsu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: goodssku.FieldUpdatedAt,
		})
	}
	if value, ok := gsu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goodssku.FieldDeletedAt,
		})
	}
	if gsu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: goodssku.FieldDeletedAt,
		})
	}
	if value, ok := gsu.mutation.SkuName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodssku.FieldSkuName,
		})
	}
	if value, ok := gsu.mutation.SkuCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodssku.FieldSkuCode,
		})
	}
	if value, ok := gsu.mutation.StockNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodssku.FieldStockNum,
		})
	}
	if value, ok := gsu.mutation.AddedStockNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodssku.FieldStockNum,
		})
	}
	if value, ok := gsu.mutation.SalesNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodssku.FieldSalesNum,
		})
	}
	if value, ok := gsu.mutation.AddedSalesNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodssku.FieldSalesNum,
		})
	}
	if value, ok := gsu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodssku.FieldPrice,
		})
	}
	if value, ok := gsu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodssku.FieldPrice,
		})
	}
	if gsu.mutation.GoodsSpuCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsu.mutation.GoodsSpuIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gsu.mutation.GoodsSpecsOptionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsu.mutation.RemovedGoodsSpecsOptionIDs(); len(nodes) > 0 && !gsu.mutation.GoodsSpecsOptionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsu.mutation.GoodsSpecsOptionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodssku.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GoodsSkuUpdateOne is the builder for updating a single GoodsSku entity.
type GoodsSkuUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GoodsSkuMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (gsuo *GoodsSkuUpdateOne) SetUpdatedAt(t time.Time) *GoodsSkuUpdateOne {
	gsuo.mutation.SetUpdatedAt(t)
	return gsuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (gsuo *GoodsSkuUpdateOne) ClearUpdatedAt() *GoodsSkuUpdateOne {
	gsuo.mutation.ClearUpdatedAt()
	return gsuo
}

// SetDeletedAt sets the "deleted_at" field.
func (gsuo *GoodsSkuUpdateOne) SetDeletedAt(t time.Time) *GoodsSkuUpdateOne {
	gsuo.mutation.SetDeletedAt(t)
	return gsuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gsuo *GoodsSkuUpdateOne) SetNillableDeletedAt(t *time.Time) *GoodsSkuUpdateOne {
	if t != nil {
		gsuo.SetDeletedAt(*t)
	}
	return gsuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gsuo *GoodsSkuUpdateOne) ClearDeletedAt() *GoodsSkuUpdateOne {
	gsuo.mutation.ClearDeletedAt()
	return gsuo
}

// SetSkuName sets the "sku_name" field.
func (gsuo *GoodsSkuUpdateOne) SetSkuName(s string) *GoodsSkuUpdateOne {
	gsuo.mutation.SetSkuName(s)
	return gsuo
}

// SetSkuCode sets the "sku_code" field.
func (gsuo *GoodsSkuUpdateOne) SetSkuCode(s string) *GoodsSkuUpdateOne {
	gsuo.mutation.SetSkuCode(s)
	return gsuo
}

// SetStockNum sets the "stock_num" field.
func (gsuo *GoodsSkuUpdateOne) SetStockNum(i int) *GoodsSkuUpdateOne {
	gsuo.mutation.ResetStockNum()
	gsuo.mutation.SetStockNum(i)
	return gsuo
}

// SetNillableStockNum sets the "stock_num" field if the given value is not nil.
func (gsuo *GoodsSkuUpdateOne) SetNillableStockNum(i *int) *GoodsSkuUpdateOne {
	if i != nil {
		gsuo.SetStockNum(*i)
	}
	return gsuo
}

// AddStockNum adds i to the "stock_num" field.
func (gsuo *GoodsSkuUpdateOne) AddStockNum(i int) *GoodsSkuUpdateOne {
	gsuo.mutation.AddStockNum(i)
	return gsuo
}

// SetSalesNum sets the "sales_num" field.
func (gsuo *GoodsSkuUpdateOne) SetSalesNum(i int) *GoodsSkuUpdateOne {
	gsuo.mutation.ResetSalesNum()
	gsuo.mutation.SetSalesNum(i)
	return gsuo
}

// SetNillableSalesNum sets the "sales_num" field if the given value is not nil.
func (gsuo *GoodsSkuUpdateOne) SetNillableSalesNum(i *int) *GoodsSkuUpdateOne {
	if i != nil {
		gsuo.SetSalesNum(*i)
	}
	return gsuo
}

// AddSalesNum adds i to the "sales_num" field.
func (gsuo *GoodsSkuUpdateOne) AddSalesNum(i int) *GoodsSkuUpdateOne {
	gsuo.mutation.AddSalesNum(i)
	return gsuo
}

// SetPrice sets the "price" field.
func (gsuo *GoodsSkuUpdateOne) SetPrice(i int) *GoodsSkuUpdateOne {
	gsuo.mutation.ResetPrice()
	gsuo.mutation.SetPrice(i)
	return gsuo
}

// AddPrice adds i to the "price" field.
func (gsuo *GoodsSkuUpdateOne) AddPrice(i int) *GoodsSkuUpdateOne {
	gsuo.mutation.AddPrice(i)
	return gsuo
}

// SetGoodsSpuID sets the "goods_spu" edge to the GoodsSpu entity by ID.
func (gsuo *GoodsSkuUpdateOne) SetGoodsSpuID(id int) *GoodsSkuUpdateOne {
	gsuo.mutation.SetGoodsSpuID(id)
	return gsuo
}

// SetNillableGoodsSpuID sets the "goods_spu" edge to the GoodsSpu entity by ID if the given value is not nil.
func (gsuo *GoodsSkuUpdateOne) SetNillableGoodsSpuID(id *int) *GoodsSkuUpdateOne {
	if id != nil {
		gsuo = gsuo.SetGoodsSpuID(*id)
	}
	return gsuo
}

// SetGoodsSpu sets the "goods_spu" edge to the GoodsSpu entity.
func (gsuo *GoodsSkuUpdateOne) SetGoodsSpu(g *GoodsSpu) *GoodsSkuUpdateOne {
	return gsuo.SetGoodsSpuID(g.ID)
}

// AddGoodsSpecsOptionIDs adds the "goods_specs_option" edge to the GoodsSpecsOption entity by IDs.
func (gsuo *GoodsSkuUpdateOne) AddGoodsSpecsOptionIDs(ids ...int) *GoodsSkuUpdateOne {
	gsuo.mutation.AddGoodsSpecsOptionIDs(ids...)
	return gsuo
}

// AddGoodsSpecsOption adds the "goods_specs_option" edges to the GoodsSpecsOption entity.
func (gsuo *GoodsSkuUpdateOne) AddGoodsSpecsOption(g ...*GoodsSpecsOption) *GoodsSkuUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gsuo.AddGoodsSpecsOptionIDs(ids...)
}

// Mutation returns the GoodsSkuMutation object of the builder.
func (gsuo *GoodsSkuUpdateOne) Mutation() *GoodsSkuMutation {
	return gsuo.mutation
}

// ClearGoodsSpu clears the "goods_spu" edge to the GoodsSpu entity.
func (gsuo *GoodsSkuUpdateOne) ClearGoodsSpu() *GoodsSkuUpdateOne {
	gsuo.mutation.ClearGoodsSpu()
	return gsuo
}

// ClearGoodsSpecsOption clears all "goods_specs_option" edges to the GoodsSpecsOption entity.
func (gsuo *GoodsSkuUpdateOne) ClearGoodsSpecsOption() *GoodsSkuUpdateOne {
	gsuo.mutation.ClearGoodsSpecsOption()
	return gsuo
}

// RemoveGoodsSpecsOptionIDs removes the "goods_specs_option" edge to GoodsSpecsOption entities by IDs.
func (gsuo *GoodsSkuUpdateOne) RemoveGoodsSpecsOptionIDs(ids ...int) *GoodsSkuUpdateOne {
	gsuo.mutation.RemoveGoodsSpecsOptionIDs(ids...)
	return gsuo
}

// RemoveGoodsSpecsOption removes "goods_specs_option" edges to GoodsSpecsOption entities.
func (gsuo *GoodsSkuUpdateOne) RemoveGoodsSpecsOption(g ...*GoodsSpecsOption) *GoodsSkuUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gsuo.RemoveGoodsSpecsOptionIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gsuo *GoodsSkuUpdateOne) Select(field string, fields ...string) *GoodsSkuUpdateOne {
	gsuo.fields = append([]string{field}, fields...)
	return gsuo
}

// Save executes the query and returns the updated GoodsSku entity.
func (gsuo *GoodsSkuUpdateOne) Save(ctx context.Context) (*GoodsSku, error) {
	var (
		err  error
		node *GoodsSku
	)
	gsuo.defaults()
	if len(gsuo.hooks) == 0 {
		node, err = gsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodsSkuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gsuo.mutation = mutation
			node, err = gsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gsuo.hooks) - 1; i >= 0; i-- {
			mut = gsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gsuo *GoodsSkuUpdateOne) SaveX(ctx context.Context) *GoodsSku {
	node, err := gsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gsuo *GoodsSkuUpdateOne) Exec(ctx context.Context) error {
	_, err := gsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gsuo *GoodsSkuUpdateOne) ExecX(ctx context.Context) {
	if err := gsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gsuo *GoodsSkuUpdateOne) defaults() {
	if _, ok := gsuo.mutation.UpdatedAt(); !ok && !gsuo.mutation.UpdatedAtCleared() {
		v := goodssku.UpdateDefaultUpdatedAt()
		gsuo.mutation.SetUpdatedAt(v)
	}
}

func (gsuo *GoodsSkuUpdateOne) sqlSave(ctx context.Context) (_node *GoodsSku, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodssku.Table,
			Columns: goodssku.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: goodssku.FieldID,
			},
		},
	}
	id, ok := gsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing GoodsSku.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := gsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodssku.FieldID)
		for _, f := range fields {
			if !goodssku.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != goodssku.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gsuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goodssku.FieldUpdatedAt,
		})
	}
	if gsuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: goodssku.FieldUpdatedAt,
		})
	}
	if value, ok := gsuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goodssku.FieldDeletedAt,
		})
	}
	if gsuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: goodssku.FieldDeletedAt,
		})
	}
	if value, ok := gsuo.mutation.SkuName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodssku.FieldSkuName,
		})
	}
	if value, ok := gsuo.mutation.SkuCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodssku.FieldSkuCode,
		})
	}
	if value, ok := gsuo.mutation.StockNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodssku.FieldStockNum,
		})
	}
	if value, ok := gsuo.mutation.AddedStockNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodssku.FieldStockNum,
		})
	}
	if value, ok := gsuo.mutation.SalesNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodssku.FieldSalesNum,
		})
	}
	if value, ok := gsuo.mutation.AddedSalesNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodssku.FieldSalesNum,
		})
	}
	if value, ok := gsuo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodssku.FieldPrice,
		})
	}
	if value, ok := gsuo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodssku.FieldPrice,
		})
	}
	if gsuo.mutation.GoodsSpuCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsuo.mutation.GoodsSpuIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gsuo.mutation.GoodsSpecsOptionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsuo.mutation.RemovedGoodsSpecsOptionIDs(); len(nodes) > 0 && !gsuo.mutation.GoodsSpecsOptionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsuo.mutation.GoodsSpecsOptionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GoodsSku{config: gsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodssku.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
