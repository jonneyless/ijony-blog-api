package middlewares

import (
	"net/http"

	"blog/common"
	"blog/enums"
	"blog/models"
	"github.com/gin-gonic/gin"
)

// AdminAuth 用户鉴权
func AdminAuth(ctx *gin.Context) {
	user, ok := ctx.Get("User")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.ErrorResponse(enums.Unauthorized, enums.ForceStatusText(enums.Unauthorized)))
		ctx.Abort()
	} else {
		model := user.(*models.Users)

		if !model.IsAdmin() {
			ctx.JSON(http.StatusForbidden, common.ErrorResponse(enums.Forbidden, enums.ForceStatusText(enums.Forbidden)))
			ctx.Abort()
		}
	}

	ctx.Next()
}
