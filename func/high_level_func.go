package main

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

func main() {
	x, y := 12, 23
	op := func(x, y int) int {
		return x + y
	}
	result, err := calculate(x, y, op)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)
	result, err = calculate(x, y, nil)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)
}
