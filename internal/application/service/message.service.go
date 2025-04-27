package service

import (
	"context"
	"gochat/internal/application/dto"
	"gochat/internal/domain/model"
	"gochat/internal/domain/repository"
	"gochat/pkg/paging"
	"gochat/pkg/storage"
	"gochat/pkg/validation"
	"gochat/utils"
	"mime/multipart"
)

type MessageService struct {
	validator              validation.Validation
	messageRepo            repository.IMessageRepository
	messageAttachmentsRepo repository.IMessageAttachmentRepository
	groupRepo              repository.IGroupRepository
	storage                storage.IUploadService
}

func NewMessageService(
	validator validation.Validation,
	messageRepo repository.IMessageRepository,
	messageAttachmentsRepo repository.IMessageAttachmentRepository,
	groupRepo repository.IGroupRepository,
	storage storage.IUploadService,
) *MessageService {
	return &MessageService{
		validator:              validator,
		messageRepo:            messageRepo,
		messageAttachmentsRepo: messageAttachmentsRepo,
		groupRepo:              groupRepo,
		storage:                storage,
	}
}

func (m *MessageService) SendMessage(ctx context.Context, req *dto.SendMessageRequest, attachments []*multipart.FileHeader, senderID string) (*model.Message, error) {
	if err := m.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	message := &model.Message{
		GroupID:        req.GroupID,
		ConversationID: req.ConversationID,
		Message:        req.Message,
		SenderID:       senderID,
	}

	if err := m.messageRepo.Create(ctx, message); err != nil {
		return nil, err
	}

	if len(attachments) > 0 {
		messageAttachments := make([]*model.MessageAttachment, 0, len(attachments)) // Khởi tạo slice rỗng
		for _, file := range attachments {
			if file != nil && file.Filename != "" {
				fileURL, err := m.storage.UploadFile(ctx, file, "messages")
				if err != nil {
					return nil, err
				}
				attachment := &model.MessageAttachment{
					MessageID: message.ID,
					Type:      utils.GetAttachmentTypeFromFilename(file.Filename),
					Filename:  file.Filename,
					URL:       fileURL,
				}
				messageAttachments = append(messageAttachments, attachment) // Thêm vào slice
			}
		}
		if err := m.messageAttachmentsRepo.CreateBatch(ctx, messageAttachments); err != nil {
			return nil, err
		}
	}

	return message, nil
}

func (m *MessageService) GetGroupMessages(ctx context.Context, req *dto.ListMessageRequest, groupID string, userID string) ([]*model.Message, *paging.Pagination, error) {
	if err := m.validator.ValidateStruct(req); err != nil {
		return nil, nil, err
	}

	isMember, err := m.groupRepo.IsMember(ctx, groupID, userID)
	if err != nil {
		return nil, nil, err
	}

	if !isMember {
		return nil, nil, nil
	}

	return m.messageRepo.FindMessagesByGroupID(ctx, req, groupID)
}

func (m *MessageService) GetConversationMessages(ctx context.Context, req *dto.ListMessageRequest, conversationID string, userID string) ([]*model.Message, *paging.Pagination, error) {
	if err := m.validator.ValidateStruct(req); err != nil {
		return nil, nil, err
	}

	return m.messageRepo.FindMessagesByConversationID(ctx, req, conversationID)
}

func (m *MessageService) GetMessageByID(ctx context.Context, messageID string, userID string) (*model.Message, error) {
	message, err := m.messageRepo.FindOne(ctx, messageID)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (m *MessageService) UpdateMessageByID(ctx context.Context, messageID string, userID string, req *dto.UpdateMessageRequest) (*model.Message, error) {
	if err := m.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	message, err := m.messageRepo.FindByID(ctx, messageID)
	if err != nil {
		return nil, err
	}

	if message.SenderID != userID {
		return nil, nil
	}

	utils.MapStruct(message, req)

	if err := m.messageRepo.Update(ctx, message); err != nil {
		return nil, err
	}

	return message, nil
}

func (m *MessageService) DeleteMessageByID(ctx context.Context, messageID string, userID string) error {
	message, err := m.messageRepo.FindOne(ctx, messageID)
	if err != nil {
		return err
	}

	if message.SenderID != userID {
		return nil
	}

	if len(message.Attachments) > 0 {
		for _, attachments := range message.Attachments {
			m.storage.DeleteFile(ctx, attachments.URL)
		}

		if err := m.messageAttachmentsRepo.DeleteByMessageID(ctx, messageID); err != nil {
			return err
		}
	}

	return m.messageRepo.Delete(ctx, message)
}
