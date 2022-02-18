package cmd

import (
	"bufio"
	"log"

	"github.com/dr-useless/gobkv/protocol"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Test the bound gobkv instance",
	Long:  "Usage: gobler ping",
	Run:   handlePing,
}

func init() {
	rootCmd.AddCommand(pingCmd)
}

func handlePing(cmd *cobra.Command, args []string) {
	conn := getConn(getBinding())

	msg := protocol.Message{
		Op: protocol.OpPing,
	}

	bw := bufio.NewWriter(conn)
	msg.Write(bw)
	bw.Flush()

	resp := protocol.Message{}
	br := bufio.NewReader(conn)
	resp.Read(br)

	conn.Close()

	log.Printf("op: %s, status: %s\r\n",
		protocol.MapOp()[resp.Op],
		protocol.MapStatus()[resp.Status])
}
