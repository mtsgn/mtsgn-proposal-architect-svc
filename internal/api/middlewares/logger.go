package middleware

import (
	"boilerplate-api/internal/common/utils"
	"boilerplate-api/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Retrieve the Request-ID from the context
		requestID := c.GetString(RequestIDKey)
		// This should not happen at all
		if requestID == "" {
			requestID = "unknown"
		}

		loggerCtx := log.With().
			Str("request_id", requestID).
			Logger()

		c.Set(logger.LoggerContext, loggerCtx)

		// Process the request
		c.Next()

		// After request processing`
		duration := time.Since(start)
		status := c.Writer.Status()

		// Get the user context and log it
		userCtx := utils.GetUserContext(c)
		if userCtx != nil {
			loggerCtx = loggerCtx.With().Str("user_id", userCtx.UserID).Logger()
		}

		// Log the request details with Request-ID
		loggerCtx.Info().
			Str("method", method).
			Str("path", path).
			Int("status", status).
			Dur("duration", duration).
			Msg("request processed")
	}
}
