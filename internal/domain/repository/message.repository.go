package repository

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/pkg/paging"
)

type IMessageRepository interface {
	Create(ctx context.Context, message *model.Message) error
	FindByID(ctx context.Context, id string) (*model.Message, error)
	FindOne(ctx context.Context, id string) (*model.Message, error)
	FindMessagesByGroupID(ctx context.Context, req *dto.ListMessageRequest, groupID string) ([]*model.Message, *paging.Pagination, error)
	FindMessagesByConversationID(ctx context.Context, req *dto.ListMessageRequest, conversationID string) ([]*model.Message, *paging.Pagination, error)
	Update(ctx context.Context, message *model.Message) error
	Delete(ctx context.Context, message *model.Message) error
}
