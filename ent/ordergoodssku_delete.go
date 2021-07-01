// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/shop-go/ent/ordergoodssku"
	"github.com/a20070322/shop-go/ent/predicate"
)

// OrderGoodsSkuDelete is the builder for deleting a OrderGoodsSku entity.
type OrderGoodsSkuDelete struct {
	config
	hooks    []Hook
	mutation *OrderGoodsSkuMutation
}

// Where adds a new predicate to the OrderGoodsSkuDelete builder.
func (ogsd *OrderGoodsSkuDelete) Where(ps ...predicate.OrderGoodsSku) *OrderGoodsSkuDelete {
	ogsd.mutation.predicates = append(ogsd.mutation.predicates, ps...)
	return ogsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ogsd *OrderGoodsSkuDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ogsd.hooks) == 0 {
		affected, err = ogsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderGoodsSkuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ogsd.mutation = mutation
			affected, err = ogsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ogsd.hooks) - 1; i >= 0; i-- {
			mut = ogsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ogsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ogsd *OrderGoodsSkuDelete) ExecX(ctx context.Context) int {
	n, err := ogsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ogsd *OrderGoodsSkuDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: ordergoodssku.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ordergoodssku.FieldID,
			},
		},
	}
	if ps := ogsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ogsd.driver, _spec)
}

// OrderGoodsSkuDeleteOne is the builder for deleting a single OrderGoodsSku entity.
type OrderGoodsSkuDeleteOne struct {
	ogsd *OrderGoodsSkuDelete
}

// Exec executes the deletion query.
func (ogsdo *OrderGoodsSkuDeleteOne) Exec(ctx context.Context) error {
	n, err := ogsdo.ogsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ordergoodssku.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ogsdo *OrderGoodsSkuDeleteOne) ExecX(ctx context.Context) {
	ogsdo.ogsd.ExecX(ctx)
}