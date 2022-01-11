package models

import (
    "crypto/md5"
    "encoding/hex"
    "log"
    "time"

    "blog/common"
    "blog/enums"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type Users struct {
    Id        uint   `gorm:"primarykey"`
    Username  string `form:"username" json:"username"`
    Password  string `form:"password" json:"password"`
    Nickname  string `form:"nickname" json:"nickname"`
    Email     string `form:"email" json:"email"`
    Website   string `form:"website" json:"website"`
    Status    uint   `json:"status"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (model *Users) BuildData() gin.H {
    return gin.H{
        "id":      model.Id,
        "name":    model.Name(),
        "email":   model.Email,
        "website": model.Website,
    }
}

func (model *Users) Name() string {
    if model.Nickname != "" {
        return model.Nickname
    }

    return model.Username
}

func (model *Users) BeforeCreate(db *gorm.DB) (err error) {
    if model.Password != "" {
        h := md5.New()
        h.Write([]byte(model.Password))
        model.Password = hex.EncodeToString(h.Sum(nil))
    }

    return
}

func (model *Users) AfterFind(db *gorm.DB) (err error) {
    model.Password = ""

    return
}

func (model *Users) Load(ctx *gin.Context) *Users {
    if err := ctx.ShouldBind(&model); err != nil {
        log.Panic(common.ErrorMsgException(enums.ParamsError, err.Error()))
    }

    return model
}

func (model *Users) Save() *gorm.DB {
    return db().Save(model)
}

func (model *Users) Delete() *gorm.DB {
    return db().Delete(model)
}
