package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	state, err := terminal.MakeRaw(0) //终端进入原始模式，把之前的状态保持下去，结束的时候把状态存下来
	if err != nil {
		log.Fatal(err)
	}
	defer terminal.Restore(0, state)
	conn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	defer conn.Close()
	go io.Copy(conn, os.Stdin)
	io.Copy(os.Stdout, conn)
}
