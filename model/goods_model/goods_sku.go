package goods_model

import (
	"context"
	"errors"
	"fmt"
	"github.com/a20070322/shop-go/ent"
	"github.com/a20070322/shop-go/ent/goodssku"
	"github.com/a20070322/shop-go/ent/goodsspecsoption"
	"github.com/a20070322/shop-go/ent/goodsspu"
	"github.com/a20070322/shop-go/ent/predicate"
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/pkg/order_utils"
	"github.com/a20070322/shop-go/pkg/utils/ent_utils"

)

func SkuInit(ctx context.Context) *Sku {
	art := &Sku{}
	art.db = global.Db.GoodsSku
	art.ctx = ctx
	return art
}

type Sku struct {
	db  *ent.GoodsSkuClient
	ctx context.Context
}
type SkuGetFormType struct {
	ID        int   `json:"id" form:"id"`
	OptionIds []int `json:"option_ids" form:"option_ids"`
}

// 获取商品sku
func (m *Sku) GetSku(form *SkuGetFormType) (*ent.GoodsSku, error) {
	var whereArr []predicate.GoodsSku
	whereArr = append(whereArr, goodssku.HasGoodsSpuWith(goodsspu.IDEQ(form.ID)))
	for _, v := range form.OptionIds {
		whereArr = append(whereArr, goodssku.HasGoodsSpecsOptionWith(goodsspecsoption.And(
			goodsspecsoption.IDEQ(v),
		)))
	}
	sku, err2 := m.db.
		Query().
		Where(whereArr...).
		WithGoodsSpecsOption(func(query *ent.GoodsSpecsOptionQuery) {
			query.WithGoodsSpecs()
		}).
		First(m.ctx)
	if err2 != nil {
		global.Logger.Error(err2.Error())
		return sku, err2
	}
	return sku, nil
}

// 校验库存
func (m Sku) CheckGetSku(products []*order_utils.OrderCreateGoodsItem) (bool, []*ent.GoodsSku, map[int]int, error) {
	if products == nil {
		return false, nil, nil, errors.New("products is nil")
	}
	ids, maps := order_utils.ProductsToIds(products)
	//获取下单中的所有sku
	skus, err := m.db.Query().
		Where(goodssku.IDIn(ids...)).
		WithGoodsSpu().
		WithGoodsSpecsOption(func(query *ent.GoodsSpecsOptionQuery) {
			query.WithGoodsSpecs()
		}).
		All(m.ctx)
	if err != nil {
		global.Logger.Error(err.Error())
		return false, skus, maps, err
	}
	for _, sku := range skus {
		if maps[sku.ID] < 1 {
			return false, nil, maps, errors.New("订单商品数据异常")
		}
		if sku.StockNum < maps[sku.ID] {
			return false, nil, maps, errors.New("订单商品中部分库存不足")
		}
	}
	return true, skus, maps, nil
}

type MoldType = string

const (
	MoldTypeAdd    MoldType = "add"
	MoldTypeReduce MoldType = "reduce"
)

// 库存操作
func (m Sku) TakeOffAmount(orderSkus []*ent.OrderGoodsSku, mold MoldType) error {
	var products []*order_utils.OrderCreateGoodsItem
	for _, orderSku := range orderSkus {
		products = append(products, &order_utils.OrderCreateGoodsItem{
			SkuID:  orderSku.SkuID,
			Amount: orderSku.Amount,
		})
	}
	errTx := ent_utils.WithTx(m.ctx, global.Db, func(tx *ent.Tx) error {
		var errR error
		ids, maps := order_utils.ProductsToIds(products)
		//获取下单中的所有sku
		skusTxs, err := tx.GoodsSku.Query().
			Where(goodssku.IDIn(ids...)).
			All(m.ctx)
		if err != nil {
			global.Logger.Error(err.Error())
			return err
		}
		for _, skusTx := range skusTxs {
			if skusTx.SalesNum < maps[skusTx.ID] && mold == MoldTypeReduce {
				errR = errors.New(fmt.Sprintf("skuId [%d] 商品库存不足", skusTx.ID))
				break
			}
			if mold == MoldTypeReduce {
				_, errR = skusTx.Update().SetStockNum(skusTx.StockNum - maps[skusTx.ID]).Save(m.ctx)
				if errR != nil {
					break
				}
			}
			if mold == MoldTypeAdd {
				_, errR = skusTx.Update().SetStockNum(skusTx.StockNum + maps[skusTx.ID]).Save(m.ctx)
				if errR != nil {
					break
				}
			}

		}
		return errR
	})
	return errTx
}