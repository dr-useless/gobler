package cmd

import (
	"log"

	"github.com/dr-useless/gobkv/common"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all keys with the given prefix",
	Long:  "Usage: gobler list key_prefix --limit 100",
	Run:   handleList,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().Int("limit", 0, "max number of keys to return")
}

func handleList(cmd *cobra.Command, args []string) {
	client, binding := getClient()

	rpcArgs := common.Args{
		AuthSecret: binding.AuthSecret,
	}

	if len(args) > 0 {
		rpcArgs.Key = args[0]
	}

	limit, err := cmd.Flags().GetInt("limit")
	if err != nil {
		log.Fatal("limit must be a valid integer")
	}
	rpcArgs.Limit = limit

	var reply common.KeysReply
	err = client.Call("Store.List", rpcArgs, &reply)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("reply", reply.Keys)
}
