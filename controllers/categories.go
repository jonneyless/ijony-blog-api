package controllers

import (
    "log"
    "net/http"

    "blog/common"
    "blog/enums"
    "blog/services"

    "github.com/gin-gonic/gin"
)

type categoryCtr struct{}

func registerCategories(router *gin.RouterGroup) {
    ctr := &categoryCtr{}

    router.GET("/categories", ctr.index)
    router.GET("/categories/:id", ctr.view)
    router.POST("/categories", ctr.create)
    router.PUT("/categories/:id", ctr.update)
    router.DELETE("/categories/:id", ctr.delete)
}

func (ctr *categoryCtr) index(ctx *gin.Context) {
    categories, err := services.GetCategories(ctx)
    if err != nil {
        log.Panic(common.ErrorMsgException(enums.ParamsError, err.Error()))
    }

    items := make([]gin.H, 0)
    for _, category := range categories {
        items = append(items, category.BuildData())
    }

    ctx.JSON(http.StatusOK, common.SuccessResponse(gin.H{"data": gin.H{"items": items}}))
}

func (ctr *categoryCtr) view(ctx *gin.Context) {
    model := services.GetCategory(ctx)

    ctx.JSON(http.StatusOK, common.SuccessResponse(gin.H{"data": model.BuildData()}))
}

func (ctr *categoryCtr) create(ctx *gin.Context) {
    model := services.NewCategory()

    err := model.Load(ctx).Save().Error
    if err != nil {
        log.Panic(common.ErrorMsgException(enums.SaveError, err.Error()))
    }

    ctx.JSON(http.StatusOK, common.SuccessResponse("已创建分类", gin.H{"data": model.BuildData()}))
}

func (ctr *categoryCtr) update(ctx *gin.Context) {
    model := services.GetCategory(ctx)

    err := model.Load(ctx).Save().Error
    if err != nil {
        log.Panic(common.ErrorMsgException(enums.SaveError, err.Error()))
    }

    ctx.JSON(http.StatusOK, common.SuccessResponse("分类更新成功", gin.H{"data": model.BuildData()}))
}

func (ctr *categoryCtr) delete(ctx *gin.Context) {
    model := services.GetCategory(ctx)

    err := model.Delete().Error
    if err != nil {
        log.Panic(common.ErrorMsgException(enums.DeleteError, err.Error()))
    }

    ctx.JSON(http.StatusOK, common.SuccessResponse("分类删除成功", gin.H{"data": model.BuildData()}))
}
