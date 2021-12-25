package dao

import (
	"gorm.io/gorm"

	"gserver/app/admin/internal/config"
	"gserver/app/admin/internal/model"
	"gserver/app/admin/internal/querypath"
	"gserver/pkg/cache/redis"
	"gserver/pkg/databse/mysql"
)

type UserDao struct {
	c     config.Config
	db    *mysql.DB
	redis *redis.Redis
}

func NewUserDao(c config.Config) *UserDao {
	return &UserDao{
		c:     c,
		db:    mysql.New(c.MySQL),
		redis: redis.New(c.Redis),
	}
}

func (u *UserDao) Transaction(fc func(tx *gorm.DB) error) error {
	return u.db.Master().Transaction(fc)
}

// 增

func (u *UserDao) Create(v interface{}) error {
	return querypath.NewUser(u.db.Master()).Create(v)
}

func (u *UserDao) TxCreate(tx *gorm.DB, v interface{}) error {
	return querypath.NewUser(tx).Create(v)
}

// 删

func (u *UserDao) DeleteById(id int) error {
	return querypath.NewUser(u.db.Master()).WhIdEq(id).Delete(&model.User{})
}

func (u *UserDao) TxDeleteById(tx *gorm.DB, id int) error {
	return querypath.NewUser(tx).WhIdEq(id).Delete(&model.User{})
}

// 改

func (u *UserDao) UpdateById(id int, v interface{}) error {
	return querypath.NewUser(u.db.Master()).WhIdEq(id).Update(v)
}

func (u *UserDao) TxUpdateById(tx *gorm.DB, id int, v interface{}) error {
	return querypath.NewUser(tx).WhIdEq(id).Update(v)
}

// 查

func (u *UserDao) FirstById(id int) (*model.User, error) {
	v := &model.User{}
	if err := querypath.NewUser(u.db.NextReader()).WhIdEq(id).First(v); err != nil {
		return nil, err
	}

	return v, nil
}

func (u *UserDao) TxFirstById(tx *gorm.DB, id int) (*model.User, error) {
	v := &model.User{}
	if err := querypath.NewUser(tx).WhIdEq(id).First(v); err != nil {
		return nil, err
	}

	return v, nil
}
