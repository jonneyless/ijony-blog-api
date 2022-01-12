package controllers

import (
	"log"
	"net/http"

	"blog/common"
	"blog/enums"
	"blog/services"

	"github.com/gin-gonic/gin"
)

type feedbackCtr struct{}

func registerFeedbacks(router *gin.RouterGroup) {
	ctr := &feedbackCtr{}

	router.GET("/entries/:entryId/feedbacks", ctr.index)
	router.GET("/entries/:entryId/feedbacks/:feedbackId", ctr.view)
	router.POST("/entries/:entryId/feedbacks", ctr.create)
	router.PUT("/entries/:entryId/feedbacks/:feedbackId", ctr.update)
	router.DELETE("/entries/:entryId/feedbacks/:feedbackId", ctr.delete)
}

func (ctr *feedbackCtr) index(ctx *gin.Context) {
	feedbacks, err := services.GetFeedbacks(ctx)
	if err != nil {
		log.Panic(common.ErrorMsgException(enums.ParamsError, err.Error()))
	}

	items := make([]gin.H, 0)
	for _, feedback := range feedbacks {
		items = append(items, feedback.BuildData())
	}

	ctx.JSON(http.StatusOK, common.SuccessResponse(gin.H{"data": gin.H{"items": items}}))
}

func (ctr *feedbackCtr) view(ctx *gin.Context) {
	model := services.GetFeedback(ctx)

	ctx.JSON(http.StatusOK, common.SuccessResponse(gin.H{"data": model.BuildData()}))
}

func (ctr *feedbackCtr) create(ctx *gin.Context) {
	model := services.NewFeedback()

	err := model.Load(ctx).Save().Error
	if err != nil {
		log.Panic(common.ErrorMsgException(enums.SaveError, err.Error()))
	}

	ctx.JSON(http.StatusOK, common.SuccessResponse("已创建日志", gin.H{"data": model.BuildData()}))
}

func (ctr *feedbackCtr) update(ctx *gin.Context) {
	model := services.GetFeedback(ctx)

	err := model.Load(ctx).Save().Error
	if err != nil {
		log.Panic(common.ErrorMsgException(enums.SaveError, err.Error()))
	}

	ctx.JSON(http.StatusOK, common.SuccessResponse("日志更新成功", gin.H{"data": model.BuildData()}))
}

func (ctr *feedbackCtr) delete(ctx *gin.Context) {
	model := services.GetFeedback(ctx)

	err := model.Delete().Error
	if err != nil {
		log.Panic(common.ErrorMsgException(enums.DeleteError, err.Error()))
	}

	ctx.JSON(http.StatusOK, common.SuccessResponse("日志删除成功", gin.H{"data": model.BuildData()}))
}
