package services

import (
	"strconv"

	"blog/models"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewEntry() *models.Entries {
	return &models.Entries{}
}

func GetEntry(ctx *gin.Context) *models.Entries {
	var item *models.Entries

	id := ctx.Param("entryId")

	db().Preload(clause.Associations).First(&item, "id = ?", id)

	return item
}

func GetEntries(ctx *gin.Context) ([]models.Entries, error) {
	var items []models.Entries

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	query := db()

	if !utils.IsBackend(ctx) {
		query = query.Where("is_published = ?", 1)
	}

	err := query.Preload(clause.Associations).Offset((page - 1) * limit).Limit(limit).Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func IncEntryView(model *models.Entries) {
	err := db().Model(&models.Entries{}).Where("id = ?", model.ID).UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).Error
	if err == nil {
		model.ViewCount += 1
	}
}
