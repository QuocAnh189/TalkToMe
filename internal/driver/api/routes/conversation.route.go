package routes

import (
	"gochat/internal/application/service"
	"gochat/internal/driver/api/handler"
	"gochat/internal/driver/api/middlewares"
	"gochat/internal/infrashstructrure/cache"
	"gochat/internal/infrashstructrure/persistence/db"
	"gochat/internal/infrashstructrure/persistence/repository"
	"gochat/pkg/token"
	"gochat/pkg/validation"

	"github.com/gin-gonic/gin"
)

func NewConversationRoutes(
	r *gin.RouterGroup,
	sqlDB db.IDatabase,
	validator validation.Validation,
	cache cache.IRedis,
	token token.IMarker,
) {
	conversationRepository := repository.NewConversationRepository(sqlDB)
	userRepository := repository.NewUserRepository(sqlDB)
	conversationService := service.NewConversationService(validator, conversationRepository, userRepository)
	conversationHandler := handler.NewConversationHandler(conversationService)

	authMiddleware := middlewares.NewAuthMiddleware(token, cache).TokenAuth()

	conversationRoutes := r.Group("/conversations").Use(authMiddleware)
	{
		conversationRoutes.GET("", conversationHandler.ListConversations)
		conversationRoutes.POST("", conversationHandler.CreateConversation)
		conversationRoutes.GET("/:conversationId", conversationHandler.GetConversationDetails)
		conversationRoutes.DELETE("/:conversationId", conversationHandler.DeleteConversation)
	}
}
