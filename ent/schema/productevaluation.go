package schema

import "entgo.io/ent"

// ProductEvaluation holds the schema definition for the ProductEvaluation entity.
type ProductEvaluation struct {
	ent.Schema
}

// Fields of the ProductEvaluation.
func (ProductEvaluation) Fields() []ent.Field {
	return nil
}

// Edges of the ProductEvaluation.
func (ProductEvaluation) Edges() []ent.Edge {
	return nil
}
