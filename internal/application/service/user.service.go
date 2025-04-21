package service

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/internal/domain/repository"
	"gochat/pkg/logger"
	"gochat/pkg/paging"
	"gochat/pkg/storage"
	"gochat/pkg/token"
	"gochat/pkg/validation"
	"gochat/utils"
)

type UserService struct {
	validator validation.Validation
	userRepo  repository.IUserRepository
	storage   storage.IUploadService
	token     token.IMarker
}

func NewUserService(
	validator validation.Validation,
	userRepo repository.IUserRepository,
	storage storage.IUploadService,
	token token.IMarker,
) *UserService {
	return &UserService{
		validator: validator,
		userRepo:  userRepo,
		storage:   storage,
		token:     token,
	}
}

func (u *UserService) GetUserProfile(ctx context.Context, userID string) (*model.User, error) {
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) UpdateUserProfile(ctx context.Context, userID string, req *dto.UpdateProfileRequest) (*model.User, error) {
	if err := u.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	utils.MapStruct(user, req)
	if req.Avatar != nil && req.Avatar.Filename != "" {
		var avatarURL string
		avatarURL, err = u.storage.UploadFile(ctx, req.Avatar, "users")
		if err != nil {
			logger.Errorf("Failed to upload avatar: %s", err)
			return nil, err
		}

		u.storage.DeleteFile(ctx, user.AvatarURL)

		user.AvatarURL = avatarURL
	}

	err = u.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) SearchUsers(ctx context.Context, req *dto.ListUserRequest) ([]*model.User, *paging.Pagination, error) {
	users, pagination, err := u.userRepo.ListUsers(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return users, pagination, nil
}
