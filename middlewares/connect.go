package middlewares

import (
    "blog/common"
    "blog/models"
    "github.com/gin-gonic/gin"
)

// Connect 数据库连接中间件
func Connect(ctx *gin.Context) {
    db := common.GetDatabase().Connect()

    _ = db.AutoMigrate(&models.Entries{})
    _ = db.AutoMigrate(&models.Tags{})
    _ = db.AutoMigrate(&models.Categories{})
    _ = db.AutoMigrate(&models.Feedbacks{})
    _ = db.AutoMigrate(&models.Users{})

    ctx.Next()
}
