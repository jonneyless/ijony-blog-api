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
    Id        uint           `gorm:"primarykey"`
    Type      uint           `form:"type" json:"type"`
    ParentId  uint           `form:"parent_id" json:"parent_id" gorm:"index:idx_parent_id"`
    EntryId   uint           `form:"entry_id" json:"entry_id" gorm:"index:idx_entry_id"`
    UserId    uint           `form:"user_id" json:"user_id" gorm:"index:idx_user_id"`
    Url       string         `form:"url" json:"url"`
    Content   string         `form:"content" json:"content"`
    User      Users          `gorm:"foreignkey:UserId;references:Id"`
    DeletedAt gorm.DeletedAt `gorm:"index"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (model *Feedbacks) BuildData() gin.H {
    return gin.H{
        "id":         model.Id,
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
