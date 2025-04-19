package service

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
)

// IUserService defines the interface for user profile management.
type IUserService interface {
	GetUserProfile(ctx context.Context, userID string) (*model.User, error)
	UpdateUserProfile(ctx context.Context, userID string, req *dto.UpdateProfileRequest) (*model.User, error)
	SearchUsers(ctx context.Context, query string, limit, page int) ([]*model.User, error)
	GetUserByID(ctx context.Context, targetUserID string) (*model.User, error)
}
