package main

import (
	"context"
	"fmt"
	"time"
)

func infiniteLoop(ctx context.Context) {
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("ENDDDDDDDDDDDDDDDDDDDDDD")
			return
		}
	}
	// 	for {
	// 		fmt.Println("Help!!")
	// 	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	defer cancel()

	go infiniteLoop(ctx)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
