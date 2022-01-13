package controllers

import (
	"log"
	"net/http"

	"blog/common"
	"blog/enums"
	"blog/middlewares"
	"blog/services"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

type entryCtr struct{}

func registerEntries(router *gin.RouterGroup) {
	ctr := &entryCtr{}

	router.GET("/entries", ctr.index)
	router.GET("/entries/:entryId", ctr.view)

	router.POST("/entries", middlewares.AdminAuth, ctr.create)
	router.PUT("/entries/:entryId", middlewares.AdminAuth, ctr.update)
	router.DELETE("/entries/:entryId", middlewares.AdminAuth, ctr.delete)
}

func (ctr *entryCtr) index(ctx *gin.Context) {
	entries, err := services.GetEntries(ctx)
	if err != nil {
		log.Panic(common.ErrorMsgException(enums.ParamsError, err.Error()))
	}

	items := make([]gin.H, 0)
	for _, entry := range entries {
		items = append(items, entry.BuildData())
	}

	ctx.JSON(http.StatusOK, common.SuccessResponse(gin.H{"data": gin.H{"items": items}}))
}

func (ctr *entryCtr) view(ctx *gin.Context) {
	var data gin.H

	model := services.GetEntry(ctx)

	if !utils.IsBackend(ctx) {
		services.IncEntryView(model)
		data = model.BuildData()
	} else {
		data = model.BuildDetail()
	}

	ctx.JSON(http.StatusOK, common.SuccessResponse(gin.H{"data": data}))
}

func (ctr *entryCtr) create(ctx *gin.Context) {
	model := services.NewEntry()

	err := model.Load(ctx).Save().Error
	if err != nil {
		log.Panic(common.ErrorMsgException(enums.SaveError, err.Error()))
	}

	ctx.JSON(http.StatusOK, common.SuccessResponse("已创建日志", gin.H{"data": model.BuildDetail()}))
}

func (ctr *entryCtr) update(ctx *gin.Context) {
	model := services.GetEntry(ctx)

	err := model.Load(ctx).Save().Error
	if err != nil {
		log.Panic(common.ErrorMsgException(enums.SaveError, err.Error()))
	}

	ctx.JSON(http.StatusOK, common.SuccessResponse("日志更新成功", gin.H{"data": model.BuildDetail()}))
}

func (ctr *entryCtr) delete(ctx *gin.Context) {
	model := services.GetEntry(ctx)

	err := model.Delete().Error
	if err != nil {
		log.Panic(common.ErrorMsgException(enums.DeleteError, err.Error()))
	}

	ctx.JSON(http.StatusOK, common.SuccessResponse("日志删除成功"))
}
