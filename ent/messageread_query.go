// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"ice-mall/ent/messageread"
	"ice-mall/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageReadQuery is the builder for querying MessageRead entities.
type MessageReadQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MessageRead
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MessageReadQuery builder.
func (mrq *MessageReadQuery) Where(ps ...predicate.MessageRead) *MessageReadQuery {
	mrq.predicates = append(mrq.predicates, ps...)
	return mrq
}

// Limit adds a limit step to the query.
func (mrq *MessageReadQuery) Limit(limit int) *MessageReadQuery {
	mrq.limit = &limit
	return mrq
}

// Offset adds an offset step to the query.
func (mrq *MessageReadQuery) Offset(offset int) *MessageReadQuery {
	mrq.offset = &offset
	return mrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mrq *MessageReadQuery) Unique(unique bool) *MessageReadQuery {
	mrq.unique = &unique
	return mrq
}

// Order adds an order step to the query.
func (mrq *MessageReadQuery) Order(o ...OrderFunc) *MessageReadQuery {
	mrq.order = append(mrq.order, o...)
	return mrq
}

// First returns the first MessageRead entity from the query.
// Returns a *NotFoundError when no MessageRead was found.
func (mrq *MessageReadQuery) First(ctx context.Context) (*MessageRead, error) {
	nodes, err := mrq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{messageread.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mrq *MessageReadQuery) FirstX(ctx context.Context) *MessageRead {
	node, err := mrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MessageRead ID from the query.
// Returns a *NotFoundError when no MessageRead ID was found.
func (mrq *MessageReadQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mrq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{messageread.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mrq *MessageReadQuery) FirstIDX(ctx context.Context) int {
	id, err := mrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MessageRead entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MessageRead entity is found.
// Returns a *NotFoundError when no MessageRead entities are found.
func (mrq *MessageReadQuery) Only(ctx context.Context) (*MessageRead, error) {
	nodes, err := mrq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{messageread.Label}
	default:
		return nil, &NotSingularError{messageread.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mrq *MessageReadQuery) OnlyX(ctx context.Context) *MessageRead {
	node, err := mrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MessageRead ID in the query.
// Returns a *NotSingularError when more than one MessageRead ID is found.
// Returns a *NotFoundError when no entities are found.
func (mrq *MessageReadQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mrq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{messageread.Label}
	default:
		err = &NotSingularError{messageread.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mrq *MessageReadQuery) OnlyIDX(ctx context.Context) int {
	id, err := mrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MessageReads.
func (mrq *MessageReadQuery) All(ctx context.Context) ([]*MessageRead, error) {
	if err := mrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mrq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mrq *MessageReadQuery) AllX(ctx context.Context) []*MessageRead {
	nodes, err := mrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MessageRead IDs.
func (mrq *MessageReadQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mrq.Select(messageread.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mrq *MessageReadQuery) IDsX(ctx context.Context) []int {
	ids, err := mrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mrq *MessageReadQuery) Count(ctx context.Context) (int, error) {
	if err := mrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mrq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mrq *MessageReadQuery) CountX(ctx context.Context) int {
	count, err := mrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mrq *MessageReadQuery) Exist(ctx context.Context) (bool, error) {
	if err := mrq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mrq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mrq *MessageReadQuery) ExistX(ctx context.Context) bool {
	exist, err := mrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MessageReadQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mrq *MessageReadQuery) Clone() *MessageReadQuery {
	if mrq == nil {
		return nil
	}
	return &MessageReadQuery{
		config:     mrq.config,
		limit:      mrq.limit,
		offset:     mrq.offset,
		order:      append([]OrderFunc{}, mrq.order...),
		predicates: append([]predicate.MessageRead{}, mrq.predicates...),
		// clone intermediate query.
		sql:    mrq.sql.Clone(),
		path:   mrq.path,
		unique: mrq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (mrq *MessageReadQuery) GroupBy(field string, fields ...string) *MessageReadGroupBy {
	group := &MessageReadGroupBy{config: mrq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mrq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (mrq *MessageReadQuery) Select(fields ...string) *MessageReadSelect {
	mrq.fields = append(mrq.fields, fields...)
	return &MessageReadSelect{MessageReadQuery: mrq}
}

func (mrq *MessageReadQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mrq.fields {
		if !messageread.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mrq.path != nil {
		prev, err := mrq.path(ctx)
		if err != nil {
			return err
		}
		mrq.sql = prev
	}
	return nil
}

func (mrq *MessageReadQuery) sqlAll(ctx context.Context) ([]*MessageRead, error) {
	var (
		nodes = []*MessageRead{}
		_spec = mrq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &MessageRead{config: mrq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, mrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mrq *MessageReadQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mrq.querySpec()
	_spec.Node.Columns = mrq.fields
	if len(mrq.fields) > 0 {
		_spec.Unique = mrq.unique != nil && *mrq.unique
	}
	return sqlgraph.CountNodes(ctx, mrq.driver, _spec)
}

func (mrq *MessageReadQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mrq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mrq *MessageReadQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messageread.Table,
			Columns: messageread.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messageread.FieldID,
			},
		},
		From:   mrq.sql,
		Unique: true,
	}
	if unique := mrq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mrq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messageread.FieldID)
		for i := range fields {
			if fields[i] != messageread.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mrq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mrq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mrq *MessageReadQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mrq.driver.Dialect())
	t1 := builder.Table(messageread.Table)
	columns := mrq.fields
	if len(columns) == 0 {
		columns = messageread.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mrq.sql != nil {
		selector = mrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mrq.unique != nil && *mrq.unique {
		selector.Distinct()
	}
	for _, p := range mrq.predicates {
		p(selector)
	}
	for _, p := range mrq.order {
		p(selector)
	}
	if offset := mrq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mrq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MessageReadGroupBy is the group-by builder for MessageRead entities.
type MessageReadGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mrgb *MessageReadGroupBy) Aggregate(fns ...AggregateFunc) *MessageReadGroupBy {
	mrgb.fns = append(mrgb.fns, fns...)
	return mrgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mrgb *MessageReadGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mrgb.path(ctx)
	if err != nil {
		return err
	}
	mrgb.sql = query
	return mrgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mrgb *MessageReadGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mrgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (mrgb *MessageReadGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mrgb.fields) > 1 {
		return nil, errors.New("ent: MessageReadGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mrgb *MessageReadGroupBy) StringsX(ctx context.Context) []string {
	v, err := mrgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mrgb *MessageReadGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mrgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messageread.Label}
	default:
		err = fmt.Errorf("ent: MessageReadGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mrgb *MessageReadGroupBy) StringX(ctx context.Context) string {
	v, err := mrgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (mrgb *MessageReadGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mrgb.fields) > 1 {
		return nil, errors.New("ent: MessageReadGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mrgb *MessageReadGroupBy) IntsX(ctx context.Context) []int {
	v, err := mrgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mrgb *MessageReadGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mrgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messageread.Label}
	default:
		err = fmt.Errorf("ent: MessageReadGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mrgb *MessageReadGroupBy) IntX(ctx context.Context) int {
	v, err := mrgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (mrgb *MessageReadGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mrgb.fields) > 1 {
		return nil, errors.New("ent: MessageReadGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mrgb *MessageReadGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mrgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mrgb *MessageReadGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mrgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messageread.Label}
	default:
		err = fmt.Errorf("ent: MessageReadGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mrgb *MessageReadGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mrgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (mrgb *MessageReadGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mrgb.fields) > 1 {
		return nil, errors.New("ent: MessageReadGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mrgb *MessageReadGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mrgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mrgb *MessageReadGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mrgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messageread.Label}
	default:
		err = fmt.Errorf("ent: MessageReadGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mrgb *MessageReadGroupBy) BoolX(ctx context.Context) bool {
	v, err := mrgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mrgb *MessageReadGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mrgb.fields {
		if !messageread.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mrgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mrgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mrgb *MessageReadGroupBy) sqlQuery() *sql.Selector {
	selector := mrgb.sql.Select()
	aggregation := make([]string, 0, len(mrgb.fns))
	for _, fn := range mrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mrgb.fields)+len(mrgb.fns))
		for _, f := range mrgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mrgb.fields...)...)
}

// MessageReadSelect is the builder for selecting fields of MessageRead entities.
type MessageReadSelect struct {
	*MessageReadQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mrs *MessageReadSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mrs.prepareQuery(ctx); err != nil {
		return err
	}
	mrs.sql = mrs.MessageReadQuery.sqlQuery(ctx)
	return mrs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mrs *MessageReadSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mrs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (mrs *MessageReadSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mrs.fields) > 1 {
		return nil, errors.New("ent: MessageReadSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mrs *MessageReadSelect) StringsX(ctx context.Context) []string {
	v, err := mrs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (mrs *MessageReadSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mrs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messageread.Label}
	default:
		err = fmt.Errorf("ent: MessageReadSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mrs *MessageReadSelect) StringX(ctx context.Context) string {
	v, err := mrs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (mrs *MessageReadSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mrs.fields) > 1 {
		return nil, errors.New("ent: MessageReadSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mrs *MessageReadSelect) IntsX(ctx context.Context) []int {
	v, err := mrs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (mrs *MessageReadSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mrs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messageread.Label}
	default:
		err = fmt.Errorf("ent: MessageReadSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mrs *MessageReadSelect) IntX(ctx context.Context) int {
	v, err := mrs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (mrs *MessageReadSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mrs.fields) > 1 {
		return nil, errors.New("ent: MessageReadSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mrs *MessageReadSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mrs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (mrs *MessageReadSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mrs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messageread.Label}
	default:
		err = fmt.Errorf("ent: MessageReadSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mrs *MessageReadSelect) Float64X(ctx context.Context) float64 {
	v, err := mrs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (mrs *MessageReadSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mrs.fields) > 1 {
		return nil, errors.New("ent: MessageReadSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mrs *MessageReadSelect) BoolsX(ctx context.Context) []bool {
	v, err := mrs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (mrs *MessageReadSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mrs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messageread.Label}
	default:
		err = fmt.Errorf("ent: MessageReadSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mrs *MessageReadSelect) BoolX(ctx context.Context) bool {
	v, err := mrs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mrs *MessageReadSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mrs.sql.Query()
	if err := mrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}