package repository

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/internal/infrashstructrure/persistence/db"
	"gochat/pkg/paging"
)

type UserRepository struct {
	db db.IDatabase
}

func NewUserRepository(db db.IDatabase) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) ListUsers(ctx context.Context, req *dto.ListUserRequest) ([]*model.User, *paging.Pagination, error) {
	query := make([]db.Query, 0)

	if req.Search != "" {
		query = append(query, db.NewQuery("name ILIKE ?", "%"+req.Search+"%"))
	}

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := ur.db.Count(ctx, &model.User{}, &total, db.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	var users []*model.User
	if err := ur.db.Find(
		ctx,
		&users,
		db.WithQuery(query...),
		db.WithLimit(int(pagination.Size)),
		db.WithOffset(int(pagination.Skip)),
		db.WithOrder(order),
	); err != nil {
		return nil, nil, err
	}

	return users, pagination, nil
}

func (ur *UserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	if err := ur.db.FindById(ctx, id, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	query := db.NewQuery("email = ?", email)
	if err := ur.db.FindOne(ctx, &user, db.WithQuery(query)); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) Create(ctx context.Context, user *model.User) error {
	return ur.db.Create(ctx, user)
}

func (ur *UserRepository) Update(ctx context.Context, user *model.User) error {
	return ur.db.Update(ctx, user)
}

func (ur *UserRepository) Delete(ctx context.Context, user *model.User) error {
	return ur.db.Delete(ctx, user)
}
