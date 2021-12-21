package dao

import (
	"gserver/app/admin/internal/config"
	"gserver/app/admin/internal/model"
	"gserver/app/admin/internal/querypath"
)

type userDAO struct {
	c *config.Config
}

func NewUserDAO(c *config.Config) *userDAO {
	return &userDAO{
		c: c,
	}
}

func (u *userDAO) Create(q *querypath.User, user *model.User) error {
	return q.Create(user)
}

func (u *userDAO) Delete(q *querypath.User) error {
	return q.Delete(&model.User{})
}

func (u *userDAO) Update(q *querypath.User, user *model.User) error {
	return q.Update(user) // OR q.Save(user)
}

func (u *userDAO) First(q *querypath.User) (*model.User, error) {
	user := &model.User{}
	if err := q.First(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userDAO) Find(q *querypath.User) (model.UserSlice, error) {
	users := make(model.UserSlice, 0)
	if err := q.Find(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userDAO) Count(q *querypath.User) (int64, error) {
	return q.Count()
}
