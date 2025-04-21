package service

import (
	"context"
	"errors"
	"fmt"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/internal/domain/repository"
	"gochat/internal/infrashstructrure/cache"
	"gochat/pkg/logger"
	"gochat/pkg/mail"
	"gochat/pkg/storage"
	"gochat/pkg/token"
	"gochat/pkg/validation"
	"gochat/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	validator validation.Validation
	userRepo  repository.IUserRepository
	storage   storage.IUploadService
	cache     cache.IRedis
	mailer    mail.IMailer
	token     token.IMarker
}

func NewAuthService(
	validator validation.Validation,
	userRepo repository.IUserRepository,
	storage storage.IUploadService,
	cache cache.IRedis,
	mailer mail.IMailer,
	token token.IMarker,
) *AuthService {
	return &AuthService{
		validator: validator,
		userRepo:  userRepo,
		storage:   storage,
		cache:     cache,
		mailer:    mailer,
		token:     token,
	}
}

func (a *AuthService) SignIn(ctx context.Context, req *dto.SignInRequest) (string, string, *model.User, error) {
	if err := a.validator.ValidateStruct(req); err != nil {
		return "", "", nil, err
	}

	user, err := a.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		logger.Errorf("Login.GetUserByEmail fail, email: %s, error: %s", req.Email, err)
		return "", "", nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", "", nil, errors.New("wrong password")
	}

	tokenData := token.AuthPayload{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		Jit:   uuid.New().String(),
		Type:  token.AccessTokenType,
	}

	accessToken := a.token.GenerateAccessToken(&tokenData)
	refreshToken := a.token.GenerateRefreshToken(&tokenData)

	return accessToken, refreshToken, user, nil
}

func (a *AuthService) SignUp(ctx context.Context, req *dto.SignUpRequest) (string, string, *model.User, error) {
	if err := a.validator.ValidateStruct(req); err != nil {
		return "", "", nil, err
	}

	var avatarUrlUpload = ""
	logger.Info("req.FileName: ", req.Avatar.Filename)
	if req.Avatar != nil && req.Avatar.Filename != "" {
		avatarURL, err := a.storage.UploadFile(ctx, req.Avatar, "users")
		if err != nil {
			logger.Errorf("Failed to upload avatar: %s", err)
			return "", "", nil, err
		}
		avatarUrlUpload = avatarURL
	}

	var user *model.User
	utils.MapStruct(&user, &req)
	user.AvatarURL = avatarUrlUpload

	err := a.userRepo.Create(ctx, user)
	if err != nil {
		logger.Errorf("Register.Create fail, email: %s, error: %s", req.Email, err)
		return "", "", nil, err
	}

	if err := a.mailer.Send(user.Email, "Hello!", "<h1>Congratulations</h1><p>Your account has been successfully created</p>", true); err != nil {
		logger.Errorf("Send mail failure: %v", err)
	}

	tokenData := token.AuthPayload{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}

	accessToken := a.token.GenerateAccessToken(&tokenData)
	refreshToken := a.token.GenerateRefreshToken(&tokenData)

	return accessToken, refreshToken, user, nil
}

func (a *AuthService) SignOut(ctx context.Context, userID string, jit string) error {
	value := `{"status": "blacklisted"}`

	// err := u.cache.Set(fmt.Sprintf("blacklist:%s", strings.ReplaceAll(token, " ", "_")), value)
	err := a.cache.Set(fmt.Sprintf("blacklist:%s_%s", userID, jit), value)
	if err != nil {
		logger.Error("Failed to blacklist token: ", err)
		return err
	}

	logger.Info("User signed out successfully")
	return nil
}

func (a *AuthService) RefreshToken(ctx context.Context, userId string, jit string) (string, error) {
	user, err := a.userRepo.FindByID(ctx, userId)
	if err != nil {
		logger.Errorf("RefreshToken.GetUserByID fail, id: %s, error: %s", userId, err)
		return "", err
	}

	tokenData := token.AuthPayload{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		Jit:   jit,
		Type:  token.AccessTokenType,
	}

	accessToken := a.token.GenerateAccessToken(&tokenData)
	return accessToken, nil
}
