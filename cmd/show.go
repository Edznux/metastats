package cmd

import (
	"fmt"
	"github.com/edznux/metastats/config"
	"github.com/edznux/metastats/events"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows you the sum, max, min or all stats from an event",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cfg = config.LoadConfig()
		events.Cfg = cfg
	},
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
		if len(args) != 1 {
			cmd.Help()
			os.Exit(1)
		}
		data, err := events.ReadFromEvent(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sum := getSum(data)
		fmt.Printf("Sum for %s events is : %v\n", args[0], sum)
	},
}

var maxCmd = &cobra.Command{
	Use:   "max",
	Short: "Shows you the max event from an event type",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			os.Exit(1)
		}
		data, err := events.ReadFromEvent(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		max := getMax(data)
		fmt.Printf("Max for %s events is : %v\n", args[0], max)
	},
}

var minCmd = &cobra.Command{
	Use:   "min",
	Short: "Shows you the min event from an event type",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			os.Exit(1)
		}
		data, err := events.ReadFromEvent(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		min := getMin(data)
		fmt.Printf("Min for %s events is : %v\n", args[0], min)
	},
}

func getSum(data [][]string) []int {
	/*
		We calculate only the max FOR EACH ELEMENTS
		We don't try to get the maximum "line" (if there is multiple value)
		So there might not be a line corresponding to the output.
	*/
	sum := []int{}
	for _, d := range data {
		for index, element := range d {
			// skip the timestamp.
			if index == 0 {
				continue
			}
			// the sum will only works on numbers. (integers for now)
			tmp, err := strconv.Atoi(element)
			if err != nil {
				continue
			}
			if len(sum) < index {
				sum = append(sum, 0)
			}
			sum[index-1] += tmp
		}
	}
	return sum
}
func getMax(data [][]string) []int {
	/*
		We calculate only the max FOR EACH ELEMENTS
		We don't try to get the maximum "line" (if there is multiple value)
		So there might not be a line corresponding to the output.
	*/
	max := []int{}
	for _, d := range data {
		for index, element := range d {
			// skip the timestamp.
			if index == 0 {
				continue
			}
			// the max will only works on numbers. (integers for now)
			tmp, err := strconv.Atoi(element)
			if err != nil {
				continue
			}
			if len(max) < index {
				max = append(max, 0)
			}
			if max[index-1] < tmp {
				max[index-1] = tmp
			}
		}
	}
	return max
}

func getMin(data [][]string) []int {
	/*
		We calculate only the min FOR EACH ELEMENTS
		We don't try to get the minimum "line" (if there is multiple value)
		So there might not be a line corresponding to the output.
	*/
	min := []int{}
	for _, d := range data {
		for index, element := range d {
			// skip the timestamp.
			if index == 0 {
				continue
			}
			// the min will only works on numbers. (integers for now)
			tmp, err := strconv.Atoi(element)
			if err != nil {
				continue
			}
			if len(min) < index {
				min = append(min, 9223372036854775807) // maxint (64)
			}
			if min[index-1] > tmp {
				min[index-1] = tmp
			}
		}
	}
	return min
}

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Shows you the sum, max and min event stats from an event",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			os.Exit(1)
		}
		data, err := events.ReadFromEvent(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		min := getMin(data)
		max := getMax(data)
		sum := getSum(data)
		fmt.Printf("Min for %s events is : %v\n", args[0], min)
		fmt.Printf("Max for %s events is : %v\n", args[0], max)
		fmt.Printf("Sum for %s events is : %v\n", args[0], sum)
	},
}

func init() {
	showCmd.AddCommand(sumCmd)
	showCmd.AddCommand(minCmd)
	showCmd.AddCommand(maxCmd)
	showCmd.AddCommand(allCmd)
	rootCmd.AddCommand(showCmd)
}
