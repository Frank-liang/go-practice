package main

import "fmt"

func sum(s []string, c chan string) {
	var sum string
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []string{"hello", "golang", "c++", "world"}

	c1 := make(chan string)
	c2 := make(chan string)
	go sum(s[:len(s)/2], c1)
	go sum(s[len(s)/2:], c2)
	x, y := <-c1, <-c2 // receive from c

	fmt.Println(x, y, x+y)
}
