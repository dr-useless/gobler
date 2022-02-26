package cmd

import (
	"fmt"
	"log"

	"github.com/intob/rocketkv/client"
	"github.com/intob/rocketkv/protocol"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Test the bound rocketkv instance",
	Long:  "Usage: rkteer ping",
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
	resp := <-client.Msgs
	log.Println("recvd")
	fmt.Println(protocol.MapStatus()[resp.Status])
}
