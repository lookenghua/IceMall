package schema

import "entgo.io/ent"

// UserIntegral holds the schema definition for the UserIntegral entity.
type UserIntegral struct {
	ent.Schema
}

// Fields of the UserIntegral.
func (UserIntegral) Fields() []ent.Field {
	return nil
}

// Edges of the UserIntegral.
func (UserIntegral) Edges() []ent.Edge {
	return nil
}
