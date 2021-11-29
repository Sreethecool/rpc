package validator

import (
	"testing"
)

func TestSend(t *testing.T) {

	a := AddressValidator{
		Address: "localhost",
		Port:    "8080",
	}
	var resp string
	msg := "hi"
	err := a.Send(msg, &resp)
	if err != nil {
		t.Errorf("Error in Handshake:%s", err.Error())
	}
	if resp != reqBlockchainStr {
		t.Errorf("Invalid Response:%s", resp)
	}

	msg = "0xE4559721E46326F18FF59e3D926E4489FE6a5162"
	err = a.Send(msg, &resp)
	if err != nil {
		t.Errorf("Error in Validation:%s", err.Error())
	}
	if resp != recValidBlockChainStr+msg {
		t.Errorf("Invalid Response:%s", resp)
	}

	msg = "ok"
	err = a.Send(msg, &resp)
	if err != nil {
		t.Errorf("Error in Ack:%s", err.Error())
	}
	if resp != "done" {
		t.Errorf("Invalid Response:%s", resp)
	}

}

func TestSendWithWrongAddress(t *testing.T) {

	a := AddressValidator{
		Address: "localhost",
		Port:    "8080",
	}
	var resp string
	msg := "hi"
	err := a.Send(msg, &resp)
	if err != nil {
		t.Errorf("Error in Handshake:%s", err.Error())
	}
	if resp != reqBlockchainStr {
		t.Errorf("Invalid Response:%s", resp)
	}

	msg = "0xE4559721E46326F18scanclkans"
	err = a.Send(msg, &resp)
	if err == nil {
		t.Errorf("Wrong in put Accepted :%s", err.Error())
	}
	if resp != errInvalidAddress {
		t.Errorf("Invalid Response:%s", resp)
	}
}

func TestSendWithWrongHandshakes(t *testing.T) {

	a := AddressValidator{
		Address: "localhost",
		Port:    "8080",
	}
	var resp string
	msg := "hi hello"
	err := a.Send(msg, &resp)
	if err == nil {
		t.Errorf("Invalid Handshake:%s", err.Error())
	}
	if resp != errInvalidRequest {
		t.Errorf("Invalid Response:%s", resp)
	}

}
