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
	"github.com/a20070322/shop-go/ent/customer"
	"github.com/a20070322/shop-go/ent/ordergoodssku"
	"github.com/a20070322/shop-go/ent/orderinfo"
	"github.com/a20070322/shop-go/ent/predicate"
)

// OrderInfoQuery is the builder for querying OrderInfo entities.
type OrderInfoQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrderInfo
	// eager-loading edges.
	withCustomer      *CustomerQuery
	withOrderGoodsSku *OrderGoodsSkuQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderInfoQuery builder.
func (oiq *OrderInfoQuery) Where(ps ...predicate.OrderInfo) *OrderInfoQuery {
	oiq.predicates = append(oiq.predicates, ps...)
	return oiq
}

// Limit adds a limit step to the query.
func (oiq *OrderInfoQuery) Limit(limit int) *OrderInfoQuery {
	oiq.limit = &limit
	return oiq
}

// Offset adds an offset step to the query.
func (oiq *OrderInfoQuery) Offset(offset int) *OrderInfoQuery {
	oiq.offset = &offset
	return oiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oiq *OrderInfoQuery) Unique(unique bool) *OrderInfoQuery {
	oiq.unique = &unique
	return oiq
}

// Order adds an order step to the query.
func (oiq *OrderInfoQuery) Order(o ...OrderFunc) *OrderInfoQuery {
	oiq.order = append(oiq.order, o...)
	return oiq
}

