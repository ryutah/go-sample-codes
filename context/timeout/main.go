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

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	go infiniteLoop(ctx)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
