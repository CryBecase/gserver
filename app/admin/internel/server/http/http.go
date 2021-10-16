package http

import (
	"gserver/app/admin/internel/config"
	"gserver/pkg/net/http/gin"
)

func Init(c *config.Config) {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	e.Use(gin.Logger(), gin.Recovery()) // TODO 用自己的代替 或者 重写 gin
	go func() {
		if err := e.Run(c.Server.Addr); err != nil {
			panic(err)
		}
	}()
}
