package types

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/zishang520/engine.io/v2/events"
)

type RedisClient struct {
	events.EventEmitter
	*redis.Client

	Context context.Context
}

func NewRedisClient(ctx context.Context, redis *redis.Client) *RedisClient {
	return &RedisClient{
		EventEmitter: events.New(),
		Client:       redis,
		Context:      ctx,
	}
}
