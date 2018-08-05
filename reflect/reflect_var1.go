package main

import (
	"fmt"
	"reflect"
)

type x32 int32

func main() {
	var a int32 = 100
	va := reflect.ValueOf(&a).Elem()
	fmt.Println("CanAddr: ", va.CanAddr(), "CanSet", va.CanSet(), va.Type())

	var b int32 = 50
	vb := reflect.ValueOf(b)
	va.Set(vb)
	fmt.Println(a)

}
