package test

import (
	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	g := r.Group("/test")
	// middleware.JwtAuth(jwt.UserGroupCustomer),
	g.POST("/test", Test{}.Test)
	g.Any("/test2", Test{}.Test2)
	g.GET("/test3", Test{}.Test3)
	g.Any("/test4", Test{}.Test4)
	g.POST("/test5", Test{}.Test5)
}
