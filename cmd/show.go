package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/edznux/metastats/config"
	"github.com/edznux/metastats/events"
	"github.com/spf13/cobra"
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
		max, tsMax := getMax(data)
		fmt.Printf("Max: %d, tsMax : %d\n", max, tsMax)
		fmt.Printf("Max for %s events is : %v (on %s)\n", args[0], max, getTimestampByIndex(data, tsMax))
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
		min, tsMin := getMin(data)
		fmt.Printf("Min for %s events is : %v (on %s)\n", args[0], min, getTimestampByIndex(data, tsMin))
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

func getMax(data [][]string) (max []int, index int) {
	/*
		We calculate only the max FOR EACH ELEMENTS
		We don't try to get the maximum "line" (if there is multiple value)
		So there might not be a line corresponding to the output.
	*/
	max = []int{}
	for _, row := range data {
		for key, value := range row {
			// the max will only works on numbers. (integers for now)
			if key == 0 {
				continue
			}
			tmp, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
			if len(max) < key {
				max = append(max, 0) // maxint (64)
			}
			if max[key-1] < tmp {
				max[key-1] = tmp
				index = key
			}
		}
	}
	return max, index
}

func getMin(data [][]string) (min []int, index int) {
	/*
		We calculate only the min FOR EACH ELEMENTS
		We don't try to get the minimum "line" (if there is multiple value)
		So there might not be a line corresponding to the output.
	*/
	min = []int{}
	for _, row := range data {
		for key, value := range row {
			// the min will only works on numbers. (integers for now)
			if key == 0 {
				continue
			}
			tmp, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
			if len(min) < key {
				min = append(min, 9223372036854775807) // maxint (64)
			}
			if min[key-1] > tmp {
				min[key-1] = tmp
				index = key
			}
		}
	}
	return min, index
}
func getTimestampByIndex(data [][]string, index int) string {
	ts, err := strconv.Atoi(data[index][0])
	if err != nil {
		fmt.Printf("Could not parse timestamp : %s", err)
	}
	fmt.Println("Ts : ", ts)
	res, err := time.Unix(int64(ts), 0).MarshalText()
	if err != nil {
		return "Could not get first entry time"
	}
	// fmt.Println(data[index][0], index, string(res))
	return string(res)
}

func getFirst(data [][]string) string {
	return getTimestampByIndex(data, 0)
}

func getLast(data [][]string) string {
	return getTimestampByIndex(data, len(data)-1)
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
		min, tsMin := getMin(data)
		max, tsMax := getMax(data)
		sum := getSum(data)
		first := getFirst(data)
		last := getFirst(data)
		fmt.Printf("First event for %s is : %v\n", args[0], first)
		fmt.Printf("Last event for %s is : %v\n", args[0], last)
		fmt.Printf("Min for %s events is : %v (on %s)\n", args[0], min, getTimestampByIndex(data, tsMin))
		fmt.Printf("Max for %s events is : %v (on %s)\n", args[0], max, getTimestampByIndex(data, tsMax))
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
