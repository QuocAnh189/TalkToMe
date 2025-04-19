package service

import (
	"context"
	"gochat/internal/domain/model"
)

// IDirectChatService defines the interface for managing direct conversations.
type IDirectChatService interface {
	GetOrCreateConversation(ctx context.Context, userIDOne, userIDTwo string) (*model.Conversation, error)
	GetConversationDetails(ctx context.Context, conversationID string, userID string) (*model.Conversation, error)
	ListUserConversations(ctx context.Context, userID string, limit, page int) ([]*model.Conversation, error)
	DeleteConversationForUser(ctx context.Context, conversationID string, userID string) error
}
