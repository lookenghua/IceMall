// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"ice-mall/ent/productcollection"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductCollectionCreate is the builder for creating a ProductCollection entity.
type ProductCollectionCreate struct {
	config
	mutation *ProductCollectionMutation
	hooks    []Hook
}

// Mutation returns the ProductCollectionMutation object of the builder.
func (pcc *ProductCollectionCreate) Mutation() *ProductCollectionMutation {
	return pcc.mutation
}

// Save creates the ProductCollection in the database.
func (pcc *ProductCollectionCreate) Save(ctx context.Context) (*ProductCollection, error) {
	var (
		err  error
		node *ProductCollection
	)
	if len(pcc.hooks) == 0 {
		if err = pcc.check(); err != nil {
			return nil, err
		}
		node, err = pcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductCollectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pcc.check(); err != nil {
				return nil, err
			}
			pcc.mutation = mutation
			if node, err = pcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pcc.hooks) - 1; i >= 0; i-- {
			if pcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pcc *ProductCollectionCreate) SaveX(ctx context.Context) *ProductCollection {
	v, err := pcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcc *ProductCollectionCreate) Exec(ctx context.Context) error {
	_, err := pcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcc *ProductCollectionCreate) ExecX(ctx context.Context) {
	if err := pcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcc *ProductCollectionCreate) check() error {
	return nil
}

func (pcc *ProductCollectionCreate) sqlSave(ctx context.Context) (*ProductCollection, error) {
	_node, _spec := pcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pcc *ProductCollectionCreate) createSpec() (*ProductCollection, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductCollection{config: pcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: productcollection.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productcollection.FieldID,
			},
		}
	)
	return _node, _spec
}

// ProductCollectionCreateBulk is the builder for creating many ProductCollection entities in bulk.
type ProductCollectionCreateBulk struct {
	config
	builders []*ProductCollectionCreate
}

// Save creates the ProductCollection entities in the database.
func (pccb *ProductCollectionCreateBulk) Save(ctx context.Context) ([]*ProductCollection, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pccb.builders))
	nodes := make([]*ProductCollection, len(pccb.builders))
	mutators := make([]Mutator, len(pccb.builders))
	for i := range pccb.builders {
		func(i int, root context.Context) {
			builder := pccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductCollectionMutation)
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
					_, err = mutators[i+1].Mutate(root, pccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pccb *ProductCollectionCreateBulk) SaveX(ctx context.Context) []*ProductCollection {
	v, err := pccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pccb *ProductCollectionCreateBulk) Exec(ctx context.Context) error {
	_, err := pccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pccb *ProductCollectionCreateBulk) ExecX(ctx context.Context) {
	if err := pccb.Exec(ctx); err != nil {
		panic(err)
	}
}
