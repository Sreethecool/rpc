package client

import (
	"log"
	"net/rpc"

	"github.com/Sreethecool/rpc/validator"
)

type Message struct {
	ClientID string
	Msg      string
}

func NewClient(a validator.AddressValidator) *rpc.Client {
	client, err := rpc.Dial("tcp", a.Url())
	if err != nil {
		log.Panicf("Error establishing connection with host: %q", err)
	}
	return client

}
