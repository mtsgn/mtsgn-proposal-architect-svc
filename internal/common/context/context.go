package context

import (
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"

	"boilerplate-api/pkg/config"
	"boilerplate-api/pkg/redis"
)

type AppContext struct {
	Cfg     *config.Config
	DB      *gorm.DB
	DBMongo *mongo.Database
	Redis   *redis.RedisClient
	Minio   *minio.Client
}
