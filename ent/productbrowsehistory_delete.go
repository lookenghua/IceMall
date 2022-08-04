// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"ice-mall/ent/predicate"
	"ice-mall/ent/productbrowsehistory"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductBrowseHistoryDelete is the builder for deleting a ProductBrowseHistory entity.
type ProductBrowseHistoryDelete struct {
	config
	hooks    []Hook
	mutation *ProductBrowseHistoryMutation
}

// Where appends a list predicates to the ProductBrowseHistoryDelete builder.
func (pbhd *ProductBrowseHistoryDelete) Where(ps ...predicate.ProductBrowseHistory) *ProductBrowseHistoryDelete {
	pbhd.mutation.Where(ps...)
	return pbhd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pbhd *ProductBrowseHistoryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pbhd.hooks) == 0 {
		affected, err = pbhd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductBrowseHistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pbhd.mutation = mutation
			affected, err = pbhd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pbhd.hooks) - 1; i >= 0; i-- {
			if pbhd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pbhd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pbhd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pbhd *ProductBrowseHistoryDelete) ExecX(ctx context.Context) int {
	n, err := pbhd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pbhd *ProductBrowseHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: productbrowsehistory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productbrowsehistory.FieldID,
			},
		},
	}
	if ps := pbhd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pbhd.driver, _spec)
}

// ProductBrowseHistoryDeleteOne is the builder for deleting a single ProductBrowseHistory entity.
type ProductBrowseHistoryDeleteOne struct {
	pbhd *ProductBrowseHistoryDelete
}

// Exec executes the deletion query.
func (pbhdo *ProductBrowseHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := pbhdo.pbhd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productbrowsehistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pbhdo *ProductBrowseHistoryDeleteOne) ExecX(ctx context.Context) {
	pbhdo.pbhd.ExecX(ctx)
}