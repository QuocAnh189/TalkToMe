package handler

import (
	"gochat/internal/application/dto"
	"gochat/internal/domain/service"
	"gochat/pkg/logger"
	"gochat/pkg/response"
	"net/http"

	"gochat/utils"

	"github.com/gin-gonic/gin"
)

// GroupHandler handles group chat related requests.
type GroupHandler struct {
	service service.IGroupChatService
}

// NewGroupHandler creates a new GroupHandler.
func NewGroupHandler(service service.IGroupChatService) *GroupHandler {
	return &GroupHandler{
		service: service,
	}
}

// CreateGroup handles requests to create a new group.
func (h *GroupHandler) CreateGroup(c *gin.Context) {
	var req dto.CreateGroupRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to bind request", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid request parameters")
		return
	}

	userID := c.GetString("userId")
	group, err := h.service.CreateGroup(c, &req, userID)
	if err != nil {
		logger.Error("Failed to create group", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to create group")
		return
	}

	var res dto.CreateGroupResponse
	utils.MapStruct(&res, group)
	response.JSON(c, http.StatusCreated, res)
}

// ListUserGroups handles requests to list groups the user is a member of.
func (h *GroupHandler) ListUserGroups(c *gin.Context) {
	userID := c.GetString("userId")
	var req dto.ListGroupRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to bind request", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid request parameters")
		return
	}

	groups, pagination, err := h.service.ListUserGroups(c, &req, userID)
	if err != nil {
		logger.Error("Failed to list groups", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to list groups")
		return
	}

	var res dto.ListGroupResponse
	utils.MapStruct(&res.Groups, groups)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

// GetGroupDetails handles requests to get details of a specific group.
func (h *GroupHandler) GetGroupDetails(c *gin.Context) {
	userID := c.GetString("userId")
	groupID := c.Param("groupId")
	if groupID == "" {
		response.Error(c, http.StatusBadRequest, nil, "Group ID is required")
		return
	}

	group, err := h.service.GetGroupDetails(c, groupID, userID)
	if err != nil {
		logger.Error("Failed to get group details", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to get group details")
		return
	}

	if group == nil {
		response.Error(c, http.StatusNotFound, nil, "Group not found or user not a member")
		return
	}

	var res dto.GroupResponse
	utils.MapStruct(&res, group)
	response.JSON(c, http.StatusOK, res)
}

// UpdateGroup handles requests to update group information.
func (h *GroupHandler) UpdateGroup(c *gin.Context) {
	userID := c.GetString("userId")
	groupID := c.Param("groupId")
	if groupID == "" {
		response.Error(c, http.StatusBadRequest, nil, "Group ID is required")
		return
	}

	var req dto.UpdateGroupRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to bind request", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid request parameters")
		return
	}

	group, err := h.service.UpdateGroupInfo(c, &req, groupID, userID)
	if err != nil {
		logger.Error("Failed to update group", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to update group")
		return
	}

	if group == nil {
		response.Error(c, http.StatusForbidden, nil, "Not authorized to update group")
		return
	}

	var res dto.GroupResponse
	utils.MapStruct(&res, group)
	response.JSON(c, http.StatusOK, res)
}

// DeleteGroup handles requests to delete a group.
func (h *GroupHandler) DeleteGroup(c *gin.Context) {
	userID := c.GetString("userId")
	groupID := c.Param("groupId")
	if groupID == "" {
		response.Error(c, http.StatusBadRequest, nil, "Group ID is required")
		return
	}

	err := h.service.DeleteGroup(c, groupID, userID)
	if err != nil {
		logger.Error("Failed to delete group", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to delete group")
		return
	}

	response.JSON(c, http.StatusOK, "Group deleted successfully")
}

// AddMember handles requests to add a member to a group.
func (h *GroupHandler) AddMember(c *gin.Context) {
	userID := c.GetString("userId")
	var req dto.AddMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to bind request", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid request parameters")
		return
	}

	err := h.service.AddGroupMember(c, &req, userID)
	if err != nil {
		logger.Error("Failed to add member", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to add member")
		return
	}

	response.JSON(c, http.StatusOK, "Member added successfully")
}

// RemoveMember handles requests to remove a member from a group.
func (h *GroupHandler) RemoveMember(c *gin.Context) {
	var req dto.RemoveMemberRequest
	req.GroupID = c.Param("groupId")
	req.UserID = c.Param("userId")

	logger.Info(req)

	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Error("Failed to bind request", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid request parameters")
		return
	}

	removerID := c.GetString("userId")
	err := h.service.RemoveGroupMember(c, &req, removerID)
	if err != nil {
		logger.Error("Failed to remove member", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to remove member")
		return
	}

	response.JSON(c, http.StatusOK, "Member removed successfully")
}

// ListMembers handles requests to list members of a group.
func (h *GroupHandler) ListMembers(c *gin.Context) {
	userID := c.GetString("userId")
	groupID := c.Param("groupId")
	if groupID == "" {
		response.Error(c, http.StatusBadRequest, nil, "Group ID is required")
		return
	}

	members, err := h.service.ListGroupMembers(c, groupID, userID)
	if err != nil {
		logger.Error("Failed to list members", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to list members")
		return
	}

	if members == nil {
		response.Error(c, http.StatusForbidden, nil, "Not authorized to view group members")
		return
	}

	response.JSON(c, http.StatusOK, members)
}
