package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about the bound rocketkv instance",
	Long:  "Usage: rkteer info",
	Run:   handleInfo,
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

func handleInfo(cmd *cobra.Command, args []string) {
	b := getBinding()
	fmt.Println("Network:", b.Network)
	fmt.Println("Address:", b.Address)
	if b.CertFile != "" {
		fmt.Println("Cert file:", b.CertFile)
		fmt.Println("Key file:", b.KeyFile)
	}
}
