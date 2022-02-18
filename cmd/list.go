package cmd

import (
	"log"

	"github.com/dr-useless/gobkv/protocol"
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
	binding := getBinding()
	conn := getConn(binding)

	msg := protocol.Message{
		Op: protocol.OpList,
	}

	if len(args) > 0 {
		msg.Key = args[0]
	}

	/*
		limit, err := cmd.Flags().GetInt("limit")
		if err != nil {
			log.Fatal("limit must be a valid integer")
		}
	*/

	msg.Write(conn)
	msg.Read(conn)
	log.Println(string(msg.Value))
	msg = protocol.Message{
		Op: protocol.OpClose,
	}
	msg.Write(conn)
	conn.Close()
}
