package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gserver/app/admin/internal/config"
	"gserver/app/admin/internal/server"

	_ "gserver/pkg/encoding/form"
	_ "gserver/pkg/encoding/json"
	_ "gserver/pkg/encoding/xml"
	_ "gserver/pkg/encoding/yaml"
)

func main() {
	flag.Parse()
	if err := config.Init(); err != nil {
		panic(err)
	}
	server.InitHTTP(config.Conf)

	closeFunc := func() {
		fmt.Printf("%s is closing ...\n", config.Conf.Server.Name)
	}

	fmt.Printf("%s START!\n", config.Conf.Server.Name)
	defer fmt.Printf("%s EXIT!\n", config.Conf.Server.Name)

	listenSignal()

	closeFunc()
}

func listenSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	fmt.Printf("%s got signal: %s\n", config.Conf.Server.Name, s.String())
}
