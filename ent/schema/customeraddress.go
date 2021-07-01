package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CustomerAddress holds the schema definition for the CustomerAddress entity.
type CustomerAddress struct {
	ent.Schema
}

// 方法混入
func (CustomerAddress) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Fields of the CustomerAddress.
func (CustomerAddress) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("姓名"),
		field.String("phone").
			Comment("联系方式"),
		field.String("province").
			Comment("省"),
		field.String("city").
			Comment("市"),
		field.String("area").
			Comment("区"),
		field.String("detailed").
			Comment("详细地址"),
		field.String("remark").
			Comment("备注").
			Optional(),
		field.Bool("is_default").
			Comment("是否默认").
			Default(false),
	}
}

// Edges of the CustomerAddress.
func (CustomerAddress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("address").
			Unique(),
	}
}
