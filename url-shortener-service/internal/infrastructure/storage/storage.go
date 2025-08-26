package storage

import "context"

type Storage interface {
	Post(ctx context.Context, longKey string) (string, error)
}
