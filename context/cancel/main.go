package main

import (
	"context"
	"fmt"
	"time"
)

func infiniteLoop(ctx context.Context) {
	for {
		fmt.Println("Help!")
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go infiniteLoop(ctx)

	time.Sleep(2 * time.Second)
	cancel()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
