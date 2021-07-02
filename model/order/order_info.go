package order

import (
	"context"
	"fmt"
	"github.com/a20070322/shop-go/ent"
	"github.com/a20070322/shop-go/ent/orderinfo"
	"github.com/a20070322/shop-go/ent/schema"
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/model/goods_model"
	"github.com/a20070322/shop-go/pkg/order_utils"
	"github.com/a20070322/shop-go/pkg/utils/ent_utils"

)

func InfoInit(ctx context.Context) *Info {
	art := &Info{}
	art.db = global.Db.OrderInfo
	art.ctx = ctx
	return art
}

type Info struct {
	db  *ent.OrderInfoClient
	ctx context.Context
}

type CreateFormType struct {
	CustomerId int                                 `json:"customer_id" binding:"required"`
	Products   []*order_utils.OrderCreateGoodsItem `json:"products" binding:"required"`
	Remarks    string                              `json:"remarks"`
}

// 下单操作
func (m Info) CreateOrder(orderInfo *CreateFormType) (string, error) {
	// 生成订单号
	orderCode := order_utils.GenerateOrderCode()
	checkBool, skus, productMap, err := goods_model.SkuInit(m.ctx).CheckGetSku(orderInfo.Products)
	if checkBool != true || err != nil {
		return orderCode, err
	}
	err2 := ent_utils.WithTx(m.ctx, global.Db, func(tx *ent.Tx) error {
		orderDb := tx.OrderInfo.Create().
			SetRemark(orderInfo.Remarks).
			SetOrderNumber(orderCode).
			SetCustomerID(orderInfo.CustomerId)
		var bulk []*ent.OrderGoodsSkuCreate
		for _, sku := range skus {
			var specsOption []*schema.SpecsOptionType
			specsOption = nil
			fmt.Println(sku)
			if sku.Edges.GoodsSpu.IsCustomSku == true {
				specsOption = order_utils.SpecsOptionToJson(sku.Edges.GoodsSpecsOption)
			}
			bulk = append(
				bulk,
				global.Db.OrderGoodsSku.
					Create().
					SetSpuName(sku.Edges.GoodsSpu.SpuName).
					SetSpuCode(sku.Edges.GoodsSpu.SpuCode).
					SetSpuHeadImg(sku.Edges.GoodsSpu.SpuHeadImg).
					SetSpuDesc(sku.Edges.GoodsSpu.SpuDesc).
					SetSpuDetails(sku.Edges.GoodsSpu.SpuDetails).
					SetIsCustomSku(sku.Edges.GoodsSpu.IsCustomSku).
					SetSpecsOption(specsOption).
					SetSkuID(sku.ID).
					SetSkuName(sku.SkuName).
					SetSkuCode(sku.SkuCode).
					SetPrice(sku.Price).
					SetAmount(productMap[sku.ID],
					),
			)
		}
		orderSkus, err3 := global.Db.OrderGoodsSku.CreateBulk(bulk...).Save(m.ctx)
		if err3 != nil {
			return err3
		}
		orderDb.AddOrderGoodsSku(orderSkus...)
		_, err4 := orderDb.Save(m.ctx)
		return err4
	})
	return orderCode, err2
}

// 库存校验 订单状态改变
func (m Info) HandleOrder(code string) bool {
	orderTmp, err := m.db.Query().
		Where(
			orderinfo.OrderNumberEQ(code),
			orderinfo.StatusEQ(0),
		).
		WithOrderGoodsSku().
		First(m.ctx)
	if err != nil {
		global.Logger.Error(err.Error())
		return false
	}
	if orderTmp == nil {
		global.Logger.Error("未找到此订单或此订单已被处理 %s", code)
		return false
	}
	err2 := goods_model.SkuInit(m.ctx).TakeOffAmount(orderTmp.Edges.OrderGoodsSku, goods_model.MoldTypeReduce)
	if err2 != nil {
		global.Logger.Error(err2.Error())
		_, errS := orderTmp.Update().SetStatus(110).Save(m.ctx)
		if errS != nil {
			global.Logger.Error(errS.Error())
		}
		return false
	}
	_, errS := orderTmp.Update().SetStatus(1).Save(m.ctx)
	if errS != nil {
		global.Logger.Error(errS.Error())
	}
	return true
}

// 订单超时处理
func (m Info) TimeOutOrder(code string) bool {
	orderTmp, err := m.db.Query().
		Where(
			orderinfo.OrderNumberEQ(code),
			orderinfo.Or(
				orderinfo.StatusEQ(0),
				orderinfo.StatusEQ(1),
			),
		).
		WithOrderGoodsSku().
		First(m.ctx)
	if err != nil {
		global.Logger.Error(err.Error())
		return false
	}
	if orderTmp == nil {
		global.Logger.Error("未找到此订单或此订单已被处理 %s", code)
		return false
	}
	errT := goods_model.SkuInit(m.ctx).TakeOffAmount(orderTmp.Edges.OrderGoodsSku, goods_model.MoldTypeAdd)
	if errT != nil {
		global.Logger.Error(errT.Error())
	}
	_, errS := orderTmp.Update().SetStatus(99).Save(m.ctx)
	if errS != nil {
		global.Logger.Error(errS.Error())
	}
	return true
}
