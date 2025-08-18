package storage

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

// ErrNotFound is returned when the key is not found in Redis.
var ErrNotFound = errors.New("short key not found")

// RedisClient is a concrete implementation of the Storage interface using Redis.
type RedisClient struct {
	client *redis.Client
}

// NewRedisClient creates and returns a new RedisClient.
func NewRedisClient(addr string, password string, db int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisClient{client: rdb}
}

// Get retrieves the long URL from Redis.
func (r *RedisClient) Get(ctx context.Context, shortKey string) (string, error) {
	longURL, err := r.client.Get(ctx, shortKey).Result()
	if err == redis.Nil {
		return "", ErrNotFound
	} else if err != nil {
		panic(err)
	}
	return longURL, nil
}
