// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/shop-go/ent/goodsclassify"
	"github.com/a20070322/shop-go/ent/goodsspu"
)

// GoodsClassifyCreate is the builder for creating a GoodsClassify entity.
type GoodsClassifyCreate struct {
	config
	mutation *GoodsClassifyMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (gcc *GoodsClassifyCreate) SetCreatedAt(t time.Time) *GoodsClassifyCreate {
	gcc.mutation.SetCreatedAt(t)
	return gcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gcc *GoodsClassifyCreate) SetNillableCreatedAt(t *time.Time) *GoodsClassifyCreate {
	if t != nil {
		gcc.SetCreatedAt(*t)
	}
	return gcc
}

// SetUpdatedAt sets the "updated_at" field.
func (gcc *GoodsClassifyCreate) SetUpdatedAt(t time.Time) *GoodsClassifyCreate {
	gcc.mutation.SetUpdatedAt(t)
	return gcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gcc *GoodsClassifyCreate) SetNillableUpdatedAt(t *time.Time) *GoodsClassifyCreate {
	if t != nil {
		gcc.SetUpdatedAt(*t)
	}
	return gcc
}

// SetDeletedAt sets the "deleted_at" field.
func (gcc *GoodsClassifyCreate) SetDeletedAt(t time.Time) *GoodsClassifyCreate {
	gcc.mutation.SetDeletedAt(t)
	return gcc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gcc *GoodsClassifyCreate) SetNillableDeletedAt(t *time.Time) *GoodsClassifyCreate {
	if t != nil {
		gcc.SetDeletedAt(*t)
	}
	return gcc
}

// SetClassifyName sets the "classify_name" field.
func (gcc *GoodsClassifyCreate) SetClassifyName(s string) *GoodsClassifyCreate {
	gcc.mutation.SetClassifyName(s)
	return gcc
}

// SetNillableClassifyName sets the "classify_name" field if the given value is not nil.
func (gcc *GoodsClassifyCreate) SetNillableClassifyName(s *string) *GoodsClassifyCreate {
	if s != nil {
		gcc.SetClassifyName(*s)
	}
	return gcc
}

// SetClassifyCode sets the "classify_code" field.
func (gcc *GoodsClassifyCreate) SetClassifyCode(s string) *GoodsClassifyCreate {
	gcc.mutation.SetClassifyCode(s)
	return gcc
}

// SetNillableClassifyCode sets the "classify_code" field if the given value is not nil.
func (gcc *GoodsClassifyCreate) SetNillableClassifyCode(s *string) *GoodsClassifyCreate {
	if s != nil {
		gcc.SetClassifyCode(*s)
	}
	return gcc
}

// SetPid sets the "pid" field.
func (gcc *GoodsClassifyCreate) SetPid(i int) *GoodsClassifyCreate {
	gcc.mutation.SetPid(i)
	return gcc
}

// SetNillablePid sets the "pid" field if the given value is not nil.
func (gcc *GoodsClassifyCreate) SetNillablePid(i *int) *GoodsClassifyCreate {
	if i != nil {
		gcc.SetPid(*i)
	}
	return gcc
}

// SetLevel sets the "level" field.
func (gcc *GoodsClassifyCreate) SetLevel(i int) *GoodsClassifyCreate {
	gcc.mutation.SetLevel(i)
	return gcc
}

// SetSort sets the "sort" field.
func (gcc *GoodsClassifyCreate) SetSort(i int) *GoodsClassifyCreate {
	gcc.mutation.SetSort(i)
	return gcc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (gcc *GoodsClassifyCreate) SetNillableSort(i *int) *GoodsClassifyCreate {
	if i != nil {
		gcc.SetSort(*i)
	}
	return gcc
}

// SetIcon sets the "icon" field.
func (gcc *GoodsClassifyCreate) SetIcon(s string) *GoodsClassifyCreate {
	gcc.mutation.SetIcon(s)
	return gcc
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (gcc *GoodsClassifyCreate) SetNillableIcon(s *string) *GoodsClassifyCreate {
	if s != nil {
		gcc.SetIcon(*s)
	}
	return gcc
}

// SetIsDisable sets the "is_disable" field.
func (gcc *GoodsClassifyCreate) SetIsDisable(b bool) *GoodsClassifyCreate {
	gcc.mutation.SetIsDisable(b)
	return gcc
}

// SetNillableIsDisable sets the "is_disable" field if the given value is not nil.
func (gcc *GoodsClassifyCreate) SetNillableIsDisable(b *bool) *GoodsClassifyCreate {
	if b != nil {
		gcc.SetIsDisable(*b)
	}
	return gcc
}

// AddGoodsSpuIDs adds the "goods_spu" edge to the GoodsSpu entity by IDs.
func (gcc *GoodsClassifyCreate) AddGoodsSpuIDs(ids ...int) *GoodsClassifyCreate {
	gcc.mutation.AddGoodsSpuIDs(ids...)
	return gcc
}

// AddGoodsSpu adds the "goods_spu" edges to the GoodsSpu entity.
func (gcc *GoodsClassifyCreate) AddGoodsSpu(g ...*GoodsSpu) *GoodsClassifyCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gcc.AddGoodsSpuIDs(ids...)
}

// Mutation returns the GoodsClassifyMutation object of the builder.
func (gcc *GoodsClassifyCreate) Mutation() *GoodsClassifyMutation {
	return gcc.mutation
}

// Save creates the GoodsClassify in the database.
func (gcc *GoodsClassifyCreate) Save(ctx context.Context) (*GoodsClassify, error) {
	var (
		err  error
		node *GoodsClassify
	)
	gcc.defaults()
	if len(gcc.hooks) == 0 {
		if err = gcc.check(); err != nil {
			return nil, err
		}
		node, err = gcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodsClassifyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gcc.check(); err != nil {
				return nil, err
			}
			gcc.mutation = mutation
			node, err = gcc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gcc.hooks) - 1; i >= 0; i-- {
			mut = gcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gcc *GoodsClassifyCreate) SaveX(ctx context.Context) *GoodsClassify {
	v, err := gcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (gcc *GoodsClassifyCreate) defaults() {
	if _, ok := gcc.mutation.CreatedAt(); !ok {
		v := goodsclassify.DefaultCreatedAt()
		gcc.mutation.SetCreatedAt(v)
	}
	if _, ok := gcc.mutation.Pid(); !ok {
		v := goodsclassify.DefaultPid
		gcc.mutation.SetPid(v)
	}
	if _, ok := gcc.mutation.Sort(); !ok {
		v := goodsclassify.DefaultSort
		gcc.mutation.SetSort(v)
	}
	if _, ok := gcc.mutation.Icon(); !ok {
		v := goodsclassify.DefaultIcon
		gcc.mutation.SetIcon(v)
	}
	if _, ok := gcc.mutation.IsDisable(); !ok {
		v := goodsclassify.DefaultIsDisable
		gcc.mutation.SetIsDisable(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gcc *GoodsClassifyCreate) check() error {
	if _, ok := gcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := gcc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New("ent: missing required field \"level\"")}
	}
	if v, ok := gcc.mutation.Sort(); ok {
		if err := goodsclassify.SortValidator(v); err != nil {
			return &ValidationError{Name: "sort", err: fmt.Errorf("ent: validator failed for field \"sort\": %w", err)}
		}
	}
	return nil
}

func (gcc *GoodsClassifyCreate) sqlSave(ctx context.Context) (*GoodsClassify, error) {
	_node, _spec := gcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gcc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gcc *GoodsClassifyCreate) createSpec() (*GoodsClassify, *sqlgraph.CreateSpec) {
	var (
		_node = &GoodsClassify{config: gcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: goodsclassify.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: goodsclassify.FieldID,
			},
		}
	)
	if value, ok := gcc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goodsclassify.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := gcc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goodsclassify.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := gcc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goodsclassify.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := gcc.mutation.ClassifyName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodsclassify.FieldClassifyName,
		})
		_node.ClassifyName = value
	}
	if value, ok := gcc.mutation.ClassifyCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodsclassify.FieldClassifyCode,
		})
		_node.ClassifyCode = value
	}
	if value, ok := gcc.mutation.Pid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodsclassify.FieldPid,
		})
		_node.Pid = value
	}
	if value, ok := gcc.mutation.Level(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodsclassify.FieldLevel,
		})
		_node.Level = value
	}
	if value, ok := gcc.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: goodsclassify.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := gcc.mutation.Icon(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodsclassify.FieldIcon,
		})
		_node.Icon = value
	}
	if value, ok := gcc.mutation.IsDisable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: goodsclassify.FieldIsDisable,
		})
		_node.IsDisable = value
	}
	if nodes := gcc.mutation.GoodsSpuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   goodsclassify.GoodsSpuTable,
			Columns: []string{goodsclassify.GoodsSpuColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GoodsClassifyCreateBulk is the builder for creating many GoodsClassify entities in bulk.
type GoodsClassifyCreateBulk struct {
	config
	builders []*GoodsClassifyCreate
}

// Save creates the GoodsClassify entities in the database.
func (gccb *GoodsClassifyCreateBulk) Save(ctx context.Context) ([]*GoodsClassify, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gccb.builders))
	nodes := make([]*GoodsClassify, len(gccb.builders))
	mutators := make([]Mutator, len(gccb.builders))
	for i := range gccb.builders {
		func(i int, root context.Context) {
			builder := gccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GoodsClassifyMutation)
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
					_, err = mutators[i+1].Mutate(root, gccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gccb *GoodsClassifyCreateBulk) SaveX(ctx context.Context) []*GoodsClassify {
	v, err := gccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}