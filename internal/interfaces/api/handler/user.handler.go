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

type UserHandler struct {
	service service.IUserService
}

func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// GetProfile handles requests to get the current user's profile.
func (h *UserHandler) GetProfile(c *gin.Context) {
	userId := c.GetString("userId")
	user, err := h.service.GetUserProfile(c, userId)
	if err != nil {
		logger.Error("Failed to get user detail: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.UserResponse
	utils.MapStruct(&res, user)
	response.JSON(c, http.StatusOK, res)
}

// UpdateProfile handles requests to update the current user's profile.
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var req dto.UpdateProfileRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to bind request: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid request")
		return
	}

	userId := c.GetString("userId")
	user, err := h.service.UpdateUserProfile(c, userId, &req)
	if err != nil {
		logger.Error("Failed to update user: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.UserResponse
	utils.MapStruct(&res, user)
	response.JSON(c, http.StatusOK, res)
}

// GetUserByID handles requests to get a user's profile by their ID.
func (h *UserHandler) GetUserByID(c *gin.Context) {
	userId := c.Param("id")
	user, err := h.service.GetUserByID(c, userId)
	if err != nil {
		logger.Error("Failed to get user detail: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.UserResponse
	utils.MapStruct(&res, user)
	response.JSON(c, http.StatusOK, res)
}

// SearchUsers handles requests to search for users.
func (h *UserHandler) SearchUsers(c *gin.Context) {
	var req dto.ListUserRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get query", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	users, pagination, err := h.service.SearchUsers(c, &req)
	if err != nil {
		logger.Error("Failed to get users", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to get users")
		return
	}

	var res dto.ListUserResponse
	utils.MapStruct(&res.Users, users)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}
