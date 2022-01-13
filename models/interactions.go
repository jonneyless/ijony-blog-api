package models

import (
	"log"
	"strconv"
	"time"

	"blog/common"
	"blog/enums"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Interactions struct {
	FeedbackId uint      `form:"feedback_id" json:"feedback_id" gorm:"comment:'反馈ID';primarykey;autoIncrement:false"`
	UserId     uint      `form:"user_id" json:"user_id" gorm:"comment:'用户ID';primarykey;autoIncrement:false"`
	Type       uint      `form:"type" json:"type" gorm:"type:tinyint(1) unsigned;not null;default:1;comment:'互动类型'"`
	CreatedAt  time.Time `gorm:"comment:'创建时间';index:idx_created_at"`
	UpdatedAt  time.Time `gorm:"comment:'更新时间';index:idx_updated_at"`
}

func (model *Interactions) AfterCreate(db *gorm.DB) (err error) {
	if model.Type == enums.InteractionTypeLiked {
		db.Model(&Feedbacks{}).Where("id = ?", model.FeedbackId).UpdateColumn("liked", gorm.Expr("liked + ?", 1))
	} else {
		db.Model(&Feedbacks{}).Where("id = ?", model.FeedbackId).UpdateColumn("trample", gorm.Expr("trample + ?", 1))
	}

	return
}

func (model *Interactions) AfterUpdate(db *gorm.DB) (err error) {
	if model.Type == enums.InteractionTypeLiked {
		db.Model(&Feedbacks{}).Where("id = ?", model.FeedbackId).UpdateColumns(map[string]interface{}{"liked": gorm.Expr("liked + ?", 1), "trample": gorm.Expr("trample - ?", 1)})
	} else {
		db.Model(&Feedbacks{}).Where("id = ?", model.FeedbackId).UpdateColumns(map[string]interface{}{"liked": gorm.Expr("liked - ?", 1), "trample": gorm.Expr("trample + ?", 1)})
	}

	return
}

func (model *Interactions) Load(ctx *gin.Context) *Interactions {
	if err := ctx.ShouldBind(&model); err != nil {
		log.Panic(common.ErrorMsgException(enums.ParamsError, err.Error()))
	}

	if model.FeedbackId == 0 {
		feedbackId, _ := strconv.Atoi(ctx.Param("feedbackId"))
		model.FeedbackId = uint(feedbackId)
	}

	return model
}

func (model *Interactions) Save() *gorm.DB {
	return db().Save(model)
}

func (model *Interactions) Delete() *gorm.DB {
	return db().Delete(model)
}
