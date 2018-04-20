package cmd

import (
	"fmt"
	"os"

	"github.com/edznux/metastats/events"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "metastats",
	Short: "Metastats is a simple CLI tool and daemon for self hosted personal life analytics",
	Run: func(cmd *cobra.Command, args []string) {
		events.Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
