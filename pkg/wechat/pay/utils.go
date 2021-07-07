package pay

import (
	"github.com/a20070322/shop-go/global"
	"github.com/go-pay/gopay/wechat/v3"
)

// 公共解密封装
type V3NotifyPub struct {
	V3NotifyReq *wechat.V3NotifyReq
	MchAPIv3Key string
}

// 解密普通支付回调中的加密订单信息
func (c V3NotifyPub) DecryptCipherText() (*wechat.V3DecryptResult, error) {
	result, err := c.V3NotifyReq.DecryptCipherText(c.MchAPIv3Key)
	if err != nil {
		global.Logger.Error(err.Error())
	}
	return result, err
}