package cmd

import (
	"github.com/PerryHuan9/learn_go/router"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start web server.",
	Run: func(cmd *cobra.Command, args []string) {
		router.RegisterRoute()
	},
}
