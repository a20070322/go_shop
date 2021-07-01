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
	"github.com/a20070322/shop-go/ent/goodsspuimgs"
	"github.com/a20070322/shop-go/ent/predicate"
)

// GoodsSpuImgsUpdate is the builder for updating GoodsSpuImgs entities.
type GoodsSpuImgsUpdate struct {
	config
	hooks    []Hook
	mutation *GoodsSpuImgsMutation
}

// Where adds a new predicate for the GoodsSpuImgsUpdate builder.
func (gsiu *GoodsSpuImgsUpdate) Where(ps ...predicate.GoodsSpuImgs) *GoodsSpuImgsUpdate {
	gsiu.mutation.predicates = append(gsiu.mutation.predicates, ps...)
	return gsiu
}

// SetUpdatedAt sets the "updated_at" field.
func (gsiu *GoodsSpuImgsUpdate) SetUpdatedAt(t time.Time) *GoodsSpuImgsUpdate {
	gsiu.mutation.SetUpdatedAt(t)
	return gsiu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (gsiu *GoodsSpuImgsUpdate) ClearUpdatedAt() *GoodsSpuImgsUpdate {
	gsiu.mutation.ClearUpdatedAt()
	return gsiu
}

// SetDeletedAt sets the "deleted_at" field.
func (gsiu *GoodsSpuImgsUpdate) SetDeletedAt(t time.Time) *GoodsSpuImgsUpdate {
	gsiu.mutation.SetDeletedAt(t)
	return gsiu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gsiu *GoodsSpuImgsUpdate) SetNillableDeletedAt(t *time.Time) *GoodsSpuImgsUpdate {
	if t != nil {
		gsiu.SetDeletedAt(*t)
	}
	return gsiu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gsiu *GoodsSpuImgsUpdate) ClearDeletedAt() *GoodsSpuImgsUpdate {
	gsiu.mutation.ClearDeletedAt()
	return gsiu
}

// SetImgName sets the "img_name" field.
func (gsiu *GoodsSpuImgsUpdate) SetImgName(s string) *GoodsSpuImgsUpdate {
	gsiu.mutation.SetImgName(s)
	return gsiu
}

// SetNillableImgName sets the "img_name" field if the given value is not nil.
func (gsiu *GoodsSpuImgsUpdate) SetNillableImgName(s *string) *GoodsSpuImgsUpdate {
	if s != nil {
		gsiu.SetImgName(*s)
	}
	return gsiu
}

// ClearImgName clears the value of the "img_name" field.
func (gsiu *GoodsSpuImgsUpdate) ClearImgName() *GoodsSpuImgsUpdate {
	gsiu.mutation.ClearImgName()
	return gsiu
}

// SetImgPath sets the "img_path" field.
func (gsiu *GoodsSpuImgsUpdate) SetImgPath(s string) *GoodsSpuImgsUpdate {
	gsiu.mutation.SetImgPath(s)
	return gsiu
}

// SetNillableImgPath sets the "img_path" field if the given value is not nil.
func (gsiu *GoodsSpuImgsUpdate) SetNillableImgPath(s *string) *GoodsSpuImgsUpdate {
	if s != nil {
		gsiu.SetImgPath(*s)
	}
	return gsiu
}

// ClearImgPath clears the value of the "img_path" field.
func (gsiu *GoodsSpuImgsUpdate) ClearImgPath() *GoodsSpuImgsUpdate {
	gsiu.mutation.ClearImgPath()
	return gsiu
}

// SetGoodsSpuID sets the "goods_spu" edge to the GoodsSpu entity by ID.
func (gsiu *GoodsSpuImgsUpdate) SetGoodsSpuID(id int) *GoodsSpuImgsUpdate {
	gsiu.mutation.SetGoodsSpuID(id)
	return gsiu
}

// SetNillableGoodsSpuID sets the "goods_spu" edge to the GoodsSpu entity by ID if the given value is not nil.
func (gsiu *GoodsSpuImgsUpdate) SetNillableGoodsSpuID(id *int) *GoodsSpuImgsUpdate {
	if id != nil {
		gsiu = gsiu.SetGoodsSpuID(*id)
	}
	return gsiu
}

// SetGoodsSpu sets the "goods_spu" edge to the GoodsSpu entity.
func (gsiu *GoodsSpuImgsUpdate) SetGoodsSpu(g *GoodsSpu) *GoodsSpuImgsUpdate {
	return gsiu.SetGoodsSpuID(g.ID)
}

// Mutation returns the GoodsSpuImgsMutation object of the builder.
func (gsiu *GoodsSpuImgsUpdate) Mutation() *GoodsSpuImgsMutation {
	return gsiu.mutation
}

// ClearGoodsSpu clears the "goods_spu" edge to the GoodsSpu entity.
func (gsiu *GoodsSpuImgsUpdate) ClearGoodsSpu() *GoodsSpuImgsUpdate {
	gsiu.mutation.ClearGoodsSpu()
	return gsiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gsiu *GoodsSpuImgsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	gsiu.defaults()
	if len(gsiu.hooks) == 0 {
		affected, err = gsiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodsSpuImgsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gsiu.mutation = mutation
			affected, err = gsiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gsiu.hooks) - 1; i >= 0; i-- {
			mut = gsiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gsiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gsiu *GoodsSpuImgsUpdate) SaveX(ctx context.Context) int {
	affected, err := gsiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gsiu *GoodsSpuImgsUpdate) Exec(ctx context.Context) error {
	_, err := gsiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gsiu *GoodsSpuImgsUpdate) ExecX(ctx context.Context) {
	if err := gsiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gsiu *GoodsSpuImgsUpdate) defaults() {
	if _, ok := gsiu.mutation.UpdatedAt(); !ok && !gsiu.mutation.UpdatedAtCleared() {
		v := goodsspuimgs.UpdateDefaultUpdatedAt()
		gsiu.mutation.SetUpdatedAt(v)
	}
}

func (gsiu *GoodsSpuImgsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodsspuimgs.Table,
			Columns: goodsspuimgs.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: goodsspuimgs.FieldID,
			},
		},
	}
	if ps := gsiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gsiu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goodsspuimgs.FieldUpdatedAt,
		})
	}
	if gsiu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: goodsspuimgs.FieldUpdatedAt,
		})
	}
	if value, ok := gsiu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goodsspuimgs.FieldDeletedAt,
		})
	}
	if gsiu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: goodsspuimgs.FieldDeletedAt,
		})
	}
	if value, ok := gsiu.mutation.ImgName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodsspuimgs.FieldImgName,
		})
	}
	if gsiu.mutation.ImgNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: goodsspuimgs.FieldImgName,
		})
	}
	if value, ok := gsiu.mutation.ImgPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodsspuimgs.FieldImgPath,
		})
	}
	if gsiu.mutation.ImgPathCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: goodsspuimgs.FieldImgPath,
		})
	}
	if gsiu.mutation.GoodsSpuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   goodsspuimgs.GoodsSpuTable,
			Columns: []string{goodsspuimgs.GoodsSpuColumn},
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
	if nodes := gsiu.mutation.GoodsSpuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   goodsspuimgs.GoodsSpuTable,
			Columns: []string{goodsspuimgs.GoodsSpuColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, gsiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodsspuimgs.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GoodsSpuImgsUpdateOne is the builder for updating a single GoodsSpuImgs entity.
type GoodsSpuImgsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GoodsSpuImgsMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (gsiuo *GoodsSpuImgsUpdateOne) SetUpdatedAt(t time.Time) *GoodsSpuImgsUpdateOne {
	gsiuo.mutation.SetUpdatedAt(t)
	return gsiuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (gsiuo *GoodsSpuImgsUpdateOne) ClearUpdatedAt() *GoodsSpuImgsUpdateOne {
	gsiuo.mutation.ClearUpdatedAt()
	return gsiuo
}

// SetDeletedAt sets the "deleted_at" field.
func (gsiuo *GoodsSpuImgsUpdateOne) SetDeletedAt(t time.Time) *GoodsSpuImgsUpdateOne {
	gsiuo.mutation.SetDeletedAt(t)
	return gsiuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gsiuo *GoodsSpuImgsUpdateOne) SetNillableDeletedAt(t *time.Time) *GoodsSpuImgsUpdateOne {
	if t != nil {
		gsiuo.SetDeletedAt(*t)
	}
	return gsiuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gsiuo *GoodsSpuImgsUpdateOne) ClearDeletedAt() *GoodsSpuImgsUpdateOne {
	gsiuo.mutation.ClearDeletedAt()
	return gsiuo
}

// SetImgName sets the "img_name" field.
func (gsiuo *GoodsSpuImgsUpdateOne) SetImgName(s string) *GoodsSpuImgsUpdateOne {
	gsiuo.mutation.SetImgName(s)
	return gsiuo
}

// SetNillableImgName sets the "img_name" field if the given value is not nil.
func (gsiuo *GoodsSpuImgsUpdateOne) SetNillableImgName(s *string) *GoodsSpuImgsUpdateOne {
	if s != nil {
		gsiuo.SetImgName(*s)
	}
	return gsiuo
}

// ClearImgName clears the value of the "img_name" field.
func (gsiuo *GoodsSpuImgsUpdateOne) ClearImgName() *GoodsSpuImgsUpdateOne {
	gsiuo.mutation.ClearImgName()
	return gsiuo
}

// SetImgPath sets the "img_path" field.
func (gsiuo *GoodsSpuImgsUpdateOne) SetImgPath(s string) *GoodsSpuImgsUpdateOne {
	gsiuo.mutation.SetImgPath(s)
	return gsiuo
}

// SetNillableImgPath sets the "img_path" field if the given value is not nil.
func (gsiuo *GoodsSpuImgsUpdateOne) SetNillableImgPath(s *string) *GoodsSpuImgsUpdateOne {
	if s != nil {
		gsiuo.SetImgPath(*s)
	}
	return gsiuo
}

// ClearImgPath clears the value of the "img_path" field.
func (gsiuo *GoodsSpuImgsUpdateOne) ClearImgPath() *GoodsSpuImgsUpdateOne {
	gsiuo.mutation.ClearImgPath()
	return gsiuo
}

// SetGoodsSpuID sets the "goods_spu" edge to the GoodsSpu entity by ID.
func (gsiuo *GoodsSpuImgsUpdateOne) SetGoodsSpuID(id int) *GoodsSpuImgsUpdateOne {
	gsiuo.mutation.SetGoodsSpuID(id)
	return gsiuo
}

// SetNillableGoodsSpuID sets the "goods_spu" edge to the GoodsSpu entity by ID if the given value is not nil.
func (gsiuo *GoodsSpuImgsUpdateOne) SetNillableGoodsSpuID(id *int) *GoodsSpuImgsUpdateOne {
	if id != nil {
		gsiuo = gsiuo.SetGoodsSpuID(*id)
	}
	return gsiuo
}

// SetGoodsSpu sets the "goods_spu" edge to the GoodsSpu entity.
func (gsiuo *GoodsSpuImgsUpdateOne) SetGoodsSpu(g *GoodsSpu) *GoodsSpuImgsUpdateOne {
	return gsiuo.SetGoodsSpuID(g.ID)
}

// Mutation returns the GoodsSpuImgsMutation object of the builder.
func (gsiuo *GoodsSpuImgsUpdateOne) Mutation() *GoodsSpuImgsMutation {
	return gsiuo.mutation
}

// ClearGoodsSpu clears the "goods_spu" edge to the GoodsSpu entity.
func (gsiuo *GoodsSpuImgsUpdateOne) ClearGoodsSpu() *GoodsSpuImgsUpdateOne {
	gsiuo.mutation.ClearGoodsSpu()
	return gsiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gsiuo *GoodsSpuImgsUpdateOne) Select(field string, fields ...string) *GoodsSpuImgsUpdateOne {
	gsiuo.fields = append([]string{field}, fields...)
	return gsiuo
}

// Save executes the query and returns the updated GoodsSpuImgs entity.
func (gsiuo *GoodsSpuImgsUpdateOne) Save(ctx context.Context) (*GoodsSpuImgs, error) {
	var (
		err  error
		node *GoodsSpuImgs
	)
	gsiuo.defaults()
	if len(gsiuo.hooks) == 0 {
		node, err = gsiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodsSpuImgsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gsiuo.mutation = mutation
			node, err = gsiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gsiuo.hooks) - 1; i >= 0; i-- {
			mut = gsiuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gsiuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gsiuo *GoodsSpuImgsUpdateOne) SaveX(ctx context.Context) *GoodsSpuImgs {
	node, err := gsiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gsiuo *GoodsSpuImgsUpdateOne) Exec(ctx context.Context) error {
	_, err := gsiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gsiuo *GoodsSpuImgsUpdateOne) ExecX(ctx context.Context) {
	if err := gsiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gsiuo *GoodsSpuImgsUpdateOne) defaults() {
	if _, ok := gsiuo.mutation.UpdatedAt(); !ok && !gsiuo.mutation.UpdatedAtCleared() {
		v := goodsspuimgs.UpdateDefaultUpdatedAt()
		gsiuo.mutation.SetUpdatedAt(v)
	}
}

func (gsiuo *GoodsSpuImgsUpdateOne) sqlSave(ctx context.Context) (_node *GoodsSpuImgs, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodsspuimgs.Table,
			Columns: goodsspuimgs.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: goodsspuimgs.FieldID,
			},
		},
	}
	id, ok := gsiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing GoodsSpuImgs.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := gsiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodsspuimgs.FieldID)
		for _, f := range fields {
			if !goodsspuimgs.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != goodsspuimgs.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gsiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gsiuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goodsspuimgs.FieldUpdatedAt,
		})
	}
	if gsiuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: goodsspuimgs.FieldUpdatedAt,
		})
	}
	if value, ok := gsiuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goodsspuimgs.FieldDeletedAt,
		})
	}
	if gsiuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: goodsspuimgs.FieldDeletedAt,
		})
	}
	if value, ok := gsiuo.mutation.ImgName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodsspuimgs.FieldImgName,
		})
	}
	if gsiuo.mutation.ImgNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: goodsspuimgs.FieldImgName,
		})
	}
	if value, ok := gsiuo.mutation.ImgPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodsspuimgs.FieldImgPath,
		})
	}
	if gsiuo.mutation.ImgPathCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: goodsspuimgs.FieldImgPath,
		})
	}
	if gsiuo.mutation.GoodsSpuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   goodsspuimgs.GoodsSpuTable,
			Columns: []string{goodsspuimgs.GoodsSpuColumn},
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
	if nodes := gsiuo.mutation.GoodsSpuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   goodsspuimgs.GoodsSpuTable,
			Columns: []string{goodsspuimgs.GoodsSpuColumn},
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
	_node = &GoodsSpuImgs{config: gsiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gsiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodsspuimgs.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
