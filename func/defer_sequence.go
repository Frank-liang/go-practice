package main

import "fmt"

func f(n int) {
	defer func() {
		fmt.Printf("defer %d\n", n)
	}()
	fmt.Printf("f(%d)\n", n+0/n) //crash when n ==0
	f(n - 1)
}
func main() {
	f(3)
}
