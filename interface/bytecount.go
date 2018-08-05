package main

import "fmt"
import "io"
import "os"

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	b := new(ByteCounter)
	io.Copy(b, os.Stdin)
	fmt.Println(*b)
}
