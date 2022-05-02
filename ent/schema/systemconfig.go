package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// SystemConfig holds the schema definition for the SystemConfig entity.
type SystemConfig struct {
	ent.Schema
}

func (SystemConfig) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("rule"),
	}
}

// Fields of the SystemConfig.
func (SystemConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("rule").Comment("规则名称").Unique(),
		field.String("value").Comment("规则值"),
		field.String("remark").Comment("备注"),
	}
}

// Edges of the SystemConfig.
func (SystemConfig) Edges() []ent.Edge {
	return nil
}
