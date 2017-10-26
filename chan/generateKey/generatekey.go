package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateKey(channel chan int) {
	fmt.Println("Generating key")
	keys := []int{3, 5, 7, 11}
	key := keys[rand.Intn(len(keys))]
	time.Sleep(3 * time.Second)
	fmt.Println("Done generating")
	channel <- key
}

func main() {
	rand.Seed(time.Now().Unix())
	channel := make(chan int)

	for i := 0; i < 4; i++ {
		go generateKey(channel)
	}
	for i := 0; i < 4; i++ {
		fmt.Println(<-channel)
	}
	fmt.Println("All done!")
}
