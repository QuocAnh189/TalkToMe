package dto

import "time"

type CreateGroupRequest struct {
	Name        string `json:"name" binding:"required,min=3"`
	Description string `json:"description"`
}

type UpdateGroupRequest struct {
	Name        *string `json:"name" binding:"omitempty,min=3"`
	Description *string `json:"description"`
}

type GroupResponse struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Owner       UserResponse     `json:"owner"`
	LastMessage *MessageResponse `json:"last_message,omitempty"`
	MemberCount int              `json:"member_count"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

type GroupMemberRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type GroupMemberResponse struct {
	User     UserResponse `json:"user"`
	JoinedAt time.Time    `json:"joined_at"`
}

type GroupQuery struct {
	Limit int `form:"limit,default=20"`
	Page  int `form:"page,default=1"`
}
