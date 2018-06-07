package main

import (
	"flag"
	"fmt"
)

var (
	num int
)

func init() {
	flag.IntVar(&num, "n", 0, "number")
	flag.Parse()
}

func fib(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = x+y, x
	}
}
func main() {
	var ch = make(chan int, num)
	go fib(num, ch)
	//Why those three lines do not work?
	//	for c := range ch {
	//		fmt.Println(c)
	//	}

	for i := 0; i < num; i++ {
		c := <-ch
		fmt.Println(c)
	}

}
