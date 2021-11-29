package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type AddressValidator struct {
	Address   string
	Port      string
	handshake bool
	complete  bool
}

var reqBlockchainStr = "Please enter your blockchain address"
var errInvalidAck = "Invalid Acknowledgement"
var recValidBlockChainStr = "Received Blockchain address + "
var errInvalidAddress = "Invalid Blockchain Address"
var errInvalidRequest = "Invalid Request"

func (a *AddressValidator) Url() string {
	addr := fmt.Sprintf("%s:%s", a.Address, a.Port)
	return addr
}
func (a *AddressValidator) reset() {
	a.handshake = false
	a.complete = false
}
func (a *AddressValidator) Send(msg string, reply *string) error {
	var err error
	*reply = ""
	fmt.Println("Server Recieved the msg : ", msg)
	if !a.handshake && strings.ToLower(msg) == "hi" {
		a.handshake = true
		*reply = reqBlockchainStr
	} else if a.handshake {
		if a.complete && strings.ToLower(msg) == "ok" {
			a.reset()
			*reply = "done"
		} else if a.complete {
			*reply = errInvalidAck
			err = errors.New("invalid acknowledgement")
			a.reset()
		} else if common.IsHexAddress(msg) {
			*reply = recValidBlockChainStr + msg
			a.complete = true
		} else {
			*reply = errInvalidAddress
			err = errors.New("invalid blockchain address")
			a.reset()
		}
	} else {
		*reply = errInvalidRequest
		err = errors.New("invalid request")
		a.reset()
	}

	return err
}
