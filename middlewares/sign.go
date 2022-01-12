package middlewares

import (
	"net/http"
	"net/url"

	"blog/common"
	"github.com/gin-gonic/gin"
)

// Sign 验签中间件
func Sign(ctx *gin.Context) {
	if gin.Mode() == gin.ReleaseMode {
		sign := ctx.Request.Header.Get("Sign")
		signTime := ctx.Request.Header.Get("SignTime")

		var params url.Values

		if ctx.Request.Method == "GET" {
			params = ctx.Request.URL.Query()
		} else {
			_ = ctx.Request.ParseMultipartForm(10e6)
			params = ctx.Request.PostForm
		}

		if !common.CheckSign(sign, signTime, params) {
			ctx.JSON(http.StatusForbidden, common.SuccessResponse("请求非法"))
			ctx.Abort()
		}
	}

	ctx.Next()
}
