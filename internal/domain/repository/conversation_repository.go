package repository

import (
	"context"
	"gochat/internal/domain/model"
)

type IConversationRepository interface {
	Create(ctx context.Context, conversation *model.Conversation) error
	FindByID(ctx context.Context, id string) (*model.Conversation, error)
	FindByUserIDs(ctx context.Context, userIDOne, userIDTwo string) (*model.Conversation, error)
	ListByUserID(ctx context.Context, userID string) ([]*model.Conversation, error)
	UpdateLastMessage(ctx context.Context, conversationID string, messageID string) error
	Delete(ctx context.Context, id string) error
}
