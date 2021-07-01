// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/a20070322/shop-go/ent/goodsspu"
	"github.com/a20070322/shop-go/ent/ordergoodssku"
	"github.com/a20070322/shop-go/ent/orderinfo"
	"github.com/a20070322/shop-go/ent/predicate"
)

// OrderGoodsSkuQuery is the builder for querying OrderGoodsSku entities.
type OrderGoodsSkuQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrderGoodsSku
	// eager-loading edges.
	withGoodsSpu  *GoodsSpuQuery
	withOrderInfo *OrderInfoQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderGoodsSkuQuery builder.
func (ogsq *OrderGoodsSkuQuery) Where(ps ...predicate.OrderGoodsSku) *OrderGoodsSkuQuery {
	ogsq.predicates = append(ogsq.predicates, ps...)
	return ogsq
}

// Limit adds a limit step to the query.
func (ogsq *OrderGoodsSkuQuery) Limit(limit int) *OrderGoodsSkuQuery {
	ogsq.limit = &limit
	return ogsq
}

// Offset adds an offset step to the query.
func (ogsq *OrderGoodsSkuQuery) Offset(offset int) *OrderGoodsSkuQuery {
	ogsq.offset = &offset
	return ogsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ogsq *OrderGoodsSkuQuery) Unique(unique bool) *OrderGoodsSkuQuery {
	ogsq.unique = &unique
	return ogsq
}

// Order adds an order step to the query.
func (ogsq *OrderGoodsSkuQuery) Order(o ...OrderFunc) *OrderGoodsSkuQuery {
	ogsq.order = append(ogsq.order, o...)
	return ogsq
}

