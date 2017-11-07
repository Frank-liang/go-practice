package main

import "fmt"

func main() {
	ages := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(ages)
	//delete(ages,"a")

	fmt.Println(ages["c"])
	for name, age := range ages {
		fmt.Println("name", name, "age", age)
	}
	for _, age := range ages {
		fmt.Println(age)
	}
}
