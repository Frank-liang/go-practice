package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	r, err := gzip.NewReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, r)
}
