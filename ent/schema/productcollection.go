package schema

import "entgo.io/ent"

// ProductCollection holds the schema definition for the ProductCollection entity.
type ProductCollection struct {
	ent.Schema
}

// Fields of the ProductCollection.
func (ProductCollection) Fields() []ent.Field {
	return nil
}

// Edges of the ProductCollection.
func (ProductCollection) Edges() []ent.Edge {
	return nil
}
