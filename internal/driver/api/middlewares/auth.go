package middlewares

import (
	"encoding/json"
	"fmt"
	"gochat/internal/infrashstructrure/cache"
	"gochat/pkg/logger"
	"gochat/pkg/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	token token.IMarker
	cache cache.IRedis
}

func NewAuthMiddleware(token token.IMarker, cache cache.IRedis) *AuthMiddleware {
	return &AuthMiddleware{
		token: token,
		cache: cache,
	}
}

func (a *AuthMiddleware) TokenAuth() gin.HandlerFunc {
	return a.Token(token.AccessTokenType, a.cache)
}

func (a *AuthMiddleware) TokenRefresh() gin.HandlerFunc {
	return a.Token(token.RefreshTokenType, a.cache)
}

func (a *AuthMiddleware) Token(tokenType string, cache cache.IRedis) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.GetHeader("Authorization")
		if tokenValue == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		payload, err := a.token.ValidateToken(tokenValue)
		if err != nil || payload == nil || payload.Type != tokenType {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			c.Abort()
			return
		}

		var rawValue string
		if err = cache.Get(fmt.Sprintf("blacklist:%s_%s", payload.ID, payload.Jit), &rawValue); err != nil {
			logger.Error("Failed to get value from Redis:", err)
		}

		var value map[string]string
		err = json.Unmarshal([]byte(rawValue), &value)
		if err != nil {
			logger.Error("Failed to unmarshal JSON:", err)
		}

		if value["status"] == "blacklisted" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is blacklisted"})
			c.Abort()
			return
		}

		c.Set("userId", payload.ID)
		c.Set("role", payload.Role)
		c.Set("jit", payload.Jit)
		c.Set("token", tokenValue)
		c.Next()
	}
}
