package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// ErrNotFound is returned when the key is not found in Redis.
var ErrNotFound = errors.New("short key not found")

// Ensure RedisClient implicitly implements Storage.
// This is a compile-time check to ensure the contract is fulfilled.
var _ Storage = (*RedisClient)(nil)

// RedisClient is a concrete implementation of the Storage interface using Redis.
type RedisClient struct {
	client *redis.Client
}

// NewRedisClient creates and returns a new RedisClient, or an error if the connection fails.
// It also accepts a context for handling timeouts and cancellations during initialization.
func NewRedisClient(ctx context.Context, addr string, password string, db int) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Use Ping to verify the connection. This is a crucial step.
	// We use the provided context to ensure the ping doesn't hang indefinitely.
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RedisClient{client: rdb}, nil
}

// Get retrieves the long URL from Redis.
func (r *RedisClient) Get(ctx context.Context, shortKey string) (string, error) {
	longURL, err := r.client.Get(ctx, shortKey).Result()
	if err == redis.Nil {
		return "", ErrNotFound
	} else if err != nil {
		return "", fmt.Errorf("failed to get key from Redis: %w", err)
	}
	return longURL, nil
}
