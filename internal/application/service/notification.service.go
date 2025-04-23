package service

import (
	"context"
	"errors"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/internal/domain/repository"
	"gochat/pkg/paging"
	"gochat/pkg/validation"
	"gochat/utils"
)

type NotificationService struct {
	validator        validation.Validation
	notificationRepo repository.INotificationRepository
}

func NewNotificationService(
	validator validation.Validation,
	notificationRepo repository.INotificationRepository,
) *NotificationService {
	return &NotificationService{
		validator:        validator,
		notificationRepo: notificationRepo,
	}
}

func (n *NotificationService) CreateNotification(ctx context.Context, req *dto.CreateNotificationRequest, userID string) (*model.Notification, error) {
	if req.FromID != userID {
		return nil, errors.New("unauthorized to create this notification")
	}

	var notification *model.Notification
	utils.MapStruct(&notification, req)

	notification, err := n.notificationRepo.Create(ctx, notification)
	if err != nil {
		return nil, err
	}

	return notification, nil
}

func (n *NotificationService) ListNotifications(ctx context.Context, req *dto.ListNotificationRequest, userID string) ([]*model.Notification, *paging.Pagination, error) {
	if err := n.validator.ValidateStruct(req); err != nil {
		return nil, nil, err
	}

	notifications, pagination, err := n.notificationRepo.FindByUserID(ctx, req, userID)
	if err != nil {
		return nil, nil, err
	}

	return notifications, pagination, nil
}

func (n *NotificationService) MarkNotificationAsRead(ctx context.Context, notificationID string, userID string) error {
	notification, err := n.notificationRepo.FindByID(ctx, notificationID)
	if err != nil {
		return err
	}

	if notification.ToID != userID {
		return errors.New("unauthorized to mark this notification as read")
	}

	return n.notificationRepo.MarkAsRead(ctx, notificationID)
}

func (n *NotificationService) DeleteNotification(ctx context.Context, notificationID string, userID string) error {
	notification, err := n.notificationRepo.FindByID(ctx, notificationID)
	if err != nil {
		return err
	}

	if notification.ToID != userID {
		return errors.New("unauthorized to delete this notification")
	}

	return n.notificationRepo.Delete(ctx, notificationID)
}

func (n *NotificationService) MarkAllNotificationsAsRead(ctx context.Context, userID string) error {
	return n.notificationRepo.MarkAllAsRead(ctx, userID)
}

func (n *NotificationService) CountUnreadNotifications(ctx context.Context, userID string) (int64, error) {
	return n.notificationRepo.CountUnread(ctx, userID)
}
