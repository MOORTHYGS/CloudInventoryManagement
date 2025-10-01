package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var UserJWTSecret = []byte("0mPanXw9ag7z7yXLMIy/8MI6kCcD9J55nM74XYVpdogNVl+Z4m8EwHxzOmxCR/N3jP/OqKdpOQFRrMSzUuSLSw==") // Supabase secret

// Claims structure for type safety
type UserClaims struct {
	UserID    string `json:"user_id"`
	TenantID  string `json:"tenant_id"`
	Role      string `json:"role"`
	TokenType string `json:"token_type"` // ✅ added
	jwt.RegisteredClaims
}

// GenerateUserToken generates JWT token for a user
func GenerateUserToken(userID, tenantID, role string) (string, error) {
	claims := UserClaims{
		UserID:    userID,
		TenantID:  tenantID,
		Role:      role,
		TokenType: "user", // ✅ enforce type
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "cloud_inventory",
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(UserJWTSecret)
}

// UserAuthMiddleware validates user JWT
func UserAuthMiddleware() gin.HandlerFunc {
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
		token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrInvalidKey
			}
			return UserJWTSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Validate claims
		if claims, ok := token.Claims.(*UserClaims); ok {
			if claims.TokenType != "user" { // ✅ enforce type
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token type"})
				c.Abort()
				return
			}
			c.Set("user_id", claims.UserID)
			c.Set("tenant_id", claims.TenantID)
			c.Set("role", claims.Role)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}
