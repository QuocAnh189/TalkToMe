package repository

import (
	"context"
	"gochat/internal/domain/model"
)

type INotificationRepository interface {
	Create(ctx context.Context, notification *model.Notification) error
	FindByID(ctx context.Context, id string) (*model.Notification, error)
	FindByUserID(ctx context.Context, userID string, limit, offset int) ([]*model.Notification, error)
	MarkAsRead(ctx context.Context, id string) error
	MarkAllAsRead(ctx context.Context, userID string) error
	CountUnread(ctx context.Context, userID string) (int64, error)
	Delete(ctx context.Context, id string) error
}
