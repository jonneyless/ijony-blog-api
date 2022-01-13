package models

import (
	"log"

	"blog/common"
	"blog/enums"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Tags struct {
	ID   uint   `gorm:"comment:'标签ID';primarykey"`
	Name string `form:"name" json:"name" gorm:"type:varchar(30);not null;comment:'标签名称';index:idx_name"`
}

func (model *Tags) BuildData() gin.H {
	return gin.H{
		"id":   model.ID,
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
