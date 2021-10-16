package mysql

import (
	"time"
)

type Config struct {
	Addr         string        `mapstructure:"addr"`
	WriteDSN     string        `mapstructure:"writeDSN"`
	ReadDSN      []string      `mapstructure:"readDSN"`
	Active       int           `mapstructure:"active"`
	Idle         int           `mapstructure:"idle"`
	IdleTimeout  time.Duration `mapstructure:"idleTimeout"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
	TranTimeout  time.Duration `mapstructure:"tranTimeout"`
}
