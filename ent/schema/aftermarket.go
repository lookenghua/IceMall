package schema

import "entgo.io/ent"

// Aftermarket holds the schema definition for the Aftermarket entity.
type Aftermarket struct {
	ent.Schema
}

// Fields of the Aftermarket.
func (Aftermarket) Fields() []ent.Field {
	return nil
}

// Edges of the Aftermarket.
func (Aftermarket) Edges() []ent.Edge {
	return nil
}
