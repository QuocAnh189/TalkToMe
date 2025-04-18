package dto

import "time"

type FriendRequest struct {
	TargetUserID string `json:"target_user_id" binding:"required"`
}

type FriendResponse struct {
	FriendshipID string       `json:"friendship_id"`
	User         UserResponse `json:"user"`
	CreatedAt    time.Time    `json:"created_at"`
}
