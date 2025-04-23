package repository

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/pkg/paging"
)

type INotificationRepository interface {
	Create(ctx context.Context, notification *model.Notification) (*model.Notification, error)
	FindByID(ctx context.Context, id string) (*model.Notification, error)
	FindByUserID(ctx context.Context, req *dto.ListNotificationRequest, userID string) ([]*model.Notification, *paging.Pagination, error)
	MarkAsRead(ctx context.Context, id string) error
	MarkAllAsRead(ctx context.Context, userID string) error
	CountUnread(ctx context.Context, userID string) (int64, error)
	Delete(ctx context.Context, id string) error
}
