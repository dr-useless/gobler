package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const (
	flagAddr = "addr"
	flagAuth = "auth"
)

const path = "gobler.json"

type Binding struct {
	Address    string
	AuthSecret string
}

var bindCmd = &cobra.Command{
	Use:   "bind",
	Short: "Bind to a gobkv instance",
	Long:  "Usage: gobler bind [ADDRESS] [AUTH_SECRET]",
	Run:   handleBind,
}

func init() {
	rootCmd.AddCommand(bindCmd)
}

func handleBind(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Println("specify an address")
		return
	}

	addr := args[0]
	var auth string
	if len(args) > 1 {
		auth = args[1]
	}

	log.Printf("binding to: %s\r\n", addr)

	if auth != "" {
		log.Printf("with auth secret: %s\r\n", auth)
	}

	b := Binding{
		Address:    addr,
		AuthSecret: auth,
	}
	b.write()
}

func (b *Binding) read() error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(b)
}

func (b *Binding) write() error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}
	r := bytes.NewReader(data)
	_, err = io.Copy(f, r)
	return err
}
