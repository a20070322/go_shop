package order_utils

import (
	"github.com/a20070322/shop-go/ent"
	"github.com/a20070322/shop-go/ent/schema"
)

type OrderCreateGoodsItem struct {
	SkuID  int `json:"sku_id" binding:"required"`
	Amount int `json:"amount" binding:"required"`
}


// 商品列表 数组对象转 对象
func ProductsToIds(products []*OrderCreateGoodsItem) (ids []int, mapsT map[int]int) {
	var maps = make(map[int]int)
	if products != nil {
		for _, p := range products {
			ids = append(ids, p.SkuID)
			maps[p.SkuID] = p.Amount
		}
	}
	return ids, maps
}


// 商品属性序列化
func SpecsOptionToJson(specsOptions []*ent.GoodsSpecsOption) (optionsJson []*schema.SpecsOptionType) {
	for _,o := range specsOptions{
		optionsJson = append(optionsJson, &schema.SpecsOptionType{
			Key: o.Edges.GoodsSpecs.SpecsName,
			Value: o.SpecsOptionValue,
		})
	}
	return optionsJson
}


