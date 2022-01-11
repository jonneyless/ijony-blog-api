package middlewares

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// Cors 跨域处理中间件
func Cors(ctx *gin.Context) {
    ctx.Writer.Header().Add("Access-Control-Allow-Origin", "*")
    ctx.Writer.Header().Add("Access-Control-Allow-Headers", "sign,signtime")
    method := ctx.Request.Method
    if method == "OPTIONS" {
        ctx.Writer.Header().Set("Allow", "GET,POST,PUT,DELETE")
        ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
        ctx.JSON(http.StatusOK, gin.H{})
        ctx.Abort()
    }
    ctx.Next()
}
