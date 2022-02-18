package cmd

import (
	"log"

	"github.com/dr-useless/gobkv/protocol"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a value from the bound gobkv instance",
	Long:  "Usage: gobler get the_key",
	Run:   handleGet,
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func handleGet(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Println("specify a key")
		return
	}

	binding := getBinding()
	conn := getConn(binding)

	msg := protocol.Message{
		Op:  protocol.OpGet,
		Key: args[0],
	}

	msg.Write(conn)

	msg.Read(conn)

	log.Printf("op: %s, status: %s, value: %s\r\n",
		protocol.MapOp()[msg.Op],
		protocol.MapStatus()[msg.Status],
		string(msg.Value))
}
