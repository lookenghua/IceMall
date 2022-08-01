package schema

import "entgo.io/ent"

// Annex holds the schema definition for the Annex entity.
type Annex struct {
	ent.Schema
}

// Fields of the Annex.
func (Annex) Fields() []ent.Field {
	return nil
}

// Edges of the Annex.
func (Annex) Edges() []ent.Edge {
	return nil
}
