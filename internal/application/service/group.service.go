package service

import (
	"context"
	"errors"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/internal/domain/repository"
	"gochat/pkg/logger"
	"gochat/pkg/paging"
	"gochat/pkg/storage"
	"gochat/pkg/validation"
	"gochat/utils"
)

type GroupService struct {
	validator     validation.Validation
	groupRepo     repository.IGroupRepository
	groupUserRepo repository.IGroupUserRepository
	storage       storage.IUploadService
}

func NewGroupService(
	validator validation.Validation,
	groupRepo repository.IGroupRepository,
	groupUserRepo repository.IGroupUserRepository,
	storage storage.IUploadService,
) *GroupService {
	return &GroupService{
		validator:     validator,
		groupRepo:     groupRepo,
		groupUserRepo: groupUserRepo,
		storage:       storage,
	}
}

func (g *GroupService) CreateGroup(ctx context.Context, req *dto.CreateGroupRequest, userID string) (*model.Group, error) {
	if err := g.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	var group *model.Group
	utils.MapStruct(&group, req)
	group.OwnerID = userID

	if req.Avatar != nil && req.Avatar.Filename != "" {
		avatarURL, err := g.storage.UploadFile(ctx, req.Avatar, "groups")
		if err != nil {
			return nil, err
		}
		group.AvatarURL = avatarURL
	}

	if err := g.groupRepo.Create(ctx, group, req.MemberIDs); err != nil {
		return nil, err
	}

	return group, nil
}

func (g *GroupService) GetGroupDetails(ctx context.Context, groupID string, userID string) (*model.Group, error) {
	isMember, err := g.groupRepo.IsMember(ctx, groupID, userID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, nil
	}

	return g.groupRepo.FindOne(ctx, groupID)
}

func (g *GroupService) UpdateGroupInfo(ctx context.Context, req *dto.UpdateGroupRequest, groupID, userID string) (*model.Group, error) {
	if err := g.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	group, err := g.groupRepo.FindByID(ctx, groupID)
	if err != nil {
		return nil, err
	}

	if group.OwnerID != userID {
		return nil, nil
	}

	utils.MapStruct(group, req)

	logger.Infof("Group: %v", group)

	if req.Avatar != nil && req.Avatar.Filename != "" {
		avatarURL, err := g.storage.UploadFile(ctx, req.Avatar, "groups")
		if err != nil {
			return nil, err
		}

		if group.AvatarURL != "" {
			g.storage.DeleteFile(ctx, group.AvatarURL)
		}

		group.AvatarURL = avatarURL
	}

	if err := g.groupRepo.Update(ctx, group); err != nil {
		return nil, err
	}

	return group, nil
}

func (g *GroupService) DeleteGroup(ctx context.Context, groupID string, userID string) error {
	group, err := g.groupRepo.FindByID(ctx, groupID)
	if err != nil {
		return err
	}

	if group.OwnerID != userID {
		return nil
	}

	return g.groupRepo.Delete(ctx, group)
}

func (g *GroupService) AddGroupMember(ctx context.Context, req *dto.AddMemberRequest, userID string) error {
	if err := g.validator.ValidateStruct(req); err != nil {
		return err
	}

	isMember, err := g.groupRepo.IsMember(ctx, req.GroupID, userID)
	if err != nil {
		return err
	}

	if !isMember {
		return nil
	}

	var group_user *model.GroupUser
	utils.MapStruct(&group_user, req)

	return g.groupRepo.AddMember(ctx, group_user)
}

func (g *GroupService) RemoveGroupMember(ctx context.Context, req *dto.RemoveMemberRequest, removerID string) error {
	if err := g.validator.ValidateStruct(req); err != nil {
		return err
	}

	if removerID == req.UserID {
		return errors.New("cannot remove yourself from the group")
	}

	isMember, err := g.groupRepo.IsMember(ctx, req.GroupID, removerID)
	if err != nil {
		return err
	}

	if !isMember {
		return nil
	}

	return g.groupRepo.RemoveMember(ctx, req.GroupID, req.UserID)
}

func (g *GroupService) ListUserGroups(ctx context.Context, req *dto.ListGroupRequest, userID string) ([]*model.Group, *paging.Pagination, error) {
	if err := g.validator.ValidateStruct(req); err != nil {
		return nil, nil, err
	}

	groups, pagination, err := g.groupRepo.ListByUserID(ctx, req, userID)

	return groups, pagination, err
}

func (g *GroupService) ListGroupMembers(ctx context.Context, groupID string, userID string) ([]*model.User, error) {
	isMember, err := g.groupRepo.IsMember(ctx, groupID, userID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, nil
	}

	return g.groupUserRepo.ListUsersByGroupID(ctx, groupID)
}

func (g *GroupService) IsGroupMember(ctx context.Context, req *dto.CheckIsMemberRequest) (bool, error) {
	if err := g.validator.ValidateStruct(req); err != nil {
		return false, err
	}

	return g.groupRepo.IsMember(ctx, req.GroupID, req.UserID)
}
