package service

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/pkg/paging"
)

// IFriendshipService defines the interface for managing friendships.
type IFriendshipService interface {
	AddFriend(ctx context.Context, req *dto.AddFriendRequest) error
	ListFriends(ctx context.Context, req *dto.ListFriendRequest, userID string) ([]*model.User, *paging.Pagination, error)
	RemoveFriend(ctx context.Context, FromID, ToID string) error
	CheckFriendship(ctx context.Context, userID1, userID2 string) (bool, error)
}
