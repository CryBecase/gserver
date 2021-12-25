package router

import (
	v1 "gserver/app/admin/internal/api/v1"
	"gserver/app/admin/internal/config"
	"gserver/pkg/net/http/gin"
)

type IUserAPI interface {
	Info(c *gin.Context)
}

func Init(e *gin.Engine, c config.Config) {
	var (
		u IUserAPI = v1.NewUser(c)
	)

	userGroup := e.Group("user")
	userGroup.GET("info", u.Info)
}
