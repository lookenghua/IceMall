package schema

import "entgo.io/ent"

// ProductBrowseHistory holds the schema definition for the ProductBrowseHistory entity.
type ProductBrowseHistory struct {
	ent.Schema
}

// Fields of the ProductBrowseHistory.
func (ProductBrowseHistory) Fields() []ent.Field {
	return nil
}

// Edges of the ProductBrowseHistory.
func (ProductBrowseHistory) Edges() []ent.Edge {
	return nil
}
