package service

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/pkg/paging"
)

// IGroupChatService defines the interface for managing group chats.
type IGroupChatService interface {
	CreateGroup(ctx context.Context, req *dto.CreateGroupRequest, userID string) (*model.Group, error)
	GetGroupDetails(ctx context.Context, groupID string, userID string) (*model.Group, error)
	UpdateGroupInfo(ctx context.Context, req *dto.UpdateGroupRequest, groupID, userID string) (*model.Group, error)
	DeleteGroup(ctx context.Context, groupID string, userID string) error
	AddGroupMember(ctx context.Context, req *dto.AddMemberRequest, userID string) error
	RemoveGroupMember(ctx context.Context, req *dto.RemoveMemberRequest, removerID string) error
	ListUserGroups(ctx context.Context, req *dto.ListGroupRequest, userID string) ([]*model.Group, *paging.Pagination, error)
	ListGroupMembers(ctx context.Context, groupID string, userID string) ([]*model.User, error)
	IsGroupMember(ctx context.Context, req *dto.CheckIsMemberRequest) (bool, error)
}
