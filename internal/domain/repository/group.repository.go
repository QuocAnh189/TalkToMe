package repository

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/pkg/paging"
)

type IGroupRepository interface {
	Create(ctx context.Context, group *model.Group, memberIds []string) error
	FindByID(ctx context.Context, id string) (*model.Group, error)
	FindOne(ctx context.Context, id string) (*model.Group, error)
	Update(ctx context.Context, group *model.Group) error
	Delete(ctx context.Context, group *model.Group) error
	AddMember(ctx context.Context, group_user *model.GroupUser) error
	RemoveMember(ctx context.Context, GroupID, UserID string) error
	IsMember(ctx context.Context, GroupID, UserID string) (bool, error)
	ListByUserID(ctx context.Context, req *dto.ListGroupRequest, userID string) ([]*model.Group, *paging.Pagination, error)
	UpdateLastMessage(ctx context.Context, groupID string, messageID string) error
}
