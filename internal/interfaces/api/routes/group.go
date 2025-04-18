package routes

import (
	"github.com/gin-gonic/gin"
)

func NewGroupRoutes(r *gin.RouterGroup) {
	groupRoutes := r.Group("/groups")
	{
		groupRoutes.POST("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "create group"})
		})
		groupRoutes.GET("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "get my group"})
		})
		groupRoutes.GET("/:groupId", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "get group id"})
		})
		groupRoutes.PUT("/:groupId", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "update group"})
		})
		groupRoutes.DELETE("/:groupId", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "delete group"})
		})

		// Member management
		groupRoutes.POST("/:groupId/members", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "add member"})
		})
		groupRoutes.DELETE("/:groupId/members/:userId", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "remove member"})
		})
		groupRoutes.GET("/:groupId/members", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "get members"})
		})
	}
}
