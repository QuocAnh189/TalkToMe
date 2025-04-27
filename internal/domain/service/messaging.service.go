package service

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/pkg/paging"
	"mime/multipart"
)

// IMessagingService defines the interface for sending and retrieving messages.
type IMessagingService interface {
	SendMessage(ctx context.Context, req *dto.SendMessageRequest, attachments []*multipart.FileHeader, senderID string) (*model.Message, error)
	GetGroupMessages(ctx context.Context, req *dto.ListMessageRequest, groupID string, userID string) ([]*model.Message, *paging.Pagination, error)
	GetConversationMessages(ctx context.Context, req *dto.ListMessageRequest, conversationID string, userID string) ([]*model.Message, *paging.Pagination, error)
	GetMessageByID(ctx context.Context, messageID string, userID string) (*model.Message, error)
	UpdateMessageByID(ctx context.Context, messageID string, userID string, req *dto.UpdateMessageRequest) (*model.Message, error)
	DeleteMessageByID(ctx context.Context, messageID string, userID string) error
}
