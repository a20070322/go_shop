package pay

import (
	"context"
	"errors"
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/types/config"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"net/http"
	"time"
)

func PayInit(ctx context.Context, config *config.WechatPay) *Pay {
	return &Pay{
		Client: global.WeChatPay.PayClient,
		Ctx:    ctx,
		Config: config,
	}
}

type Pay struct {
	Client *core.Client `json:"client"`
	Ctx    context.Context
	Config *config.WechatPay
}

type FormJsapiPrepay struct {
	OrderNumber string
	OpenID      string
	Total       int32
	NotifyUrl   string
}

// 常规下单封装
func (p Pay) JsapiPrepay(form *FormJsapiPrepay) (*jsapi.PrepayWithRequestPaymentResponse, error) {
	// 支付订单有效期20分钟
	timeExpire := time.Now().Add(time.Second * 60 * 20)
	svc := jsapi.JsapiApiService{Client: p.Client}
	// 得到prepay_id，以及调起支付所需的参数和签名
	resp, _, err := svc.PrepayWithRequestPayment(p.Ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(global.AppSetting.MinChat.AppId),
			Mchid:       core.String(global.AppSetting.WechatPay.MchID),
			Description: core.String("goShop商城-支付"),
			OutTradeNo:  core.String(form.OrderNumber),
			Attach:      core.String(""),
			TimeExpire:  &timeExpire,
			NotifyUrl:   core.String(form.NotifyUrl),
			Amount: &jsapi.Amount{
				Total: core.Int32(form.Total),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(form.OpenID),
			},
		},
	)
	if err != nil {
		global.Logger.Debug(err)
		return nil, errors.New("微信支付下单异常")
	}
	return resp, nil
}

// 消息通知解码
func (p Pay) PayNotify(req *http.Request) (*V3NotifyPub, error) {
	//解析微信回调请求
	notifyReq, errNotify := wechat.V3ParseNotify(req)
	if errNotify != nil {
		global.Logger.Error(errNotify.Error())
		return nil, errNotify
	}
	certVisitor := downloader.MgrInstance().GetCertificateVisitor(p.Config.MchID)
	certMap := certVisitor.ExportAll(p.Ctx)
	var publicKey string
	for _, v := range certMap {
		publicKey = v
	}
	errReq := notifyReq.VerifySign(publicKey)
	if errReq != nil {
		global.Logger.Error(errReq.Error())
		return nil, errReq
	}
	return &V3NotifyPub{
		V3NotifyReq: notifyReq,
		MchAPIv3Key: p.Config.MchAPIv3Key,
	}, nil
}
