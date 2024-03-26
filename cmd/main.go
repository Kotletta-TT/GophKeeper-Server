package main

import (
	"GophKeeper-Server/internal/app"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	cfg, log := Initialize()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error { return app.Run(ctx, cfg, log) })
	fmt.Println(g.Wait())
}

func Initialize()
