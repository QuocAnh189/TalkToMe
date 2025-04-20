package service

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
)

// IAuthService defines the interface for authentication related business logic.
type IAuthService interface {
	SignUp(ctx context.Context, req *dto.SignUpRequest) (accessToken, refreshToken string, user *model.User, err error)
	SignIn(ctx context.Context, req *dto.SignInRequest) (accessToken, refreshToken string, user *model.User, err error)
	SignOut(ctx context.Context, userID string, jit string) error
	RefreshToken(ctx context.Context, userID string, jit string) (accessToken string, err error)
}
