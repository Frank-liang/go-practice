package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"io"
)

func main() {
	key := "123456"
	md5sum := md5.Sum([]byte(key))
	block, err := aes.NewCipher(md5sum[:]) // aes-128 16字节
	if err != nil {
		panic(err)
	}

	iv := make([]byte, block.BlockSize())
	io.ReadFull(rand.Reader, iv)

	stream := cipher.NewCFBEncrypter(block, iv)
	buf := []byte("hello")
	stream.XORKeyStream(buf, buf)

}
