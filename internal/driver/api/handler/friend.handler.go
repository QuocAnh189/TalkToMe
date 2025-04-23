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

type FriendHandler struct {
	service service.IFriendshipService
}

func NewFriendHandler(service service.IFriendshipService) *FriendHandler {
	return &FriendHandler{
		service: service,
	}
}

func (h *FriendHandler) AddFriend(c *gin.Context) {
	var req dto.AddFriendRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	err := h.service.AddFriend(c, &req)
	if err != nil {
		logger.Error("Failed to accept friend request: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to accept friend request")
		return
	}

	response.JSON(c, http.StatusOK, "Add friends successfully")
}

func (h *FriendHandler) ListFriends(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, nil, "Unauthorized")
		return
	}

	var req dto.ListFriendRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get query", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	friends, pagination, err := h.service.ListFriends(c, &req, userID)
	if err != nil {
		logger.Error("Failed to list friends: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to list friends")
		return
	}

	var res dto.ListFriendResponse
	utils.MapStruct(&res.Users, friends)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

func (h *FriendHandler) RemoveFriend(c *gin.Context) {
	var req dto.RemoveFriendRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	err := h.service.RemoveFriend(c, req.FromID, req.ToID)
	if err != nil {
		logger.Error("Failed to remove friend: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to remove friend")
		return
	}

	response.JSON(c, http.StatusOK, "Friend removed successfully")
}
