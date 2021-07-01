// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/shop-go/ent/basiclink"
	"github.com/a20070322/shop-go/ent/predicate"
)

// BasicLinkDelete is the builder for deleting a BasicLink entity.
type BasicLinkDelete struct {
	config
	hooks    []Hook
	mutation *BasicLinkMutation
}

// Where adds a new predicate to the BasicLinkDelete builder.
func (bld *BasicLinkDelete) Where(ps ...predicate.BasicLink) *BasicLinkDelete {
	bld.mutation.predicates = append(bld.mutation.predicates, ps...)
	return bld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bld *BasicLinkDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bld.hooks) == 0 {
		affected, err = bld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BasicLinkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bld.mutation = mutation
			affected, err = bld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bld.hooks) - 1; i >= 0; i-- {
			mut = bld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bld *BasicLinkDelete) ExecX(ctx context.Context) int {
	n, err := bld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bld *BasicLinkDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: basiclink.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: basiclink.FieldID,
			},
		},
	}
	if ps := bld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bld.driver, _spec)
}

// BasicLinkDeleteOne is the builder for deleting a single BasicLink entity.
type BasicLinkDeleteOne struct {
	bld *BasicLinkDelete
}

// Exec executes the deletion query.
func (bldo *BasicLinkDeleteOne) Exec(ctx context.Context) error {
	n, err := bldo.bld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{basiclink.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bldo *BasicLinkDeleteOne) ExecX(ctx context.Context) {
	bldo.bld.ExecX(ctx)
}