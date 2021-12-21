package querypath

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"gserver/pkg/cache/redis"
)

type Paginate struct {
	Page  int   `json:"page"`
	Size  int   `json:"size"`
	Total int64 `json:"total"`
}

func (p *Paginate) OffSet() int {
	return (p.Page - 1) * p.Size
}

func (p *Paginate) Fix() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Size <= 0 {
		p.Size = 20
	}
}

type redisWrapper struct {
	*redis.Redis
	key    string
	expire time.Duration
}

type querypath struct {
	db    *gorm.DB
	redis *redisWrapper
	p     *Paginate
}

func (q *querypath) Create(value interface{}) error {
	if err := q.db.Create(value).Error; err != nil {
		return err
	}
	if q.redis != nil {
		return q.redis.Del(context.TODO(), q.redis.key).Err()
	}
	return nil
}

func (q *querypath) Delete(value interface{}) error {
	if err := q.db.Delete(value).Error; err != nil {
		return err
	}
	if q.redis != nil {
		return q.redis.Del(context.TODO(), q.redis.key).Err()
	}
	return nil
}

func (q *querypath) Update(value interface{}) error {
	if err := q.db.Updates(value).Error; err != nil {
		return err
	}
	if q.redis != nil {
		return q.redis.Del(context.TODO(), q.redis.key).Err()
	}
	return nil
}

func (q *querypath) Save(value interface{}) error {
	if err := q.db.Save(value).Error; err != nil {
		return err
	}
	if q.redis != nil {
		return q.redis.Del(context.TODO(), q.redis.key).Err()
	}
	return nil
}

func (q *querypath) First(dest interface{}) error {
	if q.redis != nil {
		data, err := q.redis.Get(context.Background(), q.redis.key).Bytes()
		if err != nil && err != redis.Nil {
			return err
		}
		if err == nil && len(data) > 0 {
			err = json.Unmarshal(data, dest)
			if err != nil {
				return err
			}
			return nil
		}
	}

	if err := q.db.First(dest).Error; err != nil {
		return err
	}

	if q.redis != nil {
		vjson, err := json.Marshal(dest)
		if err != nil {
			return err
		}

		if err = q.redis.Set(context.Background(), q.redis.key, vjson, q.redis.expire).Err(); err != nil {
			return err
		}
	}

	return nil
}

func (q *querypath) Find(dest interface{}) error {
	if q.p != nil {
		// 分页不走缓存
		return q.db.Limit(q.p.Size).Offset(q.p.OffSet()).Find(dest).Error
	}

	if q.redis != nil {
		data, err := q.redis.Get(context.Background(), q.redis.key).Bytes()
		if err != nil && err != redis.Nil {
			return err
		}
		if err == nil && len(data) > 0 {
			err = json.Unmarshal(data, dest)
			if err != nil {
				return err
			}
			return nil
		}
	}

	err := q.db.Find(dest).Error
	if err != nil {
		return err
	}

	if q.redis != nil {
		vjson, err := json.Marshal(dest)
		if err != nil {
			return err
		}

		if err = q.redis.Set(context.Background(), q.redis.key, vjson, q.redis.expire).Err(); err != nil {
			return err
		}
	}
	return nil
}

func (q *querypath) Count() (int64, error) {
	var cnt int64
	if err := q.db.Count(&cnt).Error; err != nil {
		return 0, err
	}
	return cnt, nil
}

func (q *querypath) Paginate(dest interface{}) (*Paginate, error) {
	if q.p == nil {
		return nil, errors.New("no paginate")
	}
	q.p.Fix()

	cnt, err := q.Count()
	if err != nil {
		return nil, err
	}
	q.p.Total = cnt

	return q.p, q.Find(dest)
}

func (q *querypath) ForUpdate() *querypath {
	q.db.Clauses(clause.Locking{Strength: "UPDATE"})
	return q
}

func (q *querypath) RawDB() *gorm.DB {
	return q.db
}
