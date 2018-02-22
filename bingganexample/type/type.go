package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		x byte

		x1 int8

		x2 int32
		x3 int64
		x4 uint8
		x5 uint32
		x6 uint64
	)
	x4 = 255
	x4 = x4 + 1
	fmt.Println(x, x1, x2, x3, x4, x5, x6)
	fmt.Println(unsafe.Sizeof(x1))
}
