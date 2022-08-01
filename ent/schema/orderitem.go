package schema

import "entgo.io/ent"

// OrderItem holds the schema definition for the OrderItem entity.
type OrderItem struct {
	ent.Schema
}

// Fields of the OrderItem.
func (OrderItem) Fields() []ent.Field {
	return nil
}

// Edges of the OrderItem.
func (OrderItem) Edges() []ent.Edge {
	return nil
}
