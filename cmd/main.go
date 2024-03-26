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
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error { return app.Run(ctx) })
	fmt.Println(g.Wait())
}
