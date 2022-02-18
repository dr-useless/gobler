package cmd

import (
	"crypto/tls"
	"log"
	"net"
)

func getBinding() Binding {
	b := Binding{}
	err := b.read()
	if err != nil {
		log.Fatal("failed to read binding", err)
	}
	return b
}

func getConn(b Binding) net.Conn {

	if b.CertFile == "" {
		conn, err := net.Dial(b.Network, b.Address)
		if err != nil {
			log.Fatalf("failed to connect to %s over %s", b.Address, b.Network)
		}
		return conn
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
		conn, err := tls.Dial(b.Network, b.Address, &config)
		if err != nil {
			log.Fatalf("failed to connect to %s with TLS", b.Address)
		}
		return conn
	}
}
