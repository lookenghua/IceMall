// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"ice-mall/ent/predicate"
	"ice-mall/ent/productspecs"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductSpecsDelete is the builder for deleting a ProductSpecs entity.
type ProductSpecsDelete struct {
	config
	hooks    []Hook
	mutation *ProductSpecsMutation
}

// Where appends a list predicates to the ProductSpecsDelete builder.
func (psd *ProductSpecsDelete) Where(ps ...predicate.ProductSpecs) *ProductSpecsDelete {
	psd.mutation.Where(ps...)
	return psd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (psd *ProductSpecsDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(psd.hooks) == 0 {
		affected, err = psd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductSpecsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			psd.mutation = mutation
			affected, err = psd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(psd.hooks) - 1; i >= 0; i-- {
			if psd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = psd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (psd *ProductSpecsDelete) ExecX(ctx context.Context) int {
	n, err := psd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (psd *ProductSpecsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: productspecs.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productspecs.FieldID,
			},
		},
	}
	if ps := psd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, psd.driver, _spec)
}

// ProductSpecsDeleteOne is the builder for deleting a single ProductSpecs entity.
type ProductSpecsDeleteOne struct {
	psd *ProductSpecsDelete
}

// Exec executes the deletion query.
func (psdo *ProductSpecsDeleteOne) Exec(ctx context.Context) error {
	n, err := psdo.psd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productspecs.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (psdo *ProductSpecsDeleteOne) ExecX(ctx context.Context) {
	psdo.psd.ExecX(ctx)
}
