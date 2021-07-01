package goods

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	g := r.Group("/goods")
	g.GET("/list", Spu{}.List)
	g.GET("/get", Spu{}.GetID)
	g.POST("/get_sku", Sku{}.GetSku)
	g.GET("/classify", Classify{}.ConsumerGet)
}
