package main

import (
	"bytes"
	"io"
	"os"
)

type XorWriter struct {
	w io.Writer
	x byte
}

func (x *XorWriter) Write(p []byte) (int, error) {
	p1 := make([]byte, len(p))
	copy(p1, p)
	for i, b := range p1 {
		p1[i] = b ^ x.x
	}
	return x.w.Write(p1)
}

func NewXorWriter(w io.Writer, x byte) *XorWriter {
	return &XorWriter{
		w: w,
		x: x,
	}
}

type XorReader struct {
	r io.Reader
	x byte
}

func (x *XorReader) Read(p []byte) (int, error) {
	return x.r.Read(p)
}

func NewXorReader(r io.Reader, x byte) *XorReader {
	return &XorReader{
		r: r,
		x: x,
	}
}

func main() {
	buf := new(bytes.Buffer)
	x := NewXorWriter(buf, 'a')
	io.WriteString(x, "hello")
	//fmt.Println(buf.Bytes())

	x1 := NewXorReader(buf, 'a')
	io.Copy(os.Stdout, x1)
}
