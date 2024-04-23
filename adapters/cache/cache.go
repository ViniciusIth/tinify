package cache

import (
	"context"
)

type Cache interface {
	Connect(connectionString string) error
	Close() error
	Set(ctx context.Context, key string, val string, ttl int) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}