// QueryGoodsSpu chains the current query on the "goods_spu" edge.
func (ogsq *OrderGoodsSkuQuery) QueryGoodsSpu() *GoodsSpuQuery {
	query := &GoodsSpuQuery{config: ogsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ogsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ogsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ordergoodssku.Table, ordergoodssku.FieldID, selector),
			sqlgraph.To(goodsspu.Table, goodsspu.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ordergoodssku.GoodsSpuTable, ordergoodssku.GoodsSpuColumn),
		)
		fromU = sqlgraph.SetNeighbors(ogsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderInfo chains the current query on the "order_info" edge.
func (ogsq *OrderGoodsSkuQuery) QueryOrderInfo() *OrderInfoQuery {
	query := &OrderInfoQuery{config: ogsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ogsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ogsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ordergoodssku.Table, ordergoodssku.FieldID, selector),
			sqlgraph.To(orderinfo.Table, orderinfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ordergoodssku.OrderInfoTable, ordergoodssku.OrderInfoColumn),
		)
		fromU = sqlgraph.SetNeighbors(ogsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrderGoodsSku entity from the query.
// Returns a *NotFoundError when no OrderGoodsSku was found.
func (ogsq *OrderGoodsSkuQuery) First(ctx context.Context) (*OrderGoodsSku, error) {
	nodes, err := ogsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ordergoodssku.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ogsq *OrderGoodsSkuQuery) FirstX(ctx context.Context) *OrderGoodsSku {
	node, err := ogsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderGoodsSku ID from the query.
// Returns a *NotFoundError when no OrderGoodsSku ID was found.
func (ogsq *OrderGoodsSkuQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ogsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ordergoodssku.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ogsq *OrderGoodsSkuQuery) FirstIDX(ctx context.Context) int {
	id, err := ogsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderGoodsSku entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one OrderGoodsSku entity is not found.
// Returns a *NotFoundError when no OrderGoodsSku entities are found.
func (ogsq *OrderGoodsSkuQuery) Only(ctx context.Context) (*OrderGoodsSku, error) {
	nodes, err := ogsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ordergoodssku.Label}
	default:
		return nil, &NotSingularError{ordergoodssku.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ogsq *OrderGoodsSkuQuery) OnlyX(ctx context.Context) *OrderGoodsSku {
	node, err := ogsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderGoodsSku ID in the query.
// Returns a *NotSingularError when exactly one OrderGoodsSku ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ogsq *OrderGoodsSkuQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ogsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ordergoodssku.Label}
	default:
		err = &NotSingularError{ordergoodssku.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ogsq *OrderGoodsSkuQuery) OnlyIDX(ctx context.Context) int {
	id, err := ogsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderGoodsSkus.
func (ogsq *OrderGoodsSkuQuery) All(ctx context.Context) ([]*OrderGoodsSku, error) {
	if err := ogsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ogsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ogsq *OrderGoodsSkuQuery) AllX(ctx context.Context) []*OrderGoodsSku {
	nodes, err := ogsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderGoodsSku IDs.
func (ogsq *OrderGoodsSkuQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ogsq.Select(ordergoodssku.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ogsq *OrderGoodsSkuQuery) IDsX(ctx context.Context) []int {
	ids, err := ogsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ogsq *OrderGoodsSkuQuery) Count(ctx context.Context) (int, error) {
	if err := ogsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ogsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ogsq *OrderGoodsSkuQuery) CountX(ctx context.Context) int {
	count, err := ogsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ogsq *OrderGoodsSkuQuery) Exist(ctx context.Context) (bool, error) {
	if err := ogsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ogsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ogsq *OrderGoodsSkuQuery) ExistX(ctx context.Context) bool {
	exist, err := ogsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderGoodsSkuQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ogsq *OrderGoodsSkuQuery) Clone() *OrderGoodsSkuQuery {
	if ogsq == nil {
		return nil
	}
	return &OrderGoodsSkuQuery{
		config:        ogsq.config,
		limit:         ogsq.limit,
		offset:        ogsq.offset,
		order:         append([]OrderFunc{}, ogsq.order...),
		predicates:    append([]predicate.OrderGoodsSku{}, ogsq.predicates...),
		withGoodsSpu:  ogsq.withGoodsSpu.Clone(),
		withOrderInfo: ogsq.withOrderInfo.Clone(),
		// clone intermediate query.
		sql:  ogsq.sql.Clone(),
		path: ogsq.path,
	}
}

// WithGoodsSpu tells the query-builder to eager-load the nodes that are connected to
// the "goods_spu" edge. The optional arguments are used to configure the query builder of the edge.
func (ogsq *OrderGoodsSkuQuery) WithGoodsSpu(opts ...func(*GoodsSpuQuery)) *OrderGoodsSkuQuery {
	query := &GoodsSpuQuery{config: ogsq.config}
	for _, opt := range opts {
		opt(query)
	}
	ogsq.withGoodsSpu = query
	return ogsq
}

// WithOrderInfo tells the query-builder to eager-load the nodes that are connected to
// the "order_info" edge. The optional arguments are used to configure the query builder of the edge.
func (ogsq *OrderGoodsSkuQuery) WithOrderInfo(opts ...func(*OrderInfoQuery)) *OrderGoodsSkuQuery {
	query := &OrderInfoQuery{config: ogsq.config}
	for _, opt := range opts {
		opt(query)
	}
	ogsq.withOrderInfo = query
	return ogsq
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
//	client.OrderGoodsSku.Query().
//		GroupBy(ordergoodssku.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ogsq *OrderGoodsSkuQuery) GroupBy(field string, fields ...string) *OrderGoodsSkuGroupBy {
	group := &OrderGoodsSkuGroupBy{config: ogsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ogsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ogsq.sqlQuery(ctx), nil
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
//	client.OrderGoodsSku.Query().
//		Select(ordergoodssku.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (ogsq *OrderGoodsSkuQuery) Select(field string, fields ...string) *OrderGoodsSkuSelect {
	ogsq.fields = append([]string{field}, fields...)
	return &OrderGoodsSkuSelect{OrderGoodsSkuQuery: ogsq}
}

func (ogsq *OrderGoodsSkuQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ogsq.fields {
		if !ordergoodssku.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ogsq.path != nil {
		prev, err := ogsq.path(ctx)
		if err != nil {
			return err
		}
		ogsq.sql = prev
	}
	return nil
}

func (ogsq *OrderGoodsSkuQuery) sqlAll(ctx context.Context) ([]*OrderGoodsSku, error) {
	var (
		nodes       = []*OrderGoodsSku{}
		withFKs     = ogsq.withFKs
		_spec       = ogsq.querySpec()
		loadedTypes = [2]bool{
			ogsq.withGoodsSpu != nil,
			ogsq.withOrderInfo != nil,
		}
	)
	if ogsq.withGoodsSpu != nil || ogsq.withOrderInfo != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, ordergoodssku.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &OrderGoodsSku{config: ogsq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ogsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ogsq.withGoodsSpu; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrderGoodsSku)
		for i := range nodes {
			if nodes[i].goods_spu_order_goods_sku == nil {
				continue
			}
			fk := *nodes[i].goods_spu_order_goods_sku
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(goodsspu.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "goods_spu_order_goods_sku" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.GoodsSpu = n
			}
		}
	}

	if query := ogsq.withOrderInfo; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrderGoodsSku)
		for i := range nodes {
			if nodes[i].order_info_order_goods_sku == nil {
				continue
			}
			fk := *nodes[i].order_info_order_goods_sku
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(orderinfo.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "order_info_order_goods_sku" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OrderInfo = n
			}
		}
	}

	return nodes, nil
}

func (ogsq *OrderGoodsSkuQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ogsq.querySpec()
	return sqlgraph.CountNodes(ctx, ogsq.driver, _spec)
}

func (ogsq *OrderGoodsSkuQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ogsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ogsq *OrderGoodsSkuQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ordergoodssku.Table,
			Columns: ordergoodssku.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ordergoodssku.FieldID,
			},
		},
		From:   ogsq.sql,
		Unique: true,
	}
	if unique := ogsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ogsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ordergoodssku.FieldID)
		for i := range fields {
			if fields[i] != ordergoodssku.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ogsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ogsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ogsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ogsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ogsq *OrderGoodsSkuQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ogsq.driver.Dialect())
	t1 := builder.Table(ordergoodssku.Table)
	selector := builder.Select(t1.Columns(ordergoodssku.Columns...)...).From(t1)
	if ogsq.sql != nil {
		selector = ogsq.sql
		selector.Select(selector.Columns(ordergoodssku.Columns...)...)
	}
	for _, p := range ogsq.predicates {
		p(selector)
	}
	for _, p := range ogsq.order {
		p(selector)
	}
	if offset := ogsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ogsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderGoodsSkuGroupBy is the group-by builder for OrderGoodsSku entities.
type OrderGoodsSkuGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogsgb *OrderGoodsSkuGroupBy) Aggregate(fns ...AggregateFunc) *OrderGoodsSkuGroupBy {
	ogsgb.fns = append(ogsgb.fns, fns...)
	return ogsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ogsgb *OrderGoodsSkuGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ogsgb.path(ctx)
	if err != nil {
		return err
	}
	ogsgb.sql = query
	return ogsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ogsgb *OrderGoodsSkuGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ogsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ogsgb *OrderGoodsSkuGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ogsgb.fields) > 1 {
		return nil, errors.New("ent: OrderGoodsSkuGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ogsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ogsgb *OrderGoodsSkuGroupBy) StringsX(ctx context.Context) []string {
	v, err := ogsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ogsgb *OrderGoodsSkuGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ogsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ordergoodssku.Label}
	default:
		err = fmt.Errorf("ent: OrderGoodsSkuGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ogsgb *OrderGoodsSkuGroupBy) StringX(ctx context.Context) string {
	v, err := ogsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ogsgb *OrderGoodsSkuGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ogsgb.fields) > 1 {
		return nil, errors.New("ent: OrderGoodsSkuGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ogsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ogsgb *OrderGoodsSkuGroupBy) IntsX(ctx context.Context) []int {
	v, err := ogsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ogsgb *OrderGoodsSkuGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ogsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ordergoodssku.Label}
	default:
		err = fmt.Errorf("ent: OrderGoodsSkuGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ogsgb *OrderGoodsSkuGroupBy) IntX(ctx context.Context) int {
	v, err := ogsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ogsgb *OrderGoodsSkuGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ogsgb.fields) > 1 {
		return nil, errors.New("ent: OrderGoodsSkuGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ogsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ogsgb *OrderGoodsSkuGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ogsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ogsgb *OrderGoodsSkuGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ogsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ordergoodssku.Label}
	default:
		err = fmt.Errorf("ent: OrderGoodsSkuGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ogsgb *OrderGoodsSkuGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ogsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ogsgb *OrderGoodsSkuGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ogsgb.fields) > 1 {
		return nil, errors.New("ent: OrderGoodsSkuGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ogsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ogsgb *OrderGoodsSkuGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ogsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ogsgb *OrderGoodsSkuGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ogsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ordergoodssku.Label}
	default:
		err = fmt.Errorf("ent: OrderGoodsSkuGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ogsgb *OrderGoodsSkuGroupBy) BoolX(ctx context.Context) bool {
	v, err := ogsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ogsgb *OrderGoodsSkuGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ogsgb.fields {
		if !ordergoodssku.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ogsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ogsgb *OrderGoodsSkuGroupBy) sqlQuery() *sql.Selector {
	selector := ogsgb.sql
	columns := make([]string, 0, len(ogsgb.fields)+len(ogsgb.fns))
	columns = append(columns, ogsgb.fields...)
	for _, fn := range ogsgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(ogsgb.fields...)
}

// OrderGoodsSkuSelect is the builder for selecting fields of OrderGoodsSku entities.
type OrderGoodsSkuSelect struct {
	*OrderGoodsSkuQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ogss *OrderGoodsSkuSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ogss.prepareQuery(ctx); err != nil {
		return err
	}
	ogss.sql = ogss.OrderGoodsSkuQuery.sqlQuery(ctx)
	return ogss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ogss *OrderGoodsSkuSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ogss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ogss *OrderGoodsSkuSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ogss.fields) > 1 {
		return nil, errors.New("ent: OrderGoodsSkuSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ogss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ogss *OrderGoodsSkuSelect) StringsX(ctx context.Context) []string {
	v, err := ogss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ogss *OrderGoodsSkuSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ogss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ordergoodssku.Label}
	default:
		err = fmt.Errorf("ent: OrderGoodsSkuSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ogss *OrderGoodsSkuSelect) StringX(ctx context.Context) string {
	v, err := ogss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ogss *OrderGoodsSkuSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ogss.fields) > 1 {
		return nil, errors.New("ent: OrderGoodsSkuSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ogss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ogss *OrderGoodsSkuSelect) IntsX(ctx context.Context) []int {
	v, err := ogss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ogss *OrderGoodsSkuSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ogss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ordergoodssku.Label}
	default:
		err = fmt.Errorf("ent: OrderGoodsSkuSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ogss *OrderGoodsSkuSelect) IntX(ctx context.Context) int {
	v, err := ogss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ogss *OrderGoodsSkuSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ogss.fields) > 1 {
		return nil, errors.New("ent: OrderGoodsSkuSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ogss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ogss *OrderGoodsSkuSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ogss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ogss *OrderGoodsSkuSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ogss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ordergoodssku.Label}
	default:
		err = fmt.Errorf("ent: OrderGoodsSkuSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ogss *OrderGoodsSkuSelect) Float64X(ctx context.Context) float64 {
	v, err := ogss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ogss *OrderGoodsSkuSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ogss.fields) > 1 {
		return nil, errors.New("ent: OrderGoodsSkuSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ogss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ogss *OrderGoodsSkuSelect) BoolsX(ctx context.Context) []bool {
	v, err := ogss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ogss *OrderGoodsSkuSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ogss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ordergoodssku.Label}
	default:
		err = fmt.Errorf("ent: OrderGoodsSkuSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ogss *OrderGoodsSkuSelect) BoolX(ctx context.Context) bool {
	v, err := ogss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ogss *OrderGoodsSkuSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ogss.sqlQuery().Query()
	if err := ogss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ogss *OrderGoodsSkuSelect) sqlQuery() sql.Querier {
	selector := ogss.sql
	selector.Select(selector.Columns(ogss.fields...)...)
	return selector
}