// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"ice-mall/ent/feedback"
	"ice-mall/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FeedbackQuery is the builder for querying Feedback entities.
type FeedbackQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Feedback
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FeedbackQuery builder.
func (fq *FeedbackQuery) Where(ps ...predicate.Feedback) *FeedbackQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit adds a limit step to the query.
func (fq *FeedbackQuery) Limit(limit int) *FeedbackQuery {
	fq.limit = &limit
	return fq
}

// Offset adds an offset step to the query.
func (fq *FeedbackQuery) Offset(offset int) *FeedbackQuery {
	fq.offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FeedbackQuery) Unique(unique bool) *FeedbackQuery {
	fq.unique = &unique
	return fq
}

// Order adds an order step to the query.
func (fq *FeedbackQuery) Order(o ...OrderFunc) *FeedbackQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// First returns the first Feedback entity from the query.
// Returns a *NotFoundError when no Feedback was found.
func (fq *FeedbackQuery) First(ctx context.Context) (*Feedback, error) {
	nodes, err := fq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{feedback.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FeedbackQuery) FirstX(ctx context.Context) *Feedback {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Feedback ID from the query.
// Returns a *NotFoundError when no Feedback ID was found.
func (fq *FeedbackQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{feedback.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FeedbackQuery) FirstIDX(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Feedback entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Feedback entity is found.
// Returns a *NotFoundError when no Feedback entities are found.
func (fq *FeedbackQuery) Only(ctx context.Context) (*Feedback, error) {
	nodes, err := fq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{feedback.Label}
	default:
		return nil, &NotSingularError{feedback.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FeedbackQuery) OnlyX(ctx context.Context) *Feedback {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Feedback ID in the query.
// Returns a *NotSingularError when more than one Feedback ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FeedbackQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{feedback.Label}
	default:
		err = &NotSingularError{feedback.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FeedbackQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Feedbacks.
func (fq *FeedbackQuery) All(ctx context.Context) ([]*Feedback, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fq *FeedbackQuery) AllX(ctx context.Context) []*Feedback {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Feedback IDs.
func (fq *FeedbackQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fq.Select(feedback.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FeedbackQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FeedbackQuery) Count(ctx context.Context) (int, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FeedbackQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FeedbackQuery) Exist(ctx context.Context) (bool, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FeedbackQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FeedbackQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FeedbackQuery) Clone() *FeedbackQuery {
	if fq == nil {
		return nil
	}
	return &FeedbackQuery{
		config:     fq.config,
		limit:      fq.limit,
		offset:     fq.offset,
		order:      append([]OrderFunc{}, fq.order...),
		predicates: append([]predicate.Feedback{}, fq.predicates...),
		// clone intermediate query.
		sql:    fq.sql.Clone(),
		path:   fq.path,
		unique: fq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (fq *FeedbackQuery) GroupBy(field string, fields ...string) *FeedbackGroupBy {
	group := &FeedbackGroupBy{config: fq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (fq *FeedbackQuery) Select(fields ...string) *FeedbackSelect {
	fq.fields = append(fq.fields, fields...)
	return &FeedbackSelect{FeedbackQuery: fq}
}

func (fq *FeedbackQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fq.fields {
		if !feedback.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FeedbackQuery) sqlAll(ctx context.Context) ([]*Feedback, error) {
	var (
		nodes = []*Feedback{}
		_spec = fq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Feedback{config: fq.config}
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
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (fq *FeedbackQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Node.Columns = fq.fields
	if len(fq.fields) > 0 {
		_spec.Unique = fq.unique != nil && *fq.unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FeedbackQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fq *FeedbackQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   feedback.Table,
			Columns: feedback.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: feedback.FieldID,
			},
		},
		From:   fq.sql,
		Unique: true,
	}
	if unique := fq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feedback.FieldID)
		for i := range fields {
			if fields[i] != feedback.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FeedbackQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(feedback.Table)
	columns := fq.fields
	if len(columns) == 0 {
		columns = feedback.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.unique != nil && *fq.unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FeedbackGroupBy is the group-by builder for Feedback entities.
type FeedbackGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FeedbackGroupBy) Aggregate(fns ...AggregateFunc) *FeedbackGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fgb *FeedbackGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fgb.path(ctx)
	if err != nil {
		return err
	}
	fgb.sql = query
	return fgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fgb *FeedbackGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := fgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (fgb *FeedbackGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FeedbackGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fgb *FeedbackGroupBy) StringsX(ctx context.Context) []string {
	v, err := fgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fgb *FeedbackGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{feedback.Label}
	default:
		err = fmt.Errorf("ent: FeedbackGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fgb *FeedbackGroupBy) StringX(ctx context.Context) string {
	v, err := fgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (fgb *FeedbackGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FeedbackGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fgb *FeedbackGroupBy) IntsX(ctx context.Context) []int {
	v, err := fgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fgb *FeedbackGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{feedback.Label}
	default:
		err = fmt.Errorf("ent: FeedbackGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fgb *FeedbackGroupBy) IntX(ctx context.Context) int {
	v, err := fgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (fgb *FeedbackGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FeedbackGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fgb *FeedbackGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := fgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fgb *FeedbackGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{feedback.Label}
	default:
		err = fmt.Errorf("ent: FeedbackGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fgb *FeedbackGroupBy) Float64X(ctx context.Context) float64 {
	v, err := fgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (fgb *FeedbackGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FeedbackGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fgb *FeedbackGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := fgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fgb *FeedbackGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{feedback.Label}
	default:
		err = fmt.Errorf("ent: FeedbackGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fgb *FeedbackGroupBy) BoolX(ctx context.Context) bool {
	v, err := fgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fgb *FeedbackGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fgb.fields {
		if !feedback.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fgb *FeedbackGroupBy) sqlQuery() *sql.Selector {
	selector := fgb.sql.Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fgb.fields)+len(fgb.fns))
		for _, f := range fgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fgb.fields...)...)
}

// FeedbackSelect is the builder for selecting fields of Feedback entities.
type FeedbackSelect struct {
	*FeedbackQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FeedbackSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	fs.sql = fs.FeedbackQuery.sqlQuery(ctx)
	return fs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fs *FeedbackSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (fs *FeedbackSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FeedbackSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fs *FeedbackSelect) StringsX(ctx context.Context) []string {
	v, err := fs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (fs *FeedbackSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{feedback.Label}
	default:
		err = fmt.Errorf("ent: FeedbackSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fs *FeedbackSelect) StringX(ctx context.Context) string {
	v, err := fs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (fs *FeedbackSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FeedbackSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fs *FeedbackSelect) IntsX(ctx context.Context) []int {
	v, err := fs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (fs *FeedbackSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{feedback.Label}
	default:
		err = fmt.Errorf("ent: FeedbackSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fs *FeedbackSelect) IntX(ctx context.Context) int {
	v, err := fs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (fs *FeedbackSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FeedbackSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fs *FeedbackSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (fs *FeedbackSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{feedback.Label}
	default:
		err = fmt.Errorf("ent: FeedbackSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fs *FeedbackSelect) Float64X(ctx context.Context) float64 {
	v, err := fs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (fs *FeedbackSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FeedbackSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fs *FeedbackSelect) BoolsX(ctx context.Context) []bool {
	v, err := fs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (fs *FeedbackSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{feedback.Label}
	default:
		err = fmt.Errorf("ent: FeedbackSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fs *FeedbackSelect) BoolX(ctx context.Context) bool {
	v, err := fs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fs *FeedbackSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fs.sql.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
