package service

import (
	"fmt"

	"gserver/app/admin/internel/config"
	"gserver/app/admin/internel/dao"
	"gserver/app/admin/internel/model"
	"gserver/app/admin/internel/querypath"
	"gserver/pkg/cache/redis"
	"gserver/pkg/databse/mysql"
)

type UserSvc struct {
	dao   *dao.UserDAO
	redis *redis.Redis
	db    *mysql.DB
}

func NewUserSvc(c *config.Config) *UserSvc {
	return &UserSvc{
		dao:   dao.NewUserDAO(c),
		redis: redis.New(c.Redis),
		db:    mysql.New(c.MySQL),
	}
}

const userInfoKey = "user:info:%d"

func (u *UserSvc) Info(id int) (*model.User, error) {
	return u.dao.First(querypath.NewUser(u.db.Master()).
		WithCache(u.redis, fmt.Sprintf(userInfoKey, id), 0).
		WhIdEq(id))
}
