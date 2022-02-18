package cmd

import (
	"log"
	"time"

	"github.com/dr-useless/gobkv/protocol"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value in the bound gobkv instance",
	Long:  "Usage: gobler set the_key the_value",
	Run:   handleSet,
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().Uint64("ttl", 0, "number of seconds before key expires")
}

func handleSet(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("specify a key & value")
	}

	b := getBinding()
	conn := getConn(b)

	msg := protocol.Message{
		Op:    protocol.OpSetAck,
		Key:   args[0],
		Value: []byte(args[1]),
	}

	ttl, err := cmd.Flags().GetUint64("ttl")
	if err != nil {
		log.Fatal("ttl must be a valid integer")
	}
	if ttl > 0 {
		expires := time.Now().Add(time.Duration(ttl) * time.Second)
		msg.Expires = uint64(expires.Unix())
	}

	msg.Write(conn)
	msg.Read(conn)

	conn.Close()

	log.Printf("op: %s, status: %s\r\n",
		protocol.MapOp()[msg.Op],
		protocol.MapStatus()[msg.Status])
}
