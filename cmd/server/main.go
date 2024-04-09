package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	config "github.com/Kotletta-TT/GophKeeper/config/server"
	"github.com/Kotletta-TT/GophKeeper/internal/server/app"
	l "github.com/Kotletta-TT/GophKeeper/logger"
	"golang.org/x/sync/errgroup"
)

func main() {
	cfg, logger, err := Initialize()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error { return app.Run(ctx, cfg, logger) })
	logger.Error(g.Wait().Error())
}

func Initialize() (*config.Config, l.Logger, error) {
	cfgPath, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		cfgPath = "config.yaml"
	}
	flag.StringVar(&cfgPath, "c", "config.yaml", "path to config file")
	flag.Parse()
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		return nil, nil, err
	}
	log, err := l.NewLogger(cfg)
	if err != nil {
		return nil, nil, err
	}
	return cfg, log, nil
}