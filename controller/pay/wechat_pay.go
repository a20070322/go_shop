package pay

import (
	"github.com/a20070322/shop-go/model/pay"
	"github.com/a20070322/shop-go/pkg/utils/jwt"
	"github.com/a20070322/shop-go/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"net/http"
)

type WechatPay struct {
}

// 生成微信支付
func (t WechatPay) Create(ctx *gin.Context) {
	var form pay.WeChatPayCreateFormType
	uid, errJwt := jwt.GetTokenId(ctx)
	if errJwt != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, errJwt.Error(), nil)
		return
	}
	form.CustomerId = uid
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	rep, err := pay.WeChatPayPayInit(ctx).Create(&form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

func (t WechatPay) Notify(ctx *gin.Context) {


	ctx.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
}