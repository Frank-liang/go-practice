package main

import (
	"context"
	"log"

	"github.com/51reboot/golang-01-homework/lesson15/binggan/grpc/myproto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8021", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := myproto.NewAddressBookStoreClient(conn)
	resp, err := client.AddPerson(context.TODO(), new(myproto.AddPersonRequest))
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp.GetId())
}
