package cmd

import (
	"fmt"
	"log"

	"github.com/intob/gobkv/client"
	"github.com/intob/gobkv/protocol"
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
	client := client.NewClient(conn)

	client.Ping()

	log.Println("waiting")
	resp := <-client.MsgChan
	log.Println("recvd")
	fmt.Println(protocol.MapStatus()[resp.Status])
}
