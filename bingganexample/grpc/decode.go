package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/51reboot/golang-01-homework/lesson15/binggan/grpc/myproto"
	"github.com/golang/protobuf/proto"
)

func main() {
	var p myproto.Person
	buf, _ := ioutil.ReadAll(os.Stdin)
	err := proto.Unmarshal(buf, &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p.String())
}
