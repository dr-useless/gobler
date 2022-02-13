package cmd

import (
	"log"
	"net/rpc"
)

func getClient() (*rpc.Client, Binding) {
	b := Binding{}
	b.read()
	conn, err := rpc.Dial("tcp", b.Address)
	if err != nil {
		log.Fatal(err)
	}
	return conn, b
}
