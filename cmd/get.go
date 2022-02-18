package cmd

import (
	"bufio"
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

	bw := bufio.NewWriter(conn)
	msg.Write(bw)
	bw.Flush()

	resp := protocol.Message{}
	br := bufio.NewReader(conn)
	resp.Read(br)

	log.Printf("op: %s, status: %s, value: %s\r\n",
		protocol.MapOp()[resp.Op],
		protocol.MapStatus()[resp.Status],
		string(resp.Value))
}
