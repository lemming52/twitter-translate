package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/kelseyhightower/envconfig"
	"github.com/lemming52/twitter-translate/clients/translator"
	"github.com/lemming52/twitter-translate/storage/cache"
	"github.com/lemming52/twitter-translate/translate"
	"github.com/lemming52/twitter-translate/web"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Host    string `required:"true" default:"0.0.0.0"`
	Port    int    `required:"true" default:"5000"`
	Service string `default:"twitter-translate"`
	Version string `default:"SNAPSHOT"`
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	var cfg Config
	envconfig.Process("twitter-translate", &cfg)

	translator := translator.NewTranslator()
	cache := cache.NewCache()

	p := translate.NewService(translator, cache)

	webCfg := &web.Config{
		Host:    cfg.Host,
		Port:    cfg.Port,
		Service: cfg.Service,
		Version: cfg.Version,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	err := web.ListenAndServe(ctx, webCfg, p)
	if err != nil {
		log.Fatal(err)
	}
}
