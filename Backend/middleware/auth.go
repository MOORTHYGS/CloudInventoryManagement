package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte("0mPanXw9ag7z7yXLMIy/8MI6kCcD9J55nM74XYVpdogNVl+Z4m8EwHxzOmxCR/N3jP/OqKdpOQFRrMSzUuSLSw==") // Supabase secret

// GenerateToken generates JWT token for system use
func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    userID,
		"token_type": "system",                             // ✅ added
		"exp":        time.Now().Add(time.Hour * 1).Unix(), // 1 hour expiry
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

// AuthMiddleware validates system JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrInvalidKey
			}
			return JWTSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Validate claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if claims["token_type"] != "system" { // ✅ enforce token type
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token type"})
				c.Abort()
				return
			}
			c.Set("user_id", claims["user_id"])
		}

		c.Next()
	}
}
