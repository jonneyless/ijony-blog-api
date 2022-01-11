package models

import (
    "log"

    "blog/common"
    "blog/enums"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type Tags struct {
    Id   uint   `gorm:"primarykey"`
    Name string `form:"name" json:"name"`
}

func (model *Tags) BuildData() gin.H {
    return gin.H{
        "id":   model.Id,
        "name": model.Name,
    }
}

func (model *Tags) Load(ctx *gin.Context) *Tags {
    if err := ctx.ShouldBind(&model); err != nil {
        log.Panic(common.ErrorMsgException(enums.ParamsError, err.Error()))
    }

    return model
}

func (model *Tags) Save() *gorm.DB {
    return db().Save(model)
}

func (model *Tags) Delete() *gorm.DB {
    return db().Delete(model)
}
