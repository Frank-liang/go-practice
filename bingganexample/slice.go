package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
	fmt.Println(&s[0])
	fmt.Println(&primes[1])
	var s1 []int
	s1 = s
	fmt.Println(&s1[0] == &s[0])
}
