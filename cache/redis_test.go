package cache

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestRedisConnection(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	ctx := context.Background()

	if err := rdb.Ping(ctx).Err(); err != nil {
		t.Fatalf("Failed to ping Redis: %v", err)
	}

	err := rdb.Set(ctx, "test-key", "test-value", 0).Err()
	if err != nil {
		t.Fatalf("Failed to set key: %v", err)
	}

	// Try getting the key
	val, err := rdb.Get(ctx, "test-key").Result()
	if err != nil {
		t.Fatalf("Failed to get key: %v", err)
	}

	if val != "test-value" {
		t.Fatalf("Unexpected value: got %q, want %q", val, "test-value")
	}
}
