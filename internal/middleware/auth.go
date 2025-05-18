package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Anwarjondev/fast-food/internal/repository"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		user, err := repository.GetUserByToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("user_id", user.ID)
		c.Next()
	}
}

// ConfirmationMiddleware is a middleware that gets the user ID from the X-User-ID header
// This is used for confirmation endpoints where the user isn't authenticated yet
func ConfirmationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr := c.GetHeader("X-User-ID")
		if userIDStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			c.Abort()
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
