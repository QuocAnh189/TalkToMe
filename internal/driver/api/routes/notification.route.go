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

func NewNotificationRoutes(
	r *gin.RouterGroup,
	sqlDB db.IDatabase,
	validator validation.Validation,
	cache cache.IRedis,
	token token.IMarker,
) {
	notificationRepository := repository.NewNotificationRepository(sqlDB)
	notificationService := service.NewNotificationService(validator, notificationRepository)
	notificationHandler := handler.NewNotificationHandler(notificationService)

	authMiddleware := middlewares.NewAuthMiddleware(token, cache).TokenAuth()

	notificationRoutes := r.Group("/notifications").Use(authMiddleware)
	{
		notificationRoutes.POST("", notificationHandler.CreateNotification)
		notificationRoutes.GET("", notificationHandler.ListNotifications)
		notificationRoutes.PUT("/read/:id", notificationHandler.MarkAsRead)
		notificationRoutes.PUT("/read-all", notificationHandler.MarkAllAsRead)
		notificationRoutes.DELETE("/:id", notificationHandler.DeleteNotification)
		notificationRoutes.GET("/unread/count", notificationHandler.CountUnread)
	}
}
