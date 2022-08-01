package schema

import "entgo.io/ent"

// ProductAttributeValue holds the schema definition for the ProductAttributeValue entity.
type ProductAttributeValue struct {
	ent.Schema
}

// Fields of the ProductAttributeValue.
func (ProductAttributeValue) Fields() []ent.Field {
	return nil
}

// Edges of the ProductAttributeValue.
func (ProductAttributeValue) Edges() []ent.Edge {
	return nil
}
