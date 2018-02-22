package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
	}
	fmt.Println(m["a"])
	delete(m, "a")
	fmt.Println(m["a"])

	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int)
}
