package client

import (
	"log"
	"net/rpc"

	"github.com/Sreethecool/rpc/validator"
)

//NewClient tries to establish connection with rpc server and returns pointer to rpc.client
func NewClient(a validator.AddressValidator) *rpc.Client {
	client, err := rpc.Dial("tcp", a.Url())
	if err != nil {
		log.Panicf("Error establishing connection with host: %q", err)
	}
	return client

}
