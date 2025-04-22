package repository

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/pkg/paging"
)

type IFriendRepository interface {
	Create(ctx context.Context, friend *model.Friend) error
	FindByID(ctx context.Context, id string) (*model.Friend, error)
	FindByUserIDs(ctx context.Context, userID1, userID2 string) (*model.Friend, error)
	ListByUserID(ctx context.Context, req *dto.ListFriendRequest, userID string) ([]*model.Friend, *paging.Pagination, error)
	Delete(ctx context.Context, id string) error
}
