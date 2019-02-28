package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows you the sum, max, min or all stats from an event",
	Run: func(cmd *cobra.Command, args []string) {
		// Root command does nothing
		cmd.Help()
		os.Exit(1)
	},
}

var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "Shows you the sum of events from an event type",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sum called", args)
	},
}

var maxCmd = &cobra.Command{
	Use:   "max",
	Short: "Shows you the max event from an event type",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("max called", args)
	},
}

var minCmd = &cobra.Command{
	Use:   "min",
	Short: "Shows you the min event from an event type",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("min called", args)
	},
}

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Shows you the sum, max and min event stats from an event",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("all called", args)
	},
}

func init() {
	showCmd.AddCommand(sumCmd)
	showCmd.AddCommand(minCmd)
	showCmd.AddCommand(maxCmd)
	showCmd.AddCommand(allCmd)
	rootCmd.AddCommand(showCmd)
}
