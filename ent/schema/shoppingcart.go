package schema

import "entgo.io/ent"

// ShoppingCart holds the schema definition for the ShoppingCart entity.
type ShoppingCart struct {
	ent.Schema
}

// Fields of the ShoppingCart.
func (ShoppingCart) Fields() []ent.Field {
	return nil
}

// Edges of the ShoppingCart.
func (ShoppingCart) Edges() []ent.Edge {
	return nil
}
