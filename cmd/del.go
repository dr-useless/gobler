package cmd

import (
	"fmt"
	"log"

	"github.com/intob/rocketkv/client"
	"github.com/intob/rocketkv/protocol"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Remove a key from the bound rocketkv instance",
	Long:  "Usage: rkteer del the_key",
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

	b := getBinding()
	conn := getConn(b)
	client := client.NewClient(conn)
	client.Auth(b.AuthSecret)
	authResp := <-client.Msgs
	if authResp.Status != protocol.StatusOk {
		fmt.Println(protocol.MapStatus()[authResp.Status])
	}

	err := client.Del(args[0], true)
	if err != nil {
		log.Fatal(err)
	}

	resp := <-client.Msgs

	fmt.Println(protocol.MapStatus()[resp.Status])
}
