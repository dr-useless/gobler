/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/dr-useless/gobkv/rpc"
	"github.com/spf13/cobra"
)

// putCmd represents the get command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Set a value in the bound gobkv instance",
	Long:  "Usage: gobler put the_key the_value",
	Run:   handlePut,
}

func init() {
	rootCmd.AddCommand(putCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// putCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// putCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func handlePut(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Println("specify a key & value")
		return
	}
	log.Println("handle put", args[0], args[1])

	client, binding, err := getClient()
	if err != nil {
		log.Fatal(err)
	}

	rpcArgs := rpc.Args{
		AuthSecret: binding.AuthSecret,
		Key:        args[0],
		Value:      []byte(args[1]),
	}

	var reply rpc.Result
	client.Call("Store.Put", rpcArgs, &reply)

	log.Println("reply", reply)
}
