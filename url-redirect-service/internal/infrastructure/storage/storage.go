package storage

import "context"

type Storage interface {
	Get(ctx context.Context, shortKey string) (string, error)
}
