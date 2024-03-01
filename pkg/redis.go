package pkg

import (
	"context"
	"os"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis() *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})
	return &Redis{
		Client: rdb,
	}
}

func (r *Redis) APIRateLimit() (*redis_rate.Result, error) {
	limiter := redis_rate.NewLimiter(r.Client)

	return limiter.Allow(
		context.Background(),
		"apiRateLimit",
		redis_rate.PerSecond(1),
	)
}
