package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/kr/pty"
)

//client ->server: ls | grep go\n
// server -> client : content of command EOF

func handle(conn net.Conn) {
	log.Printf("connection %s", conn.RemoteAddr())
	defer conn.Close()
	//cmd, _ := ioutil.ReadAll(conn) //此方法读到EOF才显示内容
	//fmt.Printf("%s\n", string(cmd))
	//r := bufio.NewReader(conn)
	//cmdstr, _ := r.ReadString('\n')
	//conn.Write([]byte(cmd)) //回写 到conn
	cmd := exec.Command("bash")
	//cmd.Stdout = conn
	//cmd.Stderr = conn
	//cmd.Run()              or another method

	//out, _ := cmd.CombinedOutput()
	//conn.Write(out)

	fd, err := pty.Start(cmd)
	if err != nil {
		log.Fatal(err)
	}
	f, _ := os.Create(conn.RemoteAddr().String() + ".log")
	defer f.Close()
	go io.Copy(fd, io.TeeReader(conn, f)) //分流器，同时把数据写入f，返回conn
	io.Copy(conn, fd)
}

func main() {
	l, err := net.Listen("tcp", ":8010")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, _ := l.Accept()
		go handle(conn)
	}
}
