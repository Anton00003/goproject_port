package main

import (
	"context"
	"fmt"
	"goproject_port/api"
	"goproject_port/config"
	"goproject_port/repository"
	"goproject_port/service"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	cfg, err := config.New()
	if err != nil {
		fmt.Println(err, "config init failed")
		return
	}
	lvl, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug(err, "parse level failed")
	}
	logrus.SetLevel(lvl)
	repo := repository.New(cfg.NumIn, cfg.NumOut)
	fmt.Println(*repo)
	serv := service.New(repo)
	api := api.New(serv, cfg.NumIn, cfg.NumOut, ctx)

	api.Run(cfg.Host)
	<-ctx.Done()
}
