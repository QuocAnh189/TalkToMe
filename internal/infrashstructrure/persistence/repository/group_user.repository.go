package repository

import (
	"context"
	"gochat/internal/domain/model"
	"gochat/internal/infrashstructrure/persistence/db"
)

type GroupUserRepository struct {
	db db.IDatabase
}

func NewGroupUserRepository(db db.IDatabase) *GroupUserRepository {
	return &GroupUserRepository{db: db}
}

func (gr *GroupUserRepository) Create(ctx context.Context, group_user *model.GroupUser) error {
	return gr.db.Create(ctx, group_user)
}

func (gr *GroupUserRepository) FindByUserAndGroup(ctx context.Context, userID, groupID string) (*model.GroupUser, error) {
	var groupUser model.GroupUser
	query := []db.Query{
		db.NewQuery("user_id = ? AND group_id = ?", userID, groupID),
	}

	if err := gr.db.FindOne(ctx, &groupUser, db.WithQuery(query...)); err != nil {
		return nil, err
	}

	return &groupUser, nil
}

func (gr *GroupUserRepository) Delete(ctx context.Context, groupUser *model.GroupUser) error {
	return gr.db.Delete(ctx, groupUser)
}

func (gr *GroupUserRepository) ListUsersByGroupID(ctx context.Context, groupID string) ([]*model.User, error) {
	var users []*model.User
	query := []db.Query{
		db.NewQuery("id IN (SELECT user_id FROM group_users WHERE group_id = ? AND deleted_at IS NULL)", groupID),
	}

	if err := gr.db.Find(
		ctx,
		&users,
		db.WithQuery(query...),
	); err != nil {
		return nil, err
	}

	return users, nil
}

func (gr *GroupUserRepository) ListGroupsByUserID(ctx context.Context, userID string) ([]*model.Group, error) {
	var groups []*model.Group
	query := []db.Query{
		db.NewQuery("id IN (SELECT group_id FROM group_users WHERE user_id = ?)", userID),
	}

	if err := gr.db.Find(
		ctx,
		&groups,
		db.WithQuery(query...),
		db.WithPreload([]string{"Creator", "LastMessage"}),
	); err != nil {
		return nil, err
	}

	return groups, nil
}
