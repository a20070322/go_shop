package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GoodsSku holds the schema definition for the GoodsSku entity.
type GoodsSku struct {
	ent.Schema
}

// 方法混入
func (GoodsSku) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Fields of the GoodsSku.
func (GoodsSku) Fields() []ent.Field {
	return []ent.Field{
		field.String("sku_name").
			Comment("sku名称"),
		field.String("sku_code").
			Unique().
			Comment("sku编码"),
		field.Int("stock_num").
			Default(0).
			Comment("库存"),
		field.Int("sales_num").
			Default(0).
			Comment("销量"),
		field.Int("price").
			Comment("价格"),
	}
}

// Edges of the GoodsSku.
func (GoodsSku) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("goods_spu", GoodsSpu.Type).
			Ref("goods_sku").
			Unique(),
		edge.From("goods_specs_option", GoodsSpecsOption.Type).
			Ref("goods_sku"),
	}
}
