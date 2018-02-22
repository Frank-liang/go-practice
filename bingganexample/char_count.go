package main

import (
	"bytes"
	"fmt"
	"io"
)

type ByteCounter struct {
	Sum int
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	b.Sum += len(p)
	return len(p), nil
}

type LineCounter struct {
	Sum int
}

func (l *LineCounter) Write(p []byte) (int, error) {
	for _, b := range p {
		if b == '\n' {
			l.Sum++
		}
	}
	return len(p), nil
}

func main() {
	l := new(LineCounter)
	b := new(ByteCounter)

	buf := new(bytes.Buffer)
	buf.WriteString(`
hello gopher
12334
main new
`)

	w := io.MultiWriter(l, b)
	io.Copy(w, buf)
	fmt.Println(l.Sum)
	fmt.Println(b.Sum)
}
