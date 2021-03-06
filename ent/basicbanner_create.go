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
	"github.com/a20070322/shop-go/ent/basiclink"
)

// BasicBannerCreate is the builder for creating a BasicBanner entity.
type BasicBannerCreate struct {
	config
	mutation *BasicBannerMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (bbc *BasicBannerCreate) SetCreatedAt(t time.Time) *BasicBannerCreate {
	bbc.mutation.SetCreatedAt(t)
	return bbc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bbc *BasicBannerCreate) SetNillableCreatedAt(t *time.Time) *BasicBannerCreate {
	if t != nil {
		bbc.SetCreatedAt(*t)
	}
	return bbc
}

// SetUpdatedAt sets the "updated_at" field.
func (bbc *BasicBannerCreate) SetUpdatedAt(t time.Time) *BasicBannerCreate {
	bbc.mutation.SetUpdatedAt(t)
	return bbc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bbc *BasicBannerCreate) SetNillableUpdatedAt(t *time.Time) *BasicBannerCreate {
	if t != nil {
		bbc.SetUpdatedAt(*t)
	}
	return bbc
}

// SetDeletedAt sets the "deleted_at" field.
func (bbc *BasicBannerCreate) SetDeletedAt(t time.Time) *BasicBannerCreate {
	bbc.mutation.SetDeletedAt(t)
	return bbc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (bbc *BasicBannerCreate) SetNillableDeletedAt(t *time.Time) *BasicBannerCreate {
	if t != nil {
		bbc.SetDeletedAt(*t)
	}
	return bbc
}

// SetBannerName sets the "banner_name" field.
func (bbc *BasicBannerCreate) SetBannerName(s string) *BasicBannerCreate {
	bbc.mutation.SetBannerName(s)
	return bbc
}

// SetBannerImg sets the "banner_img" field.
func (bbc *BasicBannerCreate) SetBannerImg(s string) *BasicBannerCreate {
	bbc.mutation.SetBannerImg(s)
	return bbc
}

// SetOrder sets the "order" field.
func (bbc *BasicBannerCreate) SetOrder(i int) *BasicBannerCreate {
	bbc.mutation.SetOrder(i)
	return bbc
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (bbc *BasicBannerCreate) SetNillableOrder(i *int) *BasicBannerCreate {
	if i != nil {
		bbc.SetOrder(*i)
	}
	return bbc
}

// SetStatus sets the "status" field.
func (bbc *BasicBannerCreate) SetStatus(b bool) *BasicBannerCreate {
	bbc.mutation.SetStatus(b)
	return bbc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (bbc *BasicBannerCreate) SetNillableStatus(b *bool) *BasicBannerCreate {
	if b != nil {
		bbc.SetStatus(*b)
	}
	return bbc
}

// SetBasicBannerPositionID sets the "basic_banner_position" edge to the BasicBannerPosition entity by ID.
func (bbc *BasicBannerCreate) SetBasicBannerPositionID(id int) *BasicBannerCreate {
	bbc.mutation.SetBasicBannerPositionID(id)
	return bbc
}

// SetNillableBasicBannerPositionID sets the "basic_banner_position" edge to the BasicBannerPosition entity by ID if the given value is not nil.
func (bbc *BasicBannerCreate) SetNillableBasicBannerPositionID(id *int) *BasicBannerCreate {
	if id != nil {
		bbc = bbc.SetBasicBannerPositionID(*id)
	}
	return bbc
}

// SetBasicBannerPosition sets the "basic_banner_position" edge to the BasicBannerPosition entity.
func (bbc *BasicBannerCreate) SetBasicBannerPosition(b *BasicBannerPosition) *BasicBannerCreate {
	return bbc.SetBasicBannerPositionID(b.ID)
}

// SetBasicLinkID sets the "basic_link" edge to the BasicLink entity by ID.
func (bbc *BasicBannerCreate) SetBasicLinkID(id int) *BasicBannerCreate {
	bbc.mutation.SetBasicLinkID(id)
	return bbc
}

// SetNillableBasicLinkID sets the "basic_link" edge to the BasicLink entity by ID if the given value is not nil.
func (bbc *BasicBannerCreate) SetNillableBasicLinkID(id *int) *BasicBannerCreate {
	if id != nil {
		bbc = bbc.SetBasicLinkID(*id)
	}
	return bbc
}

// SetBasicLink sets the "basic_link" edge to the BasicLink entity.
func (bbc *BasicBannerCreate) SetBasicLink(b *BasicLink) *BasicBannerCreate {
	return bbc.SetBasicLinkID(b.ID)
}

// Mutation returns the BasicBannerMutation object of the builder.
func (bbc *BasicBannerCreate) Mutation() *BasicBannerMutation {
	return bbc.mutation
}

// Save creates the BasicBanner in the database.
func (bbc *BasicBannerCreate) Save(ctx context.Context) (*BasicBanner, error) {
	var (
		err  error
		node *BasicBanner
	)
	bbc.defaults()
	if len(bbc.hooks) == 0 {
		if err = bbc.check(); err != nil {
			return nil, err
		}
		node, err = bbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BasicBannerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bbc.check(); err != nil {
				return nil, err
			}
			bbc.mutation = mutation
			node, err = bbc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bbc.hooks) - 1; i >= 0; i-- {
			mut = bbc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bbc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bbc *BasicBannerCreate) SaveX(ctx context.Context) *BasicBanner {
	v, err := bbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (bbc *BasicBannerCreate) defaults() {
	if _, ok := bbc.mutation.CreatedAt(); !ok {
		v := basicbanner.DefaultCreatedAt()
		bbc.mutation.SetCreatedAt(v)
	}
	if _, ok := bbc.mutation.Order(); !ok {
		v := basicbanner.DefaultOrder
		bbc.mutation.SetOrder(v)
	}
	if _, ok := bbc.mutation.Status(); !ok {
		v := basicbanner.DefaultStatus
		bbc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bbc *BasicBannerCreate) check() error {
	if _, ok := bbc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := bbc.mutation.BannerName(); !ok {
		return &ValidationError{Name: "banner_name", err: errors.New("ent: missing required field \"banner_name\"")}
	}
	if _, ok := bbc.mutation.BannerImg(); !ok {
		return &ValidationError{Name: "banner_img", err: errors.New("ent: missing required field \"banner_img\"")}
	}
	if _, ok := bbc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	return nil
}

func (bbc *BasicBannerCreate) sqlSave(ctx context.Context) (*BasicBanner, error) {
	_node, _spec := bbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bbc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bbc *BasicBannerCreate) createSpec() (*BasicBanner, *sqlgraph.CreateSpec) {
	var (
		_node = &BasicBanner{config: bbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: basicbanner.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: basicbanner.FieldID,
			},
		}
	)
	if value, ok := bbc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: basicbanner.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := bbc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: basicbanner.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := bbc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: basicbanner.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := bbc.mutation.BannerName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: basicbanner.FieldBannerName,
		})
		_node.BannerName = value
	}
	if value, ok := bbc.mutation.BannerImg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: basicbanner.FieldBannerImg,
		})
		_node.BannerImg = value
	}
	if value, ok := bbc.mutation.Order(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: basicbanner.FieldOrder,
		})
		_node.Order = value
	}
	if value, ok := bbc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: basicbanner.FieldStatus,
		})
		_node.Status = value
	}
	if nodes := bbc.mutation.BasicBannerPositionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   basicbanner.BasicBannerPositionTable,
			Columns: []string{basicbanner.BasicBannerPositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: basicbannerposition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.basic_banner_position_basic_banner = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bbc.mutation.BasicLinkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   basicbanner.BasicLinkTable,
			Columns: []string{basicbanner.BasicLinkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: basiclink.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.basic_link_basic_banner = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BasicBannerCreateBulk is the builder for creating many BasicBanner entities in bulk.
type BasicBannerCreateBulk struct {
	config
	builders []*BasicBannerCreate
}

// Save creates the BasicBanner entities in the database.
func (bbcb *BasicBannerCreateBulk) Save(ctx context.Context) ([]*BasicBanner, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bbcb.builders))
	nodes := make([]*BasicBanner, len(bbcb.builders))
	mutators := make([]Mutator, len(bbcb.builders))
	for i := range bbcb.builders {
		func(i int, root context.Context) {
			builder := bbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BasicBannerMutation)
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
					_, err = mutators[i+1].Mutate(root, bbcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bbcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bbcb *BasicBannerCreateBulk) SaveX(ctx context.Context) []*BasicBanner {
	v, err := bbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
