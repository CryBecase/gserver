package mysql

import (
	"sync/atomic"

	"gorm.io/gorm"
)

// DB database.
type DB struct {
	write *gorm.DB
	read  []*gorm.DB
	idx   int64
}

func (d *DB) NextReader() *gorm.DB {
	v := atomic.AddInt64(&d.idx, 1)
	return d.read[int(v)%len(d.read)]
}

func (d *DB) Master() *gorm.DB {
	return d.write
}
