package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type key struct{}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(ctx context.Context) {
			v := ctx.Value(key{})
			fmt.Println("key", v)
			wg.Done()
			<-ctx.Done()
			fmt.Println(ctx.Err(), v)
		}(context.WithValue(ctx, key{}, i))
	}
	wg.Wait()
	cancel()
	time.Sleep(time.Second)

}
