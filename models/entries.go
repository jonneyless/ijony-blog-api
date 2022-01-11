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
    Id          uint           `gorm:"primarykey"`
    CategoryId  uint           `form:"category_id" json:"category_id" gorm:"index:idx_category_id"`
    UserId      uint           `form:"user_id" json:"user_id" gorm:"index:idx_user_id"`
    Title       string         `form:"title" json:"title"`
    Content     string         `form:"content" json:"content"`
    IsPublished uint           `form:"is_published" json:"is_published"`
    User        Users          `gorm:"foreignkey:UserId;references:Id"`
    Category    Categories     `gorm:"foreignkey:CategoryId;references:Id"`
    Tags        []Tags         `gorm:"many2many:entry_tags;joinForeignKey:EntryId;joinReferences:TagId"`
    DeletedAt   gorm.DeletedAt `gorm:"index:idx_deleted_at"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

func (model *Entries) BuildListData() gin.H {
    var tags []gin.H

    if model.Tags != nil {
        for _, tag := range model.Tags {
            tags = append(tags, tag.BuildData())
        }
    }

    return gin.H{
        "id":         model.Id,
        "category":   model.Category.BuildData(),
        "author":     model.User.BuildData(),
        "title":      model.Title,
        "content":    model.Content,
        "tags":       tags,
        "created_at": model.CreatedAt.Format(enums.TimeLayout),
    }
}

func (model *Entries) BuildViewData() gin.H {
    var tagIds []uint

    if model.Tags != nil {
        for _, tag := range model.Tags {
            tagIds = append(tagIds, tag.Id)
        }
    }

    return gin.H{
        "id":          model.Id,
        "category_id": model.CategoryId,
        "title":       model.Title,
        "content":     model.Content,
        "tag_ids":     tagIds,
    }
}

func (model *Entries) Load(ctx *gin.Context) *Entries {
    if err := ctx.ShouldBind(&model); err != nil {
        log.Panic(common.ErrorMsgException(enums.ParamsError, err.Error()))
    }

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
