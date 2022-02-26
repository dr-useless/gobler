package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const path = "rkteer.cfg.json"

type Binding struct {
	Network    string
	Address    string
	AuthSecret string
	CertFile   string
	KeyFile    string
}

var bindCmd = &cobra.Command{
	Use:   "bind",
	Short: "Bind to a rocketkv instance",
	Long:  "Usage: rkteer bind [NETWORK] [ADDRESS] [AUTH_SECRET] [CERT_FILE] [KEY_FILE]",
	Run:   handleBind,
}

func init() {
	rootCmd.AddCommand(bindCmd)
	bindCmd.Flags().String("a", "", "auth secret")
	bindCmd.Flags().String("c", "", "TLS cert file")
	bindCmd.Flags().String("k", "", "TLS key file")
}

func handleBind(cmd *cobra.Command, args []string) {
	b := Binding{}

	if len(args) < 2 {
		log.Println("specify a network & address")
		return
	}

	b.Network = args[0]
	b.Address = args[1]
	log.Printf("binding to: %s %s\r\n", b.Network, b.Address)

	authSecret, _ := cmd.Flags().GetString("a")
	b.AuthSecret = authSecret
	log.Printf("with auth secret: %s\r\n", authSecret)

	certFile, _ := cmd.Flags().GetString("c")
	b.CertFile = certFile
	keyFile, _ := cmd.Flags().GetString("k")
	b.KeyFile = keyFile
	log.Printf("using cert %s & key %s\r\n", certFile, keyFile)

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
	data, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return err
	}
	r := bytes.NewReader(data)
	_, err = io.Copy(f, r)
	return err
}
