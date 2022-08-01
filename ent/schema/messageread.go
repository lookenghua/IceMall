package schema

import "entgo.io/ent"

// MessageRead holds the schema definition for the MessageRead entity.
type MessageRead struct {
	ent.Schema
}

// Fields of the MessageRead.
func (MessageRead) Fields() []ent.Field {
	return nil
}

// Edges of the MessageRead.
func (MessageRead) Edges() []ent.Edge {
	return nil
}
