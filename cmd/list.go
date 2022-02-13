package cmd

import (
	"log"
	"strconv"

	"github.com/dr-useless/gobkv/common"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all keys with the given prefix",
	Long:  "Usage: gobler list key_prefix",
	Run:   handleList,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func handleList(cmd *cobra.Command, args []string) {
	client, binding := getClient()

	rpcArgs := common.Args{
		AuthSecret: binding.AuthSecret,
		Key:        args[0],
	}

	if len(args) > 1 {
		limit, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("limit arg %s must be an int", args[1])
		}
		rpcArgs.Limit = limit
	}

	var reply common.Result
	client.Call("Store.List", rpcArgs, &reply)

	log.Println("reply", reply.Keys)
}
