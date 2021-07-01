package auth

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/auth/min_chat_login", Auth{}.MinChatLogin)
	r.POST("/auth/refresh_token", Auth{}.RefreshToken)
}



