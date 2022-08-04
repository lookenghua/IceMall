// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"ice-mall/ent/categorybanner"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoryBannerCreate is the builder for creating a CategoryBanner entity.
type CategoryBannerCreate struct {
	config
	mutation *CategoryBannerMutation
	hooks    []Hook
}

// Mutation returns the CategoryBannerMutation object of the builder.
func (cbc *CategoryBannerCreate) Mutation() *CategoryBannerMutation {
	return cbc.mutation
}

// Save creates the CategoryBanner in the database.
func (cbc *CategoryBannerCreate) Save(ctx context.Context) (*CategoryBanner, error) {
	var (
		err  error
		node *CategoryBanner
	)
	if len(cbc.hooks) == 0 {
		if err = cbc.check(); err != nil {
			return nil, err
		}
		node, err = cbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoryBannerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cbc.check(); err != nil {
				return nil, err
			}
			cbc.mutation = mutation
			if node, err = cbc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cbc.hooks) - 1; i >= 0; i-- {
			if cbc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cbc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cbc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cbc *CategoryBannerCreate) SaveX(ctx context.Context) *CategoryBanner {
	v, err := cbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cbc *CategoryBannerCreate) Exec(ctx context.Context) error {
	_, err := cbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbc *CategoryBannerCreate) ExecX(ctx context.Context) {
	if err := cbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cbc *CategoryBannerCreate) check() error {
	return nil
}

func (cbc *CategoryBannerCreate) sqlSave(ctx context.Context) (*CategoryBanner, error) {
	_node, _spec := cbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cbc *CategoryBannerCreate) createSpec() (*CategoryBanner, *sqlgraph.CreateSpec) {
	var (
		_node = &CategoryBanner{config: cbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: categorybanner.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: categorybanner.FieldID,
			},
		}
	)
	return _node, _spec
}

// CategoryBannerCreateBulk is the builder for creating many CategoryBanner entities in bulk.
type CategoryBannerCreateBulk struct {
	config
	builders []*CategoryBannerCreate
}

// Save creates the CategoryBanner entities in the database.
func (cbcb *CategoryBannerCreateBulk) Save(ctx context.Context) ([]*CategoryBanner, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cbcb.builders))
	nodes := make([]*CategoryBanner, len(cbcb.builders))
	mutators := make([]Mutator, len(cbcb.builders))
	for i := range cbcb.builders {
		func(i int, root context.Context) {
			builder := cbcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CategoryBannerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cbcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cbcb *CategoryBannerCreateBulk) SaveX(ctx context.Context) []*CategoryBanner {
	v, err := cbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cbcb *CategoryBannerCreateBulk) Exec(ctx context.Context) error {
	_, err := cbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbcb *CategoryBannerCreateBulk) ExecX(ctx context.Context) {
	if err := cbcb.Exec(ctx); err != nil {
		panic(err)
	}
}