package controllers

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册路由
func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")

	registerEntries(api)
	registerCategories(api)
	registerFeedbacks(api)
	registerInteractions(api)
	registerUsers(api)
	registerTokens(api)
}
