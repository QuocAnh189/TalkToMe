package service

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/pkg/paging"
)

// INotificationService defines the interface for managing notifications.
type INotificationService interface {
	CreateNotification(ctx context.Context, req *dto.CreateNotificationRequest, userID string) (*model.Notification, error)
	ListNotifications(ctx context.Context, req *dto.ListNotificationRequest, userID string) ([]*model.Notification, *paging.Pagination, error)
	MarkNotificationAsRead(ctx context.Context, notificationID string, userID string) error
	MarkAllNotificationsAsRead(ctx context.Context, userID string) error
	DeleteNotification(ctx context.Context, notificationID string, userID string) error
	CountUnreadNotifications(ctx context.Context, userID string) (int64, error)
}
