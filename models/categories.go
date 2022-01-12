package models

import (
	"log"

	"blog/common"
	"blog/enums"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Categories struct {
	ID   uint   `gorm:"comment:'分类ID';primarykey"`
	Name string `form:"name" json:"name" gorm:"type:varchar(30);not null;comment:'分类名称'"`
}

func (model *Categories) BuildData() gin.H {
	return gin.H{
		"id":   model.ID,
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
