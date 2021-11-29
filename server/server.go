package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/Sreethecool/rpc/utils"
	"github.com/Sreethecool/rpc/validator"
)

//Runserver Initializes RPC server and registers it with addressvalidator struct which is passed as parameter.
func RunServer(a validator.AddressValidator) {
	rpc.Register(&a)

	address, err := net.ResolveTCPAddr("tcp", utils.GetURL(a))
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
