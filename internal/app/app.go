package app

import (
	"GophKeeper-Server/config"
	"context"
	"fmt"
	"time"
)

func Run(ctx context.Context, cfg *config.Config, l logger.Logger) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Println("hello")
			time.Sleep(time.Second * 1)
		}
	}
}
