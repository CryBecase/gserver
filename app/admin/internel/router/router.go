package router

import (
	v1 "gserver/app/admin/internel/api/v1"
	"gserver/app/admin/internel/config"
	"gserver/pkg/net/http/gin"
)

func Init(e *gin.Engine, c *config.Config) {
	u := v1.NewUser(c)

	userGroup := e.Group("user")
	userGroup.GET("info", u.Info)
}
