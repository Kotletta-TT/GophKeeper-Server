package app

import (
	"GophKeeper-Server/config"
	"GophKeeper-Server/logger"
	"context"
	"time"
)

func Run(ctx context.Context, cfg *config.Config, l logger.Logger) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			l.Info("Hello, World!")
			time.Sleep(time.Second * 1)
		}
	}
}
