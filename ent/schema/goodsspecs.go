package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GoodsSpecs holds the schema definition for the GoodsSpecs entity.
type GoodsSpecs struct {
	ent.Schema
}

// 方法混入
func (GoodsSpecs) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Fields of the GoodsSpecs.
func (GoodsSpecs) Fields() []ent.Field {
	return []ent.Field{
		field.String("specs_name").
			Unique().
			Comment("规格名称"),
	}
}

// Edges of the GoodsSpecs.
func (GoodsSpecs) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("goods_specs_option",GoodsSpecsOption.Type),
	}
}
