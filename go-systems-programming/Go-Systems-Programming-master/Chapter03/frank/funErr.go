package main

import (
	"errors"
	"fmt"
	"log"
)

func division(x, y int) (int, error, error) {
	if y == 0 {
		return 0, nil, errors.New("Can't divide by zero!")
	}
	if x%y != 0 {
		remainder := errors.New("There is a remainder")
		return x / y, remainder, nil
	} else {
		return x / y, nil, nil
	}
}

func main() {
	result, rem, err := division(12, 5)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("The result is ", result)
	}
	if rem != nil {
		fmt.Println(rem)
	}

}
