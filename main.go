package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Sreethecool/rpc/client"
	"github.com/Sreethecool/rpc/server"
	"github.com/Sreethecool/rpc/validator"
)

func main() {
	fmt.Println("RPC Protocols and Exception handling with Custom Errors")

	a := validator.AddressValidator{
		Address: "localhost",
		Port:    "8080",
	}

	server.RunServer(a)

	reader := bufio.NewReader(os.Stdin)
	c := client.NewClient(a)
	fmt.Println("New Client Created...")
	for {
		reply := ""
		line, isPrefix, err := reader.ReadLine()
		for isPrefix && err != nil {
			var in []byte
			in, isPrefix, err = reader.ReadLine()
			line = append(line, in...)
		}
		if err != nil {
			log.Panicf("Error in reading input: %q", err)
		}
		err = c.Call("AddressValidator.Send", string(line), &reply)
		if err != nil {
			fmt.Println(err.Error())
		}

		if reply == "done" || err != nil {
			c.Close()
			c = client.NewClient(a)
			fmt.Println("New Client Created...")
		} else {
			fmt.Println("Server Sends :", reply)
		}

	}
}
