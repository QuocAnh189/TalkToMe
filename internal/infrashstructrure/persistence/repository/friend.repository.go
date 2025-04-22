package repository

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/internal/infrashstructrure/persistence/db"
	"gochat/pkg/paging"
)

type FriendRepository struct {
	db db.IDatabase
}

func NewFriendRepository(db db.IDatabase) *FriendRepository {
	return &FriendRepository{db: db}
}

func (fr *FriendRepository) Create(ctx context.Context, friend *model.Friend) error {
	return fr.db.Create(ctx, friend)
}

func (fr *FriendRepository) FindByID(ctx context.Context, id string) (*model.Friend, error) {
	var friend model.Friend
	if err := fr.db.FindById(ctx, id, &friend); err != nil {
		return nil, err
	}
	return &friend, nil
}

func (fr *FriendRepository) FindByUserIDs(ctx context.Context, userID1, userID2 string) (*model.Friend, error) {
	var friend model.Friend
	query := []db.Query{
		db.NewQuery("(inviter_id = ? AND accepter_id = ?) OR (inviter_id = ? AND accepter_id = ?)",
			userID1, userID2, userID2, userID1),
	}

	if err := fr.db.FindOne(ctx, &friend, db.WithQuery(query...)); err != nil {
		return nil, err
	}
	return &friend, nil
}

func (fr *FriendRepository) ListByUserID(ctx context.Context, req *dto.ListFriendRequest, userID string) ([]*model.Friend, *paging.Pagination, error) {
	query := []db.Query{
		db.NewQuery("inviter_id = ? OR accepter_id = ?", userID, userID),
	}

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := fr.db.Count(ctx, &model.Friend{}, &total, db.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	var friends []*model.Friend
	if err := fr.db.Find(
		ctx,
		&friends,
		db.WithQuery(query...),
		db.WithLimit(int(pagination.Size)),
		db.WithOffset(int(pagination.Skip)),
		db.WithOrder(order),
		db.WithPreload([]string{"Inviter", "Accepter"}),
	); err != nil {
		return nil, nil, err
	}

	return friends, pagination, nil
}

func (fr *FriendRepository) Delete(ctx context.Context, id string) error {
	friend, err := fr.FindByID(ctx, id)
	if err != nil {
		return err
	}
	return fr.db.Delete(ctx, friend)
}
