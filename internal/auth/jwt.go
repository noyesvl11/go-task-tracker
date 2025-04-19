package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

var secret = []byte("your-secret-key")

// Генерация JWT с user_id и role
func GenerateJWT(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// Извлекаем user_id и role из JWT
func ExtractUserIDAndRole(c *gin.Context) (uint, string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return 0, "", errors.New("missing Authorization header")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", errors.New("invalid claims")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, "", errors.New("invalid user_id in token")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return 0, "", errors.New("invalid role in token")
	}

	return uint(userIDFloat), role, nil
}
