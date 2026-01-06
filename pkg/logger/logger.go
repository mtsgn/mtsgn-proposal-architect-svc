package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	logs "github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func init() {
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

var LoggerContext = "context_logger"

type Logger struct {
	LogPath      string `mapstructure:"log_path"`
	PanicLogPath string `mapstructure:"panic_log_path"`
	LogLevel     string `mapstructure:"log_level"`
	LogType      string `mapstructure:"log_type"`
	MaxSizeMB    int    `mapstructure:"max_size_mb"`  // Max size in megabytes before rotation
	MaxBackups   int    `mapstructure:"max_backups"`  // Maximum number of old log files to retain
	MaxAgeDays   int    `mapstructure:"max_age_days"` // Maximum number of days to retain old log files
	Compress     bool   `mapstructure:"compress"`
}

// InitializeLogger sets up zerolog with JSON output to file and console
func InitializeLogger(conf *Logger) {
	parsedLogLevel := parseLogLevel(conf.LogLevel)
	zerolog.SetGlobalLevel(parsedLogLevel)

	var logger zerolog.Logger

	conf.LogPath = filepath.Clean(conf.LogPath)
	conf.PanicLogPath = filepath.Clean(conf.PanicLogPath)

	if conf.LogType == "file" {
		fileLogger := &lumberjack.Logger{
			Filename:   conf.LogPath,
			MaxSize:    conf.MaxSizeMB,
			MaxBackups: conf.MaxBackups,
			MaxAge:     conf.MaxAgeDays,
			Compress:   conf.Compress,
		}

		multi := zerolog.MultiLevelWriter(
			zerolog.SyncWriter(fileLogger), // File: JSON
			zerolog.SyncWriter(os.Stderr),  // Console: JSON
		)

		logger = zerolog.New(multi).With().Timestamp().Logger()

		// Skip panic redirect on Windows
		if runtime.GOOS != "windows" {
			panicLogger := &lumberjack.Logger{
				Filename:   conf.PanicLogPath,
				MaxSize:    conf.MaxSizeMB,
				MaxBackups: conf.MaxBackups,
				MaxAge:     conf.MaxAgeDays,
				Compress:   conf.Compress,
			}
			stdlogWriter := zerolog.SyncWriter(panicLogger)
			stdlog := zerolog.New(stdlogWriter).With().Timestamp().Logger()
			logs.Logger = stdlog
		}
	} else {
		// JSON to console only
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	}

	// Set logger globally
	logs.Logger = logger
	zerolog.DefaultContextLogger = &logger

	logger.Info().Msgf("current log level: %s", parsedLogLevel.String())
}

// parseLogLevel changes log level from the config to lower case and then to Level type.
func parseLogLevel(level string) zerolog.Level {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}

func GetLogger(c *gin.Context) zerolog.Logger {
	if logger, exists := c.Get(LoggerContext); exists {
		if l, ok := logger.(zerolog.Logger); ok {
			return l
		}
	}
	return zerolog.Nop()
}
