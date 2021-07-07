package pay

import (
	"github.com/a20070322/shop-go/middleware"
	"github.com/a20070322/shop-go/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup)  {
	g := r.Group("pay")
	g.POST("/wechat_pay/create",middleware.JwtAuth(jwt.UserGroupCustomer), WechatPay{}.Create)
	g.Any("/wechat_pay/notify", WechatPay{}.Notify)
	//
	//g.POST("/pre_create",middleware.JwtAuth(jwt.UserGroupCustomer), Order{}.PreCreateOrder)
	//g.POST("/state",middleware.JwtAuth(jwt.UserGroupCustomer), Order{}.StateOrder)
}
