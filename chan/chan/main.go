package main

import "fmt"
import "time"

func sum(s []int, c chan int) {
	var sum int
	for _, v := range s {
		sum += v
	}
	fmt.Println("start of %v\n", s)
	c <- sum
	fmt.Println("end of %v\n", s)
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	c1 := make(chan int, 1)
	go sum(s[:len(s)/2], c1)
	time.Sleep(time.Second)
	x, y := <-c1, 0
	fmt.Println(x, y, x+y)
	time.Sleep(time.Second)
}
