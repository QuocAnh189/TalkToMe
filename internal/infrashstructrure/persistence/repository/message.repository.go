package repository

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/internal/infrashstructrure/persistence/db"
	"gochat/pkg/paging"
)

type MessageRepository struct {
	db db.IDatabase
}

func NewMessageRepository(db db.IDatabase) *MessageRepository {
	return &MessageRepository{db: db}
}

func (mr *MessageRepository) Create(ctx context.Context, message *model.Message) error {
	handler := func() error {
		err := mr.db.Create(ctx, message)
		if err != nil {
			return err
		}

		if message.ConversationID != nil {
			var conversation model.Conversation
			err := mr.db.FindById(ctx, *message.ConversationID, &conversation)
			if err != nil {
				return err
			}

			conversation.LastMessageID = &message.ID
			if err := mr.db.Update(ctx, &conversation); err != nil {
				return err
			}
		}

		if message.GroupID != nil {
			var group model.Group
			err := mr.db.FindById(ctx, *message.GroupID, &group)
			if err != nil {
				return err
			}

			group.LastMessageID = &message.ID
			if err := mr.db.Update(ctx, &group); err != nil {
				return err
			}
		}

		return nil
	}

	err := mr.db.WithTransaction(handler)
	if err != nil {
		return err
	}

	return nil
}

func (mr *MessageRepository) FindByID(ctx context.Context, id string) (*model.Message, error) {
	var message model.Message
	if err := mr.db.FindById(ctx, id, &message); err != nil {
		return nil, err
	}
	return &message, nil
}

func (mr *MessageRepository) FindOne(ctx context.Context, id string) (*model.Message, error) {
	var message model.Message
	query := []db.Query{
		db.NewQuery("id = ?", id),
	}

	if err := mr.db.FindOne(
		ctx,
		&message,
		db.WithQuery(query...),
		db.WithPreload([]string{"Sender", "Attachments"})); err != nil {
		return nil, err
	}
	return &message, nil
}

func (mr *MessageRepository) FindMessagesByGroupID(ctx context.Context, req *dto.ListMessageRequest, groupID string) ([]*model.Message, *paging.Pagination, error) {
	var messages []*model.Message
	query := []db.Query{
		db.NewQuery("group_id = ?", groupID),
	}

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := mr.db.Count(ctx, &model.Message{}, &total, db.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if err := mr.db.Find(
		ctx,
		&messages,
		db.WithQuery(query...),
		db.WithLimit(int(pagination.Size)),
		db.WithOffset(int(pagination.Skip)),
		db.WithOrder(order),
		db.WithPreload([]string{"Sender", "Attachments"}),
	); err != nil {
		return nil, nil, err
	}

	return messages, pagination, nil
}

func (mr *MessageRepository) FindMessagesByConversationID(ctx context.Context, req *dto.ListMessageRequest, conversationID string) ([]*model.Message, *paging.Pagination, error) {
	var messages []*model.Message
	query := []db.Query{
		db.NewQuery("conversation_id = ?", conversationID),
	}

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := mr.db.Count(ctx, &model.Message{}, &total, db.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if err := mr.db.Find(
		ctx,
		&messages,
		db.WithQuery(query...),
		db.WithLimit(int(pagination.Size)),
		db.WithOffset(int(pagination.Skip)),
		db.WithOrder(order),
		db.WithPreload([]string{"Sender", "Attachments"}),
	); err != nil {
		return nil, nil, err
	}

	return messages, pagination, nil
}

func (mr *MessageRepository) Update(ctx context.Context, message *model.Message) error {
	return mr.db.Update(ctx, message)
}

func (mr *MessageRepository) Delete(ctx context.Context, message *model.Message) error {
	return mr.db.Delete(ctx, message)
}
