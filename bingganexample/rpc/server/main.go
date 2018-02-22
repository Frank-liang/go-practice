package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/51reboot/golang-01-homework/lesson15/binggan/rpc/common"
)

type MathService struct {
}

func (m *MathService) Add(request *common.AddRequest, reply *common.AddResponse) error {
	log.Printf("call add:%v", request)
	reply.Result = request.M + request.N
	return nil
}

// func (m *MathService) Mul() {

// }

func main() {
	mathService := new(MathService)
	rpc.Register(mathService)
	l, err := net.Listen("tcp", ":8021")
	if err != nil {
		log.Fatal(err)
	}
	rpc.Accept(l)
}
