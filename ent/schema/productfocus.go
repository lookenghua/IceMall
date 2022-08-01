package schema

import "entgo.io/ent"

// ProductFocus holds the schema definition for the ProductFocus entity.
type ProductFocus struct {
	ent.Schema
}

// Fields of the ProductFocus.
func (ProductFocus) Fields() []ent.Field {
	return nil
}

// Edges of the ProductFocus.
func (ProductFocus) Edges() []ent.Edge {
	return nil
}
