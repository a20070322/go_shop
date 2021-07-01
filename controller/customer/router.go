package customer

import (
	"github.com/a20070322/shop-go/middleware"
	"github.com/a20070322/shop-go/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	g := r.Group("customer",middleware.JwtAuth(jwt.UserGroupCustomer))
	addr := g.Group("address")
	addr.GET("list",Address{}.List)
	addr.POST("",Address{}.Create)
	addr.PUT("/:id",Address{}.Update)
	addr.DELETE("/:id",Address{}.Delete)
	//r.POST("/auth/min_chat_login", Auth{}.MinChatLogin)
	//r.POST("/auth/refresh_token", Auth{}.RefreshToken)
}
