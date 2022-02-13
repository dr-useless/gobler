package cmd

import (
	"log"

	"github.com/dr-useless/gobkv/common"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a value from the bound gobkv instance",
	Long:  "Usage: gobler get the_key",
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

	client, binding := getClient()

	rpcArgs := common.Args{
		AuthSecret: binding.AuthSecret,
		Key:        args[0],
	}

	var reply common.Result
	client.Call("Store.Get", rpcArgs, &reply)

	log.Println("reply", string(reply.Value))
}
