package routes

import (
	"github.com/gin-gonic/gin"
)

func NewFriendRoutes(r *gin.RouterGroup) {
	friendRoutes := r.Group("/friends")
	{
		friendRoutes.POST("/invite", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "invite"})
		})
		friendRoutes.POST("/accept/:notificationId", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "accept"})
		})
		friendRoutes.GET("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "get my friends"})
		})
		friendRoutes.DELETE("/:userId", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "unfriends"})
		})
	}
}
