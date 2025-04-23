package repository

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/internal/infrashstructrure/persistence/db"
	"gochat/pkg/paging"
)

type NotificationRepository struct {
	db db.IDatabase
}

func NewNotificationRepository(db db.IDatabase) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (nr *NotificationRepository) Create(ctx context.Context, notification *model.Notification) (*model.Notification, error) {
	err := nr.db.Create(ctx, notification)
	if err != nil {
		return nil, err
	}
	return notification, nil
}

func (nr *NotificationRepository) FindByID(ctx context.Context, id string) (*model.Notification, error) {
	var notification model.Notification
	if err := nr.db.FindById(ctx, id, &notification); err != nil {
		return nil, err
	}
	return &notification, nil
}

func (nr *NotificationRepository) FindByUserID(ctx context.Context, req *dto.ListNotificationRequest, userID string) ([]*model.Notification, *paging.Pagination, error) {
	var notifications []*model.Notification
	query := []db.Query{
		db.NewQuery("to_id = ?", userID),
	}

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := nr.db.Count(ctx, &model.Notification{}, &total, db.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if err := nr.db.Find(
		ctx,
		&notifications,
		db.WithQuery(query...),
		db.WithLimit(int(pagination.Size)),
		db.WithOffset(int(pagination.Skip)),
		db.WithOrder(order),
		db.WithPreload([]string{"FromUser"}),
	); err != nil {
		return nil, nil, err
	}

	return notifications, pagination, nil
}

func (nr *NotificationRepository) MarkAsRead(ctx context.Context, id string) error {
	notification, err := nr.FindByID(ctx, id)
	if err != nil {
		return err
	}

	notification.IsRead = true
	return nr.db.Update(ctx, notification)
}

func (nr *NotificationRepository) MarkAllAsRead(ctx context.Context, userID string) error {
	query := []db.Query{
		db.NewQuery("to_id = ? AND is_read = ?", userID, false),
	}

	var notifications []*model.Notification
	if err := nr.db.Find(ctx, &notifications, db.WithQuery(query...)); err != nil {
		return err
	}

	for _, notification := range notifications {
		notification.IsRead = true
		if err := nr.db.Update(ctx, notification); err != nil {
			return err
		}
	}

	return nil
}

func (nr *NotificationRepository) CountUnread(ctx context.Context, userID string) (int64, error) {
	var total int64
	query := []db.Query{
		db.NewQuery("to_id = ? AND is_read = ?", userID, false),
	}

	if err := nr.db.Count(ctx, &model.Notification{}, &total, db.WithQuery(query...)); err != nil {
		return 0, err
	}

	return total, nil
}

func (nr *NotificationRepository) Delete(ctx context.Context, id string) error {
	notification, err := nr.FindByID(ctx, id)
	if err != nil {
		return err
	}
	return nr.db.Delete(ctx, notification)
}
