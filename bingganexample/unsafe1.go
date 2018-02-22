package main

import (
	"fmt"
	"unsafe"
)

// data pointer
// len int64
// cap int64

type SliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

type StringHeader struct {
	Data unsafe.Pointer
	Len  int
}

func slice(s []int, b int, len int) []int {
	hdr := SliceHeader{
		Data: unsafe.Pointer(&s[b]),
		Len:  len,
		Cap:  cap(s),
	}

	s1 := *(*[]int)(unsafe.Pointer(&hdr))
	return s1
}

func stringSlice(s string, b int, len int) string {
	hdr := *(*StringHeader)(unsafe.Pointer(&s))
	hdr.Data = unsafe.Pointer(uintptr(hdr.Data) + uintptr(b))
	hdr.Len = len

	s1 := *(*string)(unsafe.Pointer(&hdr))
	return s1
}

func ZeroCopyString(buf []byte) string {
	hdr := &StringHeader{
		Data: unsafe.Pointer(&buf[0]),
		Len:  len(buf),
	}
	return *(*string)(unsafe.Pointer(hdr))
}

func main() {
	s := []int{1, 2, 3}
	fmt.Println(&s[0])

	// hdr := SliceHeader{
	// 	Data: unsafe.Pointer(&s[1]),
	// 	Len:  2,
	// 	Cap:  3,
	// }

	// s1 := *(*[]int)(unsafe.Pointer(&hdr))

	/*
		var p *SliceHeader
		p = (*SliceHeader)(unsafe.Pointer(&s1))
		fmt.Printf("%#v\n", *p)
	*/
	s1 := slice(s, 0, 2)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(s1[0], s1[1])

	str := "hello"
	str1 := stringSlice(str, 0, 2)
	fmt.Println(str1)

	buf := []byte{'h', 'e', 'l', 'l', 'o'}
	str2 := ZeroCopyString(buf)
	buf[0] = 'a'
	fmt.Println(string(buf), str2)
}
