package repository

import (
	"context"
	"gochat/internal/domain/model"
)

type IMessageRepository interface {
	Create(ctx context.Context, message *model.Message) error
	FindByID(ctx context.Context, id string) (*model.Message, error)
	FindMessagesByGroupID(ctx context.Context, groupID string, limit, offset int) ([]*model.Message, error)
	FindMessagesByConversationID(ctx context.Context, conversationID string, limit, offset int) ([]*model.Message, error)
	Update(ctx context.Context, message *model.Message) error
	Delete(ctx context.Context, id string) error
}
