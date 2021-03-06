// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/shop-go/ent/customeraddress"
	"github.com/a20070322/shop-go/ent/predicate"
)

// CustomerAddressDelete is the builder for deleting a CustomerAddress entity.
type CustomerAddressDelete struct {
	config
	hooks    []Hook
	mutation *CustomerAddressMutation
}

// Where adds a new predicate to the CustomerAddressDelete builder.
func (cad *CustomerAddressDelete) Where(ps ...predicate.CustomerAddress) *CustomerAddressDelete {
	cad.mutation.predicates = append(cad.mutation.predicates, ps...)
	return cad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cad *CustomerAddressDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cad.hooks) == 0 {
		affected, err = cad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerAddressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cad.mutation = mutation
			affected, err = cad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cad.hooks) - 1; i >= 0; i-- {
			mut = cad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cad *CustomerAddressDelete) ExecX(ctx context.Context) int {
	n, err := cad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cad *CustomerAddressDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: customeraddress.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customeraddress.FieldID,
			},
		},
	}
	if ps := cad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cad.driver, _spec)
}

// CustomerAddressDeleteOne is the builder for deleting a single CustomerAddress entity.
type CustomerAddressDeleteOne struct {
	cad *CustomerAddressDelete
}

// Exec executes the deletion query.
func (cado *CustomerAddressDeleteOne) Exec(ctx context.Context) error {
	n, err := cado.cad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{customeraddress.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cado *CustomerAddressDeleteOne) ExecX(ctx context.Context) {
	cado.cad.ExecX(ctx)
}
