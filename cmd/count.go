package cmd

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/intob/rocketkv/client"
	"github.com/intob/rocketkv/protocol"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count the number of keys with the given prefix",
	Long:  "Usage: rkteer count the_key_prefix",
	Run:   handleCount,
}

func init() {
	rootCmd.AddCommand(countCmd)
}

func handleCount(cmd *cobra.Command, args []string) {
	b := getBinding()
	conn := getConn(b)
	client := client.NewClient(conn)
	client.Auth(b.AuthSecret)
	authResp := <-client.Msgs
	if authResp.Status != protocol.StatusOk {
		log.Fatal("unauthorized")
	}

	var prefix string
	if len(args) > 0 {
		prefix = args[0]
	}

	err := client.Count(prefix)
	if err != nil {
		log.Fatal(err)
	}

	resp := <-client.Msgs

	if resp.Status == protocol.StatusOk {
		fmt.Println(binary.BigEndian.Uint64(resp.Value))
	} else {
		fmt.Println(protocol.MapStatus()[resp.Status])
	}
}
