package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	m int8
	n int8
}

func main() {
	var n int
	var m [2]int8

	fmt.Println(m)

	var p *int
	p = &n
	p = (*int)(unsafe.Pointer(&m[0]))
	*p = 0x1010
	fmt.Println(m)

}
