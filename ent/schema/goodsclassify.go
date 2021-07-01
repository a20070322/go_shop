package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GoodsClassify holds the schema definition for the GoodsClassify entity.
type GoodsClassify struct {
	ent.Schema
}

// 方法混入
func (GoodsClassify) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Fields of the GoodsClassify.
func (GoodsClassify) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id"),
		field.String("classify_name").
			Optional().
			Comment("分类名称"),
		field.String("classify_code").
			Unique().
			Optional().
			Comment("分类编码"),
		field.Int("pid").
			Comment("父级id").
			Default(0).
			Optional(),
		field.Int("level").
			Comment("级别"),
		field.Int("sort").
			Comment("排序").
			Default(0).
			Min(0).
			Optional(),
		field.Text("icon").
			Comment("icon图标").
			Default("").
			Optional(),
		field.Bool("is_disable").
			Comment("是否禁用").
			Default(false).
			Optional(),
	}
}

// Edges of the GoodsClassify.
func (GoodsClassify) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("goods_spu", GoodsSpu.Type),
	}
}
