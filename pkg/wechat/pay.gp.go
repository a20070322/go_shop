package wechat

import (
	"context"
	"fmt"
	"github.com/a20070322/shop-go/global"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"log"
	"time"
)

func PayInit(ctx context.Context) *Pay {
	return &Pay{
		client: global.PayClient,
		ctx:    ctx,
	}
}

type Pay struct {
	client *core.Client `json:"client"`
	ctx    context.Context
}

type FormJsapiPrepay struct {
	OrderNumber string
	OpenID      string
	Total       int32
}

func (p Pay) JsapiPrepay(form *FormJsapiPrepay) error {
	// 支付订单有效期20分钟
	timeExpire := time.Now().Add(time.Second * 60 * 20)
	svc := jsapi.JsapiApiService{Client: p.client}
	// 得到prepay_id，以及调起支付所需的参数和签名
	resp, result, err := svc.PrepayWithRequestPayment(p.ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(global.AppSetting.MinChat.AppId),
			Mchid:       core.String(global.AppSetting.WechatPay.MchID),
			Description: core.String("goShop商城-支付"),
			OutTradeNo:  core.String(form.OrderNumber),
			Attach:      core.String(""),
			TimeExpire:  &timeExpire,
			// 通知函数暂时不做调用
			NotifyUrl: core.String("https://www.weixin.qq.com/wxpay/pay.php"),
			Amount: &jsapi.Amount{
				Total: core.Int32(form.Total),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(form.OpenID),
			},
		},
	)
	fmt.Println(result)
	if err == nil {
		log.Println(resp)
	} else {
		log.Println(err)
		return err
	}
	return nil
}
