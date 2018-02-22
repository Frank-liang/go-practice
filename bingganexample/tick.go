package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(time.Second)
	cnt := 0
	for _ = range timer.C {
		cnt++
		if cnt > 3 {
			timer.Stop()
			return
		}
		fmt.Println("hello")
	}
}
