package service

import (
	"context"
	"gochat/internal/domain/model"
)

// INotificationService defines the interface for managing notifications.
type INotificationService interface {
	ListNotifications(ctx context.Context, userID string, limit, offset int, unreadOnly bool) ([]*model.Notification, int64, error)
	MarkNotificationAsRead(ctx context.Context, notificationID string, userID string) error
	MarkAllNotificationsAsRead(ctx context.Context, userID string) error
	DeleteNotification(ctx context.Context, notificationID string, userID string) error
	CountUnreadNotifications(ctx context.Context, userID string) (int64, error)
}
