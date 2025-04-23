package dto

import (
	"gochat/pkg/paging"
	"time"
)

type NotificationResponse struct {
	ID        string       `json:"id"`
	FromUser  UserResponse `json:"from_user"`
	Type      string       `json:"type"`
	IsRead    bool         `json:"is_read"`
	IsAccept  bool         `json:"is_accept"`
	Content   string       `json:"content,omitempty"`
	CreatedAt time.Time    `json:"created_at"`
}

type ListNotificationRequest struct {
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"size"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListNotificationResponse struct {
	Notifications []*NotificationResponse `json:"items"`
	Pagination    *paging.Pagination      `json:"metadata"`
}

type CreateNotificationRequest struct {
	FromID  string `json:"from_id" validate:"required"`
	ToID    string `json:"to_id" validate:"required"`
	Type    string `json:"type" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type CreateNotificationResponse struct {
	ID        string    `json:"id"`
	FromID    string    `json:"from_id"`
	ToID      string    `json:"to_id"`
	Type      string    `json:"type"`
	IsRead    bool      `json:"is_read"`
	IsAccept  bool      `json:"is_accept"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
