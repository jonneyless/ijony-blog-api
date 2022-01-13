package middlewares

import (
	"blog/services"
	"github.com/gin-gonic/gin"
)

// Token 获取用户
func Token(ctx *gin.Context) {
	token := ctx.DefaultQuery("token", ctx.Request.Header.Get("X-User-Token"))
	if token != "" {
		user := services.GetUserByToken(token)

		if user != nil {
			ctx.Set("User", user)
		}
	}

	ctx.Next()
}
