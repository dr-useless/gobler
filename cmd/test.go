package cmd

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"log"
	"math/rand"
	"strconv"
	"sync"

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

	h := sha1.New()

	var wg sync.WaitGroup

	for i := 0; i < limit; i++ {
		randBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(randBytes, rand.Uint32())
		h.Write(randBytes)
		rpcArgs.Key = base64.RawStdEncoding.EncodeToString(h.Sum(nil))
		h.Write([]byte("test"))
		rpcArgs.Value = h.Sum(nil)
		h.Reset()
		wg.Add(1)
		go func(rpcArgs common.Args) {
			defer wg.Done()
			var reply common.StatusReply // unused
			client.Call("Store.Set", rpcArgs, &reply)
		}(rpcArgs)
	}

	log.Println("working...")
	wg.Wait()
	log.Printf("done, set %v random keys", limit)

}
