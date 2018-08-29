package main

import "fmt"
import "time"

func namedFunction() {
	time.Sleep(10000 * time.Microsecond)
	fmt.Println("Printing from namedFunction")
}
func main() {
	fmt.Println("Chapter 09 - Goroutines.")
	go namedFunction()
	go func() {
		fmt.Println("An anonymous function!")
	}()
	go namedFunction()
	time.Sleep(100000 * time.Microsecond)
	fmt.Println("Exiting...")

}
