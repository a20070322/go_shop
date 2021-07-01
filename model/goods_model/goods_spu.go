package goods_model

import (
	"context"
	"fmt"
	"github.com/a20070322/shop-go/ent"
	"github.com/a20070322/shop-go/ent/goodsspecsoption"
	"github.com/a20070322/shop-go/ent/goodsspu"
	"github.com/a20070322/shop-go/ent/predicate"
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/pkg/model_utils"
	"github.com/a20070322/shop-go/pkg/utils"
	"sort"
)

func SpuInit(ctx context.Context) *Spu {
	art := &Spu{}
	art.db = global.Db.GoodsSpu
	art.ctx = ctx
	return art
}

type Spu struct {
	db  *ent.GoodsSpuClient
	ctx context.Context
}

type SpuListFormType struct {
	model_utils.PageOptions
	//model_utils.PageSort
	// 商品名字搜索
	SpuName string `json:"spu_name" form:"spu_name"`
}

type SpuItemTypeForm struct {
	*ent.GoodsSpu
	Specs     []*ent.GoodsSpecs `json:"specs"`
	ShowPrice string            `json:"show_price"`
	Sku       *ent.GoodsSku     `json:"sku"`
	StockNum  int               `json:"stock_num"`
	SalesNum  int               `json:"sales_num"`
}

type SpuListResponse struct {
	Data  []*SpuItemTypeForm `json:"data"`
	Total int                `json:"total"`
}

type SkuSort struct {
	p    []*ent.GoodsSku
	less func(x, y *ent.GoodsSku) bool
}

func (x SkuSort) Len() int           { return len(x.p) }
func (x SkuSort) Less(i, j int) bool { return x.less(x.p[i], x.p[j]) }
func (x SkuSort) Swap(i, j int)      { x.p[i], x.p[j] = x.p[j], x.p[i] }

func (m *Spu) List(form *SpuListFormType) (repList SpuListResponse, err error) {
	// 设置分页默认值
	model_utils.PipePagerFn(form)
	var whereArr []predicate.GoodsSpu
	if form.SpuName != "" {
		whereArr = append(whereArr, goodsspu.SpuNameContains(form.SpuName))
	}
	db := m.db.Query().Where(whereArr...).
		Select(
			goodsspu.FieldID,
			goodsspu.FieldSpuName,
			goodsspu.FieldSpuCode,
			goodsspu.FieldSpuHeadImg,
			goodsspu.FieldIsCustomSku,
		)
	total, err2 := db.Count(m.ctx)
	if err2 != nil {
		return repList, err2
	}
	repList.Total = total
	// 设置自动分页
	model_utils.PipeLimitFn(db, form)
	list, errFind := db.WithGoodsSku().All(m.ctx)
	if errFind != nil {
		return SpuListResponse{}, errFind
	}
	// 商品数据拼装
	var data []*SpuItemTypeForm
	for _, item := range list {
		var sku *ent.GoodsSku
		// 展示价格计算
		var showPrice  = ""

		sort.Sort(&SkuSort{item.Edges.GoodsSku, func(x, y *ent.GoodsSku) bool {
			return    x.Price < y.Price
		}})
		price1 := item.Edges.GoodsSku[0].Price
		price2 := item.Edges.GoodsSku[len(item.Edges.GoodsSku)-1].Price
		if price1 == price2 {
			showPrice = utils.PriceToStr(price1)
		}else{
			fmt.Println(price1/100)
			showPrice =utils.PriceToStr(price1) + "~" + utils.PriceToStr(price2)
		}
		// 是否是多sku商品
		if item.IsCustomSku {
			sku = nil
		} else {
			sku = item.Edges.GoodsSku[0]
		}
		goodsItem := &SpuItemTypeForm{
			item,
			nil,
			showPrice,
			sku,
			0,
			0,
		}
		//计算总销量总库存
		for _, v := range item.Edges.GoodsSku {
			goodsItem.StockNum += v.StockNum
			goodsItem.SalesNum += v.SalesNum
		}
		data = append(data, goodsItem)
	}

	if err != nil {
		return repList, err
	}
	repList.Data = data
	return repList, nil
}

type SpuGetIdFormType struct {
	ID int `json:"id" form:"id"`
}

func (m *Spu) GetId(form *SpuGetIdFormType) (*SpuItemTypeForm, error) {
	global.Db.GoodsSku.Query().WithGoodsSpecsOption()
	spu, err := m.db.Query().
		Where(goodsspu.IDEQ(form.ID)).
		WithGoodsSpuImgs().
		WithGoodsSku(func(query *ent.GoodsSkuQuery) {
			query.WithGoodsSpecsOption()
		}).
		First(m.ctx)
	repItem := &SpuItemTypeForm{
		spu,
		nil,
		"",
		nil,
		0,
		0,
	}
	if err != nil {
		return repItem, err
	}

	for _, v := range spu.Edges.GoodsSku {
		repItem.StockNum += v.StockNum
		repItem.SalesNum += v.SalesNum
	}

	// 如果是sku商品
	if spu.IsCustomSku {
		specs, err2 := m.GetSpecs(spu)
		if err2 != nil {
			global.Logger.Error(err2.Error())
		}
		repItem.Specs = specs
		sort.Sort(&SkuSort{spu.Edges.GoodsSku, func(x, y *ent.GoodsSku) bool {
			return x.Price < y.Price
		}})
		price1 := spu.Edges.GoodsSku[0].Price
		price2 := spu.Edges.GoodsSku[len(spu.Edges.GoodsSku)-1].Price
		if price1 == price2 {
			repItem.ShowPrice = utils.PriceToStr(price1)
		}else{
			repItem.ShowPrice = fmt.Sprintf("%s~%s",  utils.PriceToStr(price1),utils.PriceToStr(price2))
		}
	} else {
		repItem.ShowPrice =utils.PriceToStr(spu.Edges.GoodsSku[0].Price)
	}
	return repItem, nil
}

func (m *Spu) GetSpecs(spu *ent.GoodsSpu) ([]*ent.GoodsSpecs, error) {
	specs, err2 := spu.QueryGoodsSku().QueryGoodsSpecsOption().QueryGoodsSpecs().WithGoodsSpecsOption(func(query *ent.GoodsSpecsOptionQuery) {
		query.Where(goodsspecsoption.HasGoodsSku())
	}).All(m.ctx)
	if err2 != nil {
		global.Logger.Error(err2.Error())
		return specs, err2
	}
	return specs, nil
}
