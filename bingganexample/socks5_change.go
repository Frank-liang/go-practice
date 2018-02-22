package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

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
		e := recover() // interface{}
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
	r.ReadByte()

	addrtype, _ := r.ReadByte()
	log.Printf("addr type:%d", addrtype)
	if addrtype != 3 {
		return "", errors.New("bad addr type")
	}

	// 读取一个字节的数据，代表后面紧跟着的域名的长度
	// 读取n个字节得到域名,n根据上一步得到的结果来决定
	addrlen, _ := r.ReadByte()
	buf := make([]byte, addrlen)
	io.ReadFull(r, buf)
	log.Printf("addr:%s", buf)

	var port int16
	binary.Read(r, binary.BigEndian, &port)

	return fmt.Sprintf("%s:%d", addr, port), nil
}

func handshake(r *bufio.Reader, conn net.Conn) error {
	version, _ := r.ReadByte()
	log.Printf("version:%d", version)
	if version != 5 {
		return errors.New("bad version")
	}
	nmethods, _ := r.ReadByte()
	log.Printf("nmethods:%d", nmethods)

	buf := make([]byte, nmethods)
	io.ReadFull(r, buf)
	log.Printf("%v", buf)

	resp := []byte{5, 0}
	conn.Write(resp)
	return nil
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)

	handshake(r, conn)
	addr, _ := readAddr(r)
	log.Printf("addr:%s", addr)
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	conn.Write(resp)

	// 开始代理
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
