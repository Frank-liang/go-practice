package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Charter 09 - Gotoutines.")

	for i := 0; i < 10; i++ {
		go func(x int) {
			time.Sleep(10)
			fmt.Printf("%d ", x)
		}(i)
	}
	time.Sleep(10000)
	fmt.Println("Exiting...")
}
