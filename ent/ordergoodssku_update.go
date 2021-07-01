// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/shop-go/ent/goodsspu"
	"github.com/a20070322/shop-go/ent/ordergoodssku"
	"github.com/a20070322/shop-go/ent/orderinfo"
	"github.com/a20070322/shop-go/ent/predicate"
	"github.com/a20070322/shop-go/ent/schema"
)

// OrderGoodsSkuUpdate is the builder for updating OrderGoodsSku entities.
type OrderGoodsSkuUpdate struct {
	config
	hooks    []Hook
	mutation *OrderGoodsSkuMutation
}

// Where adds a new predicate for the OrderGoodsSkuUpdate builder.
func (ogsu *OrderGoodsSkuUpdate) Where(ps ...predicate.OrderGoodsSku) *OrderGoodsSkuUpdate {
	ogsu.mutation.predicates = append(ogsu.mutation.predicates, ps...)
	return ogsu
}

// SetUpdatedAt sets the "updated_at" field.
func (ogsu *OrderGoodsSkuUpdate) SetUpdatedAt(t time.Time) *OrderGoodsSkuUpdate {
	ogsu.mutation.SetUpdatedAt(t)
	return ogsu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ogsu *OrderGoodsSkuUpdate) ClearUpdatedAt() *OrderGoodsSkuUpdate {
	ogsu.mutation.ClearUpdatedAt()
	return ogsu
}

// SetDeletedAt sets the "deleted_at" field.
func (ogsu *OrderGoodsSkuUpdate) SetDeletedAt(t time.Time) *OrderGoodsSkuUpdate {
	ogsu.mutation.SetDeletedAt(t)
	return ogsu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ogsu *OrderGoodsSkuUpdate) SetNillableDeletedAt(t *time.Time) *OrderGoodsSkuUpdate {
	if t != nil {
		ogsu.SetDeletedAt(*t)
	}
	return ogsu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ogsu *OrderGoodsSkuUpdate) ClearDeletedAt() *OrderGoodsSkuUpdate {
	ogsu.mutation.ClearDeletedAt()
	return ogsu
}

// SetSpuName sets the "spu_name" field.
func (ogsu *OrderGoodsSkuUpdate) SetSpuName(s string) *OrderGoodsSkuUpdate {
	ogsu.mutation.SetSpuName(s)
	return ogsu
}

// SetSpuCode sets the "spu_code" field.
func (ogsu *OrderGoodsSkuUpdate) SetSpuCode(s string) *OrderGoodsSkuUpdate {
	ogsu.mutation.SetSpuCode(s)
	return ogsu
}

// SetSpuHeadImg sets the "spu_head_img" field.
func (ogsu *OrderGoodsSkuUpdate) SetSpuHeadImg(s string) *OrderGoodsSkuUpdate {
	ogsu.mutation.SetSpuHeadImg(s)
	return ogsu
}

// SetNillableSpuHeadImg sets the "spu_head_img" field if the given value is not nil.
func (ogsu *OrderGoodsSkuUpdate) SetNillableSpuHeadImg(s *string) *OrderGoodsSkuUpdate {
	if s != nil {
		ogsu.SetSpuHeadImg(*s)
	}
	return ogsu
}

// ClearSpuHeadImg clears the value of the "spu_head_img" field.
func (ogsu *OrderGoodsSkuUpdate) ClearSpuHeadImg() *OrderGoodsSkuUpdate {
	ogsu.mutation.ClearSpuHeadImg()
	return ogsu
}

// SetSpuDesc sets the "spu_desc" field.
func (ogsu *OrderGoodsSkuUpdate) SetSpuDesc(s string) *OrderGoodsSkuUpdate {
	ogsu.mutation.SetSpuDesc(s)
	return ogsu
}

// SetNillableSpuDesc sets the "spu_desc" field if the given value is not nil.
func (ogsu *OrderGoodsSkuUpdate) SetNillableSpuDesc(s *string) *OrderGoodsSkuUpdate {
	if s != nil {
		ogsu.SetSpuDesc(*s)
	}
	return ogsu
}

