package dao

import (
	"time"
)

type Config struct {
	ServerAddr   string
	WriteDSN     string
	ReadDSN      []string
	Active       int
	Idle         int
	IdleTimeout  time.Duration
	QueryTimeout time.Duration
	ExecTimeout  time.Duration
	TranTimeout  time.Duration
}

type Dao struct {
}

func New(cfg Config) *Dao {
	return &Dao{}
}
