// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"ice-mall/ent/predicate"
	"ice-mall/ent/systemconfig"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SystemConfigQuery is the builder for querying SystemConfig entities.
type SystemConfigQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SystemConfig
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SystemConfigQuery builder.
func (scq *SystemConfigQuery) Where(ps ...predicate.SystemConfig) *SystemConfigQuery {
	scq.predicates = append(scq.predicates, ps...)
	return scq
}

// Limit adds a limit step to the query.
func (scq *SystemConfigQuery) Limit(limit int) *SystemConfigQuery {
	scq.limit = &limit
	return scq
}

// Offset adds an offset step to the query.
func (scq *SystemConfigQuery) Offset(offset int) *SystemConfigQuery {
	scq.offset = &offset
	return scq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (scq *SystemConfigQuery) Unique(unique bool) *SystemConfigQuery {
	scq.unique = &unique
	return scq
}

// Order adds an order step to the query.
func (scq *SystemConfigQuery) Order(o ...OrderFunc) *SystemConfigQuery {
	scq.order = append(scq.order, o...)
	return scq
}

// First returns the first SystemConfig entity from the query.
// Returns a *NotFoundError when no SystemConfig was found.
func (scq *SystemConfigQuery) First(ctx context.Context) (*SystemConfig, error) {
	nodes, err := scq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{systemconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (scq *SystemConfigQuery) FirstX(ctx context.Context) *SystemConfig {
	node, err := scq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SystemConfig ID from the query.
// Returns a *NotFoundError when no SystemConfig ID was found.
func (scq *SystemConfigQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = scq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{systemconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (scq *SystemConfigQuery) FirstIDX(ctx context.Context) int {
	id, err := scq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SystemConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SystemConfig entity is found.
// Returns a *NotFoundError when no SystemConfig entities are found.
func (scq *SystemConfigQuery) Only(ctx context.Context) (*SystemConfig, error) {
	nodes, err := scq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{systemconfig.Label}
	default:
		return nil, &NotSingularError{systemconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (scq *SystemConfigQuery) OnlyX(ctx context.Context) *SystemConfig {
	node, err := scq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SystemConfig ID in the query.
// Returns a *NotSingularError when more than one SystemConfig ID is found.
// Returns a *NotFoundError when no entities are found.
func (scq *SystemConfigQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = scq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{systemconfig.Label}
	default:
		err = &NotSingularError{systemconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (scq *SystemConfigQuery) OnlyIDX(ctx context.Context) int {
	id, err := scq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SystemConfigs.
func (scq *SystemConfigQuery) All(ctx context.Context) ([]*SystemConfig, error) {
	if err := scq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return scq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (scq *SystemConfigQuery) AllX(ctx context.Context) []*SystemConfig {
	nodes, err := scq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SystemConfig IDs.
func (scq *SystemConfigQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := scq.Select(systemconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (scq *SystemConfigQuery) IDsX(ctx context.Context) []int {
	ids, err := scq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (scq *SystemConfigQuery) Count(ctx context.Context) (int, error) {
	if err := scq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return scq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (scq *SystemConfigQuery) CountX(ctx context.Context) int {
	count, err := scq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (scq *SystemConfigQuery) Exist(ctx context.Context) (bool, error) {
	if err := scq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return scq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (scq *SystemConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := scq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SystemConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (scq *SystemConfigQuery) Clone() *SystemConfigQuery {
	if scq == nil {
		return nil
	}
	return &SystemConfigQuery{
		config:     scq.config,
		limit:      scq.limit,
		offset:     scq.offset,
		order:      append([]OrderFunc{}, scq.order...),
		predicates: append([]predicate.SystemConfig{}, scq.predicates...),
		// clone intermediate query.
		sql:    scq.sql.Clone(),
		path:   scq.path,
		unique: scq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Rule string `json:"rule,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SystemConfig.Query().
//		GroupBy(systemconfig.FieldRule).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (scq *SystemConfigQuery) GroupBy(field string, fields ...string) *SystemConfigGroupBy {
	group := &SystemConfigGroupBy{config: scq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return scq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Rule string `json:"rule,omitempty"`
//	}
//
//	client.SystemConfig.Query().
//		Select(systemconfig.FieldRule).
//		Scan(ctx, &v)
//
func (scq *SystemConfigQuery) Select(fields ...string) *SystemConfigSelect {
	scq.fields = append(scq.fields, fields...)
	return &SystemConfigSelect{SystemConfigQuery: scq}
}

func (scq *SystemConfigQuery) prepareQuery(ctx context.Context) error {
	for _, f := range scq.fields {
		if !systemconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if scq.path != nil {
		prev, err := scq.path(ctx)
		if err != nil {
			return err
		}
		scq.sql = prev
	}
	return nil
}

func (scq *SystemConfigQuery) sqlAll(ctx context.Context) ([]*SystemConfig, error) {
	var (
		nodes = []*SystemConfig{}
		_spec = scq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &SystemConfig{config: scq.config}
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
	if err := sqlgraph.QueryNodes(ctx, scq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (scq *SystemConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := scq.querySpec()
	_spec.Node.Columns = scq.fields
	if len(scq.fields) > 0 {
		_spec.Unique = scq.unique != nil && *scq.unique
	}
	return sqlgraph.CountNodes(ctx, scq.driver, _spec)
}

func (scq *SystemConfigQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := scq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (scq *SystemConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   systemconfig.Table,
			Columns: systemconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: systemconfig.FieldID,
			},
		},
		From:   scq.sql,
		Unique: true,
	}
	if unique := scq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := scq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, systemconfig.FieldID)
		for i := range fields {
			if fields[i] != systemconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := scq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := scq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := scq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := scq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (scq *SystemConfigQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(scq.driver.Dialect())
	t1 := builder.Table(systemconfig.Table)
	columns := scq.fields
	if len(columns) == 0 {
		columns = systemconfig.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if scq.sql != nil {
		selector = scq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if scq.unique != nil && *scq.unique {
		selector.Distinct()
	}
	for _, p := range scq.predicates {
		p(selector)
	}
	for _, p := range scq.order {
		p(selector)
	}
	if offset := scq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := scq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SystemConfigGroupBy is the group-by builder for SystemConfig entities.
type SystemConfigGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (scgb *SystemConfigGroupBy) Aggregate(fns ...AggregateFunc) *SystemConfigGroupBy {
	scgb.fns = append(scgb.fns, fns...)
	return scgb
}

// Scan applies the group-by query and scans the result into the given value.
func (scgb *SystemConfigGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := scgb.path(ctx)
	if err != nil {
		return err
	}
	scgb.sql = query
	return scgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (scgb *SystemConfigGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := scgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (scgb *SystemConfigGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(scgb.fields) > 1 {
		return nil, errors.New("ent: SystemConfigGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := scgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (scgb *SystemConfigGroupBy) StringsX(ctx context.Context) []string {
	v, err := scgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (scgb *SystemConfigGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = scgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systemconfig.Label}
	default:
		err = fmt.Errorf("ent: SystemConfigGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (scgb *SystemConfigGroupBy) StringX(ctx context.Context) string {
	v, err := scgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (scgb *SystemConfigGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(scgb.fields) > 1 {
		return nil, errors.New("ent: SystemConfigGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := scgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (scgb *SystemConfigGroupBy) IntsX(ctx context.Context) []int {
	v, err := scgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (scgb *SystemConfigGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = scgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systemconfig.Label}
	default:
		err = fmt.Errorf("ent: SystemConfigGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (scgb *SystemConfigGroupBy) IntX(ctx context.Context) int {
	v, err := scgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (scgb *SystemConfigGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(scgb.fields) > 1 {
		return nil, errors.New("ent: SystemConfigGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := scgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (scgb *SystemConfigGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := scgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (scgb *SystemConfigGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = scgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systemconfig.Label}
	default:
		err = fmt.Errorf("ent: SystemConfigGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (scgb *SystemConfigGroupBy) Float64X(ctx context.Context) float64 {
	v, err := scgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (scgb *SystemConfigGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(scgb.fields) > 1 {
		return nil, errors.New("ent: SystemConfigGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := scgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (scgb *SystemConfigGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := scgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (scgb *SystemConfigGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = scgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systemconfig.Label}
	default:
		err = fmt.Errorf("ent: SystemConfigGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (scgb *SystemConfigGroupBy) BoolX(ctx context.Context) bool {
	v, err := scgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (scgb *SystemConfigGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range scgb.fields {
		if !systemconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := scgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (scgb *SystemConfigGroupBy) sqlQuery() *sql.Selector {
	selector := scgb.sql.Select()
	aggregation := make([]string, 0, len(scgb.fns))
	for _, fn := range scgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(scgb.fields)+len(scgb.fns))
		for _, f := range scgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(scgb.fields...)...)
}

// SystemConfigSelect is the builder for selecting fields of SystemConfig entities.
type SystemConfigSelect struct {
	*SystemConfigQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (scs *SystemConfigSelect) Scan(ctx context.Context, v interface{}) error {
	if err := scs.prepareQuery(ctx); err != nil {
		return err
	}
	scs.sql = scs.SystemConfigQuery.sqlQuery(ctx)
	return scs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (scs *SystemConfigSelect) ScanX(ctx context.Context, v interface{}) {
	if err := scs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (scs *SystemConfigSelect) Strings(ctx context.Context) ([]string, error) {
	if len(scs.fields) > 1 {
		return nil, errors.New("ent: SystemConfigSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := scs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (scs *SystemConfigSelect) StringsX(ctx context.Context) []string {
	v, err := scs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (scs *SystemConfigSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = scs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systemconfig.Label}
	default:
		err = fmt.Errorf("ent: SystemConfigSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (scs *SystemConfigSelect) StringX(ctx context.Context) string {
	v, err := scs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (scs *SystemConfigSelect) Ints(ctx context.Context) ([]int, error) {
	if len(scs.fields) > 1 {
		return nil, errors.New("ent: SystemConfigSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := scs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (scs *SystemConfigSelect) IntsX(ctx context.Context) []int {
	v, err := scs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (scs *SystemConfigSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = scs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systemconfig.Label}
	default:
		err = fmt.Errorf("ent: SystemConfigSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (scs *SystemConfigSelect) IntX(ctx context.Context) int {
	v, err := scs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (scs *SystemConfigSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(scs.fields) > 1 {
		return nil, errors.New("ent: SystemConfigSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := scs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (scs *SystemConfigSelect) Float64sX(ctx context.Context) []float64 {
	v, err := scs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (scs *SystemConfigSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = scs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systemconfig.Label}
	default:
		err = fmt.Errorf("ent: SystemConfigSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (scs *SystemConfigSelect) Float64X(ctx context.Context) float64 {
	v, err := scs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (scs *SystemConfigSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(scs.fields) > 1 {
		return nil, errors.New("ent: SystemConfigSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := scs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (scs *SystemConfigSelect) BoolsX(ctx context.Context) []bool {
	v, err := scs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (scs *SystemConfigSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = scs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{systemconfig.Label}
	default:
		err = fmt.Errorf("ent: SystemConfigSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (scs *SystemConfigSelect) BoolX(ctx context.Context) bool {
	v, err := scs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (scs *SystemConfigSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := scs.sql.Query()
	if err := scs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}