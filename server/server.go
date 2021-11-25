package server

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type AddressValidator struct {
	port     string
	shutdown chan bool
}

func (a *AddressValidator) Send(msg string, reply *string) error {
	return nil
}

func (a *AddressValidator) Recieve(msg string, reply *string) error {
	return nil
}

func RunServer(a *AddressValidator) {
	rpc.Register(a)
	rpc.HandleHTTP()

	fmt.Printf("Listening on port %s...\n", a.port)

	l, err := net.Listen("tcp", a.port)
	if err != nil {
		fmt.Printf("Can't bind port to listen. %v", err)
		return
	}

	go http.Serve(l, nil)
}
