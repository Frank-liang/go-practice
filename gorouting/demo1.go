package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(num int) {
			fmt.Println(num)
		}(i)
	}
	time.Sleep(time.Second * 1)
}
