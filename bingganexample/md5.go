package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	md5sum := md5.Sum([]byte("hello"))
	fmt.Printf("%s\n", md5sum)

	var b []byte
	b = []byte("hello")
	fmt.Printf("%v\n", b)

	md5sum1 := md5.Sum([]byte("hello1"))
	if md5sum == md5sum1 {
	}
	fmt.Println(md5sum)
}
