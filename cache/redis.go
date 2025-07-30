package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
)

func StartRedis() *redis.Client {

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	defer redisClient.Close()

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
