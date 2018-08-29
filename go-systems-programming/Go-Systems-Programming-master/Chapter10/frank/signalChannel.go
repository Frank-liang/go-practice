package main

import (
	"fmt"
	"time"
)

func A(a, b chan struct{}) {
	<-a
	fmt.Println("A!")
	time.Sleep(time.Second)
	close(b)
}

func B(b, c chan struct{}) {
	<-b
	fmt.Println("B!")
	close(c)
}

func C(a chan struct{}) {
	<-a
	fmt.Println("C!")
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	go A(x, y)
	go C(z)
	go B(y, z)
	go C(z)

	close(x)
	time.Sleep(2 * time.Second)
}
