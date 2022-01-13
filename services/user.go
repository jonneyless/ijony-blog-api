package services

import (
	"log"
	"strconv"

	"blog/common"
	"blog/enums"
	"blog/models"
	"github.com/gin-gonic/gin"
)

func NewUser() *models.Users {
	return &models.Users{}
}

func GetToken(ctx *gin.Context) string {
	var model *models.Users

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if username == "" || password == "" {
		log.Panic(common.ErrorMsgException(enums.ParamsError, "登录信息不全"))
	}

	err := db().Model(&models.Users{}).Where("username = ?", username).First(&model).Error
	if err != nil {
		log.Panic(common.ErrorMsgException(enums.NotFound, "用户不存在"))
	}

	if !model.CheckPassword(password) {
		log.Panic(common.ErrorMsgException(enums.ParamsError, "密码错误"))
	}

	return model.GenToken()
}

func GetUserByToken(token string) *models.Users {
	var item *models.Users

	err := db().First(&item, "token = ?", token).Error
	if err != nil {
		return nil
	}

	return item
}

func GetUser(ctx *gin.Context) *models.Users {
	var item *models.Users

	id := ctx.Param("userId")

	err := db().First(&item, "id = ?", id).Error
	if err != nil {
		return nil
	}

	return item
}

func GetUsers(ctx *gin.Context) ([]models.Users, error) {
	var items []models.Users

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	err := db().Offset((page - 1) * limit).Limit(limit).Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}
