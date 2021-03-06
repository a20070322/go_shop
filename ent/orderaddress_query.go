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
	"github.com/a20070322/shop-go/ent/orderaddress"
	"github.com/a20070322/shop-go/ent/orderinfo"
	"github.com/a20070322/shop-go/ent/predicate"
)

// OrderAddressQuery is the builder for querying OrderAddress entities.
type OrderAddressQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrderAddress
	// eager-loading edges.
	withOrderInfo *OrderInfoQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderAddressQuery builder.
func (oaq *OrderAddressQuery) Where(ps ...predicate.OrderAddress) *OrderAddressQuery {
	oaq.predicates = append(oaq.predicates, ps...)
	return oaq
}

// Limit adds a limit step to the query.
func (oaq *OrderAddressQuery) Limit(limit int) *OrderAddressQuery {
	oaq.limit = &limit
	return oaq
}

// Offset adds an offset step to the query.
func (oaq *OrderAddressQuery) Offset(offset int) *OrderAddressQuery {
	oaq.offset = &offset
	return oaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oaq *OrderAddressQuery) Unique(unique bool) *OrderAddressQuery {
	oaq.unique = &unique
	return oaq
}

// Order adds an order step to the query.
func (oaq *OrderAddressQuery) Order(o ...OrderFunc) *OrderAddressQuery {
	oaq.order = append(oaq.order, o...)
	return oaq
}

