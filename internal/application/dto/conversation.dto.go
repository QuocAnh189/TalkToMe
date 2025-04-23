package dto

import (
	"gochat/pkg/paging"
	"time"
)

type GetOrCreateConversationRequest struct {
	TargetUserID string `json:"target_user_id" binding:"required"`
}

type ConversationResponse struct {
	ID           string           `json:"id"`
	Participants []UserResponse   `json:"participants"`
	LastMessage  *MessageResponse `json:"last_message,omitempty"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
}

type ListConversationRequest struct {
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"size"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListConversationResponse struct {
	Users      []*ConversationResponse `json:"items"`
	Pagination *paging.Pagination      `json:"metadata"`
}
