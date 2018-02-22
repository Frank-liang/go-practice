package mycrypto

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestCrypto(t *testing.T) {
	key := "123456"
	memfile := new(bytes.Buffer)
	w := NewCryptoWriter(memfile, key)
	w.Write([]byte("hello"))

	r := NewCryptoReader(memfile, key)
	buf := make([]byte, 1024)
	n, _ := r.Read(buf)

	if string(buf[:n]) != "hello" {
		t.Errorf("not equal:#%s#, #%s#", buf[:n], "hello")
	}
}

func BenchmarkCrypto(b *testing.B) {
	buf := []byte(strings.Repeat("a", 10))

	w := NewCryptoWriter(ioutil.Discard, "123456")
	for i := 0; i < b.N; i++ {
		n, _ := w.Write(buf)
		b.SetBytes(int64(n))
	}
}