// QueryCustomer chains the current query on the "customer" edge.
func (oiq *OrderInfoQuery) QueryCustomer() *CustomerQuery {
	query := &CustomerQuery{config: oiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderinfo.Table, orderinfo.FieldID, selector),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderinfo.CustomerTable, orderinfo.CustomerColumn),
		)
		fromU = sqlgraph.SetNeighbors(oiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderGoodsSku chains the current query on the "order_goods_sku" edge.
func (oiq *OrderInfoQuery) QueryOrderGoodsSku() *OrderGoodsSkuQuery {
	query := &OrderGoodsSkuQuery{config: oiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderinfo.Table, orderinfo.FieldID, selector),
			sqlgraph.To(ordergoodssku.Table, ordergoodssku.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, orderinfo.OrderGoodsSkuTable, orderinfo.OrderGoodsSkuColumn),
		)
		fromU = sqlgraph.SetNeighbors(oiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrderInfo entity from the query.
// Returns a *NotFoundError when no OrderInfo was found.
func (oiq *OrderInfoQuery) First(ctx context.Context) (*OrderInfo, error) {
	nodes, err := oiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oiq *OrderInfoQuery) FirstX(ctx context.Context) *OrderInfo {
	node, err := oiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderInfo ID from the query.
// Returns a *NotFoundError when no OrderInfo ID was found.
func (oiq *OrderInfoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oiq *OrderInfoQuery) FirstIDX(ctx context.Context) int {
	id, err := oiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one OrderInfo entity is not found.
// Returns a *NotFoundError when no OrderInfo entities are found.
func (oiq *OrderInfoQuery) Only(ctx context.Context) (*OrderInfo, error) {
	nodes, err := oiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderinfo.Label}
	default:
		return nil, &NotSingularError{orderinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oiq *OrderInfoQuery) OnlyX(ctx context.Context) *OrderInfo {
	node, err := oiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderInfo ID in the query.
// Returns a *NotSingularError when exactly one OrderInfo ID is not found.
// Returns a *NotFoundError when no entities are found.
func (oiq *OrderInfoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderinfo.Label}
	default:
		err = &NotSingularError{orderinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oiq *OrderInfoQuery) OnlyIDX(ctx context.Context) int {
	id, err := oiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderInfos.
func (oiq *OrderInfoQuery) All(ctx context.Context) ([]*OrderInfo, error) {
	if err := oiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oiq *OrderInfoQuery) AllX(ctx context.Context) []*OrderInfo {
	nodes, err := oiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderInfo IDs.
func (oiq *OrderInfoQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oiq.Select(orderinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oiq *OrderInfoQuery) IDsX(ctx context.Context) []int {
	ids, err := oiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oiq *OrderInfoQuery) Count(ctx context.Context) (int, error) {
	if err := oiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oiq *OrderInfoQuery) CountX(ctx context.Context) int {
	count, err := oiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oiq *OrderInfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := oiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oiq *OrderInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := oiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oiq *OrderInfoQuery) Clone() *OrderInfoQuery {
	if oiq == nil {
		return nil
	}
	return &OrderInfoQuery{
		config:            oiq.config,
		limit:             oiq.limit,
		offset:            oiq.offset,
		order:             append([]OrderFunc{}, oiq.order...),
		predicates:        append([]predicate.OrderInfo{}, oiq.predicates...),
		withCustomer:      oiq.withCustomer.Clone(),
		withOrderGoodsSku: oiq.withOrderGoodsSku.Clone(),
		// clone intermediate query.
		sql:  oiq.sql.Clone(),
		path: oiq.path,
	}
}

// WithCustomer tells the query-builder to eager-load the nodes that are connected to
// the "customer" edge. The optional arguments are used to configure the query builder of the edge.
func (oiq *OrderInfoQuery) WithCustomer(opts ...func(*CustomerQuery)) *OrderInfoQuery {
	query := &CustomerQuery{config: oiq.config}
	for _, opt := range opts {
		opt(query)
	}
	oiq.withCustomer = query
	return oiq
}

// WithOrderGoodsSku tells the query-builder to eager-load the nodes that are connected to
// the "order_goods_sku" edge. The optional arguments are used to configure the query builder of the edge.
func (oiq *OrderInfoQuery) WithOrderGoodsSku(opts ...func(*OrderGoodsSkuQuery)) *OrderInfoQuery {
	query := &OrderGoodsSkuQuery{config: oiq.config}
	for _, opt := range opts {
		opt(query)
	}
	oiq.withOrderGoodsSku = query
	return oiq
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
//	client.OrderInfo.Query().
//		GroupBy(orderinfo.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (oiq *OrderInfoQuery) GroupBy(field string, fields ...string) *OrderInfoGroupBy {
	group := &OrderInfoGroupBy{config: oiq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oiq.sqlQuery(ctx), nil
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
//	client.OrderInfo.Query().
//		Select(orderinfo.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (oiq *OrderInfoQuery) Select(field string, fields ...string) *OrderInfoSelect {
	oiq.fields = append([]string{field}, fields...)
	return &OrderInfoSelect{OrderInfoQuery: oiq}
}

func (oiq *OrderInfoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oiq.fields {
		if !orderinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oiq.path != nil {
		prev, err := oiq.path(ctx)
		if err != nil {
			return err
		}
		oiq.sql = prev
	}
	return nil
}

func (oiq *OrderInfoQuery) sqlAll(ctx context.Context) ([]*OrderInfo, error) {
	var (
		nodes       = []*OrderInfo{}
		withFKs     = oiq.withFKs
		_spec       = oiq.querySpec()
		loadedTypes = [2]bool{
			oiq.withCustomer != nil,
			oiq.withOrderGoodsSku != nil,
		}
	)
	if oiq.withCustomer != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, orderinfo.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &OrderInfo{config: oiq.config}
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
	if err := sqlgraph.QueryNodes(ctx, oiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := oiq.withCustomer; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrderInfo)
		for i := range nodes {
			if nodes[i].customer_order_info == nil {
				continue
			}
			fk := *nodes[i].customer_order_info
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(customer.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "customer_order_info" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Customer = n
			}
		}
	}

	if query := oiq.withOrderGoodsSku; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*OrderInfo)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OrderGoodsSku = []*OrderGoodsSku{}
		}
		query.withFKs = true
		query.Where(predicate.OrderGoodsSku(func(s *sql.Selector) {
			s.Where(sql.InValues(orderinfo.OrderGoodsSkuColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.order_info_order_goods_sku
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "order_info_order_goods_sku" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "order_info_order_goods_sku" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OrderGoodsSku = append(node.Edges.OrderGoodsSku, n)
		}
	}

	return nodes, nil
}

func (oiq *OrderInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oiq.querySpec()
	return sqlgraph.CountNodes(ctx, oiq.driver, _spec)
}

func (oiq *OrderInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (oiq *OrderInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderinfo.Table,
			Columns: orderinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderinfo.FieldID,
			},
		},
		From:   oiq.sql,
		Unique: true,
	}
	if unique := oiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderinfo.FieldID)
		for i := range fields {
			if fields[i] != orderinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oiq *OrderInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oiq.driver.Dialect())
	t1 := builder.Table(orderinfo.Table)
	selector := builder.Select(t1.Columns(orderinfo.Columns...)...).From(t1)
	if oiq.sql != nil {
		selector = oiq.sql
		selector.Select(selector.Columns(orderinfo.Columns...)...)
	}
	for _, p := range oiq.predicates {
		p(selector)
	}
	for _, p := range oiq.order {
		p(selector)
	}
	if offset := oiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderInfoGroupBy is the group-by builder for OrderInfo entities.
type OrderInfoGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oigb *OrderInfoGroupBy) Aggregate(fns ...AggregateFunc) *OrderInfoGroupBy {
	oigb.fns = append(oigb.fns, fns...)
	return oigb
}

// Scan applies the group-by query and scans the result into the given value.
func (oigb *OrderInfoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := oigb.path(ctx)
	if err != nil {
		return err
	}
	oigb.sql = query
	return oigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (oigb *OrderInfoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := oigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (oigb *OrderInfoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(oigb.fields) > 1 {
		return nil, errors.New("ent: OrderInfoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := oigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (oigb *OrderInfoGroupBy) StringsX(ctx context.Context) []string {
	v, err := oigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oigb *OrderInfoGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = oigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderinfo.Label}
	default:
		err = fmt.Errorf("ent: OrderInfoGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (oigb *OrderInfoGroupBy) StringX(ctx context.Context) string {
	v, err := oigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (oigb *OrderInfoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(oigb.fields) > 1 {
		return nil, errors.New("ent: OrderInfoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := oigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (oigb *OrderInfoGroupBy) IntsX(ctx context.Context) []int {
	v, err := oigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oigb *OrderInfoGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = oigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderinfo.Label}
	default:
		err = fmt.Errorf("ent: OrderInfoGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (oigb *OrderInfoGroupBy) IntX(ctx context.Context) int {
	v, err := oigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (oigb *OrderInfoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(oigb.fields) > 1 {
		return nil, errors.New("ent: OrderInfoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := oigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (oigb *OrderInfoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := oigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oigb *OrderInfoGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = oigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderinfo.Label}
	default:
		err = fmt.Errorf("ent: OrderInfoGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (oigb *OrderInfoGroupBy) Float64X(ctx context.Context) float64 {
	v, err := oigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (oigb *OrderInfoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(oigb.fields) > 1 {
		return nil, errors.New("ent: OrderInfoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := oigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (oigb *OrderInfoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := oigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oigb *OrderInfoGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = oigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderinfo.Label}
	default:
		err = fmt.Errorf("ent: OrderInfoGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (oigb *OrderInfoGroupBy) BoolX(ctx context.Context) bool {
	v, err := oigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oigb *OrderInfoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range oigb.fields {
		if !orderinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := oigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oigb *OrderInfoGroupBy) sqlQuery() *sql.Selector {
	selector := oigb.sql
	columns := make([]string, 0, len(oigb.fields)+len(oigb.fns))
	columns = append(columns, oigb.fields...)
	for _, fn := range oigb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(oigb.fields...)
}

// OrderInfoSelect is the builder for selecting fields of OrderInfo entities.
type OrderInfoSelect struct {
	*OrderInfoQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ois *OrderInfoSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ois.prepareQuery(ctx); err != nil {
		return err
	}
	ois.sql = ois.OrderInfoQuery.sqlQuery(ctx)
	return ois.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ois *OrderInfoSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ois.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ois *OrderInfoSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ois.fields) > 1 {
		return nil, errors.New("ent: OrderInfoSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ois.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ois *OrderInfoSelect) StringsX(ctx context.Context) []string {
	v, err := ois.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ois *OrderInfoSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ois.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderinfo.Label}
	default:
		err = fmt.Errorf("ent: OrderInfoSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ois *OrderInfoSelect) StringX(ctx context.Context) string {
	v, err := ois.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ois *OrderInfoSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ois.fields) > 1 {
		return nil, errors.New("ent: OrderInfoSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ois.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ois *OrderInfoSelect) IntsX(ctx context.Context) []int {
	v, err := ois.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ois *OrderInfoSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ois.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderinfo.Label}
	default:
		err = fmt.Errorf("ent: OrderInfoSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ois *OrderInfoSelect) IntX(ctx context.Context) int {
	v, err := ois.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ois *OrderInfoSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ois.fields) > 1 {
		return nil, errors.New("ent: OrderInfoSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ois.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ois *OrderInfoSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ois.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ois *OrderInfoSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ois.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderinfo.Label}
	default:
		err = fmt.Errorf("ent: OrderInfoSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ois *OrderInfoSelect) Float64X(ctx context.Context) float64 {
	v, err := ois.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ois *OrderInfoSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ois.fields) > 1 {
		return nil, errors.New("ent: OrderInfoSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ois.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ois *OrderInfoSelect) BoolsX(ctx context.Context) []bool {
	v, err := ois.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ois *OrderInfoSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ois.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderinfo.Label}
	default:
		err = fmt.Errorf("ent: OrderInfoSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ois *OrderInfoSelect) BoolX(ctx context.Context) bool {
	v, err := ois.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ois *OrderInfoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ois.sqlQuery().Query()
	if err := ois.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ois *OrderInfoSelect) sqlQuery() sql.Querier {
	selector := ois.sql
	selector.Select(selector.Columns(ois.fields...)...)
	return selector
}
