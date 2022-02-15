package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"hash/fnv"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/dr-useless/gobkv/common"
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
}

func handleTest(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("specify a number of keys to set")
	}

	limit, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("specify a number of keys to set")
	}

	client, binding := getClient()
	rpcArgs := common.Args{
		AuthSecret: binding.AuthSecret,
	}

	h := fnv.New64a()

	var wg sync.WaitGroup

	log.Println("working...")
	for i := 0; i < limit; i++ {
		randBytes := make([]byte, 16)
		rand.Read(randBytes)
		h.Write(randBytes)
		rpcArgs.Key = base64.RawStdEncoding.EncodeToString(h.Sum(nil))
		h.Write([]byte("test"))
		rpcArgs.Value = h.Sum(nil)
		h.Reset()
		wg.Add(1)
		go func(rpcArgs *common.Args) {
			var reply common.StatusReply // unused
			client.Call("Store.Set", rpcArgs, &reply)
			wg.Done()
		}(&rpcArgs)
	}
	log.Println("waiting...")
	tStart := time.Now()
	wg.Wait()
	dur := time.Since(tStart)
	log.Printf("done, set %v random keys in %v seconds", limit, dur.Seconds())

}
