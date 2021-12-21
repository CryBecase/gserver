package service

import (
	"fmt"

	"gserver/app/admin/internal/config"
	"gserver/app/admin/internal/dao"
	"gserver/app/admin/internal/model"
	"gserver/app/admin/internal/querypath"
	"gserver/pkg/cache/redis"
	"gserver/pkg/databse/mysql"
)

type IUserDAO interface {
	Create(q *querypath.User, user *model.User) error
	Delete(q *querypath.User) error
	Update(q *querypath.User, user *model.User) error
	First(q *querypath.User) (*model.User, error)
	Find(q *querypath.User) (model.UserSlice, error)
	Count(q *querypath.User) (int64, error)
}

type userSvc struct {
	dao   IUserDAO
	redis *redis.Redis
	db    *mysql.DB
}

func NewUserSvc(c *config.Config) *userSvc {
	return &userSvc{
		dao:   dao.NewUserDAO(c),
		redis: redis.New(c.Redis),
		db:    mysql.New(c.MySQL),
	}
}

const userInfoKey = "user:info:%d"

func (u *userSvc) Info(id int) (*model.User, error) {
	return u.dao.First(querypath.NewUser(u.db.Master()).
		WithCache(u.redis, fmt.Sprintf(userInfoKey, id), 0).
		WhIdEq(id))
}
