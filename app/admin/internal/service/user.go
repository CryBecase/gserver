package service

import (
	"gserver/app/admin/internal/config"
	"gserver/app/admin/internal/dao"
	"gserver/app/admin/internal/model"
)

type userSvc struct {
	dao *dao.UserDao
}

func NewUserSvc(c config.Config) *userSvc {
	return &userSvc{
		dao: dao.NewUserDao(c),
	}
}

func (u *userSvc) Info(id int) (*model.User, error) {
	return u.dao.FirstById(id)
}
