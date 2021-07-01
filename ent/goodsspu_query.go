// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/shop-go/ent/goodsclassify"
	"github.com/a20070322/shop-go/ent/goodssku"
	"github.com/a20070322/shop-go/ent/goodsspu"
	"github.com/a20070322/shop-go/ent/goodsspuimgs"
	"github.com/a20070322/shop-go/ent/ordergoodssku"
	"github.com/a20070322/shop-go/ent/predicate"
)

// GoodsSpuQuery is the builder for querying GoodsSpu entities.
type GoodsSpuQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GoodsSpu
	// eager-loading edges.
	withGoodsClassify *GoodsClassifyQuery
	withGoodsSku      *GoodsSkuQuery
	withOrderGoodsSku *OrderGoodsSkuQuery
	withGoodsSpuImgs  *GoodsSpuImgsQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodsSpuQuery builder.
func (gsq *GoodsSpuQuery) Where(ps ...predicate.GoodsSpu) *GoodsSpuQuery {
	gsq.predicates = append(gsq.predicates, ps...)
	return gsq
}

// Limit adds a limit step to the query.
func (gsq *GoodsSpuQuery) Limit(limit int) *GoodsSpuQuery {
	gsq.limit = &limit
	return gsq
}

// Offset adds an offset step to the query.
func (gsq *GoodsSpuQuery) Offset(offset int) *GoodsSpuQuery {
	gsq.offset = &offset
	return gsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gsq *GoodsSpuQuery) Unique(unique bool) *GoodsSpuQuery {
	gsq.unique = &unique
	return gsq
}

// Order adds an order step to the query.
func (gsq *GoodsSpuQuery) Order(o ...OrderFunc) *GoodsSpuQuery {
	gsq.order = append(gsq.order, o...)
	return gsq
}

