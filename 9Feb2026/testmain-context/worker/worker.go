package worker

import (
	"context"
	"time"
)

func StartWorker(ctx context.Context, result chan<- string) {
	go func() {
		select {
		case <-time.After(2 * time.Second):
			result <- "Work completed!"
		case <-ctx.Done():
			result <- "Work cancelled!"
		}
	}()
}
