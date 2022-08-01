package schema

import "entgo.io/ent"

// HomeBanner holds the schema definition for the HomeBanner entity.
type HomeBanner struct {
	ent.Schema
}

// Fields of the HomeBanner.
func (HomeBanner) Fields() []ent.Field {
	return nil
}

// Edges of the HomeBanner.
func (HomeBanner) Edges() []ent.Edge {
	return nil
}
