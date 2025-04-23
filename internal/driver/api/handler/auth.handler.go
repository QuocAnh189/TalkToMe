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

type AuthHandler struct {
	service service.IAuthService
}

func NewAuthHandler(service service.IAuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

// Register handles user registration requests.
func (h *AuthHandler) SignUp(c *gin.Context) {
	var req dto.SignUpRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	accessToken, refreshToken, user, err := h.service.SignUp(c, &req)
	if err != nil {
		logger.Error("Failed to sign up ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to sign up")
		return
	}

	var res dto.SignUpResponse
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	utils.MapStruct(&res.User, user)

	response.JSON(c, http.StatusOK, res)
}

// Login handles user login requests.
func (h *AuthHandler) SignIn(c *gin.Context) {
	var req dto.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}
	accessToken, refreshToken, user, err := h.service.SignIn(c, &req)

	if err != nil {
		logger.Error("Failed to sign up ", err)
		switch err.Error() {
		case "wrong password":
			response.Error(c, http.StatusConflict, err, "Wrong password")
		case "record not found":
			response.Error(c, http.StatusConflict, err, "Email does not exist")
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to sign up")
		}
		return
	}

	var res dto.SignInResponse
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	utils.MapStruct(&res.User, user)
	response.JSON(c, http.StatusOK, res)
}

// RefreshToken handles token refresh requests.
func (h *AuthHandler) SignOut(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		response.Error(c, http.StatusBadRequest, nil, "Missing Authorization header")
		return
	}

	jit := c.GetString("jit")

	userID, exists := c.Get("userId")
	if !exists {
		response.Error(c, http.StatusNotFound, nil, "Unauthorized")
		return
	}

	err := h.service.SignOut(c, userID.(string), jit)
	if err != nil {
		logger.Error("Failed to sign out", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to sign out")
		return
	}

	response.JSON(c, http.StatusOK, "Logout successfully")
}

// Logout handles user logout requests.
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		response.Error(c, http.StatusUnauthorized, nil, "Unauthorized")
		return
	}

	jit := c.GetString("jit")
	accessToken, err := h.service.RefreshToken(c, userId, jit)
	if err != nil {
		logger.Error("Failed to refresh token", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	res := dto.RefreshTokenResponse{
		AccessToken: accessToken,
	}
	response.JSON(c, http.StatusOK, res)
}