// ClearSpuDesc clears the value of the "spu_desc" field.
func (ogsu *OrderGoodsSkuUpdate) ClearSpuDesc() *OrderGoodsSkuUpdate {
	ogsu.mutation.ClearSpuDesc()
	return ogsu
}

// SetSpuDetails sets the "spu_details" field.
func (ogsu *OrderGoodsSkuUpdate) SetSpuDetails(s string) *OrderGoodsSkuUpdate {
	ogsu.mutation.SetSpuDetails(s)
	return ogsu
}

// SetNillableSpuDetails sets the "spu_details" field if the given value is not nil.
func (ogsu *OrderGoodsSkuUpdate) SetNillableSpuDetails(s *string) *OrderGoodsSkuUpdate {
	if s != nil {
		ogsu.SetSpuDetails(*s)
	}
	return ogsu
}

// ClearSpuDetails clears the value of the "spu_details" field.
func (ogsu *OrderGoodsSkuUpdate) ClearSpuDetails() *OrderGoodsSkuUpdate {
	ogsu.mutation.ClearSpuDetails()
	return ogsu
}

// SetIsCustomSku sets the "is_custom_sku" field.
func (ogsu *OrderGoodsSkuUpdate) SetIsCustomSku(b bool) *OrderGoodsSkuUpdate {
	ogsu.mutation.SetIsCustomSku(b)
	return ogsu
}

// SetSpecsOption sets the "specs_option" field.
func (ogsu *OrderGoodsSkuUpdate) SetSpecsOption(sot []*schema.SpecsOptionType) *OrderGoodsSkuUpdate {
	ogsu.mutation.SetSpecsOption(sot)
	return ogsu
}

// SetSkuID sets the "sku_id" field.
func (ogsu *OrderGoodsSkuUpdate) SetSkuID(i int) *OrderGoodsSkuUpdate {
	ogsu.mutation.ResetSkuID()
	ogsu.mutation.SetSkuID(i)
	return ogsu
}

// AddSkuID adds i to the "sku_id" field.
func (ogsu *OrderGoodsSkuUpdate) AddSkuID(i int) *OrderGoodsSkuUpdate {
	ogsu.mutation.AddSkuID(i)
	return ogsu
}

// SetSkuName sets the "sku_name" field.
func (ogsu *OrderGoodsSkuUpdate) SetSkuName(s string) *OrderGoodsSkuUpdate {
	ogsu.mutation.SetSkuName(s)
	return ogsu
}

// SetSkuCode sets the "sku_code" field.
func (ogsu *OrderGoodsSkuUpdate) SetSkuCode(s string) *OrderGoodsSkuUpdate {
	ogsu.mutation.SetSkuCode(s)
	return ogsu
}

// SetPrice sets the "price" field.
func (ogsu *OrderGoodsSkuUpdate) SetPrice(i int) *OrderGoodsSkuUpdate {
	ogsu.mutation.ResetPrice()
	ogsu.mutation.SetPrice(i)
	return ogsu
}

// AddPrice adds i to the "price" field.
func (ogsu *OrderGoodsSkuUpdate) AddPrice(i int) *OrderGoodsSkuUpdate {
	ogsu.mutation.AddPrice(i)
	return ogsu
}

// SetAmount sets the "amount" field.
func (ogsu *OrderGoodsSkuUpdate) SetAmount(i int) *OrderGoodsSkuUpdate {
	ogsu.mutation.ResetAmount()
	ogsu.mutation.SetAmount(i)
	return ogsu
}

// AddAmount adds i to the "amount" field.
func (ogsu *OrderGoodsSkuUpdate) AddAmount(i int) *OrderGoodsSkuUpdate {
	ogsu.mutation.AddAmount(i)
	return ogsu
}

// SetGoodsSpuID sets the "goods_spu" edge to the GoodsSpu entity by ID.
func (ogsu *OrderGoodsSkuUpdate) SetGoodsSpuID(id int) *OrderGoodsSkuUpdate {
	ogsu.mutation.SetGoodsSpuID(id)
	return ogsu
}

