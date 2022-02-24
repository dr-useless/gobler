package cmd

import (
	"fmt"
	"log"

	"github.com/intob/gobkv/client"
	"github.com/intob/gobkv/protocol"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all keys with the given prefix",
	//Long:  "Usage: gobler list key_prefix --limit 100",
	Long: "Usage: gobler list key_prefix",
	Run:  handleList,
}

func init() {
	rootCmd.AddCommand(listCmd)
	//listCmd.Flags().Int("limit", 0, "max number of keys to return")
}

func handleList(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Println("specify a key prefix")
		return
	}

	b := getBinding()
	conn := getConn(b)
	client := client.NewClient(conn)
	client.Auth(b.AuthSecret)
	authResp := <-client.MsgChan
	if authResp.Status != protocol.StatusOk {
		log.Fatal("unauthorized")
	}

	err := client.List(args[0])
	if err != nil {
		log.Fatal(err)
	}

	resp := <-client.MsgChan

	if resp.Status == protocol.StatusOk {
		for _, k := range resp.Keys {
			fmt.Printf("%s, ", k)
		}
		fmt.Print("\r\n")
	} else {
		fmt.Println(protocol.MapStatus()[resp.Status])
	}
}
