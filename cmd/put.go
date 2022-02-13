package cmd

import (
	"log"

	"github.com/dr-useless/gobkv/common"
	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Set a value in the bound gobkv instance",
	Long:  "Usage: gobler put the_key the_value",
	Run:   handlePut,
}

func init() {
	rootCmd.AddCommand(putCmd)
}

func handlePut(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Println("specify a key & value")
		return
	}

	client, binding := getClient()

	rpcArgs := common.Args{
		AuthSecret: binding.AuthSecret,
		Key:        args[0],
		Value:      []byte(args[1]),
	}

	var reply common.StatusReply
	err := client.Call("Store.Put", rpcArgs, &reply)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("reply", string(reply.Status))
}
