package cmd

import (
	"encoding/base64"
	"fmt"
	"hash/fnv"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/intob/gobkv/client"
	"github.com/intob/gobkv/protocol"
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

	ttl, err := cmd.Flags().GetInt64("ttl")
	if err != nil {
		log.Fatal("ttl must be a valid integer")
	}

	b := getBinding()
	conn := getConn(b)
	client := client.NewClient(conn)
	client.Auth(b.AuthSecret)
	authResp := <-client.MsgChan
	if authResp.Status != protocol.StatusOk {
		log.Fatal("unauthorized")
	}

	var exp int64
	if ttl > 0 {
		exp = time.Now().Add(time.Duration(ttl) * time.Second).Unix()
	}

	fmt.Println("working...")
	tStart := time.Now()

	wg := new(sync.WaitGroup)

	for i := 0; i < limit; i++ {
		wg.Add(1)
		go func(exp int64) {
			rand.Seed(time.Now().UnixNano())
			randBytes := make([]byte, 16)
			rand.Read(randBytes)

			h := fnv.New128a()
			h.Write(randBytes)
			key := base64.RawStdEncoding.EncodeToString(h.Sum(nil))

			h.Write([]byte("test"))
			value := h.Sum(nil)

			err := client.Set(key, value, exp, false)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(exp)
	}

	wg.Wait()
	dur := time.Since(tStart)
	fmt.Printf("done, set %v random keys in %v seconds\r\n", limit, dur.Seconds())
}
