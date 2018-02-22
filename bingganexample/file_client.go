package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8021")
	if err != nil {
		log.Fatal(err)
	}

	conn.Write([]byte("GET /home/bingan/a.txt\n"))
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println(buf[:n])
}
