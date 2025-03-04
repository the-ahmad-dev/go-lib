package ratelimit

import (
	"github.com/the-ahmad-dev/go-lib/redis"
)

type RateLimitC struct {
	Redis *redis.RedisC
}

func Init(
	redisC *redis.RedisC,
) *RateLimitC {
	return &RateLimitC{
		Redis: redisC,
	}
}
