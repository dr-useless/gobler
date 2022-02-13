package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const path = "gobler.json"

type Binding struct {
	Address    string
	AuthSecret string
	CertFile   string
	KeyFile    string
}

var bindCmd = &cobra.Command{
	Use:   "bind",
	Short: "Bind to a gobkv instance",
	Long:  "Usage: gobler bind [ADDRESS] [AUTH_SECRET] [CERT_FILE] [KEY_FILE]",
	Run:   handleBind,
}

func init() {
	rootCmd.AddCommand(bindCmd)
}

func handleBind(cmd *cobra.Command, args []string) {
	b := Binding{}

	if len(args) < 1 {
		log.Println("specify an address")
		return
	}

	b.Address = args[0]
	log.Printf("binding to: %s\r\n", args[0])

	if len(args) >= 2 {
		b.AuthSecret = args[1]
		log.Printf("with auth secret: %s\r\n", args[1])
	}

	if len(args) >= 4 {
		b.CertFile = args[2]
		b.KeyFile = args[3]
		log.Printf("using cert %s & key %s\r\n", args[2], args[3])
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
