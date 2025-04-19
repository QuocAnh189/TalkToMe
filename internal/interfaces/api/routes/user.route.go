package routes

import (
	"github.com/gin-gonic/gin"
)

func NewUserRoutes(r *gin.RouterGroup) {
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/me", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "get me"})
		})
		userRoutes.PUT("/me", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "update me"})
		})
		userRoutes.GET("/:id", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "get by id"})
		})
		userRoutes.GET("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "get users"})
		})
	}
}
