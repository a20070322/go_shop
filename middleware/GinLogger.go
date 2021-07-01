package middleware

import (
	"fmt"
	"github.com/a20070322/shop-go/global"

	"github.com/gin-gonic/gin"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		c.Next()
		cost := time.Since(start).Milliseconds()
		global.Logger.Info(fmt.Sprintf("[%d]   %dms | %s | %s   \"%s\"",c.Writer.Status(),cost,c.ClientIP(),c.Request.Method,path))
	}
}