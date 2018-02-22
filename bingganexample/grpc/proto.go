package main

import (
	"fmt"

	"github.com/51reboot/golang-01-homework/lesson15/binggan/grpc/myproto"
	"github.com/golang/protobuf/proto"
)

func main() {
	var p myproto.Person
	p.Id = 1
	p.Name = "binggan"
	p.Email = "binggan@xx.com"
	p.Phones = []*myproto.PhoneNumber{
		{Number: "123456", Type: myproto.PhoneType_MOBILE},
	}

	buf, err := proto.Marshal(&p)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(buf))
}
