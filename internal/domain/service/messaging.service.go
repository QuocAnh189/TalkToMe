package service

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"mime/multipart"
)

// IMessagingService defines the interface for sending and retrieving messages.
type IMessagingService interface {
	SendMessage(ctx context.Context, senderID string, req *dto.SendMessageRequest, attachments []*multipart.FileHeader) (*model.Message, error)
	GetGroupMessages(ctx context.Context, groupID string, userID string, query dto.MessageQuery) ([]*model.Message, error)
	GetConversationMessages(ctx context.Context, conversationID string, userID string, query dto.MessageQuery) ([]*model.Message, error)
	GetMessageByID(ctx context.Context, messageID string, userID string) (*model.Message, error)
}
