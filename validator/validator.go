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

//Messages that will be sent by Server
var reqBlockchainStr = "Please enter your blockchain address"
var errInvalidAck = "Invalid Acknowledgement"
var recValidBlockChainStr = "Received Blockchain address + "
var errInvalidAddress = "Invalid Blockchain Address"
var errInvalidRequest = "Invalid Request"

//It resets state validation flags
func (a *AddressValidator) reset() {
	a.handshake = false
	a.complete = false
}

//Send: checks the state with the help of flags, and sends response based on the valdation.

func (a *AddressValidator) Send(msg string, reply *string) error {
	var err error
	*reply = ""
	fmt.Println("Server Recieved the msg : ", msg)
	if !a.handshake && strings.ToLower(msg) == "hi" { //Check if msg is handshake message and there s no previous handshake
		a.handshake = true
		*reply = reqBlockchainStr
	} else if a.handshake {
		if a.complete && strings.ToLower(msg) == "ok" { //check if the recieved msg is ack after address validation and resets flags
			a.reset()
			*reply = "done"
		} else if a.complete {
			*reply = errInvalidAck
			err = errors.New("invalid acknowledgement")
			a.reset()
		} else if common.IsHexAddress(msg) { //if handshake is complete then it validates address
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
