package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BasicBannerPosition holds the schema definition for the BasicBannerPosition entity.
type BasicBannerPosition struct {
	ent.Schema
}

func (BasicBannerPosition) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Fields of the BasicBannerPosition.
func (BasicBannerPosition) Fields() []ent.Field {
	return []ent.Field{
		field.String("position_name").Comment("位置名字"),
		field.String("remarks").Optional().Comment("备注"),
		field.Bool("status").Default(true),
	}
}

// Edges of the BasicBannerPosition.
func (BasicBannerPosition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("basic_banner", BasicBanner.Type),
	}
}
