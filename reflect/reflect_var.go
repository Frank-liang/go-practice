package main

import (
	"fmt"
	"reflect"
)

type x32 int32

func main() {
	var a x32 = 100
	t := reflect.TypeOf(a)
	tp := reflect.TypeOf(&a)
	fmt.Println(t.Name(), t.Kind())
	fmt.Println(tp.Kind(), tp.Elem())

	v := reflect.ValueOf(a)
	fmt.Println(v.Type(), v)
}
