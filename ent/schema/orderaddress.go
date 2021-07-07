package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrderAddress holds the schema definition for the OrderAddress entity.
type OrderAddress struct {
	ent.Schema
}

// Fields of the OrderAddress.
func (OrderAddress) Fields() []ent.Field {

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
	}
}

// Edges of the OrderAddress.
func (OrderAddress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order_info", OrderInfo.Type).
			Ref("order_address").
			Unique(),
	}
}
