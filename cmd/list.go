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
	b := getBinding()
	conn := getConn(b)
	client := client.NewClient(conn)
	client.Auth(b.AuthSecret)
	authResp := <-client.MsgChan
	if authResp.Status != protocol.StatusOk {
		log.Fatal("unauthorized")
	}

	var prefix string
	if len(args) > 0 {
		prefix = args[0]
	}

	err := client.List(prefix)
	if err != nil {
		log.Fatal(err)
	}

	for resp := range client.MsgChan {
		if resp.Status == protocol.StatusStreamEnd {
			fmt.Printf("\r\nEND")
			break
		}
		fmt.Printf("%s, ", resp.Key)
	}

	fmt.Print("\r\n")
}
