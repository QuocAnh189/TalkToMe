package repository

import (
	"context"
	"gochat/internal/domain/model"
)

type IGroupRepository interface {
	Create(ctx context.Context, group *model.Group) error
	FindByID(ctx context.Context, id string) (*model.Group, error)
	Update(ctx context.Context, group *model.Group) error
	Delete(ctx context.Context, id string) error
	AddMember(ctx context.Context, groupID string, userID string) error
	RemoveMember(ctx context.Context, groupID string, userID string) error
	IsMember(ctx context.Context, groupID string, userID string) (bool, error)
	ListByUserID(ctx context.Context, userID string) ([]*model.Group, error)
	UpdateLastMessage(ctx context.Context, groupID string, messageID string) error
}
