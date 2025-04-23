package routes

import (
	"gochat/internal/application/service"
	"gochat/internal/driver/api/handler"
	"gochat/internal/driver/api/middlewares"
	"gochat/internal/infrashstructrure/cache"
	"gochat/internal/infrashstructrure/persistence/db"
	"gochat/internal/infrashstructrure/persistence/repository"
	"gochat/pkg/mail"
	"gochat/pkg/storage"
	"gochat/pkg/token"
	"gochat/pkg/validation"

	"github.com/gin-gonic/gin"
)

func NewAuthRoutes(
	r *gin.RouterGroup,
	sqlDB db.IDatabase,
	validator validation.Validation,
	storage storage.IUploadService,
	cache cache.IRedis,
	mailer mail.IMailer,
	token token.IMarker,
) {
	userRepository := repository.NewUserRepository(sqlDB)
	authService := service.NewAuthService(validator, userRepository, storage, cache, mailer, token)
	authHandler := handler.NewAuthHandler(authService)

	authMiddleware := middlewares.NewAuthMiddleware(token, cache).TokenAuth()
	refreshMiddleware := middlewares.NewAuthMiddleware(token, cache).TokenRefresh()

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signup", authHandler.SignUp)
		authRoutes.POST("/signin", authHandler.SignIn)
		authRoutes.POST("/signout", authMiddleware, authHandler.SignOut)
		authRoutes.POST("/refresh-token", refreshMiddleware, authHandler.RefreshToken)
	}
}
