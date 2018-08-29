package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "c1 ok"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 4):
		fmt.Println("timeout c1")
	}

}
