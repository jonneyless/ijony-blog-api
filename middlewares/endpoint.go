package middlewares

import (
	"blog/enums"
	"github.com/gin-gonic/gin"
)

// Endpoint 判断前后台
func Endpoint(ctx *gin.Context) {
	endpoint := ctx.Request.Header.Get("X-Endpoint")
	if endpoint != "" {
		ctx.Set("Endpoint", endpoint)
	} else {
		ctx.Set("Endpoint", enums.EndpointFront)
	}

	ctx.Next()
}
