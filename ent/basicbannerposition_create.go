// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/shop-go/ent/basicbanner"
	"github.com/a20070322/shop-go/ent/basicbannerposition"
)

// BasicBannerPositionCreate is the builder for creating a BasicBannerPosition entity.
type BasicBannerPositionCreate struct {
	config
	mutation *BasicBannerPositionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (bbpc *BasicBannerPositionCreate) SetCreatedAt(t time.Time) *BasicBannerPositionCreate {
	bbpc.mutation.SetCreatedAt(t)
	return bbpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bbpc *BasicBannerPositionCreate) SetNillableCreatedAt(t *time.Time) *BasicBannerPositionCreate {
	if t != nil {
		bbpc.SetCreatedAt(*t)
	}
	return bbpc
}

// SetUpdatedAt sets the "updated_at" field.
func (bbpc *BasicBannerPositionCreate) SetUpdatedAt(t time.Time) *BasicBannerPositionCreate {
	bbpc.mutation.SetUpdatedAt(t)
	return bbpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bbpc *BasicBannerPositionCreate) SetNillableUpdatedAt(t *time.Time) *BasicBannerPositionCreate {
	if t != nil {
		bbpc.SetUpdatedAt(*t)
	}
	return bbpc
}

// SetDeletedAt sets the "deleted_at" field.
func (bbpc *BasicBannerPositionCreate) SetDeletedAt(t time.Time) *BasicBannerPositionCreate {
	bbpc.mutation.SetDeletedAt(t)
	return bbpc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (bbpc *BasicBannerPositionCreate) SetNillableDeletedAt(t *time.Time) *BasicBannerPositionCreate {
	if t != nil {
		bbpc.SetDeletedAt(*t)
	}
	return bbpc
}

// SetPositionName sets the "position_name" field.
func (bbpc *BasicBannerPositionCreate) SetPositionName(s string) *BasicBannerPositionCreate {
	bbpc.mutation.SetPositionName(s)
	return bbpc
}

// SetRemarks sets the "remarks" field.
func (bbpc *BasicBannerPositionCreate) SetRemarks(s string) *BasicBannerPositionCreate {
	bbpc.mutation.SetRemarks(s)
	return bbpc
}

// SetNillableRemarks sets the "remarks" field if the given value is not nil.
func (bbpc *BasicBannerPositionCreate) SetNillableRemarks(s *string) *BasicBannerPositionCreate {
	if s != nil {
		bbpc.SetRemarks(*s)
	}
	return bbpc
}

// SetStatus sets the "status" field.
func (bbpc *BasicBannerPositionCreate) SetStatus(b bool) *BasicBannerPositionCreate {
	bbpc.mutation.SetStatus(b)
	return bbpc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (bbpc *BasicBannerPositionCreate) SetNillableStatus(b *bool) *BasicBannerPositionCreate {
	if b != nil {
		bbpc.SetStatus(*b)
	}
	return bbpc
}

// AddBasicBannerIDs adds the "basic_banner" edge to the BasicBanner entity by IDs.
func (bbpc *BasicBannerPositionCreate) AddBasicBannerIDs(ids ...int) *BasicBannerPositionCreate {
	bbpc.mutation.AddBasicBannerIDs(ids...)
	return bbpc
}

// AddBasicBanner adds the "basic_banner" edges to the BasicBanner entity.
func (bbpc *BasicBannerPositionCreate) AddBasicBanner(b ...*BasicBanner) *BasicBannerPositionCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bbpc.AddBasicBannerIDs(ids...)
}

// Mutation returns the BasicBannerPositionMutation object of the builder.
func (bbpc *BasicBannerPositionCreate) Mutation() *BasicBannerPositionMutation {
	return bbpc.mutation
}

// Save creates the BasicBannerPosition in the database.
func (bbpc *BasicBannerPositionCreate) Save(ctx context.Context) (*BasicBannerPosition, error) {
	var (
		err  error
		node *BasicBannerPosition
	)
	bbpc.defaults()
	if len(bbpc.hooks) == 0 {
		if err = bbpc.check(); err != nil {
			return nil, err
		}
		node, err = bbpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BasicBannerPositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bbpc.check(); err != nil {
				return nil, err
			}
			bbpc.mutation = mutation
			node, err = bbpc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bbpc.hooks) - 1; i >= 0; i-- {
			mut = bbpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bbpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bbpc *BasicBannerPositionCreate) SaveX(ctx context.Context) *BasicBannerPosition {
	v, err := bbpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (bbpc *BasicBannerPositionCreate) defaults() {
	if _, ok := bbpc.mutation.CreatedAt(); !ok {
		v := basicbannerposition.DefaultCreatedAt()
		bbpc.mutation.SetCreatedAt(v)
	}
	if _, ok := bbpc.mutation.Status(); !ok {
		v := basicbannerposition.DefaultStatus
		bbpc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bbpc *BasicBannerPositionCreate) check() error {
	if _, ok := bbpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := bbpc.mutation.PositionName(); !ok {
		return &ValidationError{Name: "position_name", err: errors.New("ent: missing required field \"position_name\"")}
	}
	if _, ok := bbpc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	return nil
}

func (bbpc *BasicBannerPositionCreate) sqlSave(ctx context.Context) (*BasicBannerPosition, error) {
	_node, _spec := bbpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bbpc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bbpc *BasicBannerPositionCreate) createSpec() (*BasicBannerPosition, *sqlgraph.CreateSpec) {
	var (
		_node = &BasicBannerPosition{config: bbpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: basicbannerposition.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: basicbannerposition.FieldID,
			},
		}
	)
	if value, ok := bbpc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: basicbannerposition.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := bbpc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: basicbannerposition.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := bbpc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: basicbannerposition.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := bbpc.mutation.PositionName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: basicbannerposition.FieldPositionName,
		})
		_node.PositionName = value
	}
	if value, ok := bbpc.mutation.Remarks(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: basicbannerposition.FieldRemarks,
		})
		_node.Remarks = value
	}
	if value, ok := bbpc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: basicbannerposition.FieldStatus,
		})
		_node.Status = value
	}
	if nodes := bbpc.mutation.BasicBannerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   basicbannerposition.BasicBannerTable,
			Columns: []string{basicbannerposition.BasicBannerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: basicbanner.FieldID,
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

// BasicBannerPositionCreateBulk is the builder for creating many BasicBannerPosition entities in bulk.
type BasicBannerPositionCreateBulk struct {
	config
	builders []*BasicBannerPositionCreate
}

// Save creates the BasicBannerPosition entities in the database.
func (bbpcb *BasicBannerPositionCreateBulk) Save(ctx context.Context) ([]*BasicBannerPosition, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bbpcb.builders))
	nodes := make([]*BasicBannerPosition, len(bbpcb.builders))
	mutators := make([]Mutator, len(bbpcb.builders))
	for i := range bbpcb.builders {
		func(i int, root context.Context) {
			builder := bbpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BasicBannerPositionMutation)
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
					_, err = mutators[i+1].Mutate(root, bbpcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bbpcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bbpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bbpcb *BasicBannerPositionCreateBulk) SaveX(ctx context.Context) []*BasicBannerPosition {
	v, err := bbpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
