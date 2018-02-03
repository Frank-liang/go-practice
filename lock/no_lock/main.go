package main

import (
	"fmt"
	"time"
)

var i int64

func Inc(i *int64) {
	counter := 0
	for {
		counter += 1
		*i += 1
		if counter > 10000000 {
			break
		}
	}
	fmt.Println("finish")
	return
}

func main() {
	go Inc(&i)
	go Inc(&i)
	go Inc(&i)

	time.Sleep(time.Second * 20)
	fmt.Println(i)
	return
}
