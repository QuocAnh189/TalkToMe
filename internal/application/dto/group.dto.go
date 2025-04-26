package dto

import (
	"gochat/pkg/paging"
	"mime/multipart"
	"time"
)

type CreateGroupRequest struct {
	Name        string                `form:"name" validate:"required"`
	Avatar      *multipart.FileHeader `form:"avatar"`
	OwnerID     string                `form:"owner_id"`
	Description string                `form:"description"`
	MemberIDs   []string              `form:"member_ids"`
}

type CreateGroupResponse struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	AvatarURL     string    `json:"avatar_url"`
	Description   string    `json:"description,omitempty"`
	OwnerID       string    `json:"owner_id"`
	LastMessageID *string   `json:"last_message_id,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UpdateGroupRequest struct {
	Name        *string               `form:"name" validate:"required"`
	Avatar      *multipart.FileHeader `form:"avatar"`
	Description *string               `form:"description"`
}

type GroupResponse struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Owner       *UserResponse    `json:"owner"`
	LastMessage *MessageResponse `json:"last_message,omitempty"`
	Members     []*UserResponse  `json:"members"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

type ListGroupRequest struct {
	Name      string `json:"name" form:"name"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"size"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListGroupResponse struct {
	Groups     []*GroupResponse   `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}

type AddMemberRequest struct {
	GroupID string `json:"group_id" validate:"required"`
	UserID  string `json:"user_id" validate:"required"`
}

type RemoveMemberRequest struct {
	GroupID string `json:"group_id" validate:"required"`
	UserID  string `json:"user_id" validate:"required"`
}

type CheckIsMemberRequest struct {
	GroupID string `json:"group_id" validate:"required"`
	UserID  string `json:"user_id" validate:"required"`
}
