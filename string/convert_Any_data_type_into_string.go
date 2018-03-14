package main

//Strings are actually very simple: they are just read-only slices of bytes with a bit of extra syntactic support from the language
//1 First method
import (
	"fmt"
	"strconv"
)

func main() {
	i := 123
	t := strconv.Itoa(i)
	fmt.Println(t)

}

// 2 Second method
//i := 123
//	t := fmt.Sprintf("We are currently processing ticket number %d.", i)
//		fmt.Println(t)
