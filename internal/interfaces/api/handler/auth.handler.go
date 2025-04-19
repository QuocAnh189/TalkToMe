package handler

import (
	"github.com/gin-gonic/gin"
	// Import necessary service interfaces, e.g., "gochat/internal/domain/service"
)

// AuthHandler handles authentication related requests.
type AuthHandler struct {
	// Add dependencies like auth service:
	// authService service.IAuthService
}

// NewAuthHandler creates a new AuthHandler.
func NewAuthHandler( /* Pass dependencies here */ ) *AuthHandler {
	return &AuthHandler{ /* Initialize dependencies */ }
}

// Register handles user registration requests.
func (h *AuthHandler) SignIn(ctx *gin.Context)

// Login handles user login requests.
func (h *AuthHandler) SignUp(ctx *gin.Context)

// RefreshToken handles token refresh requests.
func (h *AuthHandler) RefreshToken(ctx *gin.Context)

// Logout handles user logout requests.
func (h *AuthHandler) SignOut(ctx *gin.Context)
