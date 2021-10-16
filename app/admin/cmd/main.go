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
		fmt.Printf("%s try to close...\n", config.Conf.Server.Name)
	}

	// signal handler
	fmt.Printf("%s start\n", config.Conf.Server.Name)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	fmt.Printf("%s get a signal %s\n", config.Conf.Server.Name, s.String())
	closeFunc() // TODO 限制关闭时间
	fmt.Printf("%s exit\n", config.Conf.Server.Name)
}
