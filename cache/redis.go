package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
)

func StartRedis() *redis.Client {

	addr := viper.GetString("REDIS_HOST")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return redisClient
}

func CloseRedis() {
	if err := redisClient.Close(); err != nil {
		log.Printf("Error closing Redis: %v", err)
	}
}
