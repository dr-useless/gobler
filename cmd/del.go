package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Remove a key from the bound gobkv instance",
	Long:  "Usage: gobler del the_key",
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

	/*
		client, binding := getClient()

		rpcArgs := common.Args{
			AuthSecret: binding.AuthSecret,
			Key:        args[0],
		}

		var reply common.StatusReply
		err := client.Call("Store.Del", rpcArgs, &reply)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("status:", common.MapStatus()[reply.Status])
	*/
}
