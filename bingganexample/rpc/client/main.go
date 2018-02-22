package main

import (
	"log"
	"net/rpc"

	"github.com/51reboot/golang-01-homework/lesson15/binggan/rpc/common"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8021")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	req := common.AddRequest{
		M: 10,
		N: 20,
	}
	var reply common.AddResponse
	err = client.Call("MathService.Add", &req, &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("result:%d", reply.Result)
}
