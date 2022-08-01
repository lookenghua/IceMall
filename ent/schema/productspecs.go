package schema

import "entgo.io/ent"

// ProductSpecs holds the schema definition for the ProductSpecs entity.
type ProductSpecs struct {
	ent.Schema
}

// Fields of the ProductSpecs.
func (ProductSpecs) Fields() []ent.Field {
	return nil
}

// Edges of the ProductSpecs.
func (ProductSpecs) Edges() []ent.Edge {
	return nil
}
