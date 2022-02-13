package cmd

import (
	"crypto/tls"
	"log"
	"net/rpc"
)

func getClient() (*rpc.Client, Binding) {
	b := Binding{}
	b.read()
	if b.CertFile == "" {
		// return client on open tcp connection
		client, err := rpc.Dial("tcp", b.Address)
		if err != nil {
			log.Fatal(err)
		}
		return client, b
	} else {
		// load cert & key
		cert, err := tls.LoadX509KeyPair(b.CertFile, b.KeyFile)
		if err != nil {
			log.Fatalf("failed to load key pair: %s", err)
		}
		config := tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		}
		// return client on tls connection
		conn, err := tls.Dial("tcp", b.Address, &config)
		return rpc.NewClient(conn), b
	}
}
