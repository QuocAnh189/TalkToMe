package token

import (
	"strings"
	"time"

	"gochat/config"
	"gochat/pkg/logger"
	"gochat/pkg/token"

	"github.com/golang-jwt/jwt"

	"gochat/utils"
)

const (
	AccessTokenExpiredTime  = 5 * 60 * 60    // 5 hours
	RefreshTokenExpiredTime = 30 * 24 * 3600 // 30 days
)

type JTWMarker struct {
	AccessTokenType  string
	RefreshTokenType string
}

func NewJTWMarker() (*JTWMarker, error) {
	return &JTWMarker{
		AccessTokenType:  token.AccessTokenType,
		RefreshTokenType: token.RefreshTokenType,
	}, nil
}

func (j *JTWMarker) GenerateAccessToken(payload *token.AuthPayload) string {
	cfg := config.GetConfig()
	newPayload := token.NewAuthPayload(payload.ID, payload.Email, payload.Role, time.Minute, j.AccessTokenType)

	tokenContent := jwt.MapClaims{
		"payload": newPayload,
		"exp":     time.Now().Add(time.Second * AccessTokenExpiredTime).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte(cfg.AuthSecret))
	if err != nil {
		logger.Error("Failed to generate access token: ", err)
		return ""
	}

	return token
}

func (j *JTWMarker) GenerateRefreshToken(payload *token.AuthPayload) string {
	cfg := config.GetConfig()
	newPayload := token.NewAuthPayload(payload.ID, payload.Email, payload.Role, time.Minute, j.RefreshTokenType)
	tokenContent := jwt.MapClaims{
		"payload": newPayload,
		"exp":     time.Now().Add(time.Second * RefreshTokenExpiredTime).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte(cfg.AuthSecret))
	if err != nil {
		logger.Error("Failed to generate refresh token: ", err)
		return ""
	}

	return token
}

func (j *JTWMarker) ValidateToken(jwtToken string) (*token.AuthPayload, error) {
	cfg := config.GetConfig()
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	result, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.AuthSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !result.Valid {
		return nil, jwt.ErrInvalidKey
	}

	var data *token.AuthPayload
	utils.MapStruct(&data, tokenData["payload"])

	return data, nil
}
