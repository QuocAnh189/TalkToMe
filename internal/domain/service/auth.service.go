package service

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
)

// IAuthService defines the interface for authentication related business logic.
type IAuthService interface {
	SignUp(ctx context.Context, req *dto.SignUpRequest) (*model.User, error)
	SignIn(ctx context.Context, req *dto.SignInRequest) (accessToken, refreshToken string, err error)
	RefreshToken(ctx context.Context, refreshToken string) (accessToken string, err error)
	SignOut(ctx context.Context, userID string, accessToken string) error
}
