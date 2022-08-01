package schema

import "entgo.io/ent"

// ProductAttributeKey holds the schema definition for the ProductAttributeKey entity.
type ProductAttributeKey struct {
	ent.Schema
}

// Fields of the ProductAttributeKey.
func (ProductAttributeKey) Fields() []ent.Field {
	return nil
}

// Edges of the ProductAttributeKey.
func (ProductAttributeKey) Edges() []ent.Edge {
	return nil
}
