package memcache

type Memcache struct {
	c *Config
}

func New(c *Config) *Memcache {
	return &Memcache{
		c: c,
	}
}
