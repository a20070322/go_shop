package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GoodsSpuImgs holds the schema definition for the GoodsSpuImgs entity.
type GoodsSpuImgs struct {
	ent.Schema
}

// 方法混入
func (GoodsSpuImgs) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}


// Fields of the GoodsSpuImgs.
func (GoodsSpuImgs) Fields() []ent.Field {
	return []ent.Field{
		field.String("img_name").
			Optional().
			Comment("图片名称"),
		field.String("img_path").
			Optional().
			Comment("图片路径"),
	}
}

// Edges of the GoodsSpuImgs.
func (GoodsSpuImgs) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("goods_spu", GoodsSpu.Type).
			Ref("goods_spu_imgs").
			Unique(),
	}
}
