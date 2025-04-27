package handler

import (
	"gochat/internal/application/dto"
	"gochat/internal/domain/service"
	"gochat/pkg/logger"
	"gochat/pkg/response"
	"mime/multipart"
	"net/http"

	"gochat/utils"

	"github.com/gin-gonic/gin"
)

// MessageHandler handles message related requests.
type MessageHandler struct {
	service service.IMessagingService
}

// NewMessageHandler creates a new MessageHandler.
func NewMessageHandler(service service.IMessagingService) *MessageHandler {
	return &MessageHandler{
		service: service,
	}
}

// SendMessage handles requests to send a new message.
func (h *MessageHandler) SendMessage(c *gin.Context) {
	var req dto.SendMessageRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to bind request", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid request parameters")
		return
	}

	userID := c.GetString("userId")
	form, _ := c.MultipartForm()
	var attachments []*multipart.FileHeader
	if form != nil && form.File != nil {
		if files, ok := form.File["attachments"]; ok {
			attachments = files
		}
	}

	message, err := h.service.SendMessage(c, &req, attachments, userID)
	if err != nil {
		logger.Error("Failed to send message", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to send message")
		return
	}

	var res dto.SendMessageResponse
	utils.MapStruct(&res, message)
	response.JSON(c, http.StatusCreated, res)
}

// GetGroupMessages handles requests to list messages in a group.
func (h *MessageHandler) GetGroupMessages(c *gin.Context) {
	userID := c.GetString("userId")
	groupID := c.Param("groupId")
	if groupID == "" {
		response.Error(c, http.StatusBadRequest, nil, "Group ID is required")
		return
	}

	var req dto.ListMessageRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to bind request", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid request parameters")
		return
	}

	messages, pagination, err := h.service.GetGroupMessages(c, &req, groupID, userID)
	if err != nil {
		logger.Error("Failed to get group messages", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to get group messages")
		return
	}

	if messages == nil {
		response.Error(c, http.StatusForbidden, nil, "Not authorized to view group messages")
		return
	}

	var res dto.ListMessageResponse
	utils.MapStruct(&res.Messages, messages)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

// GetConversationMessages handles requests to list messages in a conversation.
func (h *MessageHandler) GetConversationMessages(c *gin.Context) {
	userID := c.GetString("userId")
	conversationID := c.Param("conversationId")
	if conversationID == "" {
		response.Error(c, http.StatusBadRequest, nil, "Conversation ID is required")
		return
	}

	var req dto.ListMessageRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to bind request", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid request parameters")
		return
	}

	messages, pagination, err := h.service.GetConversationMessages(c, &req, conversationID, userID)
	if err != nil {
		logger.Error("Failed to get conversation messages", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to get conversation messages")
		return
	}

	var res dto.ListMessageResponse
	utils.MapStruct(&res.Messages, messages)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

// GetMessage handles requests to get a specific message.
func (h *MessageHandler) GetMessage(c *gin.Context) {
	userID := c.GetString("userId")
	messageID := c.Param("messageId")
	if messageID == "" {
		response.Error(c, http.StatusBadRequest, nil, "Message ID is required")
		return
	}

	message, err := h.service.GetMessageByID(c, messageID, userID)
	if err != nil {
		logger.Error("Failed to get message", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to get message")
		return
	}

	if message == nil {
		response.Error(c, http.StatusNotFound, nil, "Message not found")
		return
	}

	var res dto.MessageResponse
	utils.MapStruct(&res, message)
	response.JSON(c, http.StatusOK, res)
}

// UpdateMessage handles requests to update a message.
func (h *MessageHandler) UpdateMessage(c *gin.Context) {
	userID := c.GetString("userId")
	messageID := c.Param("messageId")
	if messageID == "" {
		response.Error(c, http.StatusBadRequest, nil, "Message ID is required")
		return
	}

	var req dto.UpdateMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to bind request", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid request parameters")
		return
	}

	message, err := h.service.UpdateMessageByID(c, messageID, userID, &req)
	if err != nil {
		logger.Error("Failed to update message", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to update message")
		return
	}

	if message == nil {
		response.Error(c, http.StatusForbidden, nil, "Not authorized to update message")
		return
	}

	var res dto.UpdateMessageResponse
	utils.MapStruct(&res, message)
	response.JSON(c, http.StatusOK, res)
}

// DeleteMessage handles requests to delete a message.
func (h *MessageHandler) DeleteMessage(c *gin.Context) {
	userID := c.GetString("userId")
	messageID := c.Param("messageId")
	if messageID == "" {
		response.Error(c, http.StatusBadRequest, nil, "Message ID is required")
		return
	}

	err := h.service.DeleteMessageByID(c, messageID, userID)
	if err != nil {
		logger.Error("Failed to delete message", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to delete message")
		return
	}

	response.JSON(c, http.StatusOK, "Message deleted successfully")
}
