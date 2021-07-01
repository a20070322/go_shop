package schema

import "C"
import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrderInfo holds the schema definition for the OrderInfo entity.
type OrderInfo struct {
	ent.Schema
}

// 方法混入
func (OrderInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Fields of the OrderInfo.
func (OrderInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("order_number").
			Unique().
			Comment("订单号"),
		field.String("prepay_id").
			Unique().
			Optional().
			Comment("微信支付prepay_id"),
		field.Text("remark").
			Optional().
			Comment("订单备注"),
		field.Int8("status").
			Comment("订单状态"). // 0 待创建订单 1 待支付  2 支付完成  3 退款申请 4 退款完成  99 已取消 100 已完成  110 下单失败
			Default(0),
		field.Int8("delivery_status").
			Comment("物流状态"). // 0 初始 1 待发货  2 部分发货  3 已发货 4 已签收  99 退回 100 已完成
			Default(0),
	}
}

// Edges of the OrderInfo.
func (OrderInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("order_info").
			Unique(),
		edge.To("order_goods_sku", OrderGoodsSku.Type),
	}
}
