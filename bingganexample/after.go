package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.After(time.Second * 3)
	<-c
	fmt.Println("done")
}
