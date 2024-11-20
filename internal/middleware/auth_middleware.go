package middleware

import (
	"net/http"
	"strings"

	"subscription-tracker/internal/auth"
	"subscription-tracker/internal/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.HandleHttpError(c, utils.NewUnauthorizedError("Authorization header is required"))
			c.Abort()
			return
		}

		// Check if the header starts with "Bearer"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid authorization header format"))
			return
		}

		claims, err := auth.ValidateToken(parts[1])
		if err != nil {
			utils.HandleHttpError(c, err)
			c.Abort()
			return
		}

		// Store user information in context
		c.Set("userID", claims.UserID)

		c.Next()
	}
}
