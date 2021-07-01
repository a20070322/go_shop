package goods_model

import (
	"context"
	"github.com/a20070322/shop-go/ent"
	"github.com/a20070322/shop-go/ent/goodsclassify"
	"github.com/a20070322/shop-go/global"
)

func ClassifyInit(ctx context.Context) *Classify {
	art := &Classify{}
	art.db = global.Db.GoodsClassify
	art.ctx = ctx
	return art
}

type Classify struct {
	db  *ent.GoodsClassifyClient
	ctx context.Context
}

func (m *Classify) ConsumerGet() ([]*ent.GoodsClassify, error) {
	return m.db.Query().Where(goodsclassify.LevelEQ(1)).All(m.ctx)
}

type ClassifyGetChildrenType struct {
	*ent.GoodsClassify
	Children []*ent.GoodsClassify `json:"children"`
}

func (m *Classify) ConsumerGetChildren(id int) ([]*ClassifyGetChildrenType, error) {
	list, err := m.db.Query().
		Where(goodsclassify.Pid(id)).
		All(m.ctx)
	if err != nil {
		return nil, err
	}
	var repList []*ClassifyGetChildrenType
	for _, v := range list {
		child, err2 := m.db.Query().
			Where(goodsclassify.Pid(v.ID)).
			All(m.ctx)
		if err2 != nil {
			return nil, err
		}
		repList = append(repList, &ClassifyGetChildrenType{
			GoodsClassify: v,
			Children:      child,
		})
	}
	return repList, nil
}
