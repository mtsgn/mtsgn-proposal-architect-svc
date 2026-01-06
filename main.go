// @title           boilerplate API
// @version         1.0
// @description     API server for boilerplate application
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description RoomType "Bearer" followed by a space and JWT token.
package main

import (
	_ "boilerplate-api/docs"
	"boilerplate-api/internal/app"
	"boilerplate-api/internal/common/context"
	"boilerplate-api/pkg/config"
	"boilerplate-api/pkg/database"
	"boilerplate-api/pkg/logger"
	"boilerplate-api/pkg/redis"
	"flag"
	"fmt"
	"os"

	logs "github.com/rs/zerolog/log"
)

var (
	confPath = flag.String("config", "./config/app.development.yaml", "config file path")
	CommitID string //git commit id 	 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w  -X 'main.CommitID=$(git rev-parse HEAD)'" -o $binaryName
)

func main() {
	// Load configure
	flag.Parse()

	logs.Info().Str("config_file", *confPath).Msg("loading configuration")
	conf, err := config.LoadConfig(*confPath)
	if err != nil {
		logs.Fatal().Str("config_file", *confPath).Err(err).Msg("failed to load config")
	}
	logs.Info().Str("config_file", *confPath).Msg("configuration loaded successfully")

	logger.InitializeLogger(&conf.LogConf)

	// Init database
	db, err := database.InitMySQL(conf.MySQL)
	if err != nil {
		//panic(fmt.Errorf("failed to connect to mySQL: %w", err))
		logs.Fatal().Err(err).Msg("failed to connect to MYSQL")
	}
	logs.Info().Msg("MYSQL database connected successfully")

	dbMongo, err := database.InitMongoDB(conf.MongoDB)
	if err != nil {
		//panic(fmt.Errorf("failed to connect to mongoDB: %w", err))
		logs.Fatal().Err(err).Msg("failed to connect to MongoDB")
	}
	logs.Info().Msg("MongoDB database connected successfully")

	redisClient, err := redis.InitRedis(conf.Redis)
	if err != nil {
		logs.Fatal().Err(err).Msg("failed to connect to redis")
		//panic(fmt.Errorf("failed to connect to redis: %w", err))
	}
	logs.Info().Msg("Redis database connected successfully")

	logs.Info().Msg("Minio database connected successfully")

	logs.Info().Str("port", fmt.Sprintf("%d", conf.Server.Port)).Str("commit_id", CommitID).Int("pid", os.Getpid()).Msg("boilerplate-api START")
	appContext := &context.AppContext{
		Cfg:     conf,
		DB:      db,
		DBMongo: dbMongo,
		Redis:   redisClient,
	}

	server := app.NewServer(appContext)
	server.Run()
}
