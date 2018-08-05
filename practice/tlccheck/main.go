package main

import (
	"crypto/tls"
	"fmt"
	"log"
)

func main() {
	conn, err := tls.Dial("tcp", "www.zcool.com.cn:443", nil)
	if err != nil {
		log.Fatal(err)
	}
	state := conn.ConnectionState()
	fmt.Printf("%v\n", state.PeerCertificates)
	for _, p := range state.PeerCertificates {
		fmt.Println(p.NotBefore, p.NotAfter)
	}
}
