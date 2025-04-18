package dto

import (
	"mime/multipart"
	"time"
)

type SendMessageRequest struct {
	GroupID        *string `json:"group_id" binding:"omitempty"`
	ConversationID *string `json:"conversation_id" binding:"omitempty"`
	Message        string  `form:"message"`

	// For handling file uploads with multipart/form-data
	Attachments []*multipart.FileHeader `form:"attachments"`
}

type MessageAttachmentResponse struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Filename string `json:"filename,omitempty"`
	URL      string `json:"url"`
}

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

type MessageQuery struct {
	Limit  int     `form:"limit,default=50"`
	Before *string `form:"before"`
	After  *string `form:"after"`
}
