package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BasicBanner holds the schema definition for the BasicBanner entity.
type BasicBanner struct {
	ent.Schema
}

// 方法混入
func (BasicBanner) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// 广告位置
// 跳转方式

// Fields of the BasicBanner.
func (BasicBanner) Fields() []ent.Field {
	return []ent.Field{
		field.String("banner_name").
			Comment("banner图名字"),
		field.String("banner_img").
			Comment("图片链接"),
		field.Int("order").
			Optional().
			Default(0).
			Comment("排序字段"),
		field.Bool("status").
			Default(true),
	}
}

// Edges of the BasicBanner.
func (BasicBanner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("basic_banner_position", BasicBannerPosition.Type).
			Ref("basic_banner").
			Unique(),
		edge.From("basic_link", BasicLink.Type).
			Ref("basic_banner").
			Unique(),
	}
}
