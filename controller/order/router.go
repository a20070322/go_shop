package order

import (
	"github.com/a20070322/shop-go/middleware"
	"github.com/a20070322/shop-go/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup)  {
	g := r.Group("order")
	g.POST("/create",middleware.JwtAuth(jwt.UserGroupCustomer), Order{}.Create)
	g.POST("/pre_create",middleware.JwtAuth(jwt.UserGroupCustomer), Order{}.PreCreateOrder)
	g.POST("/state",middleware.JwtAuth(jwt.UserGroupCustomer), Order{}.StateOrder)
}
