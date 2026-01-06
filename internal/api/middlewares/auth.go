package middleware

import (
	"boilerplate-api/internal/common/context"
	"boilerplate-api/internal/common/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header is required"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization header format"})
			return
		}

		tokenString := parts[1]
		userCtx, err := utils.ValidateToken(tokenString, ctx)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
			return
		}

		c.Set(utils.UserCtx, userCtx)

		//c.Set("userId", userCtx.UserID)

		c.Next()
	}
}
