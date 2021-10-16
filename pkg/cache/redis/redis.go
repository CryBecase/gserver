package redis

type Redis struct {
	c *Config
}

func New(c *Config) *Redis {
	return &Redis{
		c: c,
	}
}
