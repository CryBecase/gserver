package server

import (
	"gserver/app/admin/internal/config"
	"gserver/app/admin/internal/router"
	"gserver/pkg/net/http/gin"
)

func InitHTTP(c *config.Config) {
	e := gin.New()
	e.Use(gin.Logger(), gin.Recovery()) // TODO 用自己的代替 或者 重写 gin

	router.Init(e, c)

	go func() {
		if err := e.Run(c.Server.Addr); err != nil {
			panic(err)
		}
	}()
}
