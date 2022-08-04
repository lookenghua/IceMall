// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"ice-mall/ent/predicate"
	"ice-mall/ent/productcategory"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductCategoryDelete is the builder for deleting a ProductCategory entity.
type ProductCategoryDelete struct {
	config
	hooks    []Hook
	mutation *ProductCategoryMutation
}

// Where appends a list predicates to the ProductCategoryDelete builder.
func (pcd *ProductCategoryDelete) Where(ps ...predicate.ProductCategory) *ProductCategoryDelete {
	pcd.mutation.Where(ps...)
	return pcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pcd *ProductCategoryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pcd.hooks) == 0 {
		affected, err = pcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pcd.mutation = mutation
			affected, err = pcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pcd.hooks) - 1; i >= 0; i-- {
			if pcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcd *ProductCategoryDelete) ExecX(ctx context.Context) int {
	n, err := pcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pcd *ProductCategoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: productcategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productcategory.FieldID,
			},
		},
	}
	if ps := pcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pcd.driver, _spec)
}

// ProductCategoryDeleteOne is the builder for deleting a single ProductCategory entity.
type ProductCategoryDeleteOne struct {
	pcd *ProductCategoryDelete
}

// Exec executes the deletion query.
func (pcdo *ProductCategoryDeleteOne) Exec(ctx context.Context) error {
	n, err := pcdo.pcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productcategory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pcdo *ProductCategoryDeleteOne) ExecX(ctx context.Context) {
	pcdo.pcd.ExecX(ctx)
}