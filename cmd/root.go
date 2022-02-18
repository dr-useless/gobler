package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gobler",
	Short: "Make RPCs on a gobkv instance",
	/*
		Long: `A longer description that spans multiple lines and likely contains
		examples and usage of using your application. For example:

		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
			// Uncomment the following line if your bare application
			// has an action associated with it:
			// Run: func(cmd *cobra.Command, args []string) { },
	*/
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
