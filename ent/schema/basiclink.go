package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BasicLink holds the schema definition for the BasicLink entity.
type BasicLink struct {
	ent.Schema
}

// Fields of the BasicLink.
func (BasicLink) Fields() []ent.Field {
	return []ent.Field{
		field.String("link_name").Comment("链接名称"),
		field.String("link_type").Comment("跳转类型，0 H5页面，1 小程序页面，2 其他第三方小程序"),
		field.String("link_address").
			Comment("链接参数"),
		field.String("appid").
			Optional().
			Comment("小程序appid"),
		field.Bool("is_register").
			Default(false).
			Comment("是否需要登录"),
		field.String("remarks").
			Optional().
			Comment("备注"),
		field.Bool("status").
			Default(true),
	}
}

// 方法混入
func (BasicLink) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Edges of the BasicLink.
func (BasicLink) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("basic_banner", BasicBanner.Type),
	}
}
