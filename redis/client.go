package redis

import (
	"strings"

	"github.com/go-redis/redis/v7"
)

type RedisC struct {
	*redis.Client
}

type RedisOptions struct {
	MasterName    string
	MasterPw      string
	SentinelAddrs string
	SentinelPw    string
}

func Init(
	addr string,
	options *RedisOptions,
) *RedisC {
	c := RedisC{}

	if options != nil {
		c.Client = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:       options.MasterName,
			Password:         options.MasterPw,
			SentinelAddrs:    strings.Split(options.SentinelAddrs, ","),
			SentinelPassword: options.SentinelPw,
		})
	} else {
		c.Client = redis.NewClient(&redis.Options{
			Addr: addr,
		})
	}

	return &c
}
