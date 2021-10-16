package memcache

import (
	"time"
)

type Config struct {
	Name         string        `mapstructure:"name"`
	Proto        string        `mapstructure:"proto"`
	Addr         string        `mapstructure:"addr"`
	Active       int           `mapstructure:"active"`
	Idle         int           `mapstructure:"idle"`
	IdleTimeout  time.Duration `mapstructure:"idleTimeout"`
	DialTimeout  time.Duration `mapstructure:"dialTimeout"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
	Expire       time.Duration `mapstructure:"expire"`
}
