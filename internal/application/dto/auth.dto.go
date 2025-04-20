package dto

import "mime/multipart"

type SignUpRequest struct {
	Name     string                `form:"name" validate:"required"`
	Email    string                `form:"email" validate:"required"`
	Avatar   *multipart.FileHeader `form:"avatar"`
	Password string                `form:"password" validate:"required"`
	Role     string                `form:"role" validate:"required"`
}

type SignUpResponse struct {
	AccessToken  string        `json:"accessToken" validate:"required"`
	RefreshToken string        `json:"refreshToken" validate:"required"`
	User         *UserResponse `json:"user" validate:"required"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	AccessToken  string        `json:"accessToken" validate:"required"`
	RefreshToken string        `json:"refreshToken" validate:"required"`
	User         *UserResponse `json:"user" validate:"required"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
}
