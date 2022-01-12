package services

import (
	"strconv"

	"blog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func NewFeedback() *models.Feedbacks {
	return &models.Feedbacks{}
}

func GetFeedback(ctx *gin.Context) *models.Feedbacks {
	var item *models.Feedbacks

	id := ctx.Param("feedbackId")

	err := db().First(&item, "id = ?", id).Error
	if err != nil {
		return nil
	}

	return item
}

func GetFeedbacks(ctx *gin.Context) ([]models.Feedbacks, error) {
	var items []models.Feedbacks

	entryId, _ := strconv.Atoi(ctx.Param("entryId"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	err := db().Preload(clause.Associations).Where("entry_id = ?", entryId).Offset((page - 1) * limit).Limit(limit).Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}
