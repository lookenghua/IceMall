package schema

import "entgo.io/ent"

// CategoryBanner holds the schema definition for the CategoryBanner entity.
type CategoryBanner struct {
	ent.Schema
}

// Fields of the CategoryBanner.
func (CategoryBanner) Fields() []ent.Field {
	return nil
}

// Edges of the CategoryBanner.
func (CategoryBanner) Edges() []ent.Edge {
	return nil
}
