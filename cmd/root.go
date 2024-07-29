package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "learn-go",
	Short: "learn-go is a project whick build a web application with golang.",
	Long:  `learn-go use many usage skill`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here

	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(serverCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
