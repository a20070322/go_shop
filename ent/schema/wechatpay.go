package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// WeChatPay holds the schema definition for the WeChatPay entity.
type WeChatPay struct {
	ent.Schema
}

// 方法混入
func (WeChatPay) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Fields of the WeChatPay.
func (WeChatPay) Fields() []ent.Field {
	return []ent.Field{
		field.String("out_trade_no").
			Optional().
			Unique().
			Comment("商户订单号"),
		field.String("transaction_id").
			Optional().
			Unique().
			Comment("微信支付订单号"),
		field.Enum("trade_type").
			Values("JSAPI", "NATIVE", "APP").
			Optional().
			Comment("交易类型"),
		field.String("bank_type").
			Optional().
			Comment("付款银行"),
		field.Time("success_time").
			Optional().
			Comment("支付完成时间"),
		field.String("payer_currency").
			Optional().
			Comment("支付币种"),
		field.Int32("payer_total").
			Optional().
			Comment("支付金额"). //分为单位
			Default(0),
		field.Int8("trade_state").
			Optional().
			Comment("订单状态"). // 0 PADDING 1 SUCCESS  2 FAIL
			Default(0),
	}
}

// Edges of the WeChatPay.
func (WeChatPay) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order_info", OrderInfo.Type).
			Ref("wechat_pay").
			Unique(),
	}
}
