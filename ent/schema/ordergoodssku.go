package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrderGoodsSku holds the schema definition for the OrderGoodsSku entity.
type OrderGoodsSku struct {
	ent.Schema
}

type SpecsOptionType struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// 方法混入
func (OrderGoodsSku) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Fields of the OrderGoodsSku.
func (OrderGoodsSku) Fields() []ent.Field {
	return []ent.Field{
		// 此部分作为交易快照使用
		field.String("spu_name").
			Comment("产品名称"),
		field.String("spu_code").
			//Unique().
			Comment("产品编码"),
		field.Text("spu_head_img").
			Optional().
			Comment("简介"),
		field.Text("spu_desc").
			Optional().
			Comment("详情"),
		field.Text("spu_details").
			Optional().
			Comment("详情"),
		field.Bool("is_custom_sku").
			Comment("是否自定义sku"),
		field.JSON("specs_option", []*SpecsOptionType{}).
			Comment("产品自定义sku属性"),
		field.Int("sku_id"). //sku_id 未使用外键形式 因后续商品删除
			Comment("sku_idd"),
		field.String("sku_name").
			Comment("sku名称"),
		field.String("sku_code").
			//Unique().
			Comment("sku编码"),
		field.Int("price").
			Comment("价格"),
		field.Int("amount").
			Comment("数量"),
	}
}

// Edges of the OrderGoodsSku.
func (OrderGoodsSku) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("goods_spu", GoodsSpu.Type).
			Ref("order_goods_sku").
			Unique(),
		edge.From("order_info", OrderInfo.Type).
			Ref("order_goods_sku").
			Unique(),
	}
}
