package service

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/internal/domain/repository"
	"gochat/pkg/paging"
	"gochat/pkg/validation"
)

type FriendService struct {
	validator  validation.Validation
	friendRepo repository.IFriendRepository
	userRepo   repository.IUserRepository
}

func NewFriendService(
	validator validation.Validation,
	friendRepo repository.IFriendRepository,
	userRepo repository.IUserRepository,
) *FriendService {
	return &FriendService{
		validator:  validator,
		friendRepo: friendRepo,
		userRepo:   userRepo,
	}
}

func (f *FriendService) AddFriend(ctx context.Context, req *dto.AddFriendRequest) error {
	// Check if users exist
	_, err := f.userRepo.FindByID(ctx, req.InviterID)
	if err != nil {
		return err
	}
	_, err = f.userRepo.FindByID(ctx, req.InviterID)
	if err != nil {
		return err
	}

	// Check if friendship already exists
	existing, err := f.friendRepo.FindByUserIDs(ctx, req.InviterID, req.AccepterID)
	if err == nil && existing != nil {
		return nil
	}

	friend := &model.Friend{
		InviterID:  req.InviterID,
		AccepterID: req.AccepterID,
	}

	return f.friendRepo.Create(ctx, friend)
}

func (f *FriendService) ListFriends(ctx context.Context, req *dto.ListFriendRequest, userID string) ([]*model.User, *paging.Pagination, error) {
	friends, pagination, err := f.friendRepo.ListByUserID(ctx, req, userID)
	if err != nil {
		return nil, nil, err
	}

	var users []*model.User
	for _, friend := range friends {
		if friend.InviterID == userID {
			users = append(users, &friend.Accepter)
		} else {
			users = append(users, &friend.Inviter)
		}
	}

	return users, pagination, nil
}

func (f *FriendService) RemoveFriend(ctx context.Context, fromID, toID string) error {
	friend, err := f.friendRepo.FindByUserIDs(ctx, fromID, toID)
	if err != nil {
		return err
	}

	return f.friendRepo.Delete(ctx, friend.ID)
}

func (f *FriendService) CheckFriendship(ctx context.Context, userID1, userID2 string) (bool, error) {
	friend, err := f.friendRepo.FindByUserIDs(ctx, userID1, userID2)
	if err != nil {
		return false, nil
	}
	return friend != nil, nil
}
