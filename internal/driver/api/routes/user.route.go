package routes

import (
	"gochat/internal/application/service"
	"gochat/internal/driver/api/handler"
	"gochat/internal/driver/api/middlewares"
	"gochat/internal/infrashstructrure/cache"
	"gochat/internal/infrashstructrure/persistence/db"
	"gochat/internal/infrashstructrure/persistence/repository"
	"gochat/pkg/storage"
	"gochat/pkg/token"
	"gochat/pkg/validation"

	"github.com/gin-gonic/gin"
)

func NewUserRoutes(
	r *gin.RouterGroup,
	sqlDB db.IDatabase,
	validator validation.Validation,
	storage storage.IUploadService,
	cache cache.IRedis,
	token token.IMarker,
) {
	userRepository := repository.NewUserRepository(sqlDB)
	userService := service.NewUserService(validator, userRepository, storage, token)
	userHandler := handler.NewUserHandler(userService)

	authMiddleware := middlewares.NewAuthMiddleware(token, cache).TokenAuth()

	userRoutes := r.Group("/users").Use(authMiddleware)
	{
		userRoutes.GET("/me", userHandler.GetProfile)
		userRoutes.PUT("/me", userHandler.UpdateProfile)
		userRoutes.GET("/:id", userHandler.GetUserByID)
		userRoutes.GET("", userHandler.SearchUsers)
	}
}
