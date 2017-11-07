package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 4 {
		fmt.Println(`Usage: only support + - * / just like a+b/a-b/a*b/a\/b`)
	}
	first, _ := strconv.Atoi(os.Args[1])
	operator := os.Args[2]
	second, _ := strconv.Atoi(os.Args[3])
	switch operator {
	case "+":
		fmt.Println(first + second)
	case "-":
		fmt.Println(first - second)
	case "*":
		fmt.Println(first * second)
	case "/":
		if second == 0 {
			fmt.Println("Must be a positive number")
		}
		fmt.Println(first / second)
	default:
		fmt.Println("Do as usage!!!")
	}
}
