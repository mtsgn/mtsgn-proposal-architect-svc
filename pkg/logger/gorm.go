package logger

import (
	"context"
	"github.com/rs/zerolog"
	logs "github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
	"time"
)

// GormLogger is a custom GORM logger that uses zerolog
type GormLogger struct {
	Logger        zerolog.Logger
	SlowThreshold time.Duration
	LogLevel      logger.LogLevel
}

// LogMode changes the log level for this logger instance
func (z *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	// Return a new instance with the updated log level
	return &GormLogger{
		Logger:        z.Logger,
		SlowThreshold: z.SlowThreshold,
		LogLevel:      level,
	}
}

// Info logs an info message
func (z *GormLogger) Info(_ context.Context, msg string, data ...interface{}) {
	if z.LogLevel >= logger.Info {
		z.Logger.Info().Msgf(msg, data...)
	}
}

// Warn logs a warning message
func (z *GormLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	if z.LogLevel >= logger.Warn {
		z.Logger.Warn().Msgf(msg, data...)
	}
}

// Error logs an error message
func (z *GormLogger) Error(_ context.Context, msg string, data ...interface{}) {
	if z.LogLevel >= logger.Error {
		z.Logger.Error().Msgf(msg, data...)
	}
}

// Trace logs a SQL trace message (with time duration)
func (z *GormLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	if z.LogLevel >= logger.Info {
		elapsed := time.Since(begin)
		sql, args := fc()
		// Log SQL query and execution time
		if err != nil {
			z.Logger.Error().Err(err).Str("sql", sql).Dur("duration", elapsed).Interface("args", args).Msg("SQL Error")
		} else if elapsed > z.SlowThreshold {
			z.Logger.Warn().Str("sql", sql).Dur("duration", elapsed).Interface("args", args).Msg("Slow SQL")
		} else {
			z.Logger.Info().Str("sql", sql).Dur("duration", elapsed).Interface("args", args).Msg("SQL Query")
		}
	}
}

// NewGormLogger creates a new ormLogger instance
func NewGormLogger() *GormLogger {
	zl := logs.Logger
	return &GormLogger{
		Logger:        zl,
		SlowThreshold: 5 * time.Nanosecond, // increase threshold if needed
		LogLevel:      logger.Warn,
	}
}
