package repository

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/internal/infrashstructrure/persistence/db"
	"gochat/pkg/paging"

	"gorm.io/gorm"
)

type GroupRepository struct {
	db db.IDatabase
}

func NewGroupRepository(db db.IDatabase) *GroupRepository {
	return &GroupRepository{db: db}
}

func (gr *GroupRepository) Create(ctx context.Context, group *model.Group, memberIDs []string) error {
	handler := func() error {
		err := gr.db.Create(ctx, group)
		if err != nil {
			return err
		}

		memberIDs = append(memberIDs, group.OwnerID)
		group_users := make([]*model.GroupUser, len(memberIDs))
		for i, memberID := range memberIDs {
			var user model.User
			if err := gr.db.FindById(ctx, memberID, &user); err != nil {
				return err
			}

			group_users[i] = &model.GroupUser{
				GroupID:  group.ID,
				UserID:   memberID,
				Nickname: user.Name,
				IsAdmin:  group.OwnerID == memberID,
			}
		}

		if len(group_users) > 0 {
			if err := gr.db.CreateInBatches(ctx, group_users, len(memberIDs)); err != nil {
				return err
			}
		}

		return nil
	}

	err := gr.db.WithTransaction(handler)
	if err != nil {
		return err
	}

	return nil
}

func (gr *GroupRepository) FindByID(ctx context.Context, id string) (*model.Group, error) {
	var group model.Group
	if err := gr.db.FindById(ctx, id, &group); err != nil {
		return nil, err
	}
	return &group, nil
}

func (cr *GroupRepository) FindOne(ctx context.Context, id string) (*model.Group, error) {
	var group model.Group
	query := []db.Query{
		db.NewQuery("id = ?", id),
	}

	if err := cr.db.FindOne(
		ctx,
		&group,
		db.WithQuery(query...),
		db.WithPreload([]string{"Owner", "LastMessage", "Members"})); err != nil {
		return nil, err
	}
	return &group, nil
}

func (gr *GroupRepository) Update(ctx context.Context, group *model.Group) error {
	return gr.db.Update(ctx, group)
}

func (gr *GroupRepository) Delete(ctx context.Context, group *model.Group) error {
	handler := func() error {
		if err := gr.db.Delete(
			ctx,
			&model.GroupUser{},
			db.WithQuery(db.NewQuery("group_id =?", group.ID))); err != nil {
			return err
		}

		if err := gr.db.Delete(ctx, group); err != nil {
			return err
		}
		return nil
	}

	err := gr.db.WithTransaction(handler)
	if err != nil {
		return err
	}

	return nil
}

func (gr *GroupRepository) AddMember(ctx context.Context, group_user *model.GroupUser) error {
	return gr.db.Create(ctx, group_user)
}

func (gr *GroupRepository) RemoveMember(ctx context.Context, GroupID, UserID string) error {
	var groupUser model.GroupUser
	query := []db.Query{
		db.NewQuery("group_id = ? AND user_id = ?", GroupID, UserID),
	}

	if err := gr.db.FindOne(ctx, &groupUser, db.WithQuery(query...)); err != nil {
		return err
	}

	return gr.db.Delete(ctx, &groupUser)
}

func (gr *GroupRepository) IsMember(ctx context.Context, GroupID, UserID string) (bool, error) {
	var groupUser model.GroupUser
	query := []db.Query{
		db.NewQuery("group_id = ? AND user_id = ? AND deleted_at IS NULL", GroupID, UserID),
	}

	err := gr.db.FindOne(ctx, &groupUser, db.WithQuery(query...))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (gr *GroupRepository) ListByUserID(ctx context.Context, req *dto.ListGroupRequest, userID string) ([]*model.Group, *paging.Pagination, error) {
	var groups []*model.Group
	query := []db.Query{
		db.NewQuery("id IN (SELECT group_id FROM group_users WHERE user_id = ?)", userID),
	}

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := gr.db.Count(ctx, &model.Group{}, &total, db.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if err := gr.db.GetDB().
		WithContext(ctx).
		Model(&model.Group{}).
		Preload("Owner").
		Preload("LastMessage").
		Preload("Members", func(db *gorm.DB) *gorm.DB {
			return db.Joins("JOIN group_users ON users.id = group_users.user_id").
				Where("group_users.deleted_at IS NULL")
		}).
		Where("id IN (?)", gr.db.GetDB().Session(&gorm.Session{}).Select("group_id").Table("group_users").Where("user_id = ? AND deleted_at IS NULL", userID)).
		Order(order).
		Limit(int(pagination.Size)).
		Offset(int(pagination.Skip)).
		Find(&groups).Error; err != nil {
		return nil, nil, err
	}

	return groups, pagination, nil
}

func (gr *GroupRepository) UpdateLastMessage(ctx context.Context, groupID string, messageID string) error {
	group, err := gr.FindByID(ctx, groupID)
	if err != nil {
		return err
	}

	group.LastMessageID = &messageID
	return gr.db.Update(ctx, group)
}
