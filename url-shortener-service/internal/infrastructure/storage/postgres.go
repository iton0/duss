package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/iton0/duss/shared/domain"
)

// Ensure PostgresClient implicitly implements Storage.
var _ Storage = (*PostgresClient)(nil)

// PostgresClient is a concrete implementation of the Storage interface using PostgreSQL.
type PostgresClient struct {
	pool *pgxpool.Pool
}

// NewPostgresClient creates and returns a new PostgresClient.
func NewPostgresClient(ctx context.Context, dsn string) (*PostgresClient, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DSN: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresClient{pool: pool}, nil
}

// Save persists a domain.URL entity to the PostgreSQL database.
func (p *PostgresClient) Save(ctx context.Context, url *domain.URL) error {
	query := `
		INSERT INTO urls (short_key, long_url, created_at, redirects)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (short_key) DO NOTHING
	`
	_, err := p.pool.Exec(ctx, query, url.ShortKey, url.LongURL, url.CreatedAt, url.Redirects)
	if err != nil {
		return fmt.Errorf("failed to save URL: %w", err)
	}

	return nil
}

// IncrementRedirects increments the redirects count for a given short key.
func (p *PostgresClient) IncrementRedirects(ctx context.Context, shortKey string) error {
	query := `UPDATE urls SET redirects = redirects + 1 WHERE short_key = $1`
	_, err := p.pool.Exec(ctx, query, shortKey)
	if err != nil {
		return fmt.Errorf("failed to increment redirects: %w", err)
	}
	return nil
}
