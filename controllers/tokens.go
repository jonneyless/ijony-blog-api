package controllers

import (
	"net/http"

	"blog/common"
	"blog/services"
	"github.com/gin-gonic/gin"
)

type tokenCtr struct{}

func registerTokens(router *gin.RouterGroup) {
	ctr := &tokenCtr{}

	router.POST("/tokens", ctr.create)
}

func (ctr *tokenCtr) create(ctx *gin.Context) {
	token := services.GetToken(ctx)

	ctx.JSON(http.StatusOK, common.SuccessResponse("登录成功", gin.H{"data": gin.H{"token": token}}))
}
