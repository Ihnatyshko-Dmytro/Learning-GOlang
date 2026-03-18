package main

import (
	"context"
	"time"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second)

	parse(ctx)
}

func parse (ctx context.Context) {
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("Parsing completed")
		case <-ctx.Done():
			fmt.Println("deadline exceded")
			return

		}
	}
}