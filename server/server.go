package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/Sreethecool/rpc/validator"
)

func RunServer(a validator.AddressValidator) {
	rpc.Register(&a)

	address, err := net.ResolveTCPAddr("tcp", a.Url())
	if err != nil {
		log.Fatalf("Error :%v", err)
	}

	fmt.Printf("Listening on port %s...\n", address.String())
	conn, err := net.Listen("tcp", address.String())
	if err != nil {
		log.Fatalf("Cant listen to address : %s", address.String())
	}

	go rpc.Accept(conn)
}
