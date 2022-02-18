package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"hash/fnv"
	"log"
	"strconv"
	"time"

	"github.com/dr-useless/gobkv/protocol"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Populate the bound gobkv instance with test data",
	Long:  "Usage: gobler test [NUMBER_OF_KEYS]",
	Run:   handleTest,
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().Int64("ttl", 0, "number of seconds before key expires")
}

func handleTest(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("specify a number of keys to set")
	}

	limit, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("specify a number of keys to set")
	}

	binding := getBinding()
	conn := getConn(binding)

	ttl, err := cmd.Flags().GetInt64("ttl")
	if err != nil {
		log.Fatal("ttl must be a valid integer")
	}

	h := fnv.New64a()

	msg := protocol.Message{
		Op: protocol.OpSet,
	}

	log.Println("working...")
	tStart := time.Now()

	randBytes := make([]byte, 16)

	for i := 0; i < limit; i++ {
		if ttl > 0 {
			exp := time.Now().Add(time.Duration(ttl) * time.Second)
			msg.Expires = uint64(exp.Unix())
		}

		rand.Read(randBytes)
		h.Write(randBytes)
		msg.Key = base64.RawStdEncoding.EncodeToString(h.Sum(nil))

		h.Write([]byte("test"))
		msg.Value = h.Sum(nil)
		h.Reset()

		err = msg.Write(conn)
		if err != nil {
			log.Fatal("write msg:", err)
		}
	}

	msg = protocol.Message{
		Op: protocol.OpClose,
	}
	err = msg.Write(conn)
	if err != nil {
		log.Fatal("write msg:", err)
	}

	dur := time.Since(tStart)
	log.Printf("done, set %v random keys in %v seconds", limit, dur.Seconds())
}
