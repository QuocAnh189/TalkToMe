package handler

import (
	"gochat/internal/application/dto"
	"gochat/internal/domain/service"
	"gochat/pkg/logger"
	"gochat/pkg/response"
	"gochat/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	// Import necessary service interfaces, e.g., "gochat/internal/domain/service"
)

const (
	unauthorizeCreate = "unauthorized to create this notification"
)

// NotificationHandler handles notification related requests.
type NotificationHandler struct {
	service service.INotificationService
}

// NewNotificationHandler creates a new NotificationHandler.
func NewNotificationHandler(service service.INotificationService) *NotificationHandler {
	return &NotificationHandler{
		service: service,
	}
}

func (h *NotificationHandler) CreateNotification(c *gin.Context) {
	var req dto.CreateNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to bind notification data", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid notification data")
		return
	}

	userID := c.GetString("userId")
	notification, err := h.service.CreateNotification(c, &req, userID)
	if err != nil {
		logger.Error("Failed to create notification: ", err)
		switch err.Error() {
		case unauthorizeCreate:
			response.Error(c, http.StatusUnauthorized, err, unauthorizeCreate)
			return
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to create notification")
			return
		}
	}

	var res dto.CreateNotificationResponse
	utils.MapStruct(&res, notification)

	response.JSON(c, http.StatusCreated, res)
}

func (h *NotificationHandler) ListNotifications(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, nil, "Unauthorized")
		return
	}

	var req dto.ListNotificationRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get query", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	notifications, pagination, err := h.service.ListNotifications(c, &req, userID)
	if err != nil {
		logger.Error("Failed to list notifications: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to list notifications")
		return
	}

	var res dto.ListNotificationResponse
	utils.MapStruct(&res.Notifications, notifications)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

// MarkAsRead handles requests to mark a specific notification as read.
func (h *NotificationHandler) MarkAsRead(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, nil, "Unauthorized")
		return
	}

	notificationID := c.Param("id")
	if notificationID == "" {
		response.Error(c, http.StatusBadRequest, nil, "Invalid notification ID")
		return
	}

	err := h.service.MarkNotificationAsRead(c, notificationID, userID)
	if err != nil {
		logger.Error("Failed to mark notification as read: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to mark notification as read")
		return
	}

	response.JSON(c, http.StatusOK, "Notification marked as read successfully")
}

// MarkAllAsRead handles requests to mark all notifications as read.
func (h *NotificationHandler) MarkAllAsRead(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, nil, "Unauthorized")
		return
	}

	err := h.service.MarkAllNotificationsAsRead(c, userID)
	if err != nil {
		logger.Error("Failed to mark all notifications as read: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to mark all notifications as read")
		return
	}

	response.JSON(c, http.StatusOK, "All notifications marked as read successfully")
}

// DeleteNotification handles requests to delete a notification.
func (h *NotificationHandler) DeleteNotification(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, nil, "Unauthorized")
		return
	}

	notificationID := c.Param("id")
	if notificationID == "" {
		response.Error(c, http.StatusBadRequest, nil, "Invalid notification ID")
		return
	}

	err := h.service.DeleteNotification(c, notificationID, userID)
	if err != nil {
		logger.Error("Failed to delete notification: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to delete notification")
		return
	}

	response.JSON(c, http.StatusOK, "Notification deleted successfully")
}

func (h *NotificationHandler) CountUnread(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, nil, "Unauthorized")
		return
	}

	count, err := h.service.CountUnreadNotifications(c, userID)
	if err != nil {
		logger.Error("Failed to count unread notifications: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to count unread notifications")
		return
	}

	response.JSON(c, http.StatusOK, gin.H{"count": count})
}
