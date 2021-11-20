package redis

import (
	"github.com/go-redis/redis/v8"
)

const Nil = redis.Nil

type Redis struct {
	*redis.Client

	cfg *Config
}

func New(c *Config) *Redis {
	return &Redis{
		cfg: c,
	}
}
