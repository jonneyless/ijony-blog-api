package services

import (
    "strconv"

    "blog/models"
    "github.com/gin-gonic/gin"
)

func NewCategory() *models.Categories {
    return &models.Categories{}
}

func GetCategory(ctx *gin.Context) models.Categories {
    var item models.Categories

    id := ctx.Param("id")

    db().First(&item, "id = ?", id)

    return item
}

func GetCategories(ctx *gin.Context) ([]models.Categories, error) {
    var items []models.Categories

    page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

    err := db().Offset((page - 1) * limit).Limit(limit).Find(&items).Error
    if err != nil {
        return nil, err
    }

    return items, nil
}
