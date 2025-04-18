package routes

import (
	"github.com/gin-gonic/gin"
)

func NewConversationRoutes(r *gin.RouterGroup) {
	conversationRoutes := r.Group("/conversations")
	{
		conversationRoutes.GET("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "get conversations"})
		})

		conversationRoutes.POST("/user/:userId", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "create conversation"})
		})

		conversationRoutes.GET("/:conversationId", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "get conversation id"})
		})

		conversationRoutes.DELETE("/:conversationId", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "delete conversation"})
		})
	}
}
