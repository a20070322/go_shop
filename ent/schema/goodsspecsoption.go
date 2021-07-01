package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GoodsSpecsOption holds the schema definition for the GoodsSpecsOption entity.
type GoodsSpecsOption struct {
	ent.Schema
}

// 方法混入
func (GoodsSpecsOption) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Fields of the GoodsSpecsOption.
func (GoodsSpecsOption) Fields() []ent.Field {
	return []ent.Field{
		field.String("specs_option_value").
			Unique().
			Comment("选项值"),
	}
}

// Edges of the GoodsSpecsOption.
func (GoodsSpecsOption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("goods_specs", GoodsSpecs.Type).
			Ref("goods_specs_option").
			Unique(),
		edge.To("goods_sku",GoodsSku.Type),
	}
}
