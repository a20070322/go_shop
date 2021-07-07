package initialize

import (
	"context"
	"fmt"
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/types"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"io/ioutil"
)

func WechatPay() {
	config := global.AppSetting.WechatPay
	keyData, err := ioutil.ReadFile(config.MchPrivateKey)
	if err != nil {
		panic("微信支付秘钥缺失")
		return
	}
	privateKey, err := utils.LoadPrivateKey(string(keyData))
	if err != nil {
		panic("failed to parse DER encoded public key:" + err.Error())
	}
	ctx := context.Background()
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(config.MchID, config.MchCertificateSerialNumber, privateKey, config.MchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		panic(fmt.Sprintf("new wechat pay client err:%s", err))
	}


	certVisitor := downloader.MgrInstance().GetCertificateVisitor(config.MchID)


	global.Logger.Info(certVisitor)
	handler := notify.NewNotifyHandler(config.MchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certVisitor))

	global.WeChatPay = &types.WeChatPay{
		PayClient: client,
		Handler:   handler,
	}
	fmt.Println("微信支付请求初始化完成")
}
