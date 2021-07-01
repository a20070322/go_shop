// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/shop-go/ent/goodsclassify"
	"github.com/a20070322/shop-go/ent/predicate"
)

// GoodsClassifyDelete is the builder for deleting a GoodsClassify entity.
type GoodsClassifyDelete struct {
	config
	hooks    []Hook
	mutation *GoodsClassifyMutation
}

// Where adds a new predicate to the GoodsClassifyDelete builder.
func (gcd *GoodsClassifyDelete) Where(ps ...predicate.GoodsClassify) *GoodsClassifyDelete {
	gcd.mutation.predicates = append(gcd.mutation.predicates, ps...)
	return gcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gcd *GoodsClassifyDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gcd.hooks) == 0 {
		affected, err = gcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodsClassifyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gcd.mutation = mutation
			affected, err = gcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gcd.hooks) - 1; i >= 0; i-- {
			mut = gcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcd *GoodsClassifyDelete) ExecX(ctx context.Context) int {
	n, err := gcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gcd *GoodsClassifyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: goodsclassify.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: goodsclassify.FieldID,
			},
		},
	}
	if ps := gcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, gcd.driver, _spec)
}

// GoodsClassifyDeleteOne is the builder for deleting a single GoodsClassify entity.
type GoodsClassifyDeleteOne struct {
	gcd *GoodsClassifyDelete
}

// Exec executes the deletion query.
func (gcdo *GoodsClassifyDeleteOne) Exec(ctx context.Context) error {
	n, err := gcdo.gcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{goodsclassify.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gcdo *GoodsClassifyDeleteOne) ExecX(ctx context.Context) {
	gcdo.gcd.ExecX(ctx)
}