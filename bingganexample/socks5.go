package main

import (
	"bufio"
	"crypto/md5"
	"crypto/rc4"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type CryptoWriter struct {
	w      io.Writer
	cipher *rc4.Cipher
}

func NewCryptoWriter(w io.Writer, key string) io.Writer {
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5sum[:])
	if err != nil {
		panic(err)

	}
	return &CryptoWriter{
		w:      w,
		cipher: cipher,
	}

}

// 把b里面的数据进行加密，之后写入到w.w里面
// 调用w.w.Write进行写入
func (w *CryptoWriter) Write(b []byte) (int, error) {
	buf := make([]byte, len(b))
	w.cipher.XORKeyStream(buf, b)
	return w.w.Write(buf)

}

type CryptoReader struct {
	r      io.Reader
	cipher *rc4.Cipher
}

func NewCryptoReader(r io.Reader, key string) io.Reader {
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5sum[:])
	if err != nil {
		panic(err)

	}

	return &CryptoReader{
		r:      r,
		cipher: cipher,
	}

}

func (r *CryptoReader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	buf := b[:n]
	r.cipher.XORKeyStream(buf, buf)
	return n, err

}

func mustReadByte(r *bufio.Reader) byte {
	b, err := r.ReadByte()
	if err != nil {
		panic(err)
	}

	return b
}

// 1.握手
// 2.获取客户端代理的请求
// 3.开始代理

func readAddr(r *bufio.Reader) (addr string, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = e.(error)
		}
	}()

	version := mustReadByte(r)
	log.Printf("version:%d", version)
	if version != 5 {
		return "", errors.New("bad version")
	}
	cmd := mustReadByte(r)
	log.Printf("cmd:%d", cmd)

	if cmd != 1 {
		return "", errors.New("bad cmd")
	}

	// skip rsv字段
	mustReadByte(r)

	addrtype := mustReadByte(r)
	log.Printf("addr type:%d", addrtype)
	if addrtype != 3 {
		return "", errors.New("bad addr type")
	}

	// 读取一个字节的数据，代表后面紧跟着的域名的长度
	// 读取n个字节得到域名,n根据上一步得到的结果来决定
	addrlen := mustReadByte(r)
	buf := make([]byte, addrlen)
	io.ReadFull(r, buf)
	var port int16
	binary.Read(r, binary.BigEndian, &port)

	return fmt.Sprintf("%s:%d", buf, port), nil
}

func handshake(r *bufio.Reader, w io.Writer) error {
	version := mustReadByte(r)
	log.Printf("version:%d", version)
	if version != 5 {
		return errors.New("bad version")
	}
	nmethods := mustReadByte(r)
	log.Printf("nmethods:%d", nmethods)

	buf := make([]byte, nmethods)
	io.ReadFull(r, buf)
	log.Printf("%v", buf)

	resp := []byte{5, 0}
	w.Write(resp)
	return nil
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(NewCryptoReader(conn, "123456"))
	w := NewCryptoWriter(conn, "123456")

	err := handshake(r, w)
	if err != nil {
		log.Print(err)
		return
	}
	addr, err := readAddr(r)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("addr:%s", addr)
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	w.Write(resp)

	remote, err := net.Dial("tcp", addr)
	if err != nil {
		log.Print(err)
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)
	// go 读取(conn)的数据，发送到remote，直到conn的EOF, 关闭remote
	go func() {
		defer wg.Done()
		io.Copy(remote, r)
		remote.Close()

	}()
	// go 读取remote的数据，发送到客户端(conn)，直到remote的EOF，关闭conn
	go func() {
		defer wg.Done()
		io.Copy(w, remote)
		conn.Close()
	}()

	// 等待两个协程结束
	wg.Wait()
}

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", ":8022")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, _ := l.Accept()
		go handleConn(conn)
	}
}
