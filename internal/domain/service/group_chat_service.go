package service

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
)

// IGroupChatService defines the interface for managing group chats.
type IGroupChatService interface {
	CreateGroup(ctx context.Context, ownerID string, req *dto.CreateGroupRequest) (*model.Group, error)
	GetGroupDetails(ctx context.Context, groupID string, userID string) (*model.Group, error)
	UpdateGroupInfo(ctx context.Context, groupID string, userID string, req *dto.UpdateGroupRequest) (*model.Group, error)
	DeleteGroup(ctx context.Context, groupID string, userID string) error
	AddGroupMember(ctx context.Context, groupID string, targetUserID string, adderID string) error
	RemoveGroupMember(ctx context.Context, groupID string, targetUserID string, removerID string) error
	ListUserGroups(ctx context.Context, userID string, limit, page int) ([]*model.Group, error)
	ListGroupMembers(ctx context.Context, groupID string, userID string, limit, page int) ([]*model.User, error)
	IsGroupMember(ctx context.Context, groupID string, userID string) (bool, error)
}