// QueryGoodsClassify chains the current query on the "goods_classify" edge.
func (gsq *GoodsSpuQuery) QueryGoodsClassify() *GoodsClassifyQuery {
	query := &GoodsClassifyQuery{config: gsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(goodsspu.Table, goodsspu.FieldID, selector),
			sqlgraph.To(goodsclassify.Table, goodsclassify.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, goodsspu.GoodsClassifyTable, goodsspu.GoodsClassifyColumn),
		)
		fromU = sqlgraph.SetNeighbors(gsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGoodsSku chains the current query on the "goods_sku" edge.
func (gsq *GoodsSpuQuery) QueryGoodsSku() *GoodsSkuQuery {
	query := &GoodsSkuQuery{config: gsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(goodsspu.Table, goodsspu.FieldID, selector),
			sqlgraph.To(goodssku.Table, goodssku.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, goodsspu.GoodsSkuTable, goodsspu.GoodsSkuColumn),
		)
		fromU = sqlgraph.SetNeighbors(gsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderGoodsSku chains the current query on the "order_goods_sku" edge.
func (gsq *GoodsSpuQuery) QueryOrderGoodsSku() *OrderGoodsSkuQuery {
	query := &OrderGoodsSkuQuery{config: gsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(goodsspu.Table, goodsspu.FieldID, selector),
			sqlgraph.To(ordergoodssku.Table, ordergoodssku.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, goodsspu.OrderGoodsSkuTable, goodsspu.OrderGoodsSkuColumn),
		)
		fromU = sqlgraph.SetNeighbors(gsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGoodsSpuImgs chains the current query on the "goods_spu_imgs" edge.
func (gsq *GoodsSpuQuery) QueryGoodsSpuImgs() *GoodsSpuImgsQuery {
	query := &GoodsSpuImgsQuery{config: gsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(goodsspu.Table, goodsspu.FieldID, selector),
			sqlgraph.To(goodsspuimgs.Table, goodsspuimgs.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, goodsspu.GoodsSpuImgsTable, goodsspu.GoodsSpuImgsColumn),
		)
		fromU = sqlgraph.SetNeighbors(gsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GoodsSpu entity from the query.
// Returns a *NotFoundError when no GoodsSpu was found.
func (gsq *GoodsSpuQuery) First(ctx context.Context) (*GoodsSpu, error) {
	nodes, err := gsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodsspu.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gsq *GoodsSpuQuery) FirstX(ctx context.Context) *GoodsSpu {
	node, err := gsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodsSpu ID from the query.
// Returns a *NotFoundError when no GoodsSpu ID was found.
func (gsq *GoodsSpuQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodsspu.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gsq *GoodsSpuQuery) FirstIDX(ctx context.Context) int {
	id, err := gsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodsSpu entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one GoodsSpu entity is not found.
// Returns a *NotFoundError when no GoodsSpu entities are found.
func (gsq *GoodsSpuQuery) Only(ctx context.Context) (*GoodsSpu, error) {
	nodes, err := gsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodsspu.Label}
	default:
		return nil, &NotSingularError{goodsspu.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gsq *GoodsSpuQuery) OnlyX(ctx context.Context) *GoodsSpu {
	node, err := gsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodsSpu ID in the query.
// Returns a *NotSingularError when exactly one GoodsSpu ID is not found.
// Returns a *NotFoundError when no entities are found.
func (gsq *GoodsSpuQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodsspu.Label}
	default:
		err = &NotSingularError{goodsspu.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gsq *GoodsSpuQuery) OnlyIDX(ctx context.Context) int {
	id, err := gsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodsSpus.
func (gsq *GoodsSpuQuery) All(ctx context.Context) ([]*GoodsSpu, error) {
	if err := gsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gsq *GoodsSpuQuery) AllX(ctx context.Context) []*GoodsSpu {
	nodes, err := gsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodsSpu IDs.
func (gsq *GoodsSpuQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := gsq.Select(goodsspu.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gsq *GoodsSpuQuery) IDsX(ctx context.Context) []int {
	ids, err := gsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gsq *GoodsSpuQuery) Count(ctx context.Context) (int, error) {
	if err := gsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gsq *GoodsSpuQuery) CountX(ctx context.Context) int {
	count, err := gsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gsq *GoodsSpuQuery) Exist(ctx context.Context) (bool, error) {
	if err := gsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gsq *GoodsSpuQuery) ExistX(ctx context.Context) bool {
	exist, err := gsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodsSpuQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gsq *GoodsSpuQuery) Clone() *GoodsSpuQuery {
	if gsq == nil {
		return nil
	}
	return &GoodsSpuQuery{
		config:            gsq.config,
		limit:             gsq.limit,
		offset:            gsq.offset,
		order:             append([]OrderFunc{}, gsq.order...),
		predicates:        append([]predicate.GoodsSpu{}, gsq.predicates...),
		withGoodsClassify: gsq.withGoodsClassify.Clone(),
		withGoodsSku:      gsq.withGoodsSku.Clone(),
		withOrderGoodsSku: gsq.withOrderGoodsSku.Clone(),
		withGoodsSpuImgs:  gsq.withGoodsSpuImgs.Clone(),
		// clone intermediate query.
		sql:  gsq.sql.Clone(),
		path: gsq.path,
	}
}

// WithGoodsClassify tells the query-builder to eager-load the nodes that are connected to
// the "goods_classify" edge. The optional arguments are used to configure the query builder of the edge.
func (gsq *GoodsSpuQuery) WithGoodsClassify(opts ...func(*GoodsClassifyQuery)) *GoodsSpuQuery {
	query := &GoodsClassifyQuery{config: gsq.config}
	for _, opt := range opts {
		opt(query)
	}
	gsq.withGoodsClassify = query
	return gsq
}

// WithGoodsSku tells the query-builder to eager-load the nodes that are connected to
// the "goods_sku" edge. The optional arguments are used to configure the query builder of the edge.
func (gsq *GoodsSpuQuery) WithGoodsSku(opts ...func(*GoodsSkuQuery)) *GoodsSpuQuery {
	query := &GoodsSkuQuery{config: gsq.config}
	for _, opt := range opts {
		opt(query)
	}
	gsq.withGoodsSku = query
	return gsq
}

// WithOrderGoodsSku tells the query-builder to eager-load the nodes that are connected to
// the "order_goods_sku" edge. The optional arguments are used to configure the query builder of the edge.
func (gsq *GoodsSpuQuery) WithOrderGoodsSku(opts ...func(*OrderGoodsSkuQuery)) *GoodsSpuQuery {
	query := &OrderGoodsSkuQuery{config: gsq.config}
	for _, opt := range opts {
		opt(query)
	}
	gsq.withOrderGoodsSku = query
	return gsq
}

// WithGoodsSpuImgs tells the query-builder to eager-load the nodes that are connected to
// the "goods_spu_imgs" edge. The optional arguments are used to configure the query builder of the edge.
func (gsq *GoodsSpuQuery) WithGoodsSpuImgs(opts ...func(*GoodsSpuImgsQuery)) *GoodsSpuQuery {
	query := &GoodsSpuImgsQuery{config: gsq.config}
	for _, opt := range opts {
		opt(query)
	}
	gsq.withGoodsSpuImgs = query
	return gsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GoodsSpu.Query().
//		GroupBy(goodsspu.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gsq *GoodsSpuQuery) GroupBy(field string, fields ...string) *GoodsSpuGroupBy {
	group := &GoodsSpuGroupBy{config: gsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gsq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.GoodsSpu.Query().
//		Select(goodsspu.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (gsq *GoodsSpuQuery) Select(field string, fields ...string) *GoodsSpuSelect {
	gsq.fields = append([]string{field}, fields...)
	return &GoodsSpuSelect{GoodsSpuQuery: gsq}
}

func (gsq *GoodsSpuQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gsq.fields {
		if !goodsspu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gsq.path != nil {
		prev, err := gsq.path(ctx)
		if err != nil {
			return err
		}
		gsq.sql = prev
	}
	return nil
}

func (gsq *GoodsSpuQuery) sqlAll(ctx context.Context) ([]*GoodsSpu, error) {
	var (
		nodes       = []*GoodsSpu{}
		withFKs     = gsq.withFKs
		_spec       = gsq.querySpec()
		loadedTypes = [4]bool{
			gsq.withGoodsClassify != nil,
			gsq.withGoodsSku != nil,
			gsq.withOrderGoodsSku != nil,
			gsq.withGoodsSpuImgs != nil,
		}
	)
	if gsq.withGoodsClassify != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, goodsspu.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GoodsSpu{config: gsq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, gsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := gsq.withGoodsClassify; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*GoodsSpu)
		for i := range nodes {
			if nodes[i].goods_classify_goods_spu == nil {
				continue
			}
			fk := *nodes[i].goods_classify_goods_spu
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(goodsclassify.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "goods_classify_goods_spu" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.GoodsClassify = n
			}
		}
	}

	if query := gsq.withGoodsSku; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*GoodsSpu)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.GoodsSku = []*GoodsSku{}
		}
		query.withFKs = true
		query.Where(predicate.GoodsSku(func(s *sql.Selector) {
			s.Where(sql.InValues(goodsspu.GoodsSkuColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.goods_spu_goods_sku
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "goods_spu_goods_sku" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "goods_spu_goods_sku" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.GoodsSku = append(node.Edges.GoodsSku, n)
		}
	}

	if query := gsq.withOrderGoodsSku; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*GoodsSpu)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OrderGoodsSku = []*OrderGoodsSku{}
		}
		query.withFKs = true
		query.Where(predicate.OrderGoodsSku(func(s *sql.Selector) {
			s.Where(sql.InValues(goodsspu.OrderGoodsSkuColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.goods_spu_order_goods_sku
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "goods_spu_order_goods_sku" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "goods_spu_order_goods_sku" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OrderGoodsSku = append(node.Edges.OrderGoodsSku, n)
		}
	}

	if query := gsq.withGoodsSpuImgs; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*GoodsSpu)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.GoodsSpuImgs = []*GoodsSpuImgs{}
		}
		query.withFKs = true
		query.Where(predicate.GoodsSpuImgs(func(s *sql.Selector) {
			s.Where(sql.InValues(goodsspu.GoodsSpuImgsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.goods_spu_goods_spu_imgs
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "goods_spu_goods_spu_imgs" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "goods_spu_goods_spu_imgs" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.GoodsSpuImgs = append(node.Edges.GoodsSpuImgs, n)
		}
	}

	return nodes, nil
}

func (gsq *GoodsSpuQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gsq.querySpec()
	return sqlgraph.CountNodes(ctx, gsq.driver, _spec)
}

func (gsq *GoodsSpuQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gsq *GoodsSpuQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodsspu.Table,
			Columns: goodsspu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: goodsspu.FieldID,
			},
		},
		From:   gsq.sql,
		Unique: true,
	}
	if unique := gsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodsspu.FieldID)
		for i := range fields {
			if fields[i] != goodsspu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gsq *GoodsSpuQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gsq.driver.Dialect())
	t1 := builder.Table(goodsspu.Table)
	selector := builder.Select(t1.Columns(goodsspu.Columns...)...).From(t1)
	if gsq.sql != nil {
		selector = gsq.sql
		selector.Select(selector.Columns(goodsspu.Columns...)...)
	}
	for _, p := range gsq.predicates {
		p(selector)
	}
	for _, p := range gsq.order {
		p(selector)
	}
	if offset := gsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GoodsSpuGroupBy is the group-by builder for GoodsSpu entities.
type GoodsSpuGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gsgb *GoodsSpuGroupBy) Aggregate(fns ...AggregateFunc) *GoodsSpuGroupBy {
	gsgb.fns = append(gsgb.fns, fns...)
	return gsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (gsgb *GoodsSpuGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gsgb.path(ctx)
	if err != nil {
		return err
	}
	gsgb.sql = query
	return gsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gsgb *GoodsSpuGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (gsgb *GoodsSpuGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gsgb.fields) > 1 {
		return nil, errors.New("ent: GoodsSpuGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gsgb *GoodsSpuGroupBy) StringsX(ctx context.Context) []string {
	v, err := gsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gsgb *GoodsSpuGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodsspu.Label}
	default:
		err = fmt.Errorf("ent: GoodsSpuGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gsgb *GoodsSpuGroupBy) StringX(ctx context.Context) string {
	v, err := gsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (gsgb *GoodsSpuGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gsgb.fields) > 1 {
		return nil, errors.New("ent: GoodsSpuGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gsgb *GoodsSpuGroupBy) IntsX(ctx context.Context) []int {
	v, err := gsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gsgb *GoodsSpuGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodsspu.Label}
	default:
		err = fmt.Errorf("ent: GoodsSpuGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gsgb *GoodsSpuGroupBy) IntX(ctx context.Context) int {
	v, err := gsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (gsgb *GoodsSpuGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gsgb.fields) > 1 {
		return nil, errors.New("ent: GoodsSpuGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gsgb *GoodsSpuGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gsgb *GoodsSpuGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodsspu.Label}
	default:
		err = fmt.Errorf("ent: GoodsSpuGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gsgb *GoodsSpuGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (gsgb *GoodsSpuGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gsgb.fields) > 1 {
		return nil, errors.New("ent: GoodsSpuGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gsgb *GoodsSpuGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gsgb *GoodsSpuGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodsspu.Label}
	default:
		err = fmt.Errorf("ent: GoodsSpuGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gsgb *GoodsSpuGroupBy) BoolX(ctx context.Context) bool {
	v, err := gsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gsgb *GoodsSpuGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gsgb.fields {
		if !goodsspu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gsgb *GoodsSpuGroupBy) sqlQuery() *sql.Selector {
	selector := gsgb.sql
	columns := make([]string, 0, len(gsgb.fields)+len(gsgb.fns))
	columns = append(columns, gsgb.fields...)
	for _, fn := range gsgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(gsgb.fields...)
}

// GoodsSpuSelect is the builder for selecting fields of GoodsSpu entities.
type GoodsSpuSelect struct {
	*GoodsSpuQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gss *GoodsSpuSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gss.prepareQuery(ctx); err != nil {
		return err
	}
	gss.sql = gss.GoodsSpuQuery.sqlQuery(ctx)
	return gss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gss *GoodsSpuSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gss *GoodsSpuSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gss.fields) > 1 {
		return nil, errors.New("ent: GoodsSpuSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gss *GoodsSpuSelect) StringsX(ctx context.Context) []string {
	v, err := gss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gss *GoodsSpuSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodsspu.Label}
	default:
		err = fmt.Errorf("ent: GoodsSpuSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gss *GoodsSpuSelect) StringX(ctx context.Context) string {
	v, err := gss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gss *GoodsSpuSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gss.fields) > 1 {
		return nil, errors.New("ent: GoodsSpuSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gss *GoodsSpuSelect) IntsX(ctx context.Context) []int {
	v, err := gss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gss *GoodsSpuSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodsspu.Label}
	default:
		err = fmt.Errorf("ent: GoodsSpuSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gss *GoodsSpuSelect) IntX(ctx context.Context) int {
	v, err := gss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gss *GoodsSpuSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gss.fields) > 1 {
		return nil, errors.New("ent: GoodsSpuSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gss *GoodsSpuSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gss *GoodsSpuSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodsspu.Label}
	default:
		err = fmt.Errorf("ent: GoodsSpuSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gss *GoodsSpuSelect) Float64X(ctx context.Context) float64 {
	v, err := gss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gss *GoodsSpuSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gss.fields) > 1 {
		return nil, errors.New("ent: GoodsSpuSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gss *GoodsSpuSelect) BoolsX(ctx context.Context) []bool {
	v, err := gss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gss *GoodsSpuSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodsspu.Label}
	default:
		err = fmt.Errorf("ent: GoodsSpuSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gss *GoodsSpuSelect) BoolX(ctx context.Context) bool {
	v, err := gss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gss *GoodsSpuSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gss.sqlQuery().Query()
	if err := gss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gss *GoodsSpuSelect) sqlQuery() sql.Querier {
	selector := gss.sql
	selector.Select(selector.Columns(gss.fields...)...)
	return selector
}
