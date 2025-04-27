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

func NewMessageRoutes(
	r *gin.RouterGroup,
	sqlDB db.IDatabase,
	validator validation.Validation,
	storage storage.IUploadService,
	cache cache.IRedis,
	token token.IMarker,
) {
	messageRepository := repository.NewMessageRepository(sqlDB)
	messageAttachmentsRepository := repository.NewMessageAttachmentRepository(sqlDB)
	groupRepository := repository.NewGroupRepository(sqlDB)
	MessageService := service.NewMessageService(
		validator,
		messageRepository,
		messageAttachmentsRepository,
		groupRepository,
		storage,
	)
	messageHandler := handler.NewMessageHandler(MessageService)

	authMiddleware := middlewares.NewAuthMiddleware(token, cache).TokenAuth()

	messageRoutes := r.Group("/messages").Use(authMiddleware)
	{
		messageRoutes.POST("", messageHandler.SendMessage)
		messageRoutes.GET("/:messageId", messageHandler.GetMessage)
		messageRoutes.GET("/group/:groupId", messageHandler.GetGroupMessages)
		messageRoutes.GET("/conversation/:conversationId", messageHandler.GetConversationMessages)
		messageRoutes.PUT("/:messageId", messageHandler.UpdateMessage)
		messageRoutes.DELETE("/:messageId", messageHandler.DeleteMessage)
	}
}
