package routes

import (
	"github.com/gin-gonic/gin"
)

func NewNotificationRoutes(r *gin.RouterGroup) {
	notificationRoutes := r.Group("/notifications")
	{
		notificationRoutes.GET("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "notifications"})
		})
		notificationRoutes.POST("/read/:id", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "read notification"})
		})
		notificationRoutes.POST("/read-all", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "read all notification"})
		})
		notificationRoutes.DELETE("/:id", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "delete notification"})
		})
	}
}
