package dto

import (
	"gochat/pkg/paging"
	"mime/multipart"
	"time"
)

type MessageResponse struct {
	ID             string                      `json:"id"`
	Message        string                      `json:"message,omitempty"`
	Sender         UserResponse                `json:"sender"`
	GroupID        *string                     `json:"group_id,omitempty"`
	ConversationID *string                     `json:"conversation_id,omitempty"`
	Attachments    []MessageAttachmentResponse `json:"attachments,omitempty"`
	CreatedAt      time.Time                   `json:"created_at"`
	UpdatedAt      time.Time                   `json:"updated_at"`
}

type SendMessageRequest struct {
	GroupID        *string                 `json:"group_id,omitempty" form:"group_id" binding:"omitempty"`
	ConversationID *string                 `json:"conversation_id,omitempty" form:"conversation_id" binding:"omitempty"`
	Message        string                  `form:"message"`
	SenderID       string                  `json:"sender_id,omitempty" form:"sender_id"`
	Attachments    []*multipart.FileHeader `form:"attachments"`
}

type SendMessageResponse struct {
	ID             string    `json:"id"`
	GroupID        *string   `json:"group_id,omitempty"`
	ConversationID *string   `json:"conversation_id,omitempty"`
	SenderID       string    `json:"sender_id"`
	Message        string    `json:"message,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UpdateMessageRequest struct {
	Message string `json:"message"`
}

type UpdateMessageResponse struct {
	ID             string    `json:"id"`
	GroupID        *string   `json:"group_id,omitempty"`
	ConversationID *string   `json:"conversation_id,omitempty"`
	SenderID       string    `json:"sender_id"`
	Message        string    `json:"message,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type MessageAttachmentResponse struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Filename string `json:"filename,omitempty"`
	URL      string `json:"url"`
}

type ListMessageRequest struct {
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"size"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListMessageResponse struct {
	Messages   []*MessageResponse `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}
