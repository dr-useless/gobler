package cmd

import (
	"fmt"
	"log"

	"github.com/intob/rocketkv/client"
	"github.com/intob/rocketkv/protocol"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a value from the bound rocketkv instance",
	Long:  "Usage: rkteer get the_key",
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

	b := getBinding()
	conn := getConn(b)
	client := client.NewClient(conn)
	if b.AuthSecret != "" {
		client.Auth(b.AuthSecret)
		authResp := <-client.Msgs
		if authResp.Status != protocol.StatusOk {
			log.Fatal("unauthorized")
		}
	}

	err := client.Get(args[0])
	if err != nil {
		log.Fatal(err)
	}

	resp := <-client.Msgs

	if resp.Status == protocol.StatusOk {
		fmt.Println(string(resp.Value))
	} else {
		fmt.Println(protocol.MapStatus()[resp.Status])
	}
}
