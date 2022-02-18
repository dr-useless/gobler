package cmd

import (
	"log"

	"github.com/dr-useless/gobkv/protocol"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Remove a key from the bound gobkv instance",
	Long:  "Usage: gobler del the_key",
	Run:   handleDel,
}

func init() {
	rootCmd.AddCommand(delCmd)
}

func handleDel(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Println("specify a key")
		return
	}

	binding := getBinding()
	conn := getConn(binding)

	msg := protocol.Message{
		Op:  protocol.OpDelAck,
		Key: args[0],
	}

	msg.Write(conn)

	msg.Read(conn)

	log.Println("status:",
		protocol.MapStatus()[msg.Status])
}
