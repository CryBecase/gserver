package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(c *Config) *DB {
	w, err := open(c, c.WriteDSN)
	if err != nil {
		panic(err)
	}

	rs := make([]*gorm.DB, 0, len(c.ReadDSN))
	for _, rd := range c.ReadDSN {
		r, err := open(c, rd)
		if err != nil {
			panic(err)
		}
		rs = append(rs, r)
	}
	if len(rs) == 0 {
		rs = append(rs, w)
	}

	db := &DB{}
	db.write = w
	db.read = rs
	return db
}

func open(c *Config, dsn string) (*gorm.DB, error) {
	r, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	db, err := r.DB()
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(c.Active)
	db.SetMaxIdleConns(c.Idle)
	db.SetConnMaxIdleTime(c.IdleTimeout)
	return r, nil
}
