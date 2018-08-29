package main

import (
	"fmt"
	"reflect"
)

func main() {
	type t1 int
	type t2 int

	x1 := t1(1)
	x2 := t2(1)
	x3 := 1

	st1 := reflect.ValueOf(&x1).Elem()
	st2 := reflect.ValueOf(&x2).Elem()
	st3 := reflect.ValueOf(&x3).Elem()

	typeOfX1 := st1.Type()
	typeOfX2 := st2.Type()
	typeOfX3 := st3.Type()

	fmt.Printf("X1 Type: %s\n", typeOfX1)
	fmt.Printf("X2 Type: %s\n", typeOfX2)
	fmt.Printf("X3 Type: %s\n", typeOfX3)
}
