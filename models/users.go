package models

import (
	"fmt"
	"log"
	"time"

	"blog/common"
	"blog/enums"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Users struct {
	ID           uint      `gorm:"comment:'用户ID';primarykey"`
	Type         uint      `form:"type" json:"type" gorm:"type:tinyint(1) unsigned;not null;default:0;comment:'用户类型';after:id"`
	Username     string    `form:"username" json:"username" gorm:"<-:create;type:varchar(32);not null;comment:'登录名称';index:idx_username"`
	Password     string    `form:"password" json:"password" gorm:"type:varchar(64);not null;comment:'登录密码'"`
	PasswordHash string    `gorm:"-"`
	Token        string    `json:"token" gorm:"type:varchar(64);not null;default:'';comment:'登录Token';index:idx_token"`
	Nickname     string    `form:"nickname" json:"nickname" gorm:"type:varchar(32);not null;default:'';comment:'用户昵称';index:idx_nickname"`
	Email        string    `form:"email" json:"email" gorm:"type:varchar(128);not null;default:'';comment:'邮箱'"`
	Website      string    `form:"website" json:"website" gorm:"type:varchar(255);not null;default:'';comment:'网址'"`
	Status       uint      `form:"status" json:"status" gorm:"type:tinyint(1) unsigned;not null;default:9;comment:'状态'"`
	SigninAt     time.Time `gorm:"comment:'登录时间'"`
	CreatedAt    time.Time `gorm:"comment:'注册时间';index:idx_created_at"`
	UpdatedAt    time.Time `gorm:"comment:'更新时间';index:idx_updated_at"`
}

func (model *Users) BuildData() gin.H {
	return gin.H{
		"id":      model.ID,
		"name":    model.Name(),
		"email":   model.Email,
		"website": model.Website,
	}
}

func (model *Users) BuildListData() gin.H {
	return gin.H{
		"id":      model.ID,
		"name":    model.Name(),
		"email":   model.Email,
		"website": model.Website,
		"status":  model.Status,
	}
}

func (model *Users) BuildViewData() gin.H {
	return gin.H{
		"id":       model.ID,
		"username": model.Username,
		"nickname": model.Nickname,
		"email":    model.Email,
		"website":  model.Website,
		"status":   model.Status,
	}
}

func (model *Users) Name() string {
	if model.Nickname != "" {
		return model.Nickname
	}

	return model.Username
}

func (model *Users) IsAdmin() bool {
	return model.Type == enums.UserTypeAdmin
}

func (model *Users) BeforeSave(*gorm.DB) (err error) {
	if model.Password != "" {
		model.Password = utils.Md5(model.Password)
	} else {
		model.Password = model.PasswordHash
	}

	return
}

func (model *Users) AfterFind(*gorm.DB) (err error) {
	model.PasswordHash = model.Password
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

func (model *Users) CheckPassword(password string) bool {
	if password == "" {
		return false
	}

	return model.PasswordHash == utils.Md5(password)
}

func (model *Users) GenToken() string {
	model.Token = utils.Md5(fmt.Sprintf("%s%s", model.Username, time.Now()))

	err := model.Save().Error
	if err != nil {
		log.Panic(common.ErrorMsgException(enums.SaveError, err.Error()))
	}

	return model.Token
}

func GetIdentity(ctx *gin.Context) *Users {
	model := ctx.MustGet("User").(*Users)
	if model == nil {
		log.Panic(common.ErrorCodeException(enums.Unauthorized))
	}

	return model
}

func GetAdminIdentity(ctx *gin.Context) *Users {
	model := ctx.MustGet("User").(*Users)
	if model == nil {
		log.Panic(common.ErrorCodeException(enums.Unauthorized))
	}

	if !model.IsAdmin() {
		log.Panic(common.ErrorCodeException(enums.Forbidden))
	}

	return model
}
