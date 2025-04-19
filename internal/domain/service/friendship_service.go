package service

import (
	"context"
	"gochat/internal/domain/model"
)

// IFriendshipService defines the interface for managing friendships.
type IFriendshipService interface {
	SendFriendRequest(ctx context.Context, inviterID, accepterID string) error
	AcceptFriendRequest(ctx context.Context, notificationID string, accepterID string) (*model.Friend, error)
	RejectFriendRequest(ctx context.Context, notificationID string, userID string) error
	ListFriends(ctx context.Context, userID string) ([]*model.User, error)
	RemoveFriend(ctx context.Context, userID, friendUserID string) error
	CheckFriendship(ctx context.Context, userID1, userID2 string) (bool, error)
}