// SetNillableGoodsSpuID sets the "goods_spu" edge to the GoodsSpu entity by ID if the given value is not nil.
func (ogsu *OrderGoodsSkuUpdate) SetNillableGoodsSpuID(id *int) *OrderGoodsSkuUpdate {
	if id != nil {
		ogsu = ogsu.SetGoodsSpuID(*id)
	}
	return ogsu
}

// SetGoodsSpu sets the "goods_spu" edge to the GoodsSpu entity.
func (ogsu *OrderGoodsSkuUpdate) SetGoodsSpu(g *GoodsSpu) *OrderGoodsSkuUpdate {
	return ogsu.SetGoodsSpuID(g.ID)
}

// SetOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID.
func (ogsu *OrderGoodsSkuUpdate) SetOrderInfoID(id int) *OrderGoodsSkuUpdate {
	ogsu.mutation.SetOrderInfoID(id)
	return ogsu
}

// SetNillableOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID if the given value is not nil.
func (ogsu *OrderGoodsSkuUpdate) SetNillableOrderInfoID(id *int) *OrderGoodsSkuUpdate {
	if id != nil {
		ogsu = ogsu.SetOrderInfoID(*id)
	}
	return ogsu
}

// SetOrderInfo sets the "order_info" edge to the OrderInfo entity.
func (ogsu *OrderGoodsSkuUpdate) SetOrderInfo(o *OrderInfo) *OrderGoodsSkuUpdate {
	return ogsu.SetOrderInfoID(o.ID)
}

// Mutation returns the OrderGoodsSkuMutation object of the builder.
func (ogsu *OrderGoodsSkuUpdate) Mutation() *OrderGoodsSkuMutation {
	return ogsu.mutation
}

// ClearGoodsSpu clears the "goods_spu" edge to the GoodsSpu entity.
func (ogsu *OrderGoodsSkuUpdate) ClearGoodsSpu() *OrderGoodsSkuUpdate {
	ogsu.mutation.ClearGoodsSpu()
	return ogsu
}

