package repository

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/pkg/paging"
)

type IConversationRepository interface {
	Create(ctx context.Context, conversation *model.Conversation) error
	FindByID(ctx context.Context, id string) (*model.Conversation, error)
	FindOne(ctx context.Context, id string) (*model.Conversation, error)
	FindByUserIDs(ctx context.Context, userIDOne, userIDTwo string) (*model.Conversation, error)
	ListByUserID(ctx context.Context, req *dto.ListConversationRequest, userID string) ([]*model.Conversation, *paging.Pagination, error)
	UpdateLastMessage(ctx context.Context, conversationID string, messageID string) error
	Delete(ctx context.Context, id string) error
}
