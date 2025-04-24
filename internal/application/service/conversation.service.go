package service

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/internal/domain/repository"
	"gochat/pkg/paging"
	"gochat/pkg/validation"
	"gochat/utils"
)

type ConversationService struct {
	validator        validation.Validation
	conversationRepo repository.IConversationRepository
	userRepo         repository.IUserRepository
}

func NewConversationService(
	validator validation.Validation,
	conversationRepo repository.IConversationRepository,
	userRepo repository.IUserRepository,
) *ConversationService {
	return &ConversationService{
		validator:        validator,
		conversationRepo: conversationRepo,
		userRepo:         userRepo,
	}
}

func (c *ConversationService) CreateConversation(ctx context.Context, req *dto.CreateConversationRequest) (*model.Conversation, error) {
	_, err := c.userRepo.FindByID(ctx, req.UserIDOne)
	if err != nil {
		return nil, err
	}
	_, err = c.userRepo.FindByID(ctx, req.UserIDTwo)
	if err != nil {
		return nil, err
	}

	var conversation *model.Conversation
	utils.MapStruct(&conversation, req)

	err = c.conversationRepo.Create(ctx, conversation)
	if err != nil {
		return nil, err
	}

	return conversation, nil
}

func (c *ConversationService) GetConversationDetails(ctx context.Context, conversationID string, userID string) (*model.Conversation, error) {
	conversation, err := c.conversationRepo.FindOne(ctx, conversationID)
	if err != nil {
		return nil, err
	}

	if conversation.UserIDOne != userID && conversation.UserIDTwo != userID {
		return nil, nil
	}

	return conversation, nil
}

func (c *ConversationService) ListUserConversations(ctx context.Context, req *dto.ListConversationRequest, userID string) ([]*model.Conversation, *paging.Pagination, error) {
	conversations, pagination, err := c.conversationRepo.ListByUserID(ctx, req, userID)
	if err != nil {
		return nil, nil, err
	}

	return conversations, pagination, nil
}

func (c *ConversationService) DeleteConversationForUser(ctx context.Context, conversationID string, userID string) error {
	conversation, err := c.conversationRepo.FindByID(ctx, conversationID)
	if err != nil {
		return err
	}

	if conversation.UserIDOne != userID && conversation.UserIDTwo != userID {
		return nil
	}

	return c.conversationRepo.Delete(ctx, conversationID)
}

func (c *ConversationService) UpdateLastMessage(ctx context.Context, conversationID string, messageID string) error {
	return c.conversationRepo.UpdateLastMessage(ctx, conversationID, messageID)
}