// ClearOrderInfo clears the "order_info" edge to the OrderInfo entity.
func (ogsu *OrderGoodsSkuUpdate) ClearOrderInfo() *OrderGoodsSkuUpdate {
	ogsu.mutation.ClearOrderInfo()
	return ogsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ogsu *OrderGoodsSkuUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ogsu.defaults()
	if len(ogsu.hooks) == 0 {
		affected, err = ogsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderGoodsSkuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ogsu.mutation = mutation
			affected, err = ogsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ogsu.hooks) - 1; i >= 0; i-- {
			mut = ogsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ogsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ogsu *OrderGoodsSkuUpdate) SaveX(ctx context.Context) int {
	affected, err := ogsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ogsu *OrderGoodsSkuUpdate) Exec(ctx context.Context) error {
	_, err := ogsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ogsu *OrderGoodsSkuUpdate) ExecX(ctx context.Context) {
	if err := ogsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ogsu *OrderGoodsSkuUpdate) defaults() {
	if _, ok := ogsu.mutation.UpdatedAt(); !ok && !ogsu.mutation.UpdatedAtCleared() {
		v := ordergoodssku.UpdateDefaultUpdatedAt()
		ogsu.mutation.SetUpdatedAt(v)
	}
}

func (ogsu *OrderGoodsSkuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ordergoodssku.Table,
			Columns: ordergoodssku.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ordergoodssku.FieldID,
			},
		},
	}
	if ps := ogsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ogsu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ordergoodssku.FieldUpdatedAt,
		})
	}
	if ogsu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: ordergoodssku.FieldUpdatedAt,
		})
	}
	if value, ok := ogsu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ordergoodssku.FieldDeletedAt,
		})
	}
	if ogsu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: ordergoodssku.FieldDeletedAt,
		})
	}
	if value, ok := ogsu.mutation.SpuName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordergoodssku.FieldSpuName,
		})
	}
	if value, ok := ogsu.mutation.SpuCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordergoodssku.FieldSpuCode,
		})
	}
	if value, ok := ogsu.mutation.SpuHeadImg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordergoodssku.FieldSpuHeadImg,
		})
	}
	if ogsu.mutation.SpuHeadImgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: ordergoodssku.FieldSpuHeadImg,
		})
	}
	if value, ok := ogsu.mutation.SpuDesc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordergoodssku.FieldSpuDesc,
		})
	}
	if ogsu.mutation.SpuDescCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: ordergoodssku.FieldSpuDesc,
		})
	}
	if value, ok := ogsu.mutation.SpuDetails(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordergoodssku.FieldSpuDetails,
		})
	}
	if ogsu.mutation.SpuDetailsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: ordergoodssku.FieldSpuDetails,
		})
	}
	if value, ok := ogsu.mutation.IsCustomSku(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: ordergoodssku.FieldIsCustomSku,
		})
	}
	if value, ok := ogsu.mutation.SpecsOption(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: ordergoodssku.FieldSpecsOption,
		})
	}
	if value, ok := ogsu.mutation.SkuID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordergoodssku.FieldSkuID,
		})
	}
	if value, ok := ogsu.mutation.AddedSkuID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordergoodssku.FieldSkuID,
		})
	}
	if value, ok := ogsu.mutation.SkuName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordergoodssku.FieldSkuName,
		})
	}
	if value, ok := ogsu.mutation.SkuCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordergoodssku.FieldSkuCode,
		})
	}
	if value, ok := ogsu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordergoodssku.FieldPrice,
		})
	}
	if value, ok := ogsu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordergoodssku.FieldPrice,
		})
	}
	if value, ok := ogsu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordergoodssku.FieldAmount,
		})
	}
	if value, ok := ogsu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordergoodssku.FieldAmount,
		})
	}
	if ogsu.mutation.GoodsSpuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordergoodssku.GoodsSpuTable,
			Columns: []string{ordergoodssku.GoodsSpuColumn},
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
	if nodes := ogsu.mutation.GoodsSpuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordergoodssku.GoodsSpuTable,
			Columns: []string{ordergoodssku.GoodsSpuColumn},
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
	if ogsu.mutation.OrderInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordergoodssku.OrderInfoTable,
			Columns: []string{ordergoodssku.OrderInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ogsu.mutation.OrderInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordergoodssku.OrderInfoTable,
			Columns: []string{ordergoodssku.OrderInfoColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ogsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ordergoodssku.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// OrderGoodsSkuUpdateOne is the builder for updating a single OrderGoodsSku entity.
type OrderGoodsSkuUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderGoodsSkuMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ogsuo *OrderGoodsSkuUpdateOne) SetUpdatedAt(t time.Time) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.SetUpdatedAt(t)
	return ogsuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ogsuo *OrderGoodsSkuUpdateOne) ClearUpdatedAt() *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.ClearUpdatedAt()
	return ogsuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ogsuo *OrderGoodsSkuUpdateOne) SetDeletedAt(t time.Time) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.SetDeletedAt(t)
	return ogsuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ogsuo *OrderGoodsSkuUpdateOne) SetNillableDeletedAt(t *time.Time) *OrderGoodsSkuUpdateOne {
	if t != nil {
		ogsuo.SetDeletedAt(*t)
	}
	return ogsuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ogsuo *OrderGoodsSkuUpdateOne) ClearDeletedAt() *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.ClearDeletedAt()
	return ogsuo
}

// SetSpuName sets the "spu_name" field.
func (ogsuo *OrderGoodsSkuUpdateOne) SetSpuName(s string) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.SetSpuName(s)
	return ogsuo
}

// SetSpuCode sets the "spu_code" field.
func (ogsuo *OrderGoodsSkuUpdateOne) SetSpuCode(s string) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.SetSpuCode(s)
	return ogsuo
}

