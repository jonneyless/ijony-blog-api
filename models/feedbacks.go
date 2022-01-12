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

type Feedbacks struct {
	ID        uint           `gorm:"comment:'反馈ID';primarykey"`
	Type      uint           `form:"type" json:"type" gorm:"type:tinyint(1);not null;comment:'反馈类型'"`
	ParentId  uint           `form:"parent_id" json:"parent_id" gorm:"type:bigint(20);not null;comment:'父级ID';index:idx_parent_id"`
	EntryId   uint           `form:"entry_id" json:"entry_id" gorm:"type:bigint(20);not null;comment:'内容ID';index:idx_entry_id"`
	UserId    uint           `form:"user_id" json:"user_id" gorm:"type:bigint(20);not null;comment:'用户ID';index:idx_user_id"`
	Url       string         `form:"url" json:"url" gorm:"type:varchar(255);default:'';comment:'关联地址'"`
	Content   string         `form:"content" json:"content" gorm:"type:text;comment:'反馈内容'"`
	Liked     uint           `json:"liked" gorm:"type:int(10);not null;default:0;comment:'点赞数'"`
	Trample   uint           `json:"trample" gorm:"type:int(10);not null;default:0;comment:'点赞数'"`
	DeletedAt gorm.DeletedAt `gorm:"comment:'删除时间';index:idx_deleted_at"`
	CreatedAt time.Time      `gorm:"comment:'创建时间';index:idx_created_at"`
	UpdatedAt time.Time      `gorm:"comment:'更新时间';index:idx_updated_at"`

	// 关联数据
	User Users `gorm:"foreignkey:UserId;references:ID"`
}

func (model *Feedbacks) BuildData() gin.H {
	return gin.H{
		"id":         model.ID,
		"type":       model.Type,
		"parent_id":  model.ParentId,
		"author":     model.User.BuildData(),
		"url":        model.Url,
		"content":    model.Content,
		"created_at": model.CreatedAt.Format(enums.TimeLayout),
	}
}

func (model *Feedbacks) Load(ctx *gin.Context) *Feedbacks {
	if err := ctx.ShouldBind(&model); err != nil {
		log.Panic(common.ErrorMsgException(enums.ParamsError, err.Error()))
	}

	if model.EntryId == 0 {
		entryId, _ := strconv.Atoi(ctx.Param("entryId"))
		model.EntryId = uint(entryId)
	}

	return model
}

func (model *Feedbacks) Save() *gorm.DB {
	return db().Save(model)
}

func (model *Feedbacks) Delete() *gorm.DB {
	return db().Delete(model)
}
