package controllers

import (
	"log"
	"net/http"

	"blog/common"
	"blog/enums"
	"blog/services"

	"github.com/gin-gonic/gin"
)

type interactionCtr struct{}

func registerInteractions(router *gin.RouterGroup) {
	ctr := &interactionCtr{}

	router.POST("/entries/:entryId/feedbacks/:feedbackId/interactions", ctr.interaction)
}

func (ctr *interactionCtr) interaction(ctx *gin.Context) {
	feedback := services.GetFeedback(ctx)
	if feedback == nil {
		log.Panic(common.ErrorMsgException(enums.NotFound, "评论不存在"))
	}

	model := services.NewInteraction()
	err := model.Load(ctx).Save().Error
	if err != nil {
		log.Panic(common.ErrorMsgException(enums.SaveError, err.Error()))
	}

	ctx.JSON(http.StatusOK, common.SuccessResponse())
}
