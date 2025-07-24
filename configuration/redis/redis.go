package redis

import (
	"context"
	app "go_fiber_core_project_api/configuration/app"
	custom_logger "go_fiber_core_project_api/pkg/utils/loggers"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
	RedisExpire   int
}

var (
	once   sync.Once
	client *redis.Client
)

func InitRedis() *RedisConfig {
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found, using system environment variables")
	}
	redis_host := os.Getenv("REDIS_HOST")
	redis_port := os.Getenv("REDIS_PORT")
	redis_password := os.Getenv("REDIS_PASSWORD")
	redis_db := app.GetenvInt("REDIS_DB", 0)
	redis_exprie := app.GetenvInt("REDIS_EXPIRE", 60)
	return &RedisConfig{
		RedisHost:     redis_host,
		RedisPort:     redis_port,
		RedisPassword: redis_password,
		RedisDB:       redis_db,
		RedisExpire:   redis_exprie,
	}
}

func NewRedisClient() *redis.Client {
	redis_config := InitRedis()
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     redis_config.RedisHost + ":" + redis_config.RedisPort,
			Password: redis_config.RedisPassword,
			DB:       redis_config.RedisDB,
		})
		pong, err := client.Ping(context.Background()).Result()
		if err != nil {
			custom_logger.NewCustomLog("connect_redis_failed", err.Error(), "error")
			log.Fatalf("Could not connect to Redis: %v", err)
		}
		log.Printf("Connected to Redis successfully: %s", pong)
	})
	return client
}
