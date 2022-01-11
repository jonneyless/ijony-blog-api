package models

import (
    "log"

    "blog/common"
    "blog/enums"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type Categories struct {
    Id   uint   `gorm:"primarykey"`
    Name string `form:"name" json:"name"`
}

func (model *Categories) BuildData() gin.H {
    return gin.H{
        "id":   model.Id,
        "name": model.Name,
    }
}

func (model *Categories) Load(ctx *gin.Context) *Categories {
    if err := ctx.ShouldBind(&model); err != nil {
        log.Panic(common.ErrorMsgException(enums.ParamsError, err.Error()))
    }

    return model
}

func (model *Categories) Save() *gorm.DB {
    return db().Save(model)
}

func (model *Categories) Delete() *gorm.DB {
    return db().Delete(model)
}
