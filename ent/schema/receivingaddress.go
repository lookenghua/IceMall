package schema

import "entgo.io/ent"

// ReceivingAddress holds the schema definition for the ReceivingAddress entity.
type ReceivingAddress struct {
	ent.Schema
}

// Fields of the ReceivingAddress.
func (ReceivingAddress) Fields() []ent.Field {
	return nil
}

// Edges of the ReceivingAddress.
func (ReceivingAddress) Edges() []ent.Edge {
	return nil
}