// SetSpuHeadImg sets the "spu_head_img" field.
func (ogsuo *OrderGoodsSkuUpdateOne) SetSpuHeadImg(s string) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.SetSpuHeadImg(s)
	return ogsuo
}

// SetNillableSpuHeadImg sets the "spu_head_img" field if the given value is not nil.
func (ogsuo *OrderGoodsSkuUpdateOne) SetNillableSpuHeadImg(s *string) *OrderGoodsSkuUpdateOne {
	if s != nil {
		ogsuo.SetSpuHeadImg(*s)
	}
	return ogsuo
}

// ClearSpuHeadImg clears the value of the "spu_head_img" field.
func (ogsuo *OrderGoodsSkuUpdateOne) ClearSpuHeadImg() *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.ClearSpuHeadImg()
	return ogsuo
}

// SetSpuDesc sets the "spu_desc" field.
func (ogsuo *OrderGoodsSkuUpdateOne) SetSpuDesc(s string) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.SetSpuDesc(s)
	return ogsuo
}

// SetNillableSpuDesc sets the "spu_desc" field if the given value is not nil.
func (ogsuo *OrderGoodsSkuUpdateOne) SetNillableSpuDesc(s *string) *OrderGoodsSkuUpdateOne {
	if s != nil {
		ogsuo.SetSpuDesc(*s)
	}
	return ogsuo
}

// ClearSpuDesc clears the value of the "spu_desc" field.
func (ogsuo *OrderGoodsSkuUpdateOne) ClearSpuDesc() *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.ClearSpuDesc()
	return ogsuo
}

// SetSpuDetails sets the "spu_details" field.
func (ogsuo *OrderGoodsSkuUpdateOne) SetSpuDetails(s string) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.SetSpuDetails(s)
	return ogsuo
}

// SetNillableSpuDetails sets the "spu_details" field if the given value is not nil.
func (ogsuo *OrderGoodsSkuUpdateOne) SetNillableSpuDetails(s *string) *OrderGoodsSkuUpdateOne {
	if s != nil {
		ogsuo.SetSpuDetails(*s)
	}
	return ogsuo
}

// ClearSpuDetails clears the value of the "spu_details" field.
func (ogsuo *OrderGoodsSkuUpdateOne) ClearSpuDetails() *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.ClearSpuDetails()
	return ogsuo
}

// SetIsCustomSku sets the "is_custom_sku" field.
func (ogsuo *OrderGoodsSkuUpdateOne) SetIsCustomSku(b bool) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.SetIsCustomSku(b)
	return ogsuo
}

// SetSpecsOption sets the "specs_option" field.
func (ogsuo *OrderGoodsSkuUpdateOne) SetSpecsOption(sot []*schema.SpecsOptionType) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.SetSpecsOption(sot)
	return ogsuo
}

// SetSkuID sets the "sku_id" field.
func (ogsuo *OrderGoodsSkuUpdateOne) SetSkuID(i int) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.ResetSkuID()
	ogsuo.mutation.SetSkuID(i)
	return ogsuo
}

// AddSkuID adds i to the "sku_id" field.
func (ogsuo *OrderGoodsSkuUpdateOne) AddSkuID(i int) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.AddSkuID(i)
	return ogsuo
}

// SetSkuName sets the "sku_name" field.
func (ogsuo *OrderGoodsSkuUpdateOne) SetSkuName(s string) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.SetSkuName(s)
	return ogsuo
}

// SetSkuCode sets the "sku_code" field.
func (ogsuo *OrderGoodsSkuUpdateOne) SetSkuCode(s string) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.SetSkuCode(s)
	return ogsuo
}

// SetPrice sets the "price" field.
func (ogsuo *OrderGoodsSkuUpdateOne) SetPrice(i int) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.ResetPrice()
	ogsuo.mutation.SetPrice(i)
	return ogsuo
}

// AddPrice adds i to the "price" field.
func (ogsuo *OrderGoodsSkuUpdateOne) AddPrice(i int) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.AddPrice(i)
	return ogsuo
}

