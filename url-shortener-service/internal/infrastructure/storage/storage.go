package storage

import (
	"context"

	"github.com/iton0/duss/shared/domain"
)

type Storage interface {
	Save(ctx context.Context, url *domain.URL) error
}
