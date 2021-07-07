package order

import (
	"context"
	"errors"
	"github.com/a20070322/shop-go/ent"
	"github.com/a20070322/shop-go/ent/customer"
	"github.com/a20070322/shop-go/ent/customeraddress"
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
	AddressId  int                                 `json:"address_id" binding:"required"`
	Products   []*order_utils.OrderCreateGoodsItem `json:"products" binding:"required"`
	Remarks    string                              `json:"remarks"`
}

// 下单操作
func (m Info) CreateOrder(orderInfo *CreateFormType) (string, error) {
	// 生成订单号
	orderCode := order_utils.GenerateOrderCode()
	checkBool, skus, productMap, err := goods_model.SkuInit(m.ctx).CheckGetSku(orderInfo.Products, true)
	if checkBool != true || err != nil {
		return orderCode, err
	}
	// 获取用户下单地址
	addr, errAddr := global.Db.CustomerAddress.Query().Where(customeraddress.IDEQ(orderInfo.AddressId)).First(m.ctx)
	if addr == nil || errAddr != nil {
		return orderCode, errors.New("收货地址异常")
	}

	err2 := ent_utils.WithTx(m.ctx, global.Db, func(tx *ent.Tx) error {
		// 订单基础信息创建
		orderDb := tx.OrderInfo.Create().
			SetRemark(orderInfo.Remarks).
			SetOrderNumber(orderCode).
			SetCustomerID(orderInfo.CustomerId)
		// 订单地址创建
		orderAddr, orderAddrError := tx.OrderAddress.Create().
			SetName(addr.Name).
			SetPhone(addr.Phone).
			SetProvince(addr.Province).
			SetCity(addr.City).
			SetArea(addr.Area).
			SetDetailed(addr.Detailed).
			SetRemark(addr.Remark).
			Save(m.ctx)
		// 支付金额
		payMoney := 0
		if orderAddrError != nil {
			return orderAddrError
		}
		var bulk []*ent.OrderGoodsSkuCreate
		for _, sku := range skus {
			var specsOption []*schema.SpecsOptionType
			specsOption = nil
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
			payMoney += sku.Price * productMap[sku.ID]
		}
		orderSkus, err3 := global.Db.OrderGoodsSku.CreateBulk(bulk...).Save(m.ctx)
		if err3 != nil {
			return err3
		}
		orderDb.AddOrderGoodsSku(orderSkus...)
		orderDb.AddOrderAddress(orderAddr)
		orderDb.SetPayMoney(payMoney)
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

type PreCreateFormType struct {
	CustomerId int                                 `json:"customer_id"`
	Products   []*order_utils.OrderCreateGoodsItem `json:"products" binding:"required"`
}
type PreCreateRepType struct {
	GoodList       []*PreCreateItemType `json:"good_list"`
	DefaultAddress *ent.CustomerAddress `json:"default_address"`
}
type PreCreateItemType struct {
	StockNum int    `json:"stock_num"`
	SpuName  string `json:"spu_name"`
	Image    string `json:"image"`
	SkuName  string `json:"sku_name"`
	Price    int    `json:"price"`
	Count    int    `json:"count"`
	SkuId    int    `json:"sku_id"`
}

// 预下单下单操作
func (m Info) PreCreateOrder(orderInfo *PreCreateFormType) (*PreCreateRepType, error) {
	var rep PreCreateRepType
	_, skus, preMap, err := goods_model.SkuInit(m.ctx).CheckGetSku(orderInfo.Products, false)
	//if checkBool != true || err != nil {
	//	return nil, err
	//}
	addr, errAddr := global.Db.CustomerAddress.
		Query().
		Where(
			customeraddress.HasCustomerWith(customer.IDEQ(orderInfo.CustomerId)),
			customeraddress.IsDefaultEQ(true),
		).
		First(m.ctx)
	if errAddr == nil {
		rep.DefaultAddress = addr
	}
	for _, v := range skus {
		var tmp = &PreCreateItemType{
			StockNum: v.StockNum,
			SpuName:  v.Edges.GoodsSpu.SpuName,
			Image:    v.Edges.GoodsSpu.SpuHeadImg,
			SkuName:  "",
			Price:    v.Price,
			Count:    preMap[v.ID],
			SkuId:    v.ID,
		}
		if v.Edges.GoodsSpu.IsCustomSku {
			for _, optionSelect := range v.Edges.GoodsSpecsOption {
				tmp.SkuName += optionSelect.SpecsOptionValue + " "
			}
		}
		rep.GoodList = append(rep.GoodList, tmp)
	}
	return &rep, err
}

type StateOrderFormType struct {
	CustomerId int    `json:"customer_id" binding:"required"`
	OrderCode  string `json:"order_code" form:"option_ids" binding:"required"`
}

// 查看订单状态
func (m Info) StateOrder(form *StateOrderFormType) (int8, error) {
	orderInfo, err := m.db.Query().
		Where(
			orderinfo.OrderNumberEQ(form.OrderCode),
			orderinfo.HasCustomerWith(customer.IDEQ(form.CustomerId)),
		).
		First(m.ctx)
	if err != nil {
		global.Logger.Error(err.Error())
		return -1, errors.New("订单查询失败")
	}
	// // 0 待创建订单 1 待支付  2 支付完成   99 已取消 100 已完成  110 下单失败
	if orderInfo.Status == 110 {
		return -1, errors.New("下单失败")
	}
	if orderInfo.Status == 99 {
		return -1, errors.New("该订单已取消")
	}
	if orderInfo.Status > 1 {
		return -1, errors.New("订单状态异常")
	}
	return orderInfo.Status, nil
}
