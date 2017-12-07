package main

import "fmt"
import "log"
import "net"
import "net/url"
import "time"

func urlopen(uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	//u.Host -> www.baidu.com
	//u.PATH -> /s
	//u.RawQuery -> wd=golang
	conn, err := net.Dial("tcp", u.Host+":80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Fprint(conn, "GET %s?%s HTTP/1.1\r\nHost: %s\r\n\r\n", u.Path, u.RawQuery, u.Host)

	buf := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))

	var res []byte
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}
		res = append(res, buf[:n]...)
	}
	return string(res)
}

func main() {
	fmt.Print(urlopen("http://www.baidu.com/s?wd=golang"))
}
