package routes

import (
	"gochat/internal/application/service"
	"gochat/internal/infrashstructrure/cache"
	"gochat/internal/infrashstructrure/persistence/db"
	"gochat/internal/infrashstructrure/persistence/repository"
	"gochat/internal/interfaces/api/handler"
	"gochat/internal/interfaces/api/middlewares"
	"gochat/pkg/token"
	"gochat/pkg/validation"

	"github.com/gin-gonic/gin"
)

func NewFriendRoutes(
	r *gin.RouterGroup,
	sqlDB db.IDatabase,
	validator validation.Validation,
	cache cache.IRedis,
	token token.IMarker,
) {
	friendRepository := repository.NewFriendRepository(sqlDB)
	userRepository := repository.NewUserRepository(sqlDB)
	friendService := service.NewFriendService(validator, friendRepository, userRepository)
	friendHandler := handler.NewFriendHandler(friendService)

	authMiddleware := middlewares.NewAuthMiddleware(token, cache).TokenAuth()

	friendRoutes := r.Group("/friends").Use(authMiddleware)
	{
		friendRoutes.GET("", friendHandler.ListFriends)
		friendRoutes.POST("/add", friendHandler.AddFriend)
		friendRoutes.DELETE("/remove", friendHandler.RemoveFriend)
	}
}
