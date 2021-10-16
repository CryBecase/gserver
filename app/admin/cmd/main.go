package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gserver/app/admin/internel/config"
	"gserver/app/admin/internel/server/http"
)

func main() {
	flag.Parse()
	if err := config.Init(); err != nil {
		panic(err)
	}
	http.Init(config.Conf)

	closeFunc := func() {
		fmt.Printf("%s is closing ...\n", config.Conf.Server.Name)
	}

	fmt.Printf("%s START!\n", config.Conf.Server.Name)
	defer fmt.Printf("%s EXIT!\n", config.Conf.Server.Name)

	// signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	fmt.Printf("%s got signal: %s\n", config.Conf.Server.Name, s.String())
	closeFunc() // TODO 限制关闭时间
}
