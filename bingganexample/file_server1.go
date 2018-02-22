package main

import (
	"bufio"
	"flag"
	"log"
	"net"
)

var (
	root = flag.String("root", "/", "root of ftp server data dir")
)

// client -> GET /home/binggan/a.txt\n
// server -> content of /home/bingan/a.txt

// client -> STORE /home/bingan/a.txt\n content of file EOF
// server -> OK

// client -> LS /home/bingan\n
// server -> content of dir /home/bingan

func handleConn(conn net.Conn) {
	r := bufio.NewReader(conn)
	line, err := r.ReadString('\n')

	var content []byte
	// 读取文件内容到content

	conn.Write(content)
}

func main() {
	addr := ":8021"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}
