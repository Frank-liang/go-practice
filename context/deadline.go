package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	time.AfterFunc(time.Second*3, cancel)
	done := ctx.Done()
	for i := 0; ; i++ {
		select {
		case <-done:
			fmt.Println("exit", ctx.Err())
			return
		case <-time.After(time.Second):
			fmt.Println("tick", i)
		}
	}
}
