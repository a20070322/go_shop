package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GoodsSpu holds the schema definition for the GoodsSpu entity.
type GoodsSpu struct {
	ent.Schema
}

// 方法混入
func (GoodsSpu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Fields of the GoodsSpu.
func (GoodsSpu) Fields() []ent.Field {
	return []ent.Field{
		field.String("spu_name").
			Comment("产品名称"),
		field.String("spu_code").
			Unique().
			Comment("产品编码"),
		field.Text("spu_head_img").
			Optional().
			Comment("简介"),
		field.Int("sales_num").
			Optional().
			Comment("商品总销量"),
		field.Text("spu_desc").
			Optional().
			Comment("描述"),
		field.Text("spu_details").
			Optional().
			Comment("详情"),
		field.Bool("is_custom_sku").
			Comment("是否自定义sku"),
	}
}

// Edges of the GoodsSpu.
func (GoodsSpu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("goods_classify", GoodsClassify.Type).
			Ref("goods_spu").
			Unique(),
		edge.To("goods_sku", GoodsSku.Type),
		edge.To("order_goods_sku", OrderGoodsSku.Type),
		edge.To("goods_spu_imgs", GoodsSpuImgs.Type),
	}
}
