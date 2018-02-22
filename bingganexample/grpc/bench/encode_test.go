package main

import (
	"encoding/json"
	"testing"

	"github.com/51reboot/golang-01-homework/lesson15/binggan/grpc/myproto"
	"github.com/golang/protobuf/proto"
)

func BenchmarkProto(b *testing.B) {
	var p myproto.Person
	p.Id = 1
	p.Name = "binggan"
	p.Email = "binggan@xx.com"
	p.Phones = []*myproto.PhoneNumber{
		{Number: "123456", Type: myproto.PhoneType_MOBILE},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(&p)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkJSON(b *testing.B) {
	var p myproto.Person
	p.Id = 1
	p.Name = "binggan"
	p.Email = "binggan@xx.com"
	p.Phones = []*myproto.PhoneNumber{
		{Number: "123456", Type: myproto.PhoneType_MOBILE},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&p)
		if err != nil {
			panic(err)
		}
	}
}
