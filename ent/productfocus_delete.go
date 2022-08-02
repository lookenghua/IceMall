// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"ice-mall/ent/predicate"
	"ice-mall/ent/productfocus"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductFocusDelete is the builder for deleting a ProductFocus entity.
type ProductFocusDelete struct {
	config
	hooks    []Hook
	mutation *ProductFocusMutation
}

// Where appends a list predicates to the ProductFocusDelete builder.
func (pfd *ProductFocusDelete) Where(ps ...predicate.ProductFocus) *ProductFocusDelete {
	pfd.mutation.Where(ps...)
	return pfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pfd *ProductFocusDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pfd.hooks) == 0 {
		affected, err = pfd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductFocusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pfd.mutation = mutation
			affected, err = pfd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pfd.hooks) - 1; i >= 0; i-- {
			if pfd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pfd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pfd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pfd *ProductFocusDelete) ExecX(ctx context.Context) int {
	n, err := pfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pfd *ProductFocusDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: productfocus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productfocus.FieldID,
			},
		},
	}
	if ps := pfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pfd.driver, _spec)
}

// ProductFocusDeleteOne is the builder for deleting a single ProductFocus entity.
type ProductFocusDeleteOne struct {
	pfd *ProductFocusDelete
}

// Exec executes the deletion query.
func (pfdo *ProductFocusDeleteOne) Exec(ctx context.Context) error {
	n, err := pfdo.pfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productfocus.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pfdo *ProductFocusDeleteOne) ExecX(ctx context.Context) {
	pfdo.pfd.ExecX(ctx)
}
