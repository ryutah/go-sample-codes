package main

import (
	"context"
	"fmt"
	"time"
)

func infiniteLoop(ctx context.Context) {
	for {
		fmt.Println("Help")
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(2*time.Second))
	defer cancel()

	cancel()

	go infiniteLoop(ctx)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
