package middlewares

import (
    "encoding/json"
    "log"
    "net/http"

    "blog/common"
    "blog/enums"

    "github.com/gin-gonic/gin"
)

// CatchError 异常捕获中间件
func CatchError(ctx *gin.Context) {
    defer func() {
        if err := recover(); err != nil {
            url := ctx.Request.URL
            method := ctx.Request.Method

            log.Printf("| url [%s] | method | [%s] | error [%s] |", url, method, err)

            var exception common.Exception
            err := json.Unmarshal([]byte(string(err.(string))), &exception)
            if err != nil {
                ctx.JSON(http.StatusBadRequest, common.ErrorResponse("未知错误，请联系管理员！"))
                ctx.Abort()
                return
            }

            var errorMessage string
            var ok bool
            errorMessage = exception.Msg
            if errorMessage == "" {
                errorMessage, ok = enums.StatusText(exception.Code)
                if !ok {
                    errorMessage = "系统异常"
                }
            }

            ctx.JSON(http.StatusBadRequest, common.ErrorResponse(exception.Code, errorMessage))
            ctx.Abort()
        }
    }()
    ctx.Next()
}
