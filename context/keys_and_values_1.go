package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type key *int

func main() {
	k := new(key)
	ctx, cancle := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(ctx context.Context) {
			v := ctx.Value(k)
			fmt.Println("key", ctx.Value(new(key)))
			wg.Done()
			<-ctx.Done()
			fmt.Println(ctx.Err(), v)
		}(context.WithValue(ctx, k, i))
	}
	wg.Wait()
	cancle()
	time.Sleep(time.Second)
}
