package services

import (
    "strconv"

    "blog/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm/clause"
)

func NewEntry() *models.Entries {
    return &models.Entries{}
}

func GetEntry(ctx *gin.Context) models.Entries {
    var item models.Entries

    id := ctx.Param("id")

    db().First(&item, "id = ?", id)

    return item
}

func GetEntries(ctx *gin.Context) ([]models.Entries, error) {
    var items []models.Entries

    page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

    err := db().Preload(clause.Associations).Offset((page - 1) * limit).Limit(limit).Find(&items).Error
    if err != nil {
        return nil, err
    }

    return items, nil
}
