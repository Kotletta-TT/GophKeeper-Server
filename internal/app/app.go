package app

import (
	"context"
	"fmt"
	"time"
)

func Run(ctx context.Context) error {
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
