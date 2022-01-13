package utils

import (
	"crypto/md5"
	"encoding/hex"

	"blog/enums"
	"github.com/gin-gonic/gin"
)

// Md5 字符串MD5
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// IsBackend 判断是否后台请求
func IsBackend(ctx *gin.Context) bool {
	endpoint, ok := ctx.Get("Endpoint")
	if !ok {
		return false
	}

	return endpoint == enums.EndpointBack
}
