package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	log.Print("remote: ", conn.RemoteAddr())
	r := bufio.NewReader(conn)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Print(err)
		return
	}
	fs := strings.Fields(line) //类似于split，但是此函数能把多个空格当一个空格处理
	method := fs[0]
	path := fs[1]
	log.Printf("method:%s", method)
	log.Printf("path:%s", path)
	if method == "GET" {
		//第一种，直接全部读取
		//content, err := ioutil.ReadFile(path) //读取所有的文件，放入内存，如果是大文件请注意
		//if err != nil {
		//	log.Print(err)
		//	return
		//}
		//fmt.Fprintf(conn, "%s", content)

		//按快读取数据，4k. 对应大文件
		//buf := make([]byte, 4096)
		//f, err := os.Open(path)
		//if err != nil {
		//	log.Print(err)
		//	return
		//}
		//defer f.Close()
		//for {
		//	n, err := f.Read(buf)
		//	if err != nil {
		//		break
		//	}
		//	conn.Write(buf[:n])
		//}
		// f.Seek(0, os.SEEK_CUR) 设置读文件的游标

		//ioutil.ReadAll() socket的所有内容一次读取出来，知道EOF,注意大文件

		//io.copy,本质上按块一样，默认缓存是32k
		f, err := os.Open(path)
		if err != nil {
			log.Print(err)
			return
		}
		defer f.Close()
		io.Copy(conn, f) // 频繁请求内存，对GC有压力

		//buf := make([]byte, 4096) 从内存池里面拿一个buffer
		//io.CopyBuffer(conn, f, buf) 可以指定io.Copy的缓存大小,此例子复用内存，优化GC,减少开销

	}

}

func main() {
	listener, err := net.Listen("tcp", ":8880")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go HandleConn(conn)
	}
}
