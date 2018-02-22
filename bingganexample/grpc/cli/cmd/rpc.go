package cmd

import (
	"log"

	"github.com/51reboot/golang-01-homework/lesson15/binggan/grpc/myproto"

	"google.golang.org/grpc"
)

func newClient(addr string) myproto.AddressBookStoreClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	return myproto.NewAddressBookStoreClient(conn)
}
