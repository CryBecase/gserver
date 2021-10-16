package mysql

type MySQL struct {
	c *Config
}

func New(c *Config) *MySQL {
	return &MySQL{
		c: c,
	}
}