// SetAmount sets the "amount" field.
func (ogsuo *OrderGoodsSkuUpdateOne) SetAmount(i int) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.ResetAmount()
	ogsuo.mutation.SetAmount(i)
	return ogsuo
}

// AddAmount adds i to the "amount" field.
func (ogsuo *OrderGoodsSkuUpdateOne) AddAmount(i int) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.AddAmount(i)
	return ogsuo
}

// SetGoodsSpuID sets the "goods_spu" edge to the GoodsSpu entity by ID.
func (ogsuo *OrderGoodsSkuUpdateOne) SetGoodsSpuID(id int) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.SetGoodsSpuID(id)
	return ogsuo
}

// SetNillableGoodsSpuID sets the "goods_spu" edge to the GoodsSpu entity by ID if the given value is not nil.
func (ogsuo *OrderGoodsSkuUpdateOne) SetNillableGoodsSpuID(id *int) *OrderGoodsSkuUpdateOne {
	if id != nil {
		ogsuo = ogsuo.SetGoodsSpuID(*id)
	}
	return ogsuo
}

// SetGoodsSpu sets the "goods_spu" edge to the GoodsSpu entity.
func (ogsuo *OrderGoodsSkuUpdateOne) SetGoodsSpu(g *GoodsSpu) *OrderGoodsSkuUpdateOne {
	return ogsuo.SetGoodsSpuID(g.ID)
}

// SetOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID.
func (ogsuo *OrderGoodsSkuUpdateOne) SetOrderInfoID(id int) *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.SetOrderInfoID(id)
	return ogsuo
}

// SetNillableOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID if the given value is not nil.
func (ogsuo *OrderGoodsSkuUpdateOne) SetNillableOrderInfoID(id *int) *OrderGoodsSkuUpdateOne {
	if id != nil {
		ogsuo = ogsuo.SetOrderInfoID(*id)
	}
	return ogsuo
}

// SetOrderInfo sets the "order_info" edge to the OrderInfo entity.
func (ogsuo *OrderGoodsSkuUpdateOne) SetOrderInfo(o *OrderInfo) *OrderGoodsSkuUpdateOne {
	return ogsuo.SetOrderInfoID(o.ID)
}

// Mutation returns the OrderGoodsSkuMutation object of the builder.
func (ogsuo *OrderGoodsSkuUpdateOne) Mutation() *OrderGoodsSkuMutation {
	return ogsuo.mutation
}

// ClearGoodsSpu clears the "goods_spu" edge to the GoodsSpu entity.
func (ogsuo *OrderGoodsSkuUpdateOne) ClearGoodsSpu() *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.ClearGoodsSpu()
	return ogsuo
}

// ClearOrderInfo clears the "order_info" edge to the OrderInfo entity.
func (ogsuo *OrderGoodsSkuUpdateOne) ClearOrderInfo() *OrderGoodsSkuUpdateOne {
	ogsuo.mutation.ClearOrderInfo()
	return ogsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ogsuo *OrderGoodsSkuUpdateOne) Select(field string, fields ...string) *OrderGoodsSkuUpdateOne {
	ogsuo.fields = append([]string{field}, fields...)
	return ogsuo
}

