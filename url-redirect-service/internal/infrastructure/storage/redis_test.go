package storage_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/iton0/duss/url-redirect-service/internal/infrastructure/storage"
	"github.com/redis/go-redis/v9"
)

const redisAddr = "localhost:6379"

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rdb := redis.NewClient(&redis.Options{Addr: redisAddr})
	defer rdb.Close()

	if err := rdb.Ping(ctx).Err(); err != nil {
		os.Stderr.WriteString("Skipping Redis integration tests: could not connect to Redis at " + redisAddr + "\n")
		os.Exit(0)
	}

	code := m.Run()
	os.Exit(code)
}

func setupTest(t *testing.T) *storage.RedisClient {
	ctx := context.Background()

	client, err := storage.NewRedisClient(ctx, redisAddr, "", 0)
	if err != nil {
		t.Fatalf("setup failed: could not create Redis client: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{Addr: redisAddr})
	defer rdb.Close()

	if err := rdb.FlushAll(ctx).Err(); err != nil {
		t.Fatalf("setup failed: could not flush Redis DB: %v", err)
	}

	return client
}

func TestNewRedisClient(t *testing.T) {
	t.Run("Valid Connection", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		client, err := storage.NewRedisClient(ctx, redisAddr, "", 0)
		if err != nil {
			t.Fatalf("expected no error, but got: %v", err)
		}
		if client == nil {
			t.Fatal("expected a valid client, but got nil")
		}
	})

	t.Run("Invalid Connection", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		client, err := storage.NewRedisClient(ctx, "localhost:6380", "", 0)
		if err == nil {
			t.Fatal("expected an error, but got nil")
		}
		if client != nil {
			t.Fatal("expected a nil client, but got a value")
		}
	})
}

func TestGet(t *testing.T) {
	client := setupTest(t)
	// No client.client.Close() here, as we can't access private fields.

	t.Run("Success - Key Exists", func(t *testing.T) {
		ctx := context.Background()
		shortKey := "test_key"
		longURL := "http://example.com/test_url"

		rdb := redis.NewClient(&redis.Options{Addr: redisAddr})
		defer rdb.Close()

		err := rdb.Set(ctx, shortKey, longURL, 0).Err()
		if err != nil {
			t.Fatalf("could not set key for test: %v", err)
		}

		retrievedURL, err := client.Get(ctx, shortKey)
		if err != nil {
			t.Fatalf("expected no error, but got: %v", err)
		}
		if retrievedURL != longURL {
			t.Fatalf("expected URL %s, but got %s", longURL, retrievedURL)
		}
	})

	t.Run("Not Found - Key Does Not Exist", func(t *testing.T) {
		ctx := context.Background()
		shortKey := "non_existent_key"

		retrievedURL, err := client.Get(ctx, shortKey)
		if !errors.Is(err, storage.ErrNotFound) {
			t.Fatalf("expected ErrNotFound, but got: %v", err)
		}
		if retrievedURL != "" {
			t.Fatalf("expected empty URL, but got: %s", retrievedURL)
		}
	})
}
