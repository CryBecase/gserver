package query

import (
	"context"
	"encoding/json"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"gserver/pkg/cache/redis"
	"gserver/pkg/util"
)

type query struct {
	db       *gorm.DB
	cache    *cache
	paginate *util.Paginate
}

type cache struct {
	redis      *redis.Redis
	key        string
	expiration time.Duration
}

func (q *query) Find(dest interface{}) error {
	if q.paginate != nil {
		// 分页数据不走缓存
		return q.db.Find(dest).Error
	}

	if q.cache != nil {
		data, err := q.cache.redis.Get(context.TODO(), q.cache.key).Bytes()
		if err != nil && err != redis.Nil {
			// log 即可，缓存出问题就查库
		}
		if err == nil && len(data) > 0 {
			return json.Unmarshal(data, dest)
		}
	}

	if err := q.db.Find(dest).Error; err != nil {
		return err
	}

	data, err := json.Marshal(dest)
	if err != nil {
		return err
	}
	return q.cache.redis.Set(context.TODO(), q.cache.key, data, q.cache.expiration).Err()
}

func (q *query) First(dest interface{}) error {
	if q.cache != nil {
		data, err := q.cache.redis.Get(context.TODO(), q.cache.key).Bytes()
		if err != nil && err != redis.Nil {
			// log 即可，缓存出问题就查库
		}
		if err == nil && len(data) > 0 {
			return json.Unmarshal(data, dest)
		}
	}

	if err := q.db.First(dest).Error; err != nil {
		return err
	}

	data, err := json.Marshal(dest)
	if err != nil {
		return err
	}
	return q.cache.redis.Set(context.TODO(), q.cache.key, data, q.cache.expiration).Err()
}

func (q *query) ForUpdate() *query {
	q.db.Clauses(&clause.Locking{Strength: "UPDATE"})
	return q
}

func (q *query) Update(value interface{}) error {
	if err := q.db.Save(value).Error; err != nil {
		return err
	}
	if q.cache != nil {
		return q.cache.redis.Del(context.TODO(), q.cache.key).Err()
	}
	return nil
}

func (q *query) Insert(value interface{}) error {
	if err := q.db.Create(value).Error; err != nil {
		return err
	}
	if q.cache != nil {
		return q.cache.redis.Del(context.TODO(), q.cache.key).Err()
	}
	return nil
}

func (q *query) InsertIgnore() *query {
	q.db.Clauses(&clause.Insert{Modifier: "IGNORE"})
	return q
}

func (q *query) WithCache(redis *redis.Redis, key string, expiration time.Duration) *query {
	q.cache = &cache{
		redis:      redis,
		key:        key,
		expiration: expiration,
	}
	return q
}
