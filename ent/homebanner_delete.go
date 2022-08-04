// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"ice-mall/ent/homebanner"
	"ice-mall/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HomeBannerDelete is the builder for deleting a HomeBanner entity.
type HomeBannerDelete struct {
	config
	hooks    []Hook
	mutation *HomeBannerMutation
}

// Where appends a list predicates to the HomeBannerDelete builder.
func (hbd *HomeBannerDelete) Where(ps ...predicate.HomeBanner) *HomeBannerDelete {
	hbd.mutation.Where(ps...)
	return hbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hbd *HomeBannerDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hbd.hooks) == 0 {
		affected, err = hbd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HomeBannerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hbd.mutation = mutation
			affected, err = hbd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hbd.hooks) - 1; i >= 0; i-- {
			if hbd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hbd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hbd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (hbd *HomeBannerDelete) ExecX(ctx context.Context) int {
	n, err := hbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hbd *HomeBannerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: homebanner.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: homebanner.FieldID,
			},
		},
	}
	if ps := hbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, hbd.driver, _spec)
}

// HomeBannerDeleteOne is the builder for deleting a single HomeBanner entity.
type HomeBannerDeleteOne struct {
	hbd *HomeBannerDelete
}

// Exec executes the deletion query.
func (hbdo *HomeBannerDeleteOne) Exec(ctx context.Context) error {
	n, err := hbdo.hbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{homebanner.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hbdo *HomeBannerDeleteOne) ExecX(ctx context.Context) {
	hbdo.hbd.ExecX(ctx)
}