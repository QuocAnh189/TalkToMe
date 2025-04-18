package dto

import "time"

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	AvatarURL string    `json:"avatar_url,omitempty"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateProfileRequest struct {
	Name      *string `json:"name" binding:"omitempty,min=2"`
	AvatarURL *string `json:"avatar_url" binding:"omitempty,url"`
}

type UserSearchQuery struct {
	Query string `form:"q" binding:"required"`
	Limit int    `form:"limit,default=10"`
	Page  int    `form:"page,default=1"`
}
