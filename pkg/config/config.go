package config

import (
	"boilerplate-api/pkg/logger"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"

	glogger "gorm.io/gorm/logger"
)

var conf *Config

type Config struct {
	AppEnv string `mapstructure:"app_env"`

	LogConf logger.Logger `mapstructure:"log_conf"`
	Server  ServerConfig
	JWT     JWTConfig
	Redis   RedisConfig
	MongoDB DatabaseConfig
	MySQL   DatabaseConfig
}

//type Logger struct {
//	LogPath      string `mapstructure:"log_path"`
//	PanicLogPath string `mapstructure:"panic_log_path"`
//	LogLevel     string `mapstructure:"log_level"`
//	LogType      string `mapstructure:"log_type"`
//	MaxSizeMB    int    `mapstructure:"max_size_mb"`  // Max size in megabytes before rotation
//	MaxBackups   int    `mapstructure:"max_backups"`  // Maximum number of old log files to retain
//	MaxAgeDays   int    `mapstructure:"max_age_days"` // Maximum number of days to retain old log files
//	Compress     bool   `mapstructure:"compress"`
//}

type ServerConfig struct {
	Mode     string
	LogLevel string `mapstructure:"log_level"`
	Port     int    `mapstructure:"port"`
}

type JWTConfig struct {
	Secret                     string `mapstructure:"secret"`
	AccessTokenExpirationTime  int    `mapstructure:"access_token_expiration_time"`
	RefreshTokenExpirationTime int    `mapstructure:"refresh_token_expiration_time"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
	SSLMode  string `mapstructure:"ssl_mode"`

	Logger glogger.Interface
}

type RedisConfig struct {
	Host              string `mapstructure:"host"`
	Port              int    `mapstructure:"port"`
	Password          string `mapstructure:"password"`
	NumberChannelRoom int    `mapstructure:"number_channel_room"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	switch conf.AppEnv {
	case "development":
		conf.Server.Mode = gin.DebugMode
		conf.MySQL.Logger = glogger.Default
	case "production":
		conf.Server.Mode = gin.ReleaseMode
		conf.MySQL.Logger = logger.NewGormLogger()
	default:
		conf.Server.Mode = gin.DebugMode
		conf.MySQL.Logger = glogger.Default
	}
	return conf, nil
}
