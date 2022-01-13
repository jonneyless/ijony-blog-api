package models

import (
	"log"
	"time"

	"blog/common"
	"blog/enums"
	"github.com/chenhg5/collection"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Entries struct {
	ID            uint           `gorm:"comment:'内容ID';primarykey"`
	CategoryId    uint           `form:"category_id" json:"category_id" gorm:"comment:'分类ID';index:idx_category_id"`
	UserId        uint           `json:"user_id" gorm:"comment:'用户ID';index:idx_user_id"`
	Title         string         `form:"title" json:"title" gorm:"type:varchar(255);not null;comment:'标题'"`
	Summary       string         `form:"summary" json:"summary" gorm:"type:varchar(255);not null;default:'';comment:'摘要'"`
	Content       string         `form:"content" json:"content" gorm:"type:longtext;not null;comment:'内容'"`
	Trackback     string         `form:"trackback" json:"trackback" gorm:"type:varchar(255);not null;default:'';comment:'通告地址'"`
	IsPublished   uint           `form:"is_published" json:"is_published" gorm:"type:tinyint(1) unsigned;not null;default:0;comment:'发布状态'"`
	ViewCount     uint           `json:"view_count" gorm:"default:0;comment:'阅读数'"`
	FeedbackCount uint           `json:"feedback_count" gorm:"default:0;comment:'反馈数'"`
	DeletedAt     gorm.DeletedAt `gorm:"comment:'删除时间';index:idx_deleted_at"`
	CreatedAt     time.Time      `gorm:"comment:'创建时间';index:idx_created_at"`
	UpdatedAt     time.Time      `gorm:"comment:'更新时间';index:idx_updated_at"`

	// 关联数据
	User     Users      `gorm:"foreignkey:UserId;references:ID"`
	Category Categories `gorm:"foreignkey:CategoryId;references:ID"`
	Tags     []Tags     `gorm:"many2many:entry_tags;joinForeignKey:EntryId;joinReferences:TagId"`
}

func (model *Entries) BuildData() gin.H {
	var tags []gin.H

	if model.Tags != nil {
		for _, tag := range model.Tags {
			tags = append(tags, tag.BuildData())
		}
	}

	return gin.H{
		"id":         model.ID,
		"title":      model.Title,
		"summary":    model.Summary,
		"content":    model.Content,
		"views":      model.ViewCount,
		"feedbacks":  model.FeedbackCount,
		"created_at": model.CreatedAt.Format(enums.TimeLayout),
		"category":   model.Category.BuildData(),
		"author":     model.User.BuildData(),
		"tags":       tags,
	}
}

func (model *Entries) BuildDetail() gin.H {
	data := model.BuildData()

	data["trackback"] = model.Trackback
	data["is_published"] = model.IsPublished

	return data
}

func (model *Entries) CheckOrSetAuth(userId uint) {
	if model.UserId == 0 {
		model.UserId = userId
	}

	if model.UserId != userId {
		log.Panic(common.ErrorMsgException(enums.Forbidden, "不能修改别人的日志"))
	}
}

func (model *Entries) BeforeUpdate(db *gorm.DB) (err error) {
	db.Model(&EntryTags{}).Where("entry_id = ?", model.ID).Delete(&EntryTags{})

	return err
}

func (model *Entries) Load(ctx *gin.Context) *Entries {
	if err := ctx.ShouldBind(&model); err != nil {
		log.Panic(common.ErrorMsgException(enums.ParamsError, err.Error()))
	}

	user := GetAdminIdentity(ctx)
	model.CheckOrSetAuth(user.ID)

	tags := ctx.PostFormArray("tags[]")
	if len(tags) > 0 {
		db().Model(&Tags{}).Where("name in ?", tags).First(&model.Tags)

		var existsTags []string
		for _, tagModel := range model.Tags {
			existsTags = append(existsTags, tagModel.Name)
		}

		tags = collection.Collect(tags).Diff(existsTags).ToStringArray()

		for _, tag := range tags {
			model.Tags = append(model.Tags, Tags{Name: tag})
		}

		log.Println(model.Tags)
	}

	return model
}

func (model *Entries) Save() *gorm.DB {
	return db().Save(model)
}

func (model *Entries) Delete() *gorm.DB {
	return db().Delete(model)
}
