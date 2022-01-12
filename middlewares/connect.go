package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Connect 数据库连接中间件
func Connect(ctx *gin.Context) {

	ctx.Next()
}
