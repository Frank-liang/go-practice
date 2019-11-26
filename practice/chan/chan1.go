package main

import (
	"fmt"
	"time"
)

func main() {
	var i int
	go func(i *int) {
		for j := 0; j < 20; j++ {
			time.Sleep(time.Millisecond)
			fmt.Println(*i, j)
		}
	}(&i)

	for i = 0; i < 20; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(i)
	}
}
