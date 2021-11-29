package utils

import (
	"fmt"

	"github.com/Sreethecool/rpc/validator"
)

//Returns URL of the addressValidator
func GetURL(a validator.AddressValidator) string {
	addr := fmt.Sprintf("%s:%s", a.Address, a.Port)
	return addr
}
