package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const VERSION = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}
