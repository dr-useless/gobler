package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/intob/rocketkv/client"
	"github.com/intob/rocketkv/protocol"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value in the bound rocketkv instance",
	Long:  "Usage: rkteer set the_key the_value",
	Run:   handleSet,
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().Int64("ttl", 0, "number of seconds before key expires")
}

func handleSet(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("specify a key & value")
	}

	ttl, err := cmd.Flags().GetInt64("ttl")
	if err != nil {
		log.Fatal("ttl must be a valid integer")
	}
	var expires int64
	if ttl > 0 {
		expires = time.Now().Add(time.Duration(ttl) * time.Second).Unix()
	}

	b := getBinding()
	conn := getConn(b)
	client := client.NewClient(conn)
	if b.AuthSecret != "" {
		client.Auth(b.AuthSecret)
		authResp := <-client.Msgs
		if authResp.Status != protocol.StatusOk {
			log.Fatal("unauthorized")
		}
	}

	err = client.Set(args[0], []byte(args[1]), expires, true)
	if err != nil {
		log.Fatal(err)
	}

	resp := <-client.Msgs

	fmt.Println(protocol.MapStatus()[resp.Status])
}
