package routes

import (
	"github.com/gin-gonic/gin"
)

func NewMessageRoutes(r *gin.RouterGroup) {
	messageRoutes := r.Group("/messages")
	{
		messageRoutes.POST("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "create message"})
		})
		messageRoutes.GET("/group/:groupId", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "get messages by group Id"})
		})
		messageRoutes.GET("/conversation/:conversationId", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "get messages by conversation Id"})
		})
	}
}
