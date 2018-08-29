package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Waiting for Goroutings!")
	var waitGroup sync.WaitGroup
	waitGroup.Add(10)
	var i int
	for i = 0; i < 10; i++ {
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}
