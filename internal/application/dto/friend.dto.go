package dto

import (
	"gochat/pkg/paging"
)

type AddFriendRequest struct {
	InviterID  string `json:"inviter_id" validate:"required"`
	AccepterID string `json:"accepter_id" validate:"required"`
}

type RemoveFriendRequest struct {
	FromID string `json:"from_id" validate:"required"`
	ToID   string `json:"to_id" validate:"required"`
}

type ListFriendRequest struct {
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"size"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListFriendResponse struct {
	Users      []*UserResponse    `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}
