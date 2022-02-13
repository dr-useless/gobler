package main

import (
	"log"

	"github.com/dr-useless/gobler/cmd"
)

func main() {
	log.SetPrefix("gobler ")
	cmd.Execute()
}
