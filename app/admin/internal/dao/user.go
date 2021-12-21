package dao

import (
	"gserver/app/admin/internal/config"
	"gserver/app/admin/internal/model"
	"gserver/app/admin/internal/querypath"
)

type UserDAO struct {
	c *config.Config
}

func NewUserDAO(c *config.Config) *UserDAO {
	return &UserDAO{
		c: c,
	}
}

func (u *UserDAO) Create(q *querypath.User, user *model.User) error {
	return q.Create(user)
}

func (u *UserDAO) Delete(q *querypath.User) error {
	return q.Delete(&model.User{})
}

func (u *UserDAO) Update(q *querypath.User, user *model.User) error {
	return q.Update(user) // OR q.Save(user)
}

func (u *UserDAO) First(q *querypath.User) (*model.User, error) {
	user := &model.User{}
	if err := q.First(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserDAO) Find(q *querypath.User) ([]*model.User, error) {
	users := make(model.UserSlice, 0)
	if err := q.Find(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserDAO) Count(q *querypath.User) (int64, error) {
	return q.Count()
}
