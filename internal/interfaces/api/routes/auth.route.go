package routes

import (
	"github.com/gin-gonic/gin"
)

func NewAuthRoutes(r *gin.RouterGroup) {
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signup", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "signup"})
		})
		authRoutes.POST("/signin", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "signin"})
		})
		authRoutes.POST("/signout", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "signout"})
		})
		authRoutes.POST("/refresh-token", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "refresh-token"})
		})
	}
}
