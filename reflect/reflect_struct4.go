package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Gopher struct {
	Name string //大写的容易修改
	age  int    //小写的不容易修改
}

func main() {
	var g = Gopher{}
	//	v := reflect.ValueOf(&g)  这种写法取不出来
	v := reflect.ValueOf(&g).Elem()
	name := v.FieldByName("Name")
	age := v.FieldByName("age")
	fmt.Println("Canaddr:", name.CanAddr(), "CanSet", name.CanSet())
	fmt.Println("Canaddr:", age.CanAddr(), "CanSet", age.CanSet())
	g.Name = "frank"
	*(*string)(unsafe.Pointer(age.UnsafeAddr())) = "20"
	fmt.Println(g)

}
