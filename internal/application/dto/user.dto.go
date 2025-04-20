package dto

import (
	"gochat/pkg/paging"
	"time"
)

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	AvatarURL string    `json:"avatar_url"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

type ListUserRequest struct {
	Search    string `json:"search,omitempty" form:"search"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"size"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}
type ListUserResponse struct {
	Users      []*UserResponse    `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}
