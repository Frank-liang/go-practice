package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f, ok := w.(*os.File)
	if !ok {
		c, ok := w.(*bytes.Buffer)
		if !ok {

		}
	}
	_ = f
	_ = c
}
