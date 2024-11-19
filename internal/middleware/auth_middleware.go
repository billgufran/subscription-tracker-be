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
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Authorization header is required"))
			return
		}

		// Check if the header starts with "Bearer "
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid authorization header format"))
			return
		}

		claims, err := auth.ValidateToken(parts[1])
		if err != nil {
			switch err {
			case auth.ErrExpiredToken:
				c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Token has expired"))
			default:
				c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid token"))
			}
			return
		}

		// Store user information in context
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)

		c.Next()
	}
}