// Save executes the query and returns the updated OrderGoodsSku entity.
func (ogsuo *OrderGoodsSkuUpdateOne) Save(ctx context.Context) (*OrderGoodsSku, error) {
	var (
		err  error
		node *OrderGoodsSku
	)
	ogsuo.defaults()
	if len(ogsuo.hooks) == 0 {
		node, err = ogsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderGoodsSkuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ogsuo.mutation = mutation
			node, err = ogsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ogsuo.hooks) - 1; i >= 0; i-- {
			mut = ogsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ogsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ogsuo *OrderGoodsSkuUpdateOne) SaveX(ctx context.Context) *OrderGoodsSku {
	node, err := ogsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ogsuo *OrderGoodsSkuUpdateOne) Exec(ctx context.Context) error {
	_, err := ogsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ogsuo *OrderGoodsSkuUpdateOne) ExecX(ctx context.Context) {
	if err := ogsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ogsuo *OrderGoodsSkuUpdateOne) defaults() {
	if _, ok := ogsuo.mutation.UpdatedAt(); !ok && !ogsuo.mutation.UpdatedAtCleared() {
		v := ordergoodssku.UpdateDefaultUpdatedAt()
		ogsuo.mutation.SetUpdatedAt(v)
	}
}

func (ogsuo *OrderGoodsSkuUpdateOne) sqlSave(ctx context.Context) (_node *OrderGoodsSku, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ordergoodssku.Table,
			Columns: ordergoodssku.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ordergoodssku.FieldID,
			},
		},
	}
	id, ok := ogsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing OrderGoodsSku.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ogsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ordergoodssku.FieldID)
		for _, f := range fields {
			if !ordergoodssku.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ordergoodssku.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ogsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ogsuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ordergoodssku.FieldUpdatedAt,
		})
	}
	if ogsuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: ordergoodssku.FieldUpdatedAt,
		})
	}
	if value, ok := ogsuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ordergoodssku.FieldDeletedAt,
		})
	}
	if ogsuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: ordergoodssku.FieldDeletedAt,
		})
	}
	if value, ok := ogsuo.mutation.SpuName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordergoodssku.FieldSpuName,
		})
	}
	if value, ok := ogsuo.mutation.SpuCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordergoodssku.FieldSpuCode,
		})
	}
	if value, ok := ogsuo.mutation.SpuHeadImg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordergoodssku.FieldSpuHeadImg,
		})
	}
	if ogsuo.mutation.SpuHeadImgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: ordergoodssku.FieldSpuHeadImg,
		})
	}
	if value, ok := ogsuo.mutation.SpuDesc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordergoodssku.FieldSpuDesc,
		})
	}
	if ogsuo.mutation.SpuDescCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: ordergoodssku.FieldSpuDesc,
		})
	}
	if value, ok := ogsuo.mutation.SpuDetails(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordergoodssku.FieldSpuDetails,
		})
	}
	if ogsuo.mutation.SpuDetailsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: ordergoodssku.FieldSpuDetails,
		})
	}
	if value, ok := ogsuo.mutation.IsCustomSku(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: ordergoodssku.FieldIsCustomSku,
		})
	}
	if value, ok := ogsuo.mutation.SpecsOption(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: ordergoodssku.FieldSpecsOption,
		})
	}
	if value, ok := ogsuo.mutation.SkuID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordergoodssku.FieldSkuID,
		})
	}
	if value, ok := ogsuo.mutation.AddedSkuID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordergoodssku.FieldSkuID,
		})
	}
	if value, ok := ogsuo.mutation.SkuName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordergoodssku.FieldSkuName,
		})
	}
	if value, ok := ogsuo.mutation.SkuCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ordergoodssku.FieldSkuCode,
		})
	}
	if value, ok := ogsuo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordergoodssku.FieldPrice,
		})
	}
	if value, ok := ogsuo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordergoodssku.FieldPrice,
		})
	}
	if value, ok := ogsuo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordergoodssku.FieldAmount,
		})
	}
	if value, ok := ogsuo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: ordergoodssku.FieldAmount,
		})
	}
	if ogsuo.mutation.GoodsSpuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordergoodssku.GoodsSpuTable,
			Columns: []string{ordergoodssku.GoodsSpuColumn},
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
	if nodes := ogsuo.mutation.GoodsSpuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordergoodssku.GoodsSpuTable,
			Columns: []string{ordergoodssku.GoodsSpuColumn},
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
	if ogsuo.mutation.OrderInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordergoodssku.OrderInfoTable,
			Columns: []string{ordergoodssku.OrderInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ogsuo.mutation.OrderInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordergoodssku.OrderInfoTable,
			Columns: []string{ordergoodssku.OrderInfoColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrderGoodsSku{config: ogsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ogsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ordergoodssku.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
