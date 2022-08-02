// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"ice-mall/ent/predicate"
	"ice-mall/ent/receivingaddress"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReceivingAddressDelete is the builder for deleting a ReceivingAddress entity.
type ReceivingAddressDelete struct {
	config
	hooks    []Hook
	mutation *ReceivingAddressMutation
}

// Where appends a list predicates to the ReceivingAddressDelete builder.
func (rad *ReceivingAddressDelete) Where(ps ...predicate.ReceivingAddress) *ReceivingAddressDelete {
	rad.mutation.Where(ps...)
	return rad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rad *ReceivingAddressDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rad.hooks) == 0 {
		affected, err = rad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReceivingAddressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rad.mutation = mutation
			affected, err = rad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rad.hooks) - 1; i >= 0; i-- {
			if rad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rad *ReceivingAddressDelete) ExecX(ctx context.Context) int {
	n, err := rad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rad *ReceivingAddressDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: receivingaddress.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: receivingaddress.FieldID,
			},
		},
	}
	if ps := rad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rad.driver, _spec)
}

// ReceivingAddressDeleteOne is the builder for deleting a single ReceivingAddress entity.
type ReceivingAddressDeleteOne struct {
	rad *ReceivingAddressDelete
}

// Exec executes the deletion query.
func (rado *ReceivingAddressDeleteOne) Exec(ctx context.Context) error {
	n, err := rado.rad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{receivingaddress.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rado *ReceivingAddressDeleteOne) ExecX(ctx context.Context) {
	rado.rad.ExecX(ctx)
}
