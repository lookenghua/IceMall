package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Captcha holds the schema definition for the Captcha entity.
type Captcha struct {
	ent.Schema
}

func (Captcha) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DefaultMix{},
	}
}

// Fields of the Captcha.
func (Captcha) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Comment("类型").Values("REGISTER"),
		field.String("phone").Comment("手机号"),
		field.String("code").Comment("验证码"),
	}
}

// Edges of the Captcha.
func (Captcha) Edges() []ent.Edge {
	return nil
}
