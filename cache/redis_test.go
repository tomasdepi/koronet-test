package cache

import (
	"context"
	"testing"

	"github.com/spf13/viper"
)

func initTestViper() {
	viper.Reset() // reset to not override other tests
	viper.AutomaticEnv()
	viper.SetDefault("REDIS_HOST", "localhost:6379")
}

func TestRedisConnection(t *testing.T) {

	initTestViper()

	rdb := StartRedis()
	defer rdb.Close()

	ctx := context.Background()

	if err := rdb.Ping(ctx).Err(); err != nil {
		t.Fatalf("Failed to ping Redis: %v", err)
	}

	err := rdb.Set(ctx, "test-key", "test-value", 0).Err()
	if err != nil {
		t.Fatalf("Failed to set key: %v", err)
	}

	val, err := rdb.Get(ctx, "test-key").Result()
	if err != nil {
		t.Fatalf("Failed to get key: %v", err)
	}

	if val != "test-value" {
		t.Fatalf("Unexpected value: got %q, want %q", val, "test-value")
	}
}
