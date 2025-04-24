package repository

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/internal/infrashstructrure/persistence/db"
	"gochat/pkg/paging"
)

type ConversationRepository struct {
	db db.IDatabase
}

func NewConversationRepository(db db.IDatabase) *ConversationRepository {
	return &ConversationRepository{db: db}
}

func (cr *ConversationRepository) Create(ctx context.Context, conversation *model.Conversation) error {
	return cr.db.Create(ctx, conversation)
}

func (cr *ConversationRepository) FindByID(ctx context.Context, id string) (*model.Conversation, error) {
	var conversation model.Conversation
	if err := cr.db.FindById(ctx, id, &conversation); err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (cr *ConversationRepository) FindOne(ctx context.Context, id string) (*model.Conversation, error) {
	var conversation model.Conversation
	query := []db.Query{
		db.NewQuery("id = ?", id),
	}

	if err := cr.db.FindOne(
		ctx,
		&conversation,
		db.WithQuery(query...),
		db.WithPreload([]string{"UserOne", "UserTwo", "LastMessage"})); err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (cr *ConversationRepository) FindByUserIDs(ctx context.Context, userIDOne, userIDTwo string) (*model.Conversation, error) {
	var conversation model.Conversation
	query := []db.Query{
		db.NewQuery(
			"(user_id_one = ? AND user_id_two = ?) OR (user_id_one = ? AND user_id_two = ?)",
			userIDOne, userIDTwo, userIDTwo, userIDOne,
		),
	}

	if err := cr.db.FindOne(
		ctx,
		&conversation,
		db.WithQuery(query...),
		db.WithPreload([]string{"UserOne", "UserTwo", "LastMessage"}),
	); err != nil {
		return nil, err
	}

	return &conversation, nil
}

func (cr *ConversationRepository) ListByUserID(ctx context.Context, req *dto.ListConversationRequest, userID string) ([]*model.Conversation, *paging.Pagination, error) {
	var conversations []*model.Conversation
	query := []db.Query{
		db.NewQuery("user_id_one = ? OR user_id_two = ?", userID, userID),
	}

	var total int64
	if err := cr.db.Count(ctx, &model.Conversation{}, &total, db.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	if err := cr.db.Find(
		ctx,
		&conversations,
		db.WithQuery(query...),
		db.WithOrder("updated_at DESC"),
		db.WithPreload([]string{"UserOne", "UserTwo", "LastMessage"}),
	); err != nil {
		return nil, nil, err
	}

	return conversations, pagination, nil
}

func (cr *ConversationRepository) UpdateLastMessage(ctx context.Context, conversationID string, messageID string) error {
	conversation, err := cr.FindByID(ctx, conversationID)
	if err != nil {
		return err
	}

	conversation.LastMessageID = &messageID
	return cr.db.Update(ctx, conversation)
}

func (cr *ConversationRepository) Delete(ctx context.Context, id string) error {
	conversation, err := cr.FindByID(ctx, id)
	if err != nil {
		return err
	}
	return cr.db.Delete(ctx, conversation)
}
