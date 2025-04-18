package repository

import (
	"context"
	"gochat/internal/domain/model"
)

type IFriendRepository interface {
	Create(ctx context.Context, friend *model.Friend) error
	FindByID(ctx context.Context, id string) (*model.Friend, error)
	FindByUserIDs(ctx context.Context, userID1, userID2 string) (*model.Friend, error)
	ListByUserID(ctx context.Context, userID string) ([]*model.Friend, error)
	Delete(ctx context.Context, id string) error
}
