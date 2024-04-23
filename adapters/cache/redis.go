package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisAdapter struct {
	client *redis.Client
}

func NewRedisAdapter() *RedisAdapter {
	return &RedisAdapter{}
}

func (c *RedisAdapter) Connect(connectionString string) error {
	options, err := redis.ParseURL(connectionString)
	if err != nil {
		return err
	}

	c.client = redis.NewClient(options)
	return nil
}

func (c *RedisAdapter) Close() error {
	if c.client != nil {
		return c.client.Close()
	}
	return nil
}

func (c *RedisAdapter) Set(ctx context.Context, key string, val string, ttl int) error {
	cmd := c.client.SetNX(ctx, key, val, time.Hour * time.Duration(ttl))
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (c *RedisAdapter) Get(ctx context.Context, key string) (string, error) {
	cmd := c.client.Get(ctx, key)
	if cmd.Err() != nil {
		return "", cmd.Err()
	}

	return cmd.Val(), nil
}

func (c *RedisAdapter) Delete(ctx context.Context, key string) error {
	cmd := c.client.Del(ctx, key)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}
