package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"strings"
	"sync"
)

// 握手

func err_line(err error) error {
	_, _, line, _ := runtime.Caller(1)
	s := []string{"line: ", string(line), err.Error()}
	err_str := errors.New(strings.Join(s, " "))
	fmt.Println(err_str)
	return err_str
}

func handshake(r *bufio.Reader, conn net.Conn) error {
	version, err := r.ReadByte()
	if err != nil {
		return err_line(err)
	}

	// 读取第一个字节	位版本号
	log.Printf("version: %d", version)

	n_method, err := r.ReadByte() // 读取第二个字节 客户端请求类型 1为代理
	if err != nil {
		return err_line(err)
	}
	log.Printf("method length: %d", n_method)

	buf := make([]byte, n_method) // 读取第三个字节为 客户端支持的验证方式
	io.ReadFull(r, buf)
	log.Printf("method: %v", buf)

	resp := []byte{5, 0} // 收到客户端验证后 回应客户端 0 不需要验证
	conn.Write(resp)     // 服务端需要客户端提供哪种验证方式信息
	return nil
}

// 获取客户端请求的地址和端口
func readAddr(r *bufio.Reader) (string, error) {
	version, err := r.ReadByte() // 读取第一个字节	位版本号
	if err != nil {
		return "", err_line(err)
	}
	log.Printf(" socks version: %d", version)

	cmd, err := r.ReadByte() // 读取第二个字节 客户端请求类型 1为代理
	if err != nil {
		return "", err_line(err)
	}
	log.Printf("CMD: %d", cmd)

	r.ReadByte() // 第三个字节为保留字 RSV

	atyp, err := r.ReadByte() //第四个字节为ATYP 请求的远程服务器地址类型 ip domainname ipv6
	if err != nil {
		return "", err_line(err)
	}
	log.Printf("reomote host type: %d", atyp)

	addr_len, err := r.ReadByte() // 这个字节代表 远程服务器地址的长度
	if err != nil {
		return "", err_line(err)
	}
	log.Printf("remote addr lenght: %d", addr_len)

	addr := make([]byte, addr_len) // 根据服务器地址的长度  去读取地址
	io.ReadFull(r, addr)
	log.Printf("remote addr is %v:", string(addr))

	var port int16
	binary.Read(r, binary.BigEndian, &port)
	log.Printf("remote port: %d", port)
	return fmt.Sprintf("%s:%d", addr, port), nil
}

// 开始代理
func startProxy(addr string, client net.Conn) error {
	defer client.Close()
	remote, err := net.Dial("tcp", addr)

	if err != nil {
		return err_line(err)
	}
	defer remote.Close()
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		io.Copy(client, remote)
		//client.Close()
	}()
	go func() {
		defer wg.Done()
		io.Copy(remote, client)
		//remote.Close()
	}()
	wg.Wait()
	return nil
}
func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	err := handshake(r, conn) // 握手
	if err != nil {
		fmt.Println("handshake err:::", err)
		return
	}
	addr, err := readAddr(r)
	if err != nil {
		fmt.Println("read client Addr err:::", err)
		return
	}
	fmt.Println("remote addr is :::::", addr)
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	conn.Write(resp)

	// 开始代理
	err = startProxy(addr, conn)
	if err != nil {
		return
	}
}

func main() {
	l, err := net.Listen("tcp", ":8022")
	if err != nil {
		_, _, line, _ := runtime.Caller(0)
		log.Fatal("line: ", line, err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			_, _, line, _ := runtime.Caller(0)
			log.Fatal("line", line, err)
		}

		go handleConn(conn)
	}
}
