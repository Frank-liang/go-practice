package main

import (
	"log"
	"net"
	"sync"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/51reboot/golang-01-homework/lesson15/binggan/grpc/myproto"
)

type addressBookStoreServer struct {
	mutex sync.Mutex
	id    int32
	book  *myproto.AddressBook
}

func newAddressBookStoreServer() *addressBookStoreServer {
	return &addressBookStoreServer{
		id:   0,
		book: new(myproto.AddressBook),
	}
}

func (s *addressBookStoreServer) AddPerson(ctx context.Context, req *myproto.AddPersonRequest) (*myproto.AddPersonResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.id++
	log.Printf("add called:%v", req)
	// req.GetPerson().Id = s.id
	// s.book.People = append(s.book.People, req.GetPerson())
	return &myproto.AddPersonResponse{
		Id: s.id,
	}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8021")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	myproto.RegisterAddressBookStoreServer(server, newAddressBookStoreServer())
	server.Serve(l)
}
