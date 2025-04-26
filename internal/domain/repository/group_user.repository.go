package repository

import (
	"context"
	"gochat/internal/domain/model"
)

type IGroupUserRepository interface {
	Create(ctx context.Context, group_user *model.GroupUser) error
	FindByUserAndGroup(ctx context.Context, userID, groupID string) (*model.GroupUser, error)
	Delete(ctx context.Context, groupUser *model.GroupUser) error
	ListUsersByGroupID(ctx context.Context, groupID string) ([]*model.User, error)
	ListGroupsByUserID(ctx context.Context, userID string) ([]*model.Group, error)
}
