package handler

import (
	"gochat/internal/application/dto"
	"gochat/internal/domain/service"
	"gochat/pkg/logger"
	"gochat/pkg/response"
	"gochat/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConversationHandler struct {
	service service.IConversationService
}

// NewConversationHandler creates a new ConversationHandler.
func NewConversationHandler(service service.IConversationService) *ConversationHandler {
	return &ConversationHandler{
		service: service,
	}
}

// ListConversations handles requests to list conversations for the current user.
func (h *ConversationHandler) ListConversations(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, nil, "Unauthorized")
		return
	}

	var req dto.ListConversationRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get query", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	conversations, pagination, err := h.service.ListUserConversations(c, &req, userID)
	if err != nil {
		logger.Error("Failed to list conversations: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to list conversations")
		return
	}

	var res dto.ListConversationResponse
	utils.MapStruct(&res.Conversations, conversations)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

// GetOrCreateConversation handles requests to get or create a conversation with a specific user.
func (h *ConversationHandler) CreateConversation(c *gin.Context) {
	var req dto.CreateConversationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	conversation, err := h.service.CreateConversation(c, &req)
	if err != nil {
		logger.Error("Failed to create conversation: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to create conversation")
		return
	}

	var res dto.CreateConversationResponse
	utils.MapStruct(&res, conversation)
	response.JSON(c, http.StatusOK, res)
}

// GetConversationDetails handles requests to get details of a specific conversation.
func (h *ConversationHandler) GetConversationDetails(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, nil, "Unauthorized")
		return
	}

	conversationID := c.Param("conversationId")
	if conversationID == "" {
		response.Error(c, http.StatusBadRequest, nil, "Conversation ID is required")
		return
	}

	conversation, err := h.service.GetConversationDetails(c, conversationID, userID)
	if err != nil {
		logger.Error("Failed to get conversation details: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to get conversation details")
		return
	}

	if conversation == nil {
		response.Error(c, http.StatusNotFound, nil, "Conversation not found")
		return
	}

	var res dto.ConversationResponse
	utils.MapStruct(&res, conversation)
	response.JSON(c, http.StatusOK, res)
}

// DeleteConversation handles requests to delete (or hide) a conversation for the user.
func (h *ConversationHandler) DeleteConversation(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, nil, "Unauthorized")
		return
	}

	conversationID := c.Param("conversationId")
	if conversationID == "" {
		response.Error(c, http.StatusBadRequest, nil, "Conversation ID is required")
		return
	}

	err := h.service.DeleteConversationForUser(c, conversationID, userID)
	if err != nil {
		logger.Error("Failed to delete conversation: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to delete conversation")
		return
	}

	response.JSON(c, http.StatusOK, "Conversation deleted successfully")
}