// QueryOrderInfo chains the current query on the "order_info" edge.
func (oaq *OrderAddressQuery) QueryOrderInfo() *OrderInfoQuery {
	query := &OrderInfoQuery{config: oaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderaddress.Table, orderaddress.FieldID, selector),
			sqlgraph.To(orderinfo.Table, orderinfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderaddress.OrderInfoTable, orderaddress.OrderInfoColumn),
		)
		fromU = sqlgraph.SetNeighbors(oaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrderAddress entity from the query.
// Returns a *NotFoundError when no OrderAddress was found.
func (oaq *OrderAddressQuery) First(ctx context.Context) (*OrderAddress, error) {
	nodes, err := oaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderaddress.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oaq *OrderAddressQuery) FirstX(ctx context.Context) *OrderAddress {
	node, err := oaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderAddress ID from the query.
// Returns a *NotFoundError when no OrderAddress ID was found.
func (oaq *OrderAddressQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderaddress.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oaq *OrderAddressQuery) FirstIDX(ctx context.Context) int {
	id, err := oaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderAddress entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one OrderAddress entity is not found.
// Returns a *NotFoundError when no OrderAddress entities are found.
func (oaq *OrderAddressQuery) Only(ctx context.Context) (*OrderAddress, error) {
	nodes, err := oaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderaddress.Label}
	default:
		return nil, &NotSingularError{orderaddress.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oaq *OrderAddressQuery) OnlyX(ctx context.Context) *OrderAddress {
	node, err := oaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderAddress ID in the query.
// Returns a *NotSingularError when exactly one OrderAddress ID is not found.
// Returns a *NotFoundError when no entities are found.
func (oaq *OrderAddressQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderaddress.Label}
	default:
		err = &NotSingularError{orderaddress.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oaq *OrderAddressQuery) OnlyIDX(ctx context.Context) int {
	id, err := oaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderAddresses.
func (oaq *OrderAddressQuery) All(ctx context.Context) ([]*OrderAddress, error) {
	if err := oaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oaq *OrderAddressQuery) AllX(ctx context.Context) []*OrderAddress {
	nodes, err := oaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderAddress IDs.
func (oaq *OrderAddressQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oaq.Select(orderaddress.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oaq *OrderAddressQuery) IDsX(ctx context.Context) []int {
	ids, err := oaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oaq *OrderAddressQuery) Count(ctx context.Context) (int, error) {
	if err := oaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oaq *OrderAddressQuery) CountX(ctx context.Context) int {
	count, err := oaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oaq *OrderAddressQuery) Exist(ctx context.Context) (bool, error) {
	if err := oaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oaq *OrderAddressQuery) ExistX(ctx context.Context) bool {
	exist, err := oaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderAddressQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oaq *OrderAddressQuery) Clone() *OrderAddressQuery {
	if oaq == nil {
		return nil
	}
	return &OrderAddressQuery{
		config:        oaq.config,
		limit:         oaq.limit,
		offset:        oaq.offset,
		order:         append([]OrderFunc{}, oaq.order...),
		predicates:    append([]predicate.OrderAddress{}, oaq.predicates...),
		withOrderInfo: oaq.withOrderInfo.Clone(),
		// clone intermediate query.
		sql:  oaq.sql.Clone(),
		path: oaq.path,
	}
}

// WithOrderInfo tells the query-builder to eager-load the nodes that are connected to
// the "order_info" edge. The optional arguments are used to configure the query builder of the edge.
func (oaq *OrderAddressQuery) WithOrderInfo(opts ...func(*OrderInfoQuery)) *OrderAddressQuery {
	query := &OrderInfoQuery{config: oaq.config}
	for _, opt := range opts {
		opt(query)
	}
	oaq.withOrderInfo = query
	return oaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrderAddress.Query().
//		GroupBy(orderaddress.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (oaq *OrderAddressQuery) GroupBy(field string, fields ...string) *OrderAddressGroupBy {
	group := &OrderAddressGroupBy{config: oaq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oaq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.OrderAddress.Query().
//		Select(orderaddress.FieldName).
//		Scan(ctx, &v)
//
func (oaq *OrderAddressQuery) Select(field string, fields ...string) *OrderAddressSelect {
	oaq.fields = append([]string{field}, fields...)
	return &OrderAddressSelect{OrderAddressQuery: oaq}
}

func (oaq *OrderAddressQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oaq.fields {
		if !orderaddress.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oaq.path != nil {
		prev, err := oaq.path(ctx)
		if err != nil {
			return err
		}
		oaq.sql = prev
	}
	return nil
}

func (oaq *OrderAddressQuery) sqlAll(ctx context.Context) ([]*OrderAddress, error) {
	var (
		nodes       = []*OrderAddress{}
		withFKs     = oaq.withFKs
		_spec       = oaq.querySpec()
		loadedTypes = [1]bool{
			oaq.withOrderInfo != nil,
		}
	)
	if oaq.withOrderInfo != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, orderaddress.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &OrderAddress{config: oaq.config}
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
	if err := sqlgraph.QueryNodes(ctx, oaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := oaq.withOrderInfo; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrderAddress)
		for i := range nodes {
			if nodes[i].order_info_order_address == nil {
				continue
			}
			fk := *nodes[i].order_info_order_address
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
				return nil, fmt.Errorf(`unexpected foreign-key "order_info_order_address" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OrderInfo = n
			}
		}
	}

	return nodes, nil
}

func (oaq *OrderAddressQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oaq.querySpec()
	return sqlgraph.CountNodes(ctx, oaq.driver, _spec)
}

func (oaq *OrderAddressQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (oaq *OrderAddressQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderaddress.Table,
			Columns: orderaddress.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderaddress.FieldID,
			},
		},
		From:   oaq.sql,
		Unique: true,
	}
	if unique := oaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderaddress.FieldID)
		for i := range fields {
			if fields[i] != orderaddress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oaq *OrderAddressQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oaq.driver.Dialect())
	t1 := builder.Table(orderaddress.Table)
	selector := builder.Select(t1.Columns(orderaddress.Columns...)...).From(t1)
	if oaq.sql != nil {
		selector = oaq.sql
		selector.Select(selector.Columns(orderaddress.Columns...)...)
	}
	for _, p := range oaq.predicates {
		p(selector)
	}
	for _, p := range oaq.order {
		p(selector)
	}
	if offset := oaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderAddressGroupBy is the group-by builder for OrderAddress entities.
type OrderAddressGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oagb *OrderAddressGroupBy) Aggregate(fns ...AggregateFunc) *OrderAddressGroupBy {
	oagb.fns = append(oagb.fns, fns...)
	return oagb
}

// Scan applies the group-by query and scans the result into the given value.
func (oagb *OrderAddressGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := oagb.path(ctx)
	if err != nil {
		return err
	}
	oagb.sql = query
	return oagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (oagb *OrderAddressGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := oagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (oagb *OrderAddressGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(oagb.fields) > 1 {
		return nil, errors.New("ent: OrderAddressGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := oagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (oagb *OrderAddressGroupBy) StringsX(ctx context.Context) []string {
	v, err := oagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oagb *OrderAddressGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = oagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderaddress.Label}
	default:
		err = fmt.Errorf("ent: OrderAddressGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (oagb *OrderAddressGroupBy) StringX(ctx context.Context) string {
	v, err := oagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (oagb *OrderAddressGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(oagb.fields) > 1 {
		return nil, errors.New("ent: OrderAddressGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := oagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (oagb *OrderAddressGroupBy) IntsX(ctx context.Context) []int {
	v, err := oagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oagb *OrderAddressGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = oagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderaddress.Label}
	default:
		err = fmt.Errorf("ent: OrderAddressGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (oagb *OrderAddressGroupBy) IntX(ctx context.Context) int {
	v, err := oagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (oagb *OrderAddressGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(oagb.fields) > 1 {
		return nil, errors.New("ent: OrderAddressGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := oagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (oagb *OrderAddressGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := oagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oagb *OrderAddressGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = oagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderaddress.Label}
	default:
		err = fmt.Errorf("ent: OrderAddressGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (oagb *OrderAddressGroupBy) Float64X(ctx context.Context) float64 {
	v, err := oagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (oagb *OrderAddressGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(oagb.fields) > 1 {
		return nil, errors.New("ent: OrderAddressGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := oagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (oagb *OrderAddressGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := oagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oagb *OrderAddressGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = oagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderaddress.Label}
	default:
		err = fmt.Errorf("ent: OrderAddressGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (oagb *OrderAddressGroupBy) BoolX(ctx context.Context) bool {
	v, err := oagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oagb *OrderAddressGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range oagb.fields {
		if !orderaddress.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := oagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oagb *OrderAddressGroupBy) sqlQuery() *sql.Selector {
	selector := oagb.sql
	columns := make([]string, 0, len(oagb.fields)+len(oagb.fns))
	columns = append(columns, oagb.fields...)
	for _, fn := range oagb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(oagb.fields...)
}

// OrderAddressSelect is the builder for selecting fields of OrderAddress entities.
type OrderAddressSelect struct {
	*OrderAddressQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (oas *OrderAddressSelect) Scan(ctx context.Context, v interface{}) error {
	if err := oas.prepareQuery(ctx); err != nil {
		return err
	}
	oas.sql = oas.OrderAddressQuery.sqlQuery(ctx)
	return oas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (oas *OrderAddressSelect) ScanX(ctx context.Context, v interface{}) {
	if err := oas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (oas *OrderAddressSelect) Strings(ctx context.Context) ([]string, error) {
	if len(oas.fields) > 1 {
		return nil, errors.New("ent: OrderAddressSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := oas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (oas *OrderAddressSelect) StringsX(ctx context.Context) []string {
	v, err := oas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (oas *OrderAddressSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = oas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderaddress.Label}
	default:
		err = fmt.Errorf("ent: OrderAddressSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (oas *OrderAddressSelect) StringX(ctx context.Context) string {
	v, err := oas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (oas *OrderAddressSelect) Ints(ctx context.Context) ([]int, error) {
	if len(oas.fields) > 1 {
		return nil, errors.New("ent: OrderAddressSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := oas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (oas *OrderAddressSelect) IntsX(ctx context.Context) []int {
	v, err := oas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (oas *OrderAddressSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = oas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderaddress.Label}
	default:
		err = fmt.Errorf("ent: OrderAddressSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (oas *OrderAddressSelect) IntX(ctx context.Context) int {
	v, err := oas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (oas *OrderAddressSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(oas.fields) > 1 {
		return nil, errors.New("ent: OrderAddressSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := oas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (oas *OrderAddressSelect) Float64sX(ctx context.Context) []float64 {
	v, err := oas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (oas *OrderAddressSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = oas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderaddress.Label}
	default:
		err = fmt.Errorf("ent: OrderAddressSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (oas *OrderAddressSelect) Float64X(ctx context.Context) float64 {
	v, err := oas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (oas *OrderAddressSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(oas.fields) > 1 {
		return nil, errors.New("ent: OrderAddressSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := oas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (oas *OrderAddressSelect) BoolsX(ctx context.Context) []bool {
	v, err := oas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (oas *OrderAddressSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = oas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderaddress.Label}
	default:
		err = fmt.Errorf("ent: OrderAddressSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (oas *OrderAddressSelect) BoolX(ctx context.Context) bool {
	v, err := oas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oas *OrderAddressSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oas.sqlQuery().Query()
	if err := oas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oas *OrderAddressSelect) sqlQuery() sql.Querier {
	selector := oas.sql
	selector.Select(selector.Columns(oas.fields...)...)
	return selector
}
