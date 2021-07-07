package pay

import (
	"context"
	"errors"
	"github.com/a20070322/shop-go/ent"
	"github.com/a20070322/shop-go/ent/customer"
	"github.com/a20070322/shop-go/ent/orderinfo"
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/pkg/wechat/pay"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

func WeChatPayPayInit(ctx context.Context) *WeChatPay {
	art := &WeChatPay{}
	art.db = global.Db.OrderInfo
	art.ctx = ctx
	return art
}

type WeChatPay struct {
	db  *ent.OrderInfoClient
	ctx context.Context
}

type WeChatPayCreateFormType struct {
	CustomerId int    `json:"customer_id" binding:"required"`
	OrderCode  string `json:"order_code" form:"option_ids" binding:"required"`
}

// 下单操作
// @TODO 修改商品价格
func (m WeChatPay) Create(form *WeChatPayCreateFormType) (*jsapi.PrepayWithRequestPaymentResponse, error) {
	info, err := m.db.Query().Where(
		orderinfo.HasCustomerWith(customer.IDEQ(form.CustomerId)),
		orderinfo.Status(1),
		orderinfo.OrderNumberEQ(form.OrderCode),
	).WithCustomer().First(m.ctx)
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errors.New("未找到该订单")
	}
	//0 微信支付  1 余额支付 2 支付宝支付
	_, errSave := info.Update().SetPayMethod(0).Save(m.ctx)

	if errSave != nil {
		global.Logger.Error(errSave.Error())
		return nil, errors.New("支付失败-0001")
	}

	if info.Edges.Customer.MiniOpenid == "" {
		return nil, errors.New("支付失败-002")
	}
	var payTotal int32
	if global.AppSetting.Env == "local" {
		payTotal = 1
	} else {
		payTotal = int32(info.PayMoney)
	}
	payInfo, errPay := pay.PayInit(m.ctx,global.AppSetting.WechatPay).JsapiPrepay(&pay.FormJsapiPrepay{
		OrderNumber: info.OrderNumber,
		OpenID:      info.Edges.Customer.MiniOpenid,
		Total:       payTotal,
		NotifyUrl:   "https://api.ahh5.com/go_shop_test/api/pay/wechat_pay/notify",
	})
	if errPay != nil {
		global.Logger.Error(errPay.Error())
		return nil, errors.New("支付失败-0003")
	}
	return payInfo, nil
}
