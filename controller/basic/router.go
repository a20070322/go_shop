package basic

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	g := r.Group("/banner")
	g.GET("/list/:id", Banner{}.List)
}
