package cmd

import "log"

func getClient() {
	b := Binding{}
	b.read()

	log.Println("binding", b)
}
