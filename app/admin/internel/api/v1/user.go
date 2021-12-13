package v1

import (
	"gserver/app/admin/internel/config"
	"gserver/app/admin/internel/service"
	"gserver/pkg/net/http/gin"
)

type User struct {
	svc *service.UserSvc
}

func NewUser(c *config.Config) *User {
	return &User{
		svc: service.NewUserSvc(c),
	}
}

func (u *User) Info(c *gin.Context) {
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
