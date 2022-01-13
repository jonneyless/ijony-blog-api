package middlewares

import (
	"net/http"

	"blog/common"
	"blog/enums"
	"github.com/gin-gonic/gin"
)

// NormalAuth 用户鉴权
func NormalAuth(ctx *gin.Context) {
	_, ok := ctx.Get("User")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.ErrorResponse(enums.Unauthorized, enums.ForceStatusText(enums.Unauthorized)))
		ctx.Abort()
	}

	ctx.Next()
}
