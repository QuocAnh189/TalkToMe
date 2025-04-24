package dto

import (
	"gochat/pkg/paging"
	"time"
)

type ConversationResponse struct {
	ID          string           `json:"id"`
	UserOne     UserResponse     `json:"user_one"`
	UserTwo     UserResponse     `json:"user_two"`
	LastMessage *MessageResponse `json:"last_message,omitempty"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

type ListConversationRequest struct {
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"size"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListConversationResponse struct {
	Conversations []*ConversationResponse `json:"items"`
	Pagination    *paging.Pagination      `json:"metadata"`
}

type CreateConversationRequest struct {
	UserIDOne string `json:"user_id_one"`
	UserIDTwo string `json:"user_id_two"`
}

type CreateConversationResponse struct {
	ID            string    `json:"id"`
	UserIDOne     string    `json:"user_id_one"`
	UserIDTwo     string    `json:"user_id_two"`
	LastMessageID *string   `json:"last_message_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
