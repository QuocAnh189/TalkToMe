package service

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/pkg/paging"
)

// IDirectChatService defines the interface for managing direct conversations.
type IConversationService interface {
	CreateConversation(ctx context.Context, req *dto.CreateConversationRequest) (*model.Conversation, error)
	GetConversationDetails(ctx context.Context, conversationID string, userID string) (*model.Conversation, error)
	ListUserConversations(ctx context.Context, req *dto.ListConversationRequest, userID string) ([]*model.Conversation, *paging.Pagination, error)
	DeleteConversationForUser(ctx context.Context, conversationID string, userID string) error
}
