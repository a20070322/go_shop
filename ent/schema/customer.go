package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// 方法混入
func (Customer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("mini_openid").
			Unique().Optional().
			Comment("小程序openid"),
		field.String("wechat_openid").
			Unique().Optional().
			Comment("公众号openid"),
		field.String("union_id").
			Unique().Optional().
			Comment("开放平台"),
		field.String("phone").
			Unique().Optional().
			Comment("手机号"),
		field.String("avatar").
			Optional().
			Comment("头像"),
		field.Bool("is_disable").
			Comment("是否禁用").
			Default(false).
			Optional(),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("address", CustomerAddress.Type),
		edge.To("order_info", OrderInfo.Type),
	}
}
