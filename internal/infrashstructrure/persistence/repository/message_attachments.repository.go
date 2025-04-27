package repository

import (
	"context"
	"gochat/internal/domain/model"
	"gochat/internal/infrashstructrure/persistence/db"
)

type MessageAttachmentRepository struct {
	db db.IDatabase
}

func NewMessageAttachmentRepository(db db.IDatabase) *MessageAttachmentRepository {
	return &MessageAttachmentRepository{db: db}
}

func (mar *MessageAttachmentRepository) Create(ctx context.Context, attachment *model.MessageAttachment) error {
	return mar.db.Create(ctx, attachment)
}

func (mar *MessageAttachmentRepository) CreateBatch(ctx context.Context, attachments []*model.MessageAttachment) error {
	if len(attachments) == 0 {
		return nil
	}
	return mar.db.CreateInBatches(ctx, attachments, len(attachments))
}

func (mar *MessageAttachmentRepository) FindByMessageID(ctx context.Context, messageID string) ([]*model.MessageAttachment, error) {
	var attachments []*model.MessageAttachment
	query := []db.Query{
		db.NewQuery("message_id = ?", messageID),
	}

	if err := mar.db.Find(
		ctx,
		&attachments,
		db.WithQuery(query...),
	); err != nil {
		return nil, err
	}

	return attachments, nil
}

func (mar *MessageAttachmentRepository) FindByID(ctx context.Context, id string) (*model.MessageAttachment, error) {
	var attachment model.MessageAttachment
	if err := mar.db.FindById(ctx, id, &attachment); err != nil {
		return nil, err
	}
	return &attachment, nil
}

func (mar *MessageAttachmentRepository) Delete(ctx context.Context, id string) error {
	attachment, err := mar.FindByID(ctx, id)
	if err != nil {
		return err
	}
	return mar.db.Delete(ctx, attachment)
}

func (mar *MessageAttachmentRepository) DeleteByMessageID(ctx context.Context, messageID string) error {
	query := []db.Query{
		db.NewQuery("message_id = ?", messageID),
	}

	return mar.db.Delete(ctx, &model.MessageAttachment{}, db.WithQuery(query...))
}
