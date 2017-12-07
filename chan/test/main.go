package main

import "fmt"
import "time"

func main() {
	s := "我hello\n"
	fmt.Printf("%T %v\n", s[0], s[0])
	first := make(chan int)
	in := first
	for i, n := range s {
		out := make(chan int)
		go func(i int, n rune, in chan int, out chan int) {
			in_num := <-in
			fmt.Printf("%v read from in %v\n", string(n), in_num)
			fmt.Println(string(n))
			fmt.Printf("%v write to next %v\n", string(n), i)
			out <- i
		}(i, n, in, out)
		in = out
	}
	fmt.Printf("%v write to begin\n", -1)
	first <- -1
	time.Sleep(1 * time.Second)
}
