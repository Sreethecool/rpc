package utils

import (
	"testing"

	"github.com/Sreethecool/rpc/validator"
)

func TestGetURL(t *testing.T) {
	a := validator.AddressValidator{
		Address: "localhost",
		Port:    "8080",
	}

	url := GetURL(a)
	if url != "localhost:8080" {
		t.Errorf("Error in GETURL Method. Got %s instead of localhost:8080", url)
	}
}
