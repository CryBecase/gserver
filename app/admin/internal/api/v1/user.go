package v1

import (
	"gserver/app/admin/internal/config"
	"gserver/app/admin/internal/model"
	"gserver/app/admin/internal/service"
	"gserver/pkg/net/http/gin"
)

type IUserSvc interface {
	Info(id int) (*model.User, error)
}

type userAPI struct {
	svc IUserSvc
}

func NewUser(c config.Config) *userAPI {
	return &userAPI{
		svc: service.NewUserSvc(c),
	}
}

func (u *userAPI) Info(c *gin.Context) {
	type r struct {
		Id int `json:"id" form:"id" binding:"required"`
	}
	req := &r{}
	if err := c.ShouldBind(req); err != nil {
		_ = c.AbortWithError(400, err)
		return
	}

	userInfo, err := u.svc.Info(req.Id)
	if err != nil {
		_ = c.AbortWithError(500, err)
		return
	}

	c.JSON(200, userInfo) // TODO 封装 gin.Context
}
