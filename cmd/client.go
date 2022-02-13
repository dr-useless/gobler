package cmd

import (
	"net/rpc"
)

func getClient() (*rpc.Client, Binding, error) {
	b := Binding{}
	b.read()
	conn, err := rpc.Dial("tcp", b.Address)
	return conn, b, err
}
