package repository

import (
	"context"
	"gochat/internal/domain/model"
)

type IMessageAttachmentRepository interface {
	Create(ctx context.Context, attachment *model.MessageAttachment) error
	CreateBatch(ctx context.Context, attachments []*model.MessageAttachment) error
	FindByMessageID(ctx context.Context, messageID string) ([]*model.MessageAttachment, error)
	FindByID(ctx context.Context, id string) (*model.MessageAttachment, error)
	Delete(ctx context.Context, id string) error
}
