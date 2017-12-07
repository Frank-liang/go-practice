package main

import "fmt"
import "log"
import "net"

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Print(conn.LocalAddr())
	log.Print(conn.RemoteAddr())
	fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: www.baidu.com\r\n\r\n")
	//conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}
		fmt.Print(string(buf[:n]))
	}
}
