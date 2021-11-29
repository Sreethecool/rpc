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
	//Initialise server
	server.RunServer(a)

	reader := bufio.NewReader(os.Stdin)
	c := client.NewClient(a)
	fmt.Println("New Client Created...")
	for {
		reply := ""
		//Reads the message to be sent via client
		line, isPrefix, err := reader.ReadLine()
		for isPrefix && err != nil { //handles large input
			var in []byte
			in, isPrefix, err = reader.ReadLine()
			line = append(line, in...)
		}
		if err != nil {
			log.Panicf("Error in reading input: %q", err)
		}

		//calls the Method via RPC client
		err = c.Call("AddressValidator.Send", string(line), &reply)
		if err != nil {
			fmt.Println(err.Error())
		}

		if reply == "done" || err != nil {
			//if there is error or validation complete close the connection. and new connection is created for further address need to be validated
			c.Close()
			c = client.NewClient(a)
			fmt.Println("New Client Created...")
		} else {
			fmt.Println("Server Sends :", reply)
		}

	}
}
