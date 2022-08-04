// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"ice-mall/ent/messageread"
	"ice-mall/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageReadUpdate is the builder for updating MessageRead entities.
type MessageReadUpdate struct {
	config
	hooks    []Hook
	mutation *MessageReadMutation
}

// Where appends a list predicates to the MessageReadUpdate builder.
func (mru *MessageReadUpdate) Where(ps ...predicate.MessageRead) *MessageReadUpdate {
	mru.mutation.Where(ps...)
	return mru
}

// Mutation returns the MessageReadMutation object of the builder.
func (mru *MessageReadUpdate) Mutation() *MessageReadMutation {
	return mru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mru *MessageReadUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mru.hooks) == 0 {
		affected, err = mru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageReadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mru.mutation = mutation
			affected, err = mru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mru.hooks) - 1; i >= 0; i-- {
			if mru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mru *MessageReadUpdate) SaveX(ctx context.Context) int {
	affected, err := mru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mru *MessageReadUpdate) Exec(ctx context.Context) error {
	_, err := mru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mru *MessageReadUpdate) ExecX(ctx context.Context) {
	if err := mru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mru *MessageReadUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messageread.Table,
			Columns: messageread.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messageread.FieldID,
			},
		},
	}
	if ps := mru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messageread.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MessageReadUpdateOne is the builder for updating a single MessageRead entity.
type MessageReadUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageReadMutation
}

// Mutation returns the MessageReadMutation object of the builder.
func (mruo *MessageReadUpdateOne) Mutation() *MessageReadMutation {
	return mruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mruo *MessageReadUpdateOne) Select(field string, fields ...string) *MessageReadUpdateOne {
	mruo.fields = append([]string{field}, fields...)
	return mruo
}

// Save executes the query and returns the updated MessageRead entity.
func (mruo *MessageReadUpdateOne) Save(ctx context.Context) (*MessageRead, error) {
	var (
		err  error
		node *MessageRead
	)
	if len(mruo.hooks) == 0 {
		node, err = mruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageReadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mruo.mutation = mutation
			node, err = mruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mruo.hooks) - 1; i >= 0; i-- {
			if mruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mruo *MessageReadUpdateOne) SaveX(ctx context.Context) *MessageRead {
	node, err := mruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mruo *MessageReadUpdateOne) Exec(ctx context.Context) error {
	_, err := mruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mruo *MessageReadUpdateOne) ExecX(ctx context.Context) {
	if err := mruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mruo *MessageReadUpdateOne) sqlSave(ctx context.Context) (_node *MessageRead, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messageread.Table,
			Columns: messageread.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messageread.FieldID,
			},
		},
	}
	id, ok := mruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MessageRead.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messageread.FieldID)
		for _, f := range fields {
			if !messageread.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != messageread.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &MessageRead{config: mruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messageread.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}