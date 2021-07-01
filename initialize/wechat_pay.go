package initialize

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/a20070322/shop-go/global"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"io/ioutil"
)

func WechatPay()  {
	config := global.AppSetting.WechatPay
	certData, err := ioutil.ReadFile(config.MchPrivateKey)
	if err != nil {
		panic("微信支付秘钥缺失")
		return
	}
	block, _ := pem.Decode([]byte(certData))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded public key:" + err.Error())
	}
	ctx := context.Background()
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(config.MchID, config.MchCertificateSerialNumber, privateKey.(* rsa.PrivateKey), config.MchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		panic(fmt.Sprintf("new wechat pay client err:%s", err))
	}
	global.PayClient = client
	fmt.Println("微信支付请求初始化完成")
}