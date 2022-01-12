package controllers

import (
	"log"
	"net/http"

	"blog/common"
	"blog/enums"
	"blog/services"
	"github.com/gin-gonic/gin"
)

type userCtr struct{}

func registerUsers(router *gin.RouterGroup) {
	ctr := &userCtr{}

	router.GET("/users", ctr.index)
	router.GET("/users/:userId", ctr.view)
	router.POST("/users", ctr.create)
	router.PUT("/users/:userId", ctr.update)
	router.DELETE("/users/:userId", ctr.delete)
}

func (ctr *userCtr) index(ctx *gin.Context) {
	entries, err := services.GetUsers(ctx)
	if err != nil {
		log.Panic(common.ErrorMsgException(enums.ParamsError, err.Error()))
	}

	items := make([]gin.H, 0)
	for _, user := range entries {
		items = append(items, user.BuildListData())
	}

	ctx.JSON(http.StatusOK, common.SuccessResponse(gin.H{"data": gin.H{"items": items}}))
}

func (ctr *userCtr) view(ctx *gin.Context) {
	model := services.GetUser(ctx)

	ctx.JSON(http.StatusOK, common.SuccessResponse(gin.H{"data": model.BuildViewData()}))
}

func (ctr *userCtr) create(ctx *gin.Context) {
	model := services.NewUser()

	err := model.Load(ctx).Save().Error
	if err != nil {
		log.Panic(common.ErrorMsgException(enums.SaveError, err.Error()))
	}

	ctx.JSON(http.StatusOK, common.SuccessResponse("已创建用户", gin.H{"data": model.BuildViewData()}))
}

func (ctr *userCtr) update(ctx *gin.Context) {
	model := services.GetUser(ctx)

	err := model.Load(ctx).Save().Error
	if err != nil {
		log.Panic(common.ErrorMsgException(enums.SaveError, err.Error()))
	}

	ctx.JSON(http.StatusOK, common.SuccessResponse("用户更新成功", gin.H{"data": model.BuildViewData()}))
}

func (ctr *userCtr) delete(ctx *gin.Context) {
	model := services.GetUser(ctx)

	err := model.Delete().Error
	if err != nil {
		log.Panic(common.ErrorMsgException(enums.DeleteError, err.Error()))
	}

	ctx.JSON(http.StatusOK, common.SuccessResponse("用户删除成功", gin.H{"data": model.BuildViewData()}))
}
