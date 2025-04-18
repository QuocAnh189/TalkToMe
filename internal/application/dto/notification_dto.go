package dto

import "time"

type NotificationResponse struct {
	ID        string       `json:"id"`
	FromUser  UserResponse `json:"from_user"`
	Type      string       `json:"type"`
	IsRead    bool         `json:"is_read"`
	IsAccept  bool         `json:"is_accept"`
	Content   string       `json:"content,omitempty"`
	CreatedAt time.Time    `json:"created_at"`
}

type NotificationQuery struct {
	Limit      int  `form:"limit,default=20"`
	Page       int  `form:"page,default=1"`
	UnreadOnly bool `form:"unread_only,default=false"`
}
