package dto

import "time"

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

type ConversationQuery struct {
	Limit int `form:"limit,default=20"`
	Page  int `form:"page,default=1"`
}
