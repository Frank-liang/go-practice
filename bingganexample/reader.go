package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

type MyReader struct {
	r   io.Reader
	cnt int
}

func NewMyReader(r io.Reader) *MyReader {
	return &MyReader{
		r: r,
	}
}

func (r *MyReader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	r.cnt += n
	return n, err
}

func main() {
	uncompress, err := gzip.NewReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	//r := NewMyReader(uncompress)
	io.Copy(os.Stdout, uncompress)
	fmt.Println("size", r.cnt)
}
