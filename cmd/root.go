package cmd

import (
	"fmt"
	"os"

	"github.com/edznux/metastats/config"
	"github.com/edznux/metastats/events"
	"github.com/spf13/cobra"
)

var ConfigPath string
var cfg config.Config

var rootCmd = &cobra.Command{
	Use:   "metastats",
	Short: "Metastats is a simple CLI tool and daemon for self hosted personal life analytics",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.ListOfConfigPath = append([]string{ConfigPath}, config.ListOfConfigPath...)
	},
	Run: func(cmd *cobra.Command, args []string) {
		if os.Geteuid() != 0 {
			fmt.Println("Could not execute this program without root.")
			return
		}
		events.Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&ConfigPath, "config", "", "config path")
}
