package main

import "fmt"
import "log"
import "net"
import "time"

func GetTimeNow() time.Time {
	time.Sleep(time.Second * 2)
	return time.Now()
}

func HandleConn(conn net.Conn) {
	log.Println(conn.RemoteAddr())
	log.Println(conn.LocalAddr())
	fmt.Fprint(conn, "test\n", GetTimeNow())
	conn.Close()

}
func main() {
	listener, err := net.Listen("tcp", ":8880")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept() //net.Dial
		if err != nil {
			log.Fatal(err)
		}
		go HandleConn(conn)
	}
}
